// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resourcediscovery

import (
	"testing"
)

// issue-routing-tag: terraform/default
func TestUnit_getValidDbVersion(t *testing.T) {
	type args struct {
		dbVersion string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"version_pre_18",
			args{"11.2.0.4.200414"},
			"11.2.0.4.200414",
		},
		{
			"version_18_no_psu",
			args{"18.8.0.0"},
			"18.8.0.0",
		},
		{
			"version_19_no_psu",
			args{"19.8.0.0"},
			"19.8.0.0",
		},
		{
			"version_19_with_psu",
			args{"19.5.0.0.200414"},
			"19.5.0.0",
		},
		{
			"version_20_with_psu",
			args{"20.0.0.0.200414"},
			"20.0.0.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getValidDbVersion(tt.args.dbVersion); got != tt.want {
				t.Errorf("getValidDbVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
