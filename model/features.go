// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type FeatureFlags struct {
	MyTestFeature bool
}

func (f *FeatureFlags) SetDefaults() {
	f.MyTestFeature = false
}
