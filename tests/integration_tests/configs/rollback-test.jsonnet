local config = import 'default.jsonnet';

config {
  'maalchain_7862-1'+: {
    validators: super.validators[0:1] + [{
      name: 'fullnode',
    }],
  },
}
