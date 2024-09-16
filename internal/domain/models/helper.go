package models

func IsActionValid(action string) bool {
	for _, validAction := range ValidActions {
		if action == validAction {
			return true
		}
	}
	return false
}
