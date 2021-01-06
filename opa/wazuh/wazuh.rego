package wazuh

default ignore = false

ignore {
  full_log := input.full_log
  contains(full_log, "test.service")
}

