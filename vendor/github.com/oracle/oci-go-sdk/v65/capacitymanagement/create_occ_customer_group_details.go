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

// CreateOccCustomerGroupDetails Details about the create request for the customer group.
type CreateOccCustomerGroupDetails struct {

	// Since all resources are at tenancy level hence this will be the ocid of the tenancy where operation is to be performed.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name of the customer group.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A description about the customer group.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// To determine whether the customer group is enabled/disabled.
	Status CreateOccCustomerGroupDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`

	// A list containing all the customers that belong to this customer group.
	CustomersList []CreateOccCustomerDetails `mandatory:"false" json:"customersList"`
}

func (m CreateOccCustomerGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOccCustomerGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateOccCustomerGroupDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCreateOccCustomerGroupDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOccCustomerGroupDetailsStatusEnum Enum with underlying type: string
type CreateOccCustomerGroupDetailsStatusEnum string

// Set of constants representing the allowable values for CreateOccCustomerGroupDetailsStatusEnum
const (
	CreateOccCustomerGroupDetailsStatusEnabled  CreateOccCustomerGroupDetailsStatusEnum = "ENABLED"
	CreateOccCustomerGroupDetailsStatusDisabled CreateOccCustomerGroupDetailsStatusEnum = "DISABLED"
)

var mappingCreateOccCustomerGroupDetailsStatusEnum = map[string]CreateOccCustomerGroupDetailsStatusEnum{
	"ENABLED":  CreateOccCustomerGroupDetailsStatusEnabled,
	"DISABLED": CreateOccCustomerGroupDetailsStatusDisabled,
}

var mappingCreateOccCustomerGroupDetailsStatusEnumLowerCase = map[string]CreateOccCustomerGroupDetailsStatusEnum{
	"enabled":  CreateOccCustomerGroupDetailsStatusEnabled,
	"disabled": CreateOccCustomerGroupDetailsStatusDisabled,
}

// GetCreateOccCustomerGroupDetailsStatusEnumValues Enumerates the set of values for CreateOccCustomerGroupDetailsStatusEnum
func GetCreateOccCustomerGroupDetailsStatusEnumValues() []CreateOccCustomerGroupDetailsStatusEnum {
	values := make([]CreateOccCustomerGroupDetailsStatusEnum, 0)
	for _, v := range mappingCreateOccCustomerGroupDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOccCustomerGroupDetailsStatusEnumStringValues Enumerates the set of values in String for CreateOccCustomerGroupDetailsStatusEnum
func GetCreateOccCustomerGroupDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingCreateOccCustomerGroupDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOccCustomerGroupDetailsStatusEnum(val string) (CreateOccCustomerGroupDetailsStatusEnum, bool) {
	enum, ok := mappingCreateOccCustomerGroupDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
