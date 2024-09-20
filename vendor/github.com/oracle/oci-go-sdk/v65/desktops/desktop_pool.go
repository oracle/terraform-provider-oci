// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DesktopPool Provides information about a desktop pool including all configuration parameters.
type DesktopPool struct {

	// The OCID of the desktop pool.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment of the desktop pool.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user friendly display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the desktop pool.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The maximum number of desktops permitted in the desktop pool.
	MaximumSize *int `mandatory:"true" json:"maximumSize"`

	// The maximum number of standby desktops available in the desktop pool.
	StandbySize *int `mandatory:"true" json:"standbySize"`

	// The shape of the desktop pool.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// Indicates whether storage is enabled for the desktop pool.
	IsStorageEnabled *bool `mandatory:"true" json:"isStorageEnabled"`

	// The size in GBs of the storage for the desktop pool.
	StorageSizeInGBs *int `mandatory:"true" json:"storageSizeInGBs"`

	// The backup policy OCID of the storage.
	StorageBackupPolicyId *string `mandatory:"true" json:"storageBackupPolicyId"`

	DevicePolicy *DesktopDevicePolicy `mandatory:"true" json:"devicePolicy"`

	AvailabilityPolicy *DesktopAvailabilityPolicy `mandatory:"true" json:"availabilityPolicy"`

	Image *DesktopImage `mandatory:"true" json:"image"`

	NetworkConfiguration *DesktopNetworkConfiguration `mandatory:"true" json:"networkConfiguration"`

	// Contact information of the desktop pool administrator.
	// Avoid entering confidential information.
	ContactDetails *string `mandatory:"true" json:"contactDetails"`

	// Indicates whether desktop pool users have administrative privileges on their desktop.
	ArePrivilegedUsers *bool `mandatory:"true" json:"arePrivilegedUsers"`

	// The availability domain of the desktop pool.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// A user friendly description providing additional information about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	ShapeConfig *DesktopPoolShapeConfig `mandatory:"false" json:"shapeConfig"`

	// Indicates whether the desktop pool uses dedicated virtual machine hosts.
	UseDedicatedVmHost DesktopPoolUseDedicatedVmHostEnum `mandatory:"false" json:"useDedicatedVmHost,omitempty"`

	SessionLifecycleActions *DesktopSessionLifecycleActions `mandatory:"false" json:"sessionLifecycleActions"`

	// The start time of the desktop pool.
	TimeStartScheduled *common.SDKTime `mandatory:"false" json:"timeStartScheduled"`

	// The stop time of the desktop pool.
	TimeStopScheduled *common.SDKTime `mandatory:"false" json:"timeStopScheduled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A list of network security groups for the network.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	PrivateAccessDetails *DesktopPoolPrivateAccessDetails `mandatory:"false" json:"privateAccessDetails"`
}

func (m DesktopPool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopPool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDesktopPoolUseDedicatedVmHostEnum(string(m.UseDedicatedVmHost)); !ok && m.UseDedicatedVmHost != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UseDedicatedVmHost: %s. Supported values are: %s.", m.UseDedicatedVmHost, strings.Join(GetDesktopPoolUseDedicatedVmHostEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DesktopPoolUseDedicatedVmHostEnum Enum with underlying type: string
type DesktopPoolUseDedicatedVmHostEnum string

// Set of constants representing the allowable values for DesktopPoolUseDedicatedVmHostEnum
const (
	DesktopPoolUseDedicatedVmHostTrue  DesktopPoolUseDedicatedVmHostEnum = "TRUE"
	DesktopPoolUseDedicatedVmHostFalse DesktopPoolUseDedicatedVmHostEnum = "FALSE"
	DesktopPoolUseDedicatedVmHostAuto  DesktopPoolUseDedicatedVmHostEnum = "AUTO"
)

var mappingDesktopPoolUseDedicatedVmHostEnum = map[string]DesktopPoolUseDedicatedVmHostEnum{
	"TRUE":  DesktopPoolUseDedicatedVmHostTrue,
	"FALSE": DesktopPoolUseDedicatedVmHostFalse,
	"AUTO":  DesktopPoolUseDedicatedVmHostAuto,
}

var mappingDesktopPoolUseDedicatedVmHostEnumLowerCase = map[string]DesktopPoolUseDedicatedVmHostEnum{
	"true":  DesktopPoolUseDedicatedVmHostTrue,
	"false": DesktopPoolUseDedicatedVmHostFalse,
	"auto":  DesktopPoolUseDedicatedVmHostAuto,
}

// GetDesktopPoolUseDedicatedVmHostEnumValues Enumerates the set of values for DesktopPoolUseDedicatedVmHostEnum
func GetDesktopPoolUseDedicatedVmHostEnumValues() []DesktopPoolUseDedicatedVmHostEnum {
	values := make([]DesktopPoolUseDedicatedVmHostEnum, 0)
	for _, v := range mappingDesktopPoolUseDedicatedVmHostEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopPoolUseDedicatedVmHostEnumStringValues Enumerates the set of values in String for DesktopPoolUseDedicatedVmHostEnum
func GetDesktopPoolUseDedicatedVmHostEnumStringValues() []string {
	return []string{
		"TRUE",
		"FALSE",
		"AUTO",
	}
}

// GetMappingDesktopPoolUseDedicatedVmHostEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopPoolUseDedicatedVmHostEnum(val string) (DesktopPoolUseDedicatedVmHostEnum, bool) {
	enum, ok := mappingDesktopPoolUseDedicatedVmHostEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
