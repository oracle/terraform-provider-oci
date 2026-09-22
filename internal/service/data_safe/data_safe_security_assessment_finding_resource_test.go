// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package data_safe

import "testing"

func TestUnitSecurityAssessmentFindingReferencesAreResponseOnly(t *testing.T) {
	references := DataSafeSecurityAssessmentFindingResource().Schema["references"]
	if references == nil {
		t.Fatal("references schema is missing")
	}
	if !references.Computed || references.Optional || references.Required {
		t.Fatalf("references schema must remain response-only: computed=%t optional=%t required=%t", references.Computed, references.Optional, references.Required)
	}
}
