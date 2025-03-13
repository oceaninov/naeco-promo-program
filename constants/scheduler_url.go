package constants

import "fmt"

const server = `gcp-nb-sbox01`

var UpdateProgramToActiveURL = fmt.Sprintf("http://%s:50009/v1/program/change-status/", server)
