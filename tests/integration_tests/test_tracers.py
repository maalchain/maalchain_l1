import itertools
import json
from concurrent.futures import ThreadPoolExecutor, as_completed

from web3 import Web3

from .expected_constants import (
    EXPECTED_CALLTRACERS,
    EXPECTED_CONTRACT_CREATE_TRACER,
    EXPECTED_STRUCT_TRACER,
)
from .utils import (
    ADDRS,
    CONTRACTS,
    deploy_contract,
    derive_new_account,
    send_transaction,
    w3_wait_for_new_blocks,
)


def test_tracers(ethermint_rpc_ws):
    w3: Web3 = ethermint_rpc_ws.w3
    eth_rpc = w3.provider
    gas_price = w3.eth.gas_price
    tx = {"to": ADDRS["community"], "value": 100, "gasPrice": gas_price}
    tx_hash = send_transaction(w3, tx)["transactionHash"].hex()
    method = "debug_traceTransaction"
    tracer = {"tracer": "callTracer"}
    tx_res = eth_rpc.make_request(method, [tx_hash])
    assert tx_res["result"] == EXPECTED_STRUCT_TRACER, ""
    tx_res = eth_rpc.make_request(method, [tx_hash, tracer])
    assert tx_res["result"] == EXPECTED_CALLTRACERS, ""
    tx_res = eth_rpc.make_request(
        method,
        [tx_hash, tracer | {"tracerConfig": {"onlyTopCall": True}}],
    )
    assert tx_res["result"] == EXPECTED_CALLTRACERS, ""
    _, tx = deploy_contract(w3, CONTRACTS["TestERC20A"])
    tx_hash = tx["transactionHash"].hex()
    w3_wait_for_new_blocks(w3, 1)
    tx_res = eth_rpc.make_request(method, [tx_hash, tracer])
    tx_res["result"]["to"] = EXPECTED_CONTRACT_CREATE_TRACER["to"]
    assert tx_res["result"] == EXPECTED_CONTRACT_CREATE_TRACER, ""


def test_crosscheck(ethermint, geth):
    method = "debug_traceTransaction"
    tracer = {"tracer": "callTracer"}
    acc = derive_new_account(4)
    sender = acc.address
    # fund new sender to deploy contract with same address
    fund = 3000000000000000000
    tracers = [
        [],
        [tracer],
        [tracer | {"tracerConfig": {"onlyTopCall": True}}],
        [tracer | {"tracerConfig": {"withLog": True}}],
        [tracer | {"tracerConfig": {"diffMode": True}}],
    ]
    iterations = 1

    def process(w3):
        tx = {"to": sender, "value": fund, "gasPrice": w3.eth.gas_price}
        send_transaction(w3, tx)
        assert w3.eth.get_balance(sender, "latest") == fund
        contract, _ = deploy_contract(w3, CONTRACTS["TestMessageCall"], key=acc.key)
        tx = contract.functions.test(iterations).build_transaction()
        tx_hash = send_transaction(w3, tx)["transactionHash"].hex()
        res = []
        call = w3.provider.make_request
        with ThreadPoolExecutor(len(tracers)) as exec:
            params = [([tx_hash] + cfg) for cfg in tracers]
            exec_map = exec.map(call, itertools.repeat(method), params)
            res = [json.dumps(resp["result"], sort_keys=True) for resp in exec_map]
        return res

    providers = [ethermint.w3, geth.w3]
    with ThreadPoolExecutor(len(providers)) as exec:
        tasks = [
            exec.submit(process, w3) for w3 in providers
        ]
        res = [future.result() for future in as_completed(tasks)]
        assert len(res) == len(providers)
        assert res[0] == res[1], res
