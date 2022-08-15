package StashImportCycle

type RegisterLoto6Predictions struct {
	Body struct {
		UserID      string              `json:"user_id"`
		Predictions []map[string]string `json:"predictions"`
	} `json:"body"`
}
