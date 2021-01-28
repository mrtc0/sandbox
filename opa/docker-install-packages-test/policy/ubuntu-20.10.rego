package main

import data.packages

# expect_packages[p] { p := packages.allowed_packages[_] }
# actual_packages[p] { p := input.packages[_] }
# 
# p := expect_packages - actual_packages
contains(packages, name) {
  packages[_] = name
}

deny[msg] {
  not contains(packages.allowed_packages, input)
  msg := sprintf("(%v)", [input])
}
