package StashImportCycle

type RegisterLoto7Expectations struct {
	Body struct {
		UserID       string              `json:"user_id"`
		Expectations []map[string]string `json:"expectations"`
	} `json:"body"`
}
