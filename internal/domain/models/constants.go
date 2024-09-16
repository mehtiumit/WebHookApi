package models

const (
	HookStatusInitial = "initial"
	HookStatusSent    = "sent"
	HookStatusFailed  = "failed"
	ActionStart       = "start"
	ActionEnd         = "stop"
)

var (
	ValidActions = []string{ActionStart, ActionEnd}
)
