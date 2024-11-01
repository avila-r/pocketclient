package pocketclient_test

import (
	"testing"

	"github.com/avila-r/pocketclient"
	"github.com/avila-r/pocketclient/internal/tests"
)

func Test_ListLogs(t *testing.T) {
	client := tests.PocketClient

	logs, err := client.GetLogs()

	if err != nil {
		t.Errorf("error while listing logs - %v", err.Error())
	}

	t.Log(logs)
}

func Test_GetLogs(t *testing.T) {
	client := tests.PocketClient

	logs, err := client.GetLogs()

	if err != nil {
		t.Errorf("error while listing logs - %v", err.Error())
	}

	var target pocketclient.Log
	if len(logs.Items) > 0 {
		target = logs.Items[0]
	}

	if _, err := client.GetLogByID(target.ID); err != nil {
		t.Errorf("error while getting log - %v", err.Error())
	}
}

func Test_GetLogsStats(t *testing.T) {
	client := tests.PocketClient

	if _, err := client.GetLogggingStats(); err != nil {
		t.Errorf("error while getting logs' stats - %v", err.Error())
	}
}
