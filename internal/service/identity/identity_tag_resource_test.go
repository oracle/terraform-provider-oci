// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestIdentityTagResourceValidatorValues(t *testing.T) {
	tests := []struct {
		name           string
		stateValidator []interface{}
		validator      []interface{}
		wantErrSubstr  string
	}{
		{
			name: "validator omitted",
		},
		{
			name: "ENUM with values",
			validator: []interface{}{map[string]interface{}{
				"validator_type": "ENUM",
				"values":         []interface{}{"value1", "value2"},
			}},
		},
		{
			name: "ENUM without values",
			validator: []interface{}{map[string]interface{}{
				"validator_type": "ENUM",
			}},
			wantErrSubstr: "must contain at least one value",
		},
		{
			name: "DEFAULT without values",
			validator: []interface{}{map[string]interface{}{
				"validator_type": "DEFAULT",
			}},
		},
		{
			name: "DEFAULT with values",
			validator: []interface{}{map[string]interface{}{
				"validator_type": "DEFAULT",
				"values":         []interface{}{"value1"},
			}},
			wantErrSubstr: "must be omitted",
		},
		{
			name: "ENUM to DEFAULT without values",
			stateValidator: []interface{}{map[string]interface{}{
				"validator_type": "ENUM",
				"values":         []interface{}{"value1"},
			}},
			validator: []interface{}{map[string]interface{}{
				"validator_type": "DEFAULT",
			}},
		},
		{
			name: "DEFAULT to ENUM without values",
			stateValidator: []interface{}{map[string]interface{}{
				"validator_type": "DEFAULT",
			}},
			validator: []interface{}{map[string]interface{}{
				"validator_type": "ENUM",
			}},
			wantErrSubstr: "must contain at least one value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := map[string]interface{}{
				"description":      "test tag",
				"name":             "test-tag",
				"tag_namespace_id": "ocid1.tagnamespace.oc1..exampleuniqueID",
			}
			if tt.validator != nil {
				config["validator"] = tt.validator
			}

			resource := IdentityTagResource()
			var state *terraform.InstanceState
			if tt.stateValidator != nil {
				stateData := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
					"description":      "test tag",
					"name":             "test-tag",
					"tag_namespace_id": "ocid1.tagnamespace.oc1..exampleuniqueID",
					"validator":        tt.stateValidator,
				})
				stateData.SetId("ocid1.tagdefinition.oc1..exampleuniqueID")
				state = stateData.State()
			}

			_, err := resource.Diff(context.Background(), state, terraform.NewResourceConfigRaw(config), nil)

			if tt.wantErrSubstr == "" {
				if err != nil {
					t.Fatalf("Resource.Diff() unexpected error: %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("Resource.Diff() expected error containing %q", tt.wantErrSubstr)
			}
			if !strings.Contains(err.Error(), tt.wantErrSubstr) {
				t.Fatalf("Resource.Diff() error = %q, want substring %q", err, tt.wantErrSubstr)
			}
		})
	}
}
