local config = import 'default.jsonnet';

config {
  'maalchain_7862-1'+: {
    'app-config'+: {
      'minimum-gas-prices': '100000000000aphoton',
    },
    genesis+: {
      app_state+: {
        feemarket+: {
          params+: {
            base_fee:: super.base_fee,
          },
        },
        gov: {
          voting_params: {
            voting_period: '10s',
          },
          deposit_params: {
            max_deposit_period: '10s',
            min_deposit: [
              {
                denom: 'aphoton',
                amount: '1',
              },
            ],
          },
        },
      },
    },
  },
}
