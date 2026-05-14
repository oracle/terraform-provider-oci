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

// DatabaseToolsMcpToolsetVersionTool A tool definition for a specific MCP toolset version.
type DatabaseToolsMcpToolsetVersionTool struct {

	// The unique name of the tool.
	Name *string `mandatory:"true" json:"name"`

	// The display name of the tool.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The description of the tool.
	Description *string `mandatory:"true" json:"description"`

	// The default status of the tool in this version.
	DefaultStatus DatabaseToolsMcpToolsetToolStatusEnum `mandatory:"true" json:"defaultStatus"`

	// The roles granted access to this tool by default.
	DefaultAllowedRoles []string `mandatory:"false" json:"defaultAllowedRoles"`
}

func (m DatabaseToolsMcpToolsetVersionTool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsMcpToolsetVersionTool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsMcpToolsetToolStatusEnum(string(m.DefaultStatus)); !ok && m.DefaultStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DefaultStatus: %s. Supported values are: %s.", m.DefaultStatus, strings.Join(GetDatabaseToolsMcpToolsetToolStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
