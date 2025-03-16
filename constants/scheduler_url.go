package constants

import "fmt"

const server = `gcp-nb-sbox01`

const (
	ProgramInactive = 0
	ProgramActive = 1
	ProgramTerminate = 2
)

var UpdateProgramStatusURL = fmt.Sprintf("http://%s:50009/v1/program/change-status/", server)
var RefreshDailyQuotaURL = fmt.Sprintf("http://%s:50009/v1/whitelist/change-max-limit/",server)