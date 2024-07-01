local default = import 'default.jsonnet';

default {
  'maalchain_7862-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
