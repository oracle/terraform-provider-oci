// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePropertyDetails The information about new Property.
type CreatePropertyDetails struct {

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Text selection of the property.
	Selection SelectionEnum `mandatory:"true" json:"selection"`

	// Format of the value.
	ValueType ValueTypeEnum `mandatory:"true" json:"valueType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Values of the property (must be a single value if selection = 'SINGLE_CHOICE').
	Values []string `mandatory:"false" json:"values"`
}

func (m CreatePropertyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePropertyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSelectionEnum(string(m.Selection)); !ok && m.Selection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Selection: %s. Supported values are: %s.", m.Selection, strings.Join(GetSelectionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
