// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils_test

import (
	"testing"

	"github.com/crspeller/mattermost-server/testlib"
)

var mainHelper *testlib.MainHelper

func TestMain(m *testing.M) {
	mainHelper = testlib.NewMainHelperWithOptions(nil)
	defer mainHelper.Close()

	mainHelper.Main(m)
}
