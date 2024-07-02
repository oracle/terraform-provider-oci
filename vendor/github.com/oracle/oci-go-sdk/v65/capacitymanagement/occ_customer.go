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

// OccCustomer The details about the customer.
type OccCustomer struct {

	// The OCID of the customer group.
	OccCustomerGroupId *string `mandatory:"true" json:"occCustomerGroupId"`

	// The OCID of the tenancy belonging to the customer.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// The display name for the customer
	DisplayName *string `mandatory:"true" json:"displayName"`

	// To determine whether the customer is enabled/disabled.`
	Status OccCustomerStatusEnum `mandatory:"true" json:"status"`

	// The description about the customer group.
	Description *string `mandatory:"false" json:"description"`
}

func (m OccCustomer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccCustomer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccCustomerStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetOccCustomerStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccCustomerStatusEnum Enum with underlying type: string
type OccCustomerStatusEnum string

// Set of constants representing the allowable values for OccCustomerStatusEnum
const (
	OccCustomerStatusEnabled  OccCustomerStatusEnum = "ENABLED"
	OccCustomerStatusDisabled OccCustomerStatusEnum = "DISABLED"
)

var mappingOccCustomerStatusEnum = map[string]OccCustomerStatusEnum{
	"ENABLED":  OccCustomerStatusEnabled,
	"DISABLED": OccCustomerStatusDisabled,
}

var mappingOccCustomerStatusEnumLowerCase = map[string]OccCustomerStatusEnum{
	"enabled":  OccCustomerStatusEnabled,
	"disabled": OccCustomerStatusDisabled,
}

// GetOccCustomerStatusEnumValues Enumerates the set of values for OccCustomerStatusEnum
func GetOccCustomerStatusEnumValues() []OccCustomerStatusEnum {
	values := make([]OccCustomerStatusEnum, 0)
	for _, v := range mappingOccCustomerStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetOccCustomerStatusEnumStringValues Enumerates the set of values in String for OccCustomerStatusEnum
func GetOccCustomerStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingOccCustomerStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccCustomerStatusEnum(val string) (OccCustomerStatusEnum, bool) {
	enum, ok := mappingOccCustomerStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
