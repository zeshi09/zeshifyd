package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zeshi09/zeshifyd/internal/model"
)

func SaveNotification(n model.Notification) error {
	path := os.Getenv("HOME") + "/.cache/zeshifyd/zeshifyd.json"

	// if dir exists
	if err := os.MkdirAll(os.Getenv("HOME")+"/.cache/zeshifyd", 0755); err != nil {
		return fmt.Errorf("failed to create dir: %w", err)
	}

	// serialize data to json
	data, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	// writing to file
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
