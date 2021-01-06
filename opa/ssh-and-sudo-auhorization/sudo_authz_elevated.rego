package sudo.authz

import data.elevate
import input.sysinfo
import input.display_responses

# Allow this user if the elevation ticket they provided matches our mock API
# of an internal elevation system.
allow {
  elevate.tickets[sysinfo.pam_username] == display_responses.ticket
}
