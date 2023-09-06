package main

import (
	"encoding/json"
	"os"
)

type TTConfig struct {
	GoogleDrive  GoogleDriveConfig  `json:"google_drive"`
	VideoStorage VideoStorageConfig `json:"video_storage"`
}

type GoogleDriveConfig struct {
	CredentialsPath string `json:"credentials_path"`
}

type VideoStorageConfig struct {
	Path string `json:"path"`
}

// LoadSettings : Return server settings
func LoadConfig(config_path string) (*TTConfig, error) {
	var config TTConfig
	confBytes, err := os.ReadFile(config_path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(confBytes, &config)
	if err != nil {
		return nil, err
	}
	//Settings = &serverSettings
	return &config, nil
}
