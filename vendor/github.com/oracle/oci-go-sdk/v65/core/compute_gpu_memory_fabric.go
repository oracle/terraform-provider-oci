// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ComputeGpuMemoryFabric The customer facing object includes GPU memory fabric details.
type ComputeGpuMemoryFabric struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Customer-unique GPU memory fabric
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the compartment. This should always be the root
	// compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique HPC Island
	ComputeHpcIslandId *string `mandatory:"true" json:"computeHpcIslandId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique Network Block
	ComputeNetworkBlockId *string `mandatory:"true" json:"computeNetworkBlockId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for Customer-unique Local Block
	ComputeLocalBlockId *string `mandatory:"true" json:"computeLocalBlockId"`

	// The lifecycle state of the GPU memory fabric
	LifecycleState ComputeGpuMemoryFabricLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The health state of the GPU memory fabric
	FabricHealth ComputeGpuMemoryFabricFabricHealthEnum `mandatory:"true" json:"fabricHealth"`

	// The total number of healthy bare metal hosts located in this compute GPU memory fabric.
	HealthyHostCount *int64 `mandatory:"true" json:"healthyHostCount"`

	// The total number of bare metal hosts located in this compute GPU memory fabric.
	TotalHostCount *int64 `mandatory:"true" json:"totalHostCount"`

	// The date and time that the compute GPU memory fabric record was created, in the format defined by RFC3339
	//  (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Additional data that can be exposed to the customer. Right now it will include the switch tray ids.
	AdditionalData map[string]interface{} `mandatory:"false" json:"additionalData"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m ComputeGpuMemoryFabric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeGpuMemoryFabric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComputeGpuMemoryFabricLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComputeGpuMemoryFabricLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingComputeGpuMemoryFabricFabricHealthEnum(string(m.FabricHealth)); !ok && m.FabricHealth != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FabricHealth: %s. Supported values are: %s.", m.FabricHealth, strings.Join(GetComputeGpuMemoryFabricFabricHealthEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComputeGpuMemoryFabricLifecycleStateEnum Enum with underlying type: string
type ComputeGpuMemoryFabricLifecycleStateEnum string

// Set of constants representing the allowable values for ComputeGpuMemoryFabricLifecycleStateEnum
const (
	ComputeGpuMemoryFabricLifecycleStateAvailable    ComputeGpuMemoryFabricLifecycleStateEnum = "AVAILABLE"
	ComputeGpuMemoryFabricLifecycleStateOccupied     ComputeGpuMemoryFabricLifecycleStateEnum = "OCCUPIED"
	ComputeGpuMemoryFabricLifecycleStateProvisioning ComputeGpuMemoryFabricLifecycleStateEnum = "PROVISIONING"
	ComputeGpuMemoryFabricLifecycleStateDegraded     ComputeGpuMemoryFabricLifecycleStateEnum = "DEGRADED"
	ComputeGpuMemoryFabricLifecycleStateUnavailable  ComputeGpuMemoryFabricLifecycleStateEnum = "UNAVAILABLE"
)

var mappingComputeGpuMemoryFabricLifecycleStateEnum = map[string]ComputeGpuMemoryFabricLifecycleStateEnum{
	"AVAILABLE":    ComputeGpuMemoryFabricLifecycleStateAvailable,
	"OCCUPIED":     ComputeGpuMemoryFabricLifecycleStateOccupied,
	"PROVISIONING": ComputeGpuMemoryFabricLifecycleStateProvisioning,
	"DEGRADED":     ComputeGpuMemoryFabricLifecycleStateDegraded,
	"UNAVAILABLE":  ComputeGpuMemoryFabricLifecycleStateUnavailable,
}

var mappingComputeGpuMemoryFabricLifecycleStateEnumLowerCase = map[string]ComputeGpuMemoryFabricLifecycleStateEnum{
	"available":    ComputeGpuMemoryFabricLifecycleStateAvailable,
	"occupied":     ComputeGpuMemoryFabricLifecycleStateOccupied,
	"provisioning": ComputeGpuMemoryFabricLifecycleStateProvisioning,
	"degraded":     ComputeGpuMemoryFabricLifecycleStateDegraded,
	"unavailable":  ComputeGpuMemoryFabricLifecycleStateUnavailable,
}

// GetComputeGpuMemoryFabricLifecycleStateEnumValues Enumerates the set of values for ComputeGpuMemoryFabricLifecycleStateEnum
func GetComputeGpuMemoryFabricLifecycleStateEnumValues() []ComputeGpuMemoryFabricLifecycleStateEnum {
	values := make([]ComputeGpuMemoryFabricLifecycleStateEnum, 0)
	for _, v := range mappingComputeGpuMemoryFabricLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeGpuMemoryFabricLifecycleStateEnumStringValues Enumerates the set of values in String for ComputeGpuMemoryFabricLifecycleStateEnum
func GetComputeGpuMemoryFabricLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"OCCUPIED",
		"PROVISIONING",
		"DEGRADED",
		"UNAVAILABLE",
	}
}

// GetMappingComputeGpuMemoryFabricLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeGpuMemoryFabricLifecycleStateEnum(val string) (ComputeGpuMemoryFabricLifecycleStateEnum, bool) {
	enum, ok := mappingComputeGpuMemoryFabricLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ComputeGpuMemoryFabricFabricHealthEnum Enum with underlying type: string
type ComputeGpuMemoryFabricFabricHealthEnum string

// Set of constants representing the allowable values for ComputeGpuMemoryFabricFabricHealthEnum
const (
	ComputeGpuMemoryFabricFabricHealthHealthy   ComputeGpuMemoryFabricFabricHealthEnum = "HEALTHY"
	ComputeGpuMemoryFabricFabricHealthUnhealthy ComputeGpuMemoryFabricFabricHealthEnum = "UNHEALTHY"
)

var mappingComputeGpuMemoryFabricFabricHealthEnum = map[string]ComputeGpuMemoryFabricFabricHealthEnum{
	"HEALTHY":   ComputeGpuMemoryFabricFabricHealthHealthy,
	"UNHEALTHY": ComputeGpuMemoryFabricFabricHealthUnhealthy,
}

var mappingComputeGpuMemoryFabricFabricHealthEnumLowerCase = map[string]ComputeGpuMemoryFabricFabricHealthEnum{
	"healthy":   ComputeGpuMemoryFabricFabricHealthHealthy,
	"unhealthy": ComputeGpuMemoryFabricFabricHealthUnhealthy,
}

// GetComputeGpuMemoryFabricFabricHealthEnumValues Enumerates the set of values for ComputeGpuMemoryFabricFabricHealthEnum
func GetComputeGpuMemoryFabricFabricHealthEnumValues() []ComputeGpuMemoryFabricFabricHealthEnum {
	values := make([]ComputeGpuMemoryFabricFabricHealthEnum, 0)
	for _, v := range mappingComputeGpuMemoryFabricFabricHealthEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeGpuMemoryFabricFabricHealthEnumStringValues Enumerates the set of values in String for ComputeGpuMemoryFabricFabricHealthEnum
func GetComputeGpuMemoryFabricFabricHealthEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"UNHEALTHY",
	}
}

// GetMappingComputeGpuMemoryFabricFabricHealthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeGpuMemoryFabricFabricHealthEnum(val string) (ComputeGpuMemoryFabricFabricHealthEnum, bool) {
	enum, ok := mappingComputeGpuMemoryFabricFabricHealthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
