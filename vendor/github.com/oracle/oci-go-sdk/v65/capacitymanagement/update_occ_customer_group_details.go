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

// UpdateOccCustomerGroupDetails Details about the update request for updating the customer group.
type UpdateOccCustomerGroupDetails struct {

	// The display name of the customer group.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Additional information about the customer group.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// To determine whether the customer group is enabled/disabled.
	Status UpdateOccCustomerGroupDetailsStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m UpdateOccCustomerGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOccCustomerGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateOccCustomerGroupDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetUpdateOccCustomerGroupDetailsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOccCustomerGroupDetailsStatusEnum Enum with underlying type: string
type UpdateOccCustomerGroupDetailsStatusEnum string

// Set of constants representing the allowable values for UpdateOccCustomerGroupDetailsStatusEnum
const (
	UpdateOccCustomerGroupDetailsStatusEnabled  UpdateOccCustomerGroupDetailsStatusEnum = "ENABLED"
	UpdateOccCustomerGroupDetailsStatusDisabled UpdateOccCustomerGroupDetailsStatusEnum = "DISABLED"
)

var mappingUpdateOccCustomerGroupDetailsStatusEnum = map[string]UpdateOccCustomerGroupDetailsStatusEnum{
	"ENABLED":  UpdateOccCustomerGroupDetailsStatusEnabled,
	"DISABLED": UpdateOccCustomerGroupDetailsStatusDisabled,
}

var mappingUpdateOccCustomerGroupDetailsStatusEnumLowerCase = map[string]UpdateOccCustomerGroupDetailsStatusEnum{
	"enabled":  UpdateOccCustomerGroupDetailsStatusEnabled,
	"disabled": UpdateOccCustomerGroupDetailsStatusDisabled,
}

// GetUpdateOccCustomerGroupDetailsStatusEnumValues Enumerates the set of values for UpdateOccCustomerGroupDetailsStatusEnum
func GetUpdateOccCustomerGroupDetailsStatusEnumValues() []UpdateOccCustomerGroupDetailsStatusEnum {
	values := make([]UpdateOccCustomerGroupDetailsStatusEnum, 0)
	for _, v := range mappingUpdateOccCustomerGroupDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOccCustomerGroupDetailsStatusEnumStringValues Enumerates the set of values in String for UpdateOccCustomerGroupDetailsStatusEnum
func GetUpdateOccCustomerGroupDetailsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingUpdateOccCustomerGroupDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOccCustomerGroupDetailsStatusEnum(val string) (UpdateOccCustomerGroupDetailsStatusEnum, bool) {
	enum, ok := mappingUpdateOccCustomerGroupDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
