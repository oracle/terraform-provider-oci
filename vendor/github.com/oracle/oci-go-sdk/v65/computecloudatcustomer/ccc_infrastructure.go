// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccInfrastructure The Oracle Cloud Infrastructure resource representing the connection to the hardware and
// software located in a customer's data center running the Compute Cloud@Customer IaaS services.
type CccInfrastructure struct {

	// The Compute Cloud@Customer infrastructure OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This cannot be changed once created.
	Id *string `mandatory:"true" json:"id"`

	// The name that will be used to display the Compute Cloud@Customer infrastructure
	// in the Oracle Cloud Infrastructure console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The infrastructure compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network subnet that is
	// used to communicate with Compute Cloud@Customer infrastructure.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Compute Cloud@Customer infrastructure creation date and time, using an RFC3339 formatted
	// datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Compute Cloud@Customer infrastructure.
	LifecycleState CccInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Compute Cloud@Customer infrastructure short name.
	// This cannot be changed once created. The
	// short name is used to refer to the infrastructure in several contexts and is unique.
	ShortName *string `mandatory:"false" json:"shortName"`

	// A mutable client-meaningful text description of the Compute Cloud@Customer infrastructure.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The current connection state of the infrastructure. A user can only update
	// it from REQUEST to READY or from any state back to REJECT. The system automatically
	// handles the REJECT to REQUEST, READY to CONNECTED, or CONNECTED to DISCONNECTED transitions.
	ConnectionState CccInfrastructureConnectionStateEnum `mandatory:"false" json:"connectionState,omitempty"`

	// A message describing the current connection state in more detail.
	ConnectionDetails *string `mandatory:"false" json:"connectionDetails"`

	// Schedule used for upgrades. If no schedule is associated with the infrastructure,
	// it can be updated at any time.
	CccUpgradeScheduleId *string `mandatory:"false" json:"cccUpgradeScheduleId"`

	// Fingerprint of a Compute Cloud@Customer infrastructure in a data center generated
	// during the initial connection to this resource. The fingerprint should be verified
	// by the administrator when changing the connectionState from REQUEST to READY.
	ProvisioningFingerprint *string `mandatory:"false" json:"provisioningFingerprint"`

	// Code that is required for service personnel to connect a
	// Compute Cloud@Customer infrastructure in a data center to this resource.
	// This code will only be available when the connectionState is REJECT (usually
	// at create time of the Compute Cloud@Customer infrastructure).
	ProvisioningPin *string `mandatory:"false" json:"provisioningPin"`

	// Compute Cloud@Customer infrastructure updated date and time, using an RFC3339 formatted
	// datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current lifecycle state in more detail.
	// For example, this can be used to provide actionable information for a resource that is in
	// a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	InfrastructureInventory *CccInfrastructureInventory `mandatory:"false" json:"infrastructureInventory"`

	InfrastructureNetworkConfiguration *CccInfrastructureNetworkConfiguration `mandatory:"false" json:"infrastructureNetworkConfiguration"`

	UpgradeInformation *CccUpgradeInformation `mandatory:"false" json:"upgradeInformation"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CccInfrastructure) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccInfrastructure) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccInfrastructureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCccInfrastructureLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCccInfrastructureConnectionStateEnum(string(m.ConnectionState)); !ok && m.ConnectionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionState: %s. Supported values are: %s.", m.ConnectionState, strings.Join(GetCccInfrastructureConnectionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccInfrastructureConnectionStateEnum Enum with underlying type: string
type CccInfrastructureConnectionStateEnum string

// Set of constants representing the allowable values for CccInfrastructureConnectionStateEnum
const (
	CccInfrastructureConnectionStateReject       CccInfrastructureConnectionStateEnum = "REJECT"
	CccInfrastructureConnectionStateRequest      CccInfrastructureConnectionStateEnum = "REQUEST"
	CccInfrastructureConnectionStateReady        CccInfrastructureConnectionStateEnum = "READY"
	CccInfrastructureConnectionStateConnected    CccInfrastructureConnectionStateEnum = "CONNECTED"
	CccInfrastructureConnectionStateDisconnected CccInfrastructureConnectionStateEnum = "DISCONNECTED"
)

var mappingCccInfrastructureConnectionStateEnum = map[string]CccInfrastructureConnectionStateEnum{
	"REJECT":       CccInfrastructureConnectionStateReject,
	"REQUEST":      CccInfrastructureConnectionStateRequest,
	"READY":        CccInfrastructureConnectionStateReady,
	"CONNECTED":    CccInfrastructureConnectionStateConnected,
	"DISCONNECTED": CccInfrastructureConnectionStateDisconnected,
}

var mappingCccInfrastructureConnectionStateEnumLowerCase = map[string]CccInfrastructureConnectionStateEnum{
	"reject":       CccInfrastructureConnectionStateReject,
	"request":      CccInfrastructureConnectionStateRequest,
	"ready":        CccInfrastructureConnectionStateReady,
	"connected":    CccInfrastructureConnectionStateConnected,
	"disconnected": CccInfrastructureConnectionStateDisconnected,
}

// GetCccInfrastructureConnectionStateEnumValues Enumerates the set of values for CccInfrastructureConnectionStateEnum
func GetCccInfrastructureConnectionStateEnumValues() []CccInfrastructureConnectionStateEnum {
	values := make([]CccInfrastructureConnectionStateEnum, 0)
	for _, v := range mappingCccInfrastructureConnectionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCccInfrastructureConnectionStateEnumStringValues Enumerates the set of values in String for CccInfrastructureConnectionStateEnum
func GetCccInfrastructureConnectionStateEnumStringValues() []string {
	return []string{
		"REJECT",
		"REQUEST",
		"READY",
		"CONNECTED",
		"DISCONNECTED",
	}
}

// GetMappingCccInfrastructureConnectionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccInfrastructureConnectionStateEnum(val string) (CccInfrastructureConnectionStateEnum, bool) {
	enum, ok := mappingCccInfrastructureConnectionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CccInfrastructureLifecycleStateEnum Enum with underlying type: string
type CccInfrastructureLifecycleStateEnum string

// Set of constants representing the allowable values for CccInfrastructureLifecycleStateEnum
const (
	CccInfrastructureLifecycleStateActive         CccInfrastructureLifecycleStateEnum = "ACTIVE"
	CccInfrastructureLifecycleStateNeedsAttention CccInfrastructureLifecycleStateEnum = "NEEDS_ATTENTION"
	CccInfrastructureLifecycleStateDeleted        CccInfrastructureLifecycleStateEnum = "DELETED"
	CccInfrastructureLifecycleStateFailed         CccInfrastructureLifecycleStateEnum = "FAILED"
)

var mappingCccInfrastructureLifecycleStateEnum = map[string]CccInfrastructureLifecycleStateEnum{
	"ACTIVE":          CccInfrastructureLifecycleStateActive,
	"NEEDS_ATTENTION": CccInfrastructureLifecycleStateNeedsAttention,
	"DELETED":         CccInfrastructureLifecycleStateDeleted,
	"FAILED":          CccInfrastructureLifecycleStateFailed,
}

var mappingCccInfrastructureLifecycleStateEnumLowerCase = map[string]CccInfrastructureLifecycleStateEnum{
	"active":          CccInfrastructureLifecycleStateActive,
	"needs_attention": CccInfrastructureLifecycleStateNeedsAttention,
	"deleted":         CccInfrastructureLifecycleStateDeleted,
	"failed":          CccInfrastructureLifecycleStateFailed,
}

// GetCccInfrastructureLifecycleStateEnumValues Enumerates the set of values for CccInfrastructureLifecycleStateEnum
func GetCccInfrastructureLifecycleStateEnumValues() []CccInfrastructureLifecycleStateEnum {
	values := make([]CccInfrastructureLifecycleStateEnum, 0)
	for _, v := range mappingCccInfrastructureLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCccInfrastructureLifecycleStateEnumStringValues Enumerates the set of values in String for CccInfrastructureLifecycleStateEnum
func GetCccInfrastructureLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCccInfrastructureLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccInfrastructureLifecycleStateEnum(val string) (CccInfrastructureLifecycleStateEnum, bool) {
	enum, ok := mappingCccInfrastructureLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
