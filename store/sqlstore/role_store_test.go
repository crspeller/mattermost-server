// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/crspeller/mattermost-server/store/storetest"
)

func TestRoleStore(t *testing.T) {
	StoreTest(t, storetest.TestRoleStore)
}
