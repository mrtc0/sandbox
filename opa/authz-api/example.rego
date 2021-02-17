package httpapi.authz

import input

default allow = false

allow {
  some username
  input.method == "GET"
  input.path = ["user", username]
  input.user == username
}
