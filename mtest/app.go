package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

const configFilePath = "config.json"

// App struct
type App struct {
	ctx context.Context
}

// Mlib creates a new App application struct
func Mlib() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// 設定の保存
func (a *App) SaveConfig(key string, values []string, limit int) error {
	// Unmarshal the JSON data into a map
	println(configFilePath)
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		data = []byte("{}")
	}

	// Unmarshal the JSON data into a map
	config := make(map[string]interface{})
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	// Check if the value is nil and initialize it as an empty slice of strings
	if config[key] == nil {
		config[key] = []string{}
	}

	// Convert the interface{} type to []string
	existingValues, ok := config[key].([]string)
	if !ok {
		return fmt.Errorf("unexpected type for config[%s]", key)
	}

	// Add key-value elements to the config map
	config[key] = append(values, existingValues...)[:limit]

	// Marshal the updated config map back to JSON
	updatedData, err := json.MarshalIndent(config, "", "")
	if err != nil {
		return err
	}

	// Overwrite the config.json file with the updated JSON data
	err = os.WriteFile(configFilePath, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// 設定の読み込み
func (a *App) LoadConfig(key string) []string {
	// Read the config.json file
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		data = []byte("{}")
	}

	// Unmarshal the JSON data into a map
	config := make(map[string]interface{})
	err = json.Unmarshal(data, &config)
	if err != nil {
		return []string{}
	}

	// Check if the value is nil and initialize it as an empty slice of strings
	if config[key] == nil {
		return []string{}
	}

	// Convert the interface{} type to []interface{}
	values, ok := config[key].([]interface{})
	if !ok {
		return []string{}
	}

	// Convert each element to string
	result := make([]string, len(values))
	for i, v := range values {
		if str, ok := v.(string); ok {
			result[i] = str
		} else {
			return []string{}
		}
	}

	return result
}
