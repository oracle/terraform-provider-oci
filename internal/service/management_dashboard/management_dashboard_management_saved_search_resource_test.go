// Copyright (c) 2026, Oracle and/or its affiliates.
// Licensed under the Mozilla Public License Version 2.0

package management_dashboard

import "testing"

func TestUnitManagementSavedSearchIDWithoutResourceDataReturnsEmpty(t *testing.T) {
	crud := &ManagementDashboardManagementSavedSearchResourceCrud{}
	if got := crud.ID(); got != "" {
		t.Fatalf("ID() = %q, want empty ID", got)
	}
}
