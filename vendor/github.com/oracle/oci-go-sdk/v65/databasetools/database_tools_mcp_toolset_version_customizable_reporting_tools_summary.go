// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary Summary of versions available for the CUSTOMIZABLE_REPORTING_TOOLS MCP toolset type
type DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary struct {

	// The default version for this toolset type.
	DefaultVersion *int `mandatory:"true" json:"defaultVersion"`

	// The version configurations available for this toolset type.
	Versions []DatabaseToolsMcpToolsetCustomizableReportingToolsVersion `mandatory:"true" json:"versions"`
}

// GetDefaultVersion returns DefaultVersion
func (m DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary) GetDefaultVersion() *int {
	return m.DefaultVersion
}

func (m DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary DatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary
	}{
		"CUSTOMIZABLE_REPORTING_TOOLS",
		(MarshalTypeDatabaseToolsMcpToolsetVersionCustomizableReportingToolsSummary)(m),
	}

	return json.Marshal(&s)
}
