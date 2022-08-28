package StashImportCycle

type DeleteLoto6PredictionsDetailStash struct {
	Body struct {
		UserID      string              `json:"user_id"`
		Predictions []map[string]string `json:"predictions"`
	} `json:"body"`
}
