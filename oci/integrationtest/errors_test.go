// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/oci/globalvar"
	"github.com/terraform-providers/terraform-provider-oci/oci/tfresource"

	"github.com/stretchr/testify/assert"
)

// issue-routing-tag: terraform/default
func TestUnitGetVersionAndDateError(t *testing.T) {
	versionError := tfresource.GetVersionAndDateError()
	assert.Contains(t, versionError, "Provider version: ")
	assert.Contains(t, versionError, globalvar.Version)
	assert.Contains(t, versionError, globalvar.ReleaseDate)
	assert.NotContains(t, versionError, "Update(s) behind to current")
}
