package wazuh

default ignore = false

ignore {
  path := input.syscheck.path
  path == "/etc/test.txt"
}
