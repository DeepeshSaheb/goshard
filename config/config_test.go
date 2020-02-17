package config

import "testing"

func TestValidateConfig(t *testing.T) {
	var nilConfig []Config
	err := validateConfig(nilConfig)
	if err == nil {
		t.Error("validateConfig should have returned error")
	}

}

func TestLoadConfig(t *testing.T) {
	LoadConfig("example_config.json")
	loadedConfig := GetConfig()
	if len(loadedConfig) == 0 {
		t.Error("Config values should be loaded")
	}

}
