// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedCredentialFieldDefinition A named credential field metadata definition
type NamedCredentialFieldDefinition struct {

	// The field name
	Name *string `mandatory:"true" json:"name"`

	// The field display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// List of value categories of field allowed for this property
	ValueCategory []ValueCategoryTypeEnum `mandatory:"true" json:"valueCategory"`

	// The default value which will be used if no value is set.  If defaultValue is empty, then no default will be set.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// Optional regular expression definition which will be applied to the value when valueCategory is CLEAR_TEXT
	Regex *string `mandatory:"false" json:"regex"`

	// List of values which can be applied to the value when valueCategory is ALLOWED_VALUES
	AllowedValues []string `mandatory:"false" json:"allowedValues"`

	// Set to true if the field must be defined
	IsRequired *bool `mandatory:"false" json:"isRequired"`
}

func (m NamedCredentialFieldDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedCredentialFieldDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.ValueCategory {
		if _, ok := GetMappingValueCategoryTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueCategory: %s. Supported values are: %s.", val, strings.Join(GetValueCategoryTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
