package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/DmitryVesenniy/auth-microservice/internal/config"
)

func TestConfig(t *testing.T) {
	os.Setenv("CONFIG", "../configs/test.yaml")

	cfg, err := config.ConfigLoad()

	if err != nil {
		t.Error(err)
	}

	if cfg.Database.EngineType == "" {
		t.Error("error init: EngineType")
	}

	fmt.Printf("config: %+v\n", cfg)
}
