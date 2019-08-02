package api

import (
	"encoding/json"
	"io"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/radiohive/zimt/pkg/structs"
)

// BridgeConfig represents message of `#/bridge/config` topic
type BridgeConfig struct {
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Coordinator int    `json:"coordinator"`
	LogLevel    string `json:"log_level"`
	PermitJoin  bool   `json:"permit_join"`
}

// Print prints bridge config to standard output
func (bc BridgeConfig) Print(writer io.Writer) {
	structs.Print(&bc, "json", writer)
}

// GetBridgeConfig subscribes and returns message  from `#/bridge/config` topic
func GetBridgeConfig(client mqtt.Client) BridgeConfig {
	topic := topic("bridge/config")
	msg := getSubscribedOnce(client, topic)
	config := BridgeConfig{}
	json.Unmarshal(msg.Payload(), &config)
	return config
}
