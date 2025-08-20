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

// CreateAttributeSetDetails The details for an attribute set.
type CreateAttributeSetDetails struct {

	// The display name of the attribute set. The name is unique and changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment that contains the attribute set.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of attribute set.
	AttributeSetType AttributeSetAttributeSetTypeEnum `mandatory:"true" json:"attributeSetType"`

	// The list of values in an attribute set
	AttributeSetValues []string `mandatory:"true" json:"attributeSetValues"`

	// Description of the attribute set.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateAttributeSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAttributeSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAttributeSetAttributeSetTypeEnum(string(m.AttributeSetType)); !ok && m.AttributeSetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeSetType: %s. Supported values are: %s.", m.AttributeSetType, strings.Join(GetAttributeSetAttributeSetTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
