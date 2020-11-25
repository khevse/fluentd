## Test environment for fluentd

1. Run test environment with kibana and elasticsearch
```bash
make run-env
```
2. Upgrade fluentd config in fluent.conf
3. Run test with new fluentd config
```bash
make
```