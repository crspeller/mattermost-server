// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package feature

/*import (
	"github.com/splitio/go-client/splitio/client"
	"github.com/splitio/go-client/splitio/conf"
)

const (
	EnabledState  = "on"
	DisabledState = "off"
)

type Features struct {
	MyTestFeature bool `access:"write_restrictable"`
}

type Manager struct {
	serverID string
	values   map[string]interface{}
	client   *client.SplitClient
}

func NewManager(id string) (*Manager, error) {
	cfg := conf.Default()
	factory, err := client.NewSplitFactory("eqnesn4om7269uf9d4naj7gh67mign2sftck", cfg)
	if err != nil {
		return nil, err
	}

	client := factory.Client()
	err = client.BlockUntilReady(10)
	if err != nil {
		return nil, err
	}

	features := &Manager{
		serverID: id,
		values:   make(map[string]interface{}),
		client:   client,
	}

	return features, nil
}*/

/*func (f *FeaturesManager) Enabled(key string) bool {
	treatment := f.client.Treatment(f.serverID, key, nil)
	if treatment == "on" {
		return true
	}
	return false
}

func (f *FeaturesManager) EnabledWithFallback(key string, fallback bool) bool {
	return false
}

func (f *FeaturesManager) State(key string) string {
	return ""
}

func (f *FeaturesManager) StateWithFallback(key, fallback string) string {
	return ""
}*/
