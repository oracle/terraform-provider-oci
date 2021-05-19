// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitGetVersionAndDateError(t *testing.T) {
	versionError := getVersionAndDateError()
	assert.Contains(t, versionError, "Provider version: ")
	assert.Contains(t, versionError, Version)
	assert.Contains(t, versionError, ReleaseDate)
	assert.NotContains(t, versionError, "update(s) behind to current")
}
