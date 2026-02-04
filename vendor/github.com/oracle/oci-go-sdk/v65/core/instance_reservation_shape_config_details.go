// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceReservationShapeConfigDetails The shape configuration requested when launching instances in a compute capacity reservation.
// If the parameter is provided, the reservation is created with the resources that you specify. If some
// properties are missing or the parameter is not provided, the reservation is created
// with the default configuration values for the `shape` that you specify.
// Each shape only supports certain configurable values. If the values that you provide are not valid for the
// specified `shape`, an error is returned.
// For more information about customizing the resources that are allocated to flexible shapes,
// see Flexible Shapes (https://docs.oracle.com/iaas/Content/Compute/References/computeshapes.htm#flexible).
type InstanceReservationShapeConfigDetails struct {

	// The total number of OCPUs available to the instance.
	Ocpus *float32 `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the instance, in gigabytes.
	MemoryInGBs *float32 `mandatory:"false" json:"memoryInGBs"`

	// This field is reserved for internal use.
	ResourceManagement InstanceReservationShapeConfigDetailsResourceManagementEnum `mandatory:"false" json:"resourceManagement,omitempty"`
}

func (m InstanceReservationShapeConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceReservationShapeConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstanceReservationShapeConfigDetailsResourceManagementEnum(string(m.ResourceManagement)); !ok && m.ResourceManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceManagement: %s. Supported values are: %s.", m.ResourceManagement, strings.Join(GetInstanceReservationShapeConfigDetailsResourceManagementEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceReservationShapeConfigDetailsResourceManagementEnum Enum with underlying type: string
type InstanceReservationShapeConfigDetailsResourceManagementEnum string

// Set of constants representing the allowable values for InstanceReservationShapeConfigDetailsResourceManagementEnum
const (
	InstanceReservationShapeConfigDetailsResourceManagementDynamic InstanceReservationShapeConfigDetailsResourceManagementEnum = "DYNAMIC"
	InstanceReservationShapeConfigDetailsResourceManagementStatic  InstanceReservationShapeConfigDetailsResourceManagementEnum = "STATIC"
)

var mappingInstanceReservationShapeConfigDetailsResourceManagementEnum = map[string]InstanceReservationShapeConfigDetailsResourceManagementEnum{
	"DYNAMIC": InstanceReservationShapeConfigDetailsResourceManagementDynamic,
	"STATIC":  InstanceReservationShapeConfigDetailsResourceManagementStatic,
}

var mappingInstanceReservationShapeConfigDetailsResourceManagementEnumLowerCase = map[string]InstanceReservationShapeConfigDetailsResourceManagementEnum{
	"dynamic": InstanceReservationShapeConfigDetailsResourceManagementDynamic,
	"static":  InstanceReservationShapeConfigDetailsResourceManagementStatic,
}

// GetInstanceReservationShapeConfigDetailsResourceManagementEnumValues Enumerates the set of values for InstanceReservationShapeConfigDetailsResourceManagementEnum
func GetInstanceReservationShapeConfigDetailsResourceManagementEnumValues() []InstanceReservationShapeConfigDetailsResourceManagementEnum {
	values := make([]InstanceReservationShapeConfigDetailsResourceManagementEnum, 0)
	for _, v := range mappingInstanceReservationShapeConfigDetailsResourceManagementEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceReservationShapeConfigDetailsResourceManagementEnumStringValues Enumerates the set of values in String for InstanceReservationShapeConfigDetailsResourceManagementEnum
func GetInstanceReservationShapeConfigDetailsResourceManagementEnumStringValues() []string {
	return []string{
		"DYNAMIC",
		"STATIC",
	}
}

// GetMappingInstanceReservationShapeConfigDetailsResourceManagementEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceReservationShapeConfigDetailsResourceManagementEnum(val string) (InstanceReservationShapeConfigDetailsResourceManagementEnum, bool) {
	enum, ok := mappingInstanceReservationShapeConfigDetailsResourceManagementEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
