// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Compartment internal compartment
type Compartment struct {

	// The OCID of the compartment.
	Id *string `mandatory:"true" json:"id"`

	// The compartment enablement status.
	Enablement CompartmentEnablementEnum `mandatory:"true" json:"enablement"`
}

func (m Compartment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Compartment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompartmentEnablementEnum(string(m.Enablement)); !ok && m.Enablement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Enablement: %s. Supported values are: %s.", m.Enablement, strings.Join(GetCompartmentEnablementEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CompartmentEnablementEnum Enum with underlying type: string
type CompartmentEnablementEnum string

// Set of constants representing the allowable values for CompartmentEnablementEnum
const (
	CompartmentEnablementEnabling  CompartmentEnablementEnum = "ENABLING"
	CompartmentEnablementEnabled   CompartmentEnablementEnum = "ENABLED"
	CompartmentEnablementDisabling CompartmentEnablementEnum = "DISABLING"
	CompartmentEnablementDisabled  CompartmentEnablementEnum = "DISABLED"
)

var mappingCompartmentEnablementEnum = map[string]CompartmentEnablementEnum{
	"ENABLING":  CompartmentEnablementEnabling,
	"ENABLED":   CompartmentEnablementEnabled,
	"DISABLING": CompartmentEnablementDisabling,
	"DISABLED":  CompartmentEnablementDisabled,
}

var mappingCompartmentEnablementEnumLowerCase = map[string]CompartmentEnablementEnum{
	"enabling":  CompartmentEnablementEnabling,
	"enabled":   CompartmentEnablementEnabled,
	"disabling": CompartmentEnablementDisabling,
	"disabled":  CompartmentEnablementDisabled,
}

// GetCompartmentEnablementEnumValues Enumerates the set of values for CompartmentEnablementEnum
func GetCompartmentEnablementEnumValues() []CompartmentEnablementEnum {
	values := make([]CompartmentEnablementEnum, 0)
	for _, v := range mappingCompartmentEnablementEnum {
		values = append(values, v)
	}
	return values
}

// GetCompartmentEnablementEnumStringValues Enumerates the set of values in String for CompartmentEnablementEnum
func GetCompartmentEnablementEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
	}
}

// GetMappingCompartmentEnablementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCompartmentEnablementEnum(val string) (CompartmentEnablementEnum, bool) {
	enum, ok := mappingCompartmentEnablementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
