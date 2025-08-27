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

// NamedCredentialProperty Property item in name/value pair
type NamedCredentialProperty struct {

	// Name of the property
	Name *string `mandatory:"true" json:"name"`

	// Value of the property
	Value *string `mandatory:"true" json:"value"`

	// The category of the Named credential property value.
	// CLEAR_TEXT indicates the value field contains a clear text value.
	// SECRET_IDENTIFIER indicates the value field contains a vault secret ocid identifier.
	// ADB_IDENTIFIER indicates the value field contains an Autonomous database ocid identifier.
	// ALLOWED_VALUE indicates the value should be selected from the options in the allowedValues field.
	ValueCategory ValueCategoryTypeEnum `mandatory:"true" json:"valueCategory"`
}

func (m NamedCredentialProperty) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedCredentialProperty) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingValueCategoryTypeEnum(string(m.ValueCategory)); !ok && m.ValueCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueCategory: %s. Supported values are: %s.", m.ValueCategory, strings.Join(GetValueCategoryTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
