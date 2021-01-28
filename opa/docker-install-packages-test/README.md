# Verify the packages in Docker Image with conftest

```shell
> trivy -q i --list-all-pkgs -f json ubuntu:20.10 | jq -r '[.[].Packages[].Name]' > ubuntu-20.10-packages.json

> conftest test ubuntu-20.10-packages.json

100 tests, 100 passed, 0 warnings, 0 failures, 0 exceptions

> conftest test ubuntu-20.10-disallowed-packages.json
FAIL - ubuntu-20.10-disallowed-packages.json - main - (sudo)

101 tests, 100 passed, 0 warnings, 1 failure, 0 exceptions
```
