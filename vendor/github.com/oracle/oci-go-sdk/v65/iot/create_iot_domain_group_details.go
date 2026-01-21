// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateIotDomainGroupDetails The information about new IoT domain group to be created.
type CreateIotDomainGroupDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment corresponding to the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the domain group. LIGHTWEIGHT uses fewer resources and has a higher Recovery Time Objective (RTO),
	// making it suitable for development and testing. STANDARD is recommended for production.
	Type CreateIotDomainGroupDetailsTypeEnum `mandatory:"false" json:"type,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateIotDomainGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateIotDomainGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateIotDomainGroupDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCreateIotDomainGroupDetailsTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateIotDomainGroupDetailsTypeEnum Enum with underlying type: string
type CreateIotDomainGroupDetailsTypeEnum string

// Set of constants representing the allowable values for CreateIotDomainGroupDetailsTypeEnum
const (
	CreateIotDomainGroupDetailsTypeStandard    CreateIotDomainGroupDetailsTypeEnum = "STANDARD"
	CreateIotDomainGroupDetailsTypeLightweight CreateIotDomainGroupDetailsTypeEnum = "LIGHTWEIGHT"
)

var mappingCreateIotDomainGroupDetailsTypeEnum = map[string]CreateIotDomainGroupDetailsTypeEnum{
	"STANDARD":    CreateIotDomainGroupDetailsTypeStandard,
	"LIGHTWEIGHT": CreateIotDomainGroupDetailsTypeLightweight,
}

var mappingCreateIotDomainGroupDetailsTypeEnumLowerCase = map[string]CreateIotDomainGroupDetailsTypeEnum{
	"standard":    CreateIotDomainGroupDetailsTypeStandard,
	"lightweight": CreateIotDomainGroupDetailsTypeLightweight,
}

// GetCreateIotDomainGroupDetailsTypeEnumValues Enumerates the set of values for CreateIotDomainGroupDetailsTypeEnum
func GetCreateIotDomainGroupDetailsTypeEnumValues() []CreateIotDomainGroupDetailsTypeEnum {
	values := make([]CreateIotDomainGroupDetailsTypeEnum, 0)
	for _, v := range mappingCreateIotDomainGroupDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateIotDomainGroupDetailsTypeEnumStringValues Enumerates the set of values in String for CreateIotDomainGroupDetailsTypeEnum
func GetCreateIotDomainGroupDetailsTypeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"LIGHTWEIGHT",
	}
}

// GetMappingCreateIotDomainGroupDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateIotDomainGroupDetailsTypeEnum(val string) (CreateIotDomainGroupDetailsTypeEnum, bool) {
	enum, ok := mappingCreateIotDomainGroupDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
