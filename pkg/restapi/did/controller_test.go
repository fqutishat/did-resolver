/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package did

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/trustbloc/did-resolver/pkg/restapi/did/operation"
)

func TestProxy_New(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		controller, err := New(&operation.Config{
			RuleProvider: &mockRuleProvider{}})
		require.NoError(t, err)
		require.NotNil(t, controller)
	})
}

func TestProxy_GetOperations(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		controller, err := New(&operation.Config{
			RuleProvider: &mockRuleProvider{}})
		require.NotNil(t, controller)
		require.NoError(t, err)

		ops := controller.GetOperations()

		require.Equal(t, 1, len(ops))
	})
}

type mockRuleProvider struct {
	Err error
}

func (p *mockRuleProvider) Transform(uri string) (string, error) {
	if p.Err != nil {
		return "", p.Err
	}

	return uri, nil
}
