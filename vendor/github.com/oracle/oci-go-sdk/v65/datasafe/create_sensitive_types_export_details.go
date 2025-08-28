// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateSensitiveTypesExportDetails Details to create a new sensitive types export. To specify some sensitive types for export, use
// sensitiveTypeIdsForExport attribute. But if you want to include all sensitive types, you can set
// isIncludeAllSensitiveTypes attributes to true. In the latter case, you do not need to list all
// sensitive types.
type CreateSensitiveTypesExportDetails struct {

	// The OCID of the compartment where the sensitive types export should be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the sensitive types export. The name does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the sensitive types export.
	Description *string `mandatory:"false" json:"description"`

	// The OCIDs of the sensitive types used to create sensitive types export.
	SensitiveTypeIdsForExport []string `mandatory:"false" json:"sensitiveTypeIdsForExport"`

	// Indicates if all the existing user-defined sensitive types are used for export. If it's set to true, the
	// sensitiveTypeIdsForExport attribute is ignored and all user-defined sensitive types are used.
	IsIncludeAllSensitiveTypes *bool `mandatory:"false" json:"isIncludeAllSensitiveTypes"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSensitiveTypesExportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSensitiveTypesExportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
