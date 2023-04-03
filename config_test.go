package main

import (
	"SSPS/config"
	"fmt"
	"testing"
)

func TestReadPanelConfig(t *testing.T) {
	pc := config.PanelConfig{}

	pc.ReadPanelConfig()

	fmt.Println(pc)
}
