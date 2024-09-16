package hook

type CreateHookRequest struct {
	To        string `json:"to"`
	ContentId string `json:"contentId"`
	Action    string `json:"action"`
}
