package main

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/yrisob/journal_log/config"
	"github.com/yrisob/journal_log/utils"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func TestCommandRun(t *testing.T) {
	result, err := utils.CallSystemCommand("id", "-u", "yrisob")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestConfig(t *testing.T) {
	config, err := config.GetServerConfig()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(config.PPAS)
	t.Log(config.PPASEHED)
	t.Log(config.PPASNOTIFIER)

}
