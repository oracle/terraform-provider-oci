// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// SensitiveTypeSummary Summary of a sensitive type.
type SensitiveTypeSummary struct {

	// The OCID of the sensitive type.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the sensitive type.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the sensitive type.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the sensitive type.
	LifecycleState DiscoveryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies whether the sensitive type is user-defined or predefined.
	Source SensitiveTypeSourceEnum `mandatory:"true" json:"source"`

	// The date and time the sensitive type was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the sensitive type was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The entity type. It can be either a sensitive type with regular expressions or a sensitive category used for
	// grouping similar sensitive types.
	EntityType SensitiveTypeEntityEnum `mandatory:"true" json:"entityType"`

	// The short name of the sensitive type.
	ShortName *string `mandatory:"false" json:"shortName"`

	// The description of the sensitive type.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the parent sensitive category.
	ParentCategoryId *string `mandatory:"false" json:"parentCategoryId"`

	// The OCID of the library masking format that should be used to mask the sensitive columns associated with the sensitive type.
	DefaultMaskingFormatId *string `mandatory:"false" json:"defaultMaskingFormatId"`

	// Specifies whether the sensitive type is common. Common sensitive types belong to
	// library sensitive types which are frequently used to perform sensitive data discovery.
	IsCommon *bool `mandatory:"false" json:"isCommon"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SensitiveTypeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SensitiveTypeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDiscoveryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveTypeSourceEnum(string(m.Source)); !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetSensitiveTypeSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSensitiveTypeEntityEnum(string(m.EntityType)); !ok && m.EntityType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntityType: %s. Supported values are: %s.", m.EntityType, strings.Join(GetSensitiveTypeEntityEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
