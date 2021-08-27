package host

import (
	"testing"

	"github.com/hashicorp/boundary/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_TestPlugins(t *testing.T) {
	assert, require := assert.New(t), require.New(t)
	conn, _ := db.TestSetup(t, "postgres")

	plg := TestPlugin(t, conn, "test", "prefix")
	require.NotNil(plg)
	assert.NotEmpty(plg.GetPublicId())
}
