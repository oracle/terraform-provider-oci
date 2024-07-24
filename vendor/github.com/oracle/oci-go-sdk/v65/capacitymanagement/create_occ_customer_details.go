// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOccCustomerDetails The details about the customer.
type CreateOccCustomerDetails struct {

	// The OCID of the tenancy belonging to the customer.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The display name for the customer.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The description about the customer group.
	Description *string `mandatory:"false" json:"description"`

	// To determine whether the customer is enabled/disabled.
	Status CreateOccCustomerDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m CreateOccCustomerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOccCustomerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateOccCustomerDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCreateOccCustomerDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOccCustomerDetailsStatusEnum Enum with underlying type: string
type CreateOccCustomerDetailsStatusEnum string

// Set of constants representing the allowable values for CreateOccCustomerDetailsStatusEnum
const (
	CreateOccCustomerDetailsStatusEnabled  CreateOccCustomerDetailsStatusEnum = "ENABLED"
	CreateOccCustomerDetailsStatusDisabled CreateOccCustomerDetailsStatusEnum = "DISABLED"
)

var mappingCreateOccCustomerDetailsStatusEnum = map[string]CreateOccCustomerDetailsStatusEnum{
	"ENABLED":  CreateOccCustomerDetailsStatusEnabled,
	"DISABLED": CreateOccCustomerDetailsStatusDisabled,
}

var mappingCreateOccCustomerDetailsStatusEnumLowerCase = map[string]CreateOccCustomerDetailsStatusEnum{
	"enabled":  CreateOccCustomerDetailsStatusEnabled,
	"disabled": CreateOccCustomerDetailsStatusDisabled,
}

// GetCreateOccCustomerDetailsStatusEnumValues Enumerates the set of values for CreateOccCustomerDetailsStatusEnum
func GetCreateOccCustomerDetailsStatusEnumValues() []CreateOccCustomerDetailsStatusEnum {
	values := make([]CreateOccCustomerDetailsStatusEnum, 0)
	for _, v := range mappingCreateOccCustomerDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOccCustomerDetailsStatusEnumStringValues Enumerates the set of values in String for CreateOccCustomerDetailsStatusEnum
func GetCreateOccCustomerDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCreateOccCustomerDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOccCustomerDetailsStatusEnum(val string) (CreateOccCustomerDetailsStatusEnum, bool) {
	enum, ok := mappingCreateOccCustomerDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
