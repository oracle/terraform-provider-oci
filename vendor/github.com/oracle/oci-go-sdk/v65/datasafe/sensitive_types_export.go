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

// SensitiveTypesExport The resource represents sensitive types to be exported in Data Safe.
type SensitiveTypesExport struct {

	// The OCID of the sensitive types export.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the sensitive types export.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the sensitive types export.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the sensitive types export.
	LifecycleState SensitiveTypesExportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the sensitive types export was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Indicates if all the existing user-defined sensitive types are used for export. If it's set to true, the
	// sensitiveTypeIdsForExport attribute is ignored and all user-defined sensitive types are exported.
	IsIncludeAllSensitiveTypes *bool `mandatory:"true" json:"isIncludeAllSensitiveTypes"`

	// The description of the sensitive types export.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the sensitive types export was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The OCIDs of the sensitive types used to create sensitive types export.
	SensitiveTypeIdsForExport []string `mandatory:"false" json:"sensitiveTypeIdsForExport"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SensitiveTypesExport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveTypesExport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSensitiveTypesExportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSensitiveTypesExportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
