package sudo.authz

# By default, users are not authorized.
default allow = false

# Allow access to any user that has the "admin" role.
allow {
    data.roles["admin"][_] == input.sysinfo.pam_username
}

# If the user is not authorized, then include an error message in the response.
errors["Request denied by administrative policy"] {
    not allow
}
