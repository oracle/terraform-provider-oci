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

// UpdateOccCustomerDetails Details about the update request for updating the customer.
type UpdateOccCustomerDetails struct {

	// The display name of the customer.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional information about the customer.
	Description *string `mandatory:"false" json:"description"`

	// To determine whether the customer group is enabled/disabled.
	Status UpdateOccCustomerDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m UpdateOccCustomerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOccCustomerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateOccCustomerDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateOccCustomerDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOccCustomerDetailsStatusEnum Enum with underlying type: string
type UpdateOccCustomerDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateOccCustomerDetailsStatusEnum
const (
	UpdateOccCustomerDetailsStatusEnabled  UpdateOccCustomerDetailsStatusEnum = "ENABLED"
	UpdateOccCustomerDetailsStatusDisabled UpdateOccCustomerDetailsStatusEnum = "DISABLED"
)

var mappingUpdateOccCustomerDetailsStatusEnum = map[string]UpdateOccCustomerDetailsStatusEnum{
	"ENABLED":  UpdateOccCustomerDetailsStatusEnabled,
	"DISABLED": UpdateOccCustomerDetailsStatusDisabled,
}

var mappingUpdateOccCustomerDetailsStatusEnumLowerCase = map[string]UpdateOccCustomerDetailsStatusEnum{
	"enabled":  UpdateOccCustomerDetailsStatusEnabled,
	"disabled": UpdateOccCustomerDetailsStatusDisabled,
}

// GetUpdateOccCustomerDetailsStatusEnumValues Enumerates the set of values for UpdateOccCustomerDetailsStatusEnum
func GetUpdateOccCustomerDetailsStatusEnumValues() []UpdateOccCustomerDetailsStatusEnum {
	values := make([]UpdateOccCustomerDetailsStatusEnum, 0)
	for _, v := range mappingUpdateOccCustomerDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOccCustomerDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateOccCustomerDetailsStatusEnum
func GetUpdateOccCustomerDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateOccCustomerDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOccCustomerDetailsStatusEnum(val string) (UpdateOccCustomerDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateOccCustomerDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
