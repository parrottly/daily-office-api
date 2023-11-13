package internal

import (
	"encoding/json"
	"os"

	"dolapi/models"
)

func ReadJSONFile(filePath string) ([]models.LiturgicalData, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data []models.LiturgicalData
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data, nil
}
