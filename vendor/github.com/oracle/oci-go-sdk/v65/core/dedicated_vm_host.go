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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DedicatedVmHost A dedicated virtual machine host lets you host multiple VM instances
// on a dedicated server that is not shared with other tenancies.
type DedicatedVmHost struct {

	// The availability domain the dedicated virtual machine host is running in.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment that contains the dedicated virtual machine host.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The dedicated virtual machine host shape. The shape determines the number of CPUs and
	// other resources available for VMs.
	DedicatedVmHostShape *string `mandatory:"true" json:"dedicatedVmHostShape"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated VM host.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the dedicated VM host.
	LifecycleState DedicatedVmHostLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the dedicated VM host was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The total OCPUs of the dedicated VM host.
	TotalOcpus *float32 `mandatory:"true" json:"totalOcpus"`

	// The available OCPUs of the dedicated VM host.
	RemainingOcpus *float32 `mandatory:"true" json:"remainingOcpus"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The fault domain for the dedicated virtual machine host's assigned instances.
	// For more information, see Fault Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm#fault).
	// If you do not specify the fault domain, the system selects one for you. To change the fault domain for a dedicated virtual machine host,
	// delete it, and then create a new dedicated virtual machine host in the preferred fault domain.
	// To get a list of fault domains, use the `ListFaultDomains` operation in the Identity and Access Management Service API (https://docs.oracle.com/iaas/api/#/en/identity/20160918/).
	// Example: `FAULT-DOMAIN-1`
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	PlacementConstraintDetails PlacementConstraintDetails `mandatory:"false" json:"placementConstraintDetails"`

	// The total memory of the dedicated VM host, in GBs.
	TotalMemoryInGBs *float32 `mandatory:"false" json:"totalMemoryInGBs"`

	// The remaining memory of the dedicated VM host, in GBs.
	RemainingMemoryInGBs *float32 `mandatory:"false" json:"remainingMemoryInGBs"`

	// A list of total and remaining CPU & memory per capacity bucket.
	CapacityBins []CapacityBin `mandatory:"false" json:"capacityBins"`

	// The compute bare metal host OCID of the dedicated virtual machine host.
	ComputeBareMetalHostId *string `mandatory:"false" json:"computeBareMetalHostId"`
}

func (m DedicatedVmHost) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedVmHost) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedVmHostLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDedicatedVmHostLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DedicatedVmHost) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		FaultDomain                *string                           `json:"faultDomain"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		PlacementConstraintDetails placementconstraintdetails        `json:"placementConstraintDetails"`
		TotalMemoryInGBs           *float32                          `json:"totalMemoryInGBs"`
		RemainingMemoryInGBs       *float32                          `json:"remainingMemoryInGBs"`
		CapacityBins               []CapacityBin                     `json:"capacityBins"`
		ComputeBareMetalHostId     *string                           `json:"computeBareMetalHostId"`
		AvailabilityDomain         *string                           `json:"availabilityDomain"`
		CompartmentId              *string                           `json:"compartmentId"`
		DedicatedVmHostShape       *string                           `json:"dedicatedVmHostShape"`
		DisplayName                *string                           `json:"displayName"`
		Id                         *string                           `json:"id"`
		LifecycleState             DedicatedVmHostLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated                *common.SDKTime                   `json:"timeCreated"`
		TotalOcpus                 *float32                          `json:"totalOcpus"`
		RemainingOcpus             *float32                          `json:"remainingOcpus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.FaultDomain = model.FaultDomain

	m.FreeformTags = model.FreeformTags

	nn, e = model.PlacementConstraintDetails.UnmarshalPolymorphicJSON(model.PlacementConstraintDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PlacementConstraintDetails = nn.(PlacementConstraintDetails)
	} else {
		m.PlacementConstraintDetails = nil
	}

	m.TotalMemoryInGBs = model.TotalMemoryInGBs

	m.RemainingMemoryInGBs = model.RemainingMemoryInGBs

	m.CapacityBins = make([]CapacityBin, len(model.CapacityBins))
	copy(m.CapacityBins, model.CapacityBins)
	m.ComputeBareMetalHostId = model.ComputeBareMetalHostId

	m.AvailabilityDomain = model.AvailabilityDomain

	m.CompartmentId = model.CompartmentId

	m.DedicatedVmHostShape = model.DedicatedVmHostShape

	m.DisplayName = model.DisplayName

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TotalOcpus = model.TotalOcpus

	m.RemainingOcpus = model.RemainingOcpus

	return
}

// DedicatedVmHostLifecycleStateEnum Enum with underlying type: string
type DedicatedVmHostLifecycleStateEnum string

// Set of constants representing the allowable values for DedicatedVmHostLifecycleStateEnum
const (
	DedicatedVmHostLifecycleStateCreating DedicatedVmHostLifecycleStateEnum = "CREATING"
	DedicatedVmHostLifecycleStateActive   DedicatedVmHostLifecycleStateEnum = "ACTIVE"
	DedicatedVmHostLifecycleStateUpdating DedicatedVmHostLifecycleStateEnum = "UPDATING"
	DedicatedVmHostLifecycleStateDeleting DedicatedVmHostLifecycleStateEnum = "DELETING"
	DedicatedVmHostLifecycleStateDeleted  DedicatedVmHostLifecycleStateEnum = "DELETED"
	DedicatedVmHostLifecycleStateFailed   DedicatedVmHostLifecycleStateEnum = "FAILED"
)

var mappingDedicatedVmHostLifecycleStateEnum = map[string]DedicatedVmHostLifecycleStateEnum{
	"CREATING": DedicatedVmHostLifecycleStateCreating,
	"ACTIVE":   DedicatedVmHostLifecycleStateActive,
	"UPDATING": DedicatedVmHostLifecycleStateUpdating,
	"DELETING": DedicatedVmHostLifecycleStateDeleting,
	"DELETED":  DedicatedVmHostLifecycleStateDeleted,
	"FAILED":   DedicatedVmHostLifecycleStateFailed,
}

var mappingDedicatedVmHostLifecycleStateEnumLowerCase = map[string]DedicatedVmHostLifecycleStateEnum{
	"creating": DedicatedVmHostLifecycleStateCreating,
	"active":   DedicatedVmHostLifecycleStateActive,
	"updating": DedicatedVmHostLifecycleStateUpdating,
	"deleting": DedicatedVmHostLifecycleStateDeleting,
	"deleted":  DedicatedVmHostLifecycleStateDeleted,
	"failed":   DedicatedVmHostLifecycleStateFailed,
}

// GetDedicatedVmHostLifecycleStateEnumValues Enumerates the set of values for DedicatedVmHostLifecycleStateEnum
func GetDedicatedVmHostLifecycleStateEnumValues() []DedicatedVmHostLifecycleStateEnum {
	values := make([]DedicatedVmHostLifecycleStateEnum, 0)
	for _, v := range mappingDedicatedVmHostLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedVmHostLifecycleStateEnumStringValues Enumerates the set of values in String for DedicatedVmHostLifecycleStateEnum
func GetDedicatedVmHostLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDedicatedVmHostLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedVmHostLifecycleStateEnum(val string) (DedicatedVmHostLifecycleStateEnum, bool) {
	enum, ok := mappingDedicatedVmHostLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
