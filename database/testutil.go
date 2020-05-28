package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

// ResetTable User data replay repo to clear the table and copy the schema from test environment
func TestResetTable() error {
	dataResetPath := os.Getenv("GIERKINET_DATA_REPLAY_PATH")
	cmd := exec.Command("make", "-C", dataResetPath, "start-dev-database")
	out, err := cmd.CombinedOutput()
	fmt.Printf("\n%s\n%s\n", err, string(out))
	return err
}

// SetTableState uses data replay repo to inject data into the table
func TestSetTableState(state string) error {
	dataResetPath := os.Getenv("GIERKINET_DATA_REPLAY_PATH")
	cmd := exec.Command("make", "-C", dataResetPath, fmt.Sprintf("data-replay-%s", state))
	out, err := cmd.CombinedOutput()
	fmt.Printf("errror:\n%s\ncombined out:\n%s\n", err, string(out))
	return err
}

func IntegrationTest(t *testing.T, state string) (err error) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	err = TestResetTable()

	assert.NoError(t, err, "Failed to reset table")
	if err != nil {
		t.FailNow()
		return
	}

	if state != "" {
		err = TestSetTableState(state)

		assert.NoError(t, err, "Failed to set table state")
		if err != nil {
			t.FailNow()
			return
		}
	}

	return
}
