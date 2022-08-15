package StashImportCycle

type RegisterLoto7Predictions struct {
	Body struct {
		UserID      string              `json:"user_id"`
		Predictions []map[string]string `json:"predictions"`
	} `json:"body"`
}
