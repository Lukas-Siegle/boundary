// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package apptoken

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_getOpts(t *testing.T) {
	t.Parallel()
	testCtx := context.Background()
	t.Run("WithName", func(t *testing.T) {
		assert := assert.New(t)
		opts, err := getOpts(WithName(testCtx, "test"))
		require.NoError(t, err)
		testOpts := getDefaultOptions()
		assert.NotEqual(opts, testOpts)
		testOpts.withName = "test"
		assert.Equal(opts, testOpts)
	})
	t.Run("WithDescription", func(t *testing.T) {
		assert := assert.New(t)
		opts, err := getOpts(WithDescription(testCtx, "test"))
		require.NoError(t, err)
		testOpts := getDefaultOptions()
		assert.NotEqual(opts, testOpts)
		testOpts.withDescription = "test"
		assert.Equal(opts, testOpts)
	})
	t.Run("WithLimit", func(t *testing.T) {
		assert := assert.New(t)
		opts, err := getOpts(WithLimit(testCtx, 50))
		require.NoError(t, err)
		testOpts := getDefaultOptions()
		assert.NotEqual(opts, testOpts)
		testOpts.withLimit = 50
		assert.Equal(opts, testOpts)
	})
	t.Run("WithExpirationInterval", func(t *testing.T) {
		assert, require := assert.New(t), require.New(t)
		opts, err := getOpts(WithExpirationInterval(testCtx, 10))
		require.NoError(err)
		testOpts := getDefaultOptions()
		assert.NotEqual(opts, testOpts)
		testOpts.withExpirationInterval = 10
		assert.Equal(testOpts, opts)
	})
	t.Run("invalid-WithExpirationInterval", func(t *testing.T) {
		assert, require := assert.New(t), require.New(t)
		opts, err := getOpts(WithExpirationInterval(testCtx, 0))
		require.Error(err)
		assert.Empty(opts)
	})
}
