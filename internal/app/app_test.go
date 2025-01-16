package app

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

const wantJSON = `[{"gid":"1209167689838486","name":"Example Project 2","resource_type":"project"},{"gid":"1209167685622643","name":"Example Project","resource_type":"project"}]`

func TestAppSuccess(t *testing.T) {
	app, err := NewApp()
	require.NoError(t, err)

	projects, err := app.projectSvc.Get()
	require.NoError(t, err)

	marshalled, err := json.Marshal(projects)
	require.NoError(t, err)

	require.JSONEq(t, wantJSON, string(marshalled))
}
