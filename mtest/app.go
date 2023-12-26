package main

import (
	"context"
	"encoding/json"
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

	// Add key-value elements to the config map
	config[key] = append(values, config[key].([]string)...)[:limit]

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

	return config[key].([]string)
}
