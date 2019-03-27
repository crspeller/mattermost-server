// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/crspeller/mattermost-server/store/storetest"
)

func TestCommandWebhookStore(t *testing.T) {
	StoreTest(t, storetest.TestCommandWebhookStore)
}
