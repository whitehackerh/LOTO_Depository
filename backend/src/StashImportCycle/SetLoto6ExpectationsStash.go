package StashImportCycle

type RegisterLoto6Expectations struct {
	Body struct {
		UserID       string              `json:"user_id"`
		Expectations []map[string]string `json:"expectations"`
	} `json:"body"`
}
