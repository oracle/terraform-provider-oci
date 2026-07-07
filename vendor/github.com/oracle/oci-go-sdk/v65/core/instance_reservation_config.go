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

// InstanceReservationConfig Data that defines the capacity configuration.
type InstanceReservationConfig struct {

	// The shape to use when launching instances using compute capacity reservations. The shape determines the number of CPUs, the amount of memory,
	// and other resources allocated to the instance.
	// You can list all available shapes by calling ListComputeCapacityReservationInstanceShapes.
	InstanceShape *string `mandatory:"true" json:"instanceShape"`

	// The total number of instances that can be launched from the capacity configuration.
	ReservedCount *int64 `mandatory:"true" json:"reservedCount"`

	// The amount of capacity in use out of the total capacity reserved in this capacity configuration.
	UsedCount *int64 `mandatory:"true" json:"usedCount"`

	// The fault domain of this capacity configuration.
	// If a value is not supplied, this capacity configuration is applicable to all fault domains in the specified availability domain.
	// For more information, see Capacity Reservations (https://docs.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm).
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	ClusterConfig *ClusterConfigDetails `mandatory:"false" json:"clusterConfig"`

	RecycleConfig *RecycleConfigDetails `mandatory:"false" json:"recycleConfig"`

	InstanceShapeConfig *InstanceReservationShapeConfigDetails `mandatory:"false" json:"instanceShapeConfig"`

	// The OCID of the cluster placement group for this instance reservation capacity configuration.
	ClusterPlacementGroupId *string `mandatory:"false" json:"clusterPlacementGroupId"`

	// This field is reserved for internal use.
	InternalOverlayEnclave InstanceReservationConfigInternalOverlayEnclaveEnum `mandatory:"false" json:"internalOverlayEnclave,omitempty"`
}

func (m InstanceReservationConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceReservationConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInstanceReservationConfigInternalOverlayEnclaveEnum(string(m.InternalOverlayEnclave)); !ok && m.InternalOverlayEnclave != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InternalOverlayEnclave: %s. Supported values are: %s.", m.InternalOverlayEnclave, strings.Join(GetInstanceReservationConfigInternalOverlayEnclaveEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceReservationConfigInternalOverlayEnclaveEnum Enum with underlying type: string
type InstanceReservationConfigInternalOverlayEnclaveEnum string

// Set of constants representing the allowable values for InstanceReservationConfigInternalOverlayEnclaveEnum
const (
	InstanceReservationConfigInternalOverlayEnclaveSecurity InstanceReservationConfigInternalOverlayEnclaveEnum = "SECURITY"
)

var mappingInstanceReservationConfigInternalOverlayEnclaveEnum = map[string]InstanceReservationConfigInternalOverlayEnclaveEnum{
	"SECURITY": InstanceReservationConfigInternalOverlayEnclaveSecurity,
}

var mappingInstanceReservationConfigInternalOverlayEnclaveEnumLowerCase = map[string]InstanceReservationConfigInternalOverlayEnclaveEnum{
	"security": InstanceReservationConfigInternalOverlayEnclaveSecurity,
}

// GetInstanceReservationConfigInternalOverlayEnclaveEnumValues Enumerates the set of values for InstanceReservationConfigInternalOverlayEnclaveEnum
func GetInstanceReservationConfigInternalOverlayEnclaveEnumValues() []InstanceReservationConfigInternalOverlayEnclaveEnum {
	values := make([]InstanceReservationConfigInternalOverlayEnclaveEnum, 0)
	for _, v := range mappingInstanceReservationConfigInternalOverlayEnclaveEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceReservationConfigInternalOverlayEnclaveEnumStringValues Enumerates the set of values in String for InstanceReservationConfigInternalOverlayEnclaveEnum
func GetInstanceReservationConfigInternalOverlayEnclaveEnumStringValues() []string {
	return []string{
		"SECURITY",
	}
}

// GetMappingInstanceReservationConfigInternalOverlayEnclaveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceReservationConfigInternalOverlayEnclaveEnum(val string) (InstanceReservationConfigInternalOverlayEnclaveEnum, bool) {
	enum, ok := mappingInstanceReservationConfigInternalOverlayEnclaveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
