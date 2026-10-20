// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"testing"

	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func TestSubsettingRuleFieldMatchesCanonicalizesAPIValues(t *testing.T) {
	schemaName := " HR_TEST "
	objectName := " employees "
	actualSchemaName := "HR_TEST"
	actualObjectName := "EMPLOYEES"

	if !subsettingRuleFieldMatches(
		oci_data_safe.SubsetScopeForSpecificObjects{SchemaName: &schemaName, ObjectName: &objectName},
		oci_data_safe.SubsetScopeForSpecificObjects{SchemaName: &actualSchemaName, ObjectName: &actualObjectName},
	) {
		t.Fatal("expected normalized scopes to match")
	}

	if !subsettingRuleFieldMatches(
		oci_data_safe.PartitionSubsetRuleEntry{PartitionsList: []string{" P1 ", "P1", "P2"}},
		oci_data_safe.PartitionSubsetRuleEntry{PartitionsList: []string{"P2", "P1"}},
	) {
		t.Fatal("expected duplicate and reordered partitions to match")
	}
}
