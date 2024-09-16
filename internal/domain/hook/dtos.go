package hook

type CreateHookRequestDto struct {
	To        string `json:"to"`
	ContentId string `json:"content"`
	Action    string `json:"action"`
}

type HookDto struct {
	To      string `json:"to"`
	Content string `json:"content"`
}
