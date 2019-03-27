package sqlstore

import (
	"testing"

	"github.com/crspeller/mattermost-server/store/storetest"
)

func TestTermsOfServiceStore(t *testing.T) {
	StoreTest(t, storetest.TestTermsOfServiceStore)
}
