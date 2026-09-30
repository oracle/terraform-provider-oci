// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCreateSubsettingSchemaRelationRequestIncludesObjectKeys(t *testing.T) {
	d := schema.TestResourceDataRaw(t, DataSafeSubsettingPolicySubsettingSchemaRelationResource().Schema, map[string]interface{}{
		"child_columns":        []interface{}{"CHILD_ID"},
		"child_object_key":     "child-key",
		"child_object_name":    "CHILD_TABLE",
		"child_schema_name":    "CHILD_SCHEMA",
		"parent_columns":       []interface{}{"PARENT_ID"},
		"parent_object_key":    "parent-key",
		"parent_object_name":   "PARENT_TABLE",
		"parent_schema_name":   "PARENT_SCHEMA",
		"subsetting_policy_id": "policy-id",
	})

	crud := &DataSafeSubsettingPolicySubsettingSchemaRelationResourceCrud{}
	crud.D = d
	request, err := crud.createSubsettingSchemaRelationRequest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if request.ChildObjectKey == nil || *request.ChildObjectKey != "child-key" {
		t.Fatalf("child object key was not mapped: %#v", request.ChildObjectKey)
	}
	if request.ParentObjectKey == nil || *request.ParentObjectKey != "parent-key" {
		t.Fatalf("parent object key was not mapped: %#v", request.ParentObjectKey)
	}
}
