// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"
)

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

func Test_resourceDiscoveryBaseStep_mergeGeneratedStateFile(t *testing.T) {
	outputDir := fmt.Sprintf("/Users/papakaur/go/src/github.com/terraform-providers/terraform-provider-oci/bin")

	tfHclVersion = &TfHclVersion12{}
	args := &ExportCommandArgs{
		OutputDir: &outputDir,
	}

	// verify executable from system path
	if _, _, err := createTerraformStruct(args); err != nil {
		t.Errorf("createTerraformStruct() error = %v", err)
		t.Fail()
	}

	type fields struct {
		ctx                 *resourceDiscoveryContext
		name                string
		discoveredResources []*OCIResource
		omittedResources    []*OCIResource
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				name: "core",
				ctx: &resourceDiscoveryContext{
					ExportCommandArgs: args,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &resourceDiscoveryBaseStep{
				ctx:  tt.fields.ctx,
				name: tt.fields.name,
				//discoveredResources: tt.fields.discoveredResources,
				//omittedResources:    tt.fields.omittedResources,
			}
			if err := r.mergeGeneratedStateFile(); (err != nil) != tt.wantErr {
				t.Errorf("resourceDiscoveryBaseStep.mergeGeneratedStateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
