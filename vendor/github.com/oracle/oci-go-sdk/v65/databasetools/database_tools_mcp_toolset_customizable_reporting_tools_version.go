// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsMcpToolsetCustomizableReportingToolsVersion A specific version entry for a CUSTOMIZABLE_REPORTING_TOOLS MCP toolset version
type DatabaseToolsMcpToolsetCustomizableReportingToolsVersion struct {

	// The version number.
	Version *int `mandatory:"true" json:"version"`

	// A description of this version.
	Description *string `mandatory:"true" json:"description"`

	// The roles granted access to this toolset version by default.
	DefaultReportAllowedRoles []string `mandatory:"true" json:"defaultReportAllowedRoles"`

	// The tools available in this version.
	Tools []DatabaseToolsMcpToolsetVersionTool `mandatory:"true" json:"tools"`

	// Optional feature flags or attributes for this version.
	Features []string `mandatory:"false" json:"features"`
}

func (m DatabaseToolsMcpToolsetCustomizableReportingToolsVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsMcpToolsetCustomizableReportingToolsVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
