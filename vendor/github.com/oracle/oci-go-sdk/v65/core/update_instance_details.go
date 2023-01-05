// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInstanceDetails The representation of UpdateInstanceDetails
type UpdateInstanceDetails struct {

	// The OCID of the compute capacity reservation this instance is launched under.
	// You can remove the instance from a reservation by specifying an empty string as input for this field.
	// For more information, see Capacity Reservations (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/reserve-capacity.htm#default).
	CapacityReservationId *string `mandatory:"false" json:"capacityReservationId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	AgentConfig *UpdateInstanceAgentConfigDetails `mandatory:"false" json:"agentConfig"`

	// Custom metadata key/value string pairs that you provide. Any set of key/value pairs
	// provided here will completely replace the current set of key/value pairs in the `metadata`
	// field on the instance.
	// The "user_data" field and the "ssh_authorized_keys" field cannot be changed after an instance
	// has launched. Any request that updates, removes, or adds either of these fields will be
	// rejected. You must provide the same values for "user_data" and "ssh_authorized_keys" that
	// already exist on the instance.
	// The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of
	// 32,000 bytes.
	Metadata map[string]string `mandatory:"false" json:"metadata"`

	// Additional metadata key/value pairs that you provide. They serve the same purpose and
	// functionality as fields in the `metadata` object.
	// They are distinguished from `metadata` fields in that these can be nested JSON objects
	// (whereas `metadata` fields are string/string maps only).
	// The "user_data" field and the "ssh_authorized_keys" field cannot be changed after an instance
	// has launched. Any request that updates, removes, or adds either of these fields will be
	// rejected. You must provide the same values for "user_data" and "ssh_authorized_keys" that
	// already exist on the instance.
	// The combined size of the `metadata` and `extendedMetadata` objects can be a maximum of
	// 32,000 bytes.
	ExtendedMetadata map[string]interface{} `mandatory:"false" json:"extendedMetadata"`

	// The shape of the instance. The shape determines the number of CPUs and the amount of memory
	// allocated to the instance. For more information about how to change shapes, and a list of
	// shapes that are supported, see
	// Editing an Instance (https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/resizinginstances.htm).
	// For details about the CPUs, memory, and other properties of each shape, see
	// Compute Shapes (https://docs.cloud.oracle.com/iaas/Content/Compute/References/computeshapes.htm).
	// The new shape must be compatible with the image that was used to launch the instance. You
	// can enumerate all available shapes and determine image compatibility by calling
	// ListShapes.
	// If the instance is running when you change the shape, the instance is rebooted.
	// Example: `VM.Standard2.1`
	Shape *string `mandatory:"false" json:"shape"`

	ShapeConfig *UpdateInstanceShapeConfigDetails `mandatory:"false" json:"shapeConfig"`

	SourceDetails UpdateInstanceSourceDetails `mandatory:"false" json:"sourceDetails"`

	// The preferred maintenance action for an instance. The default is LIVE_MIGRATE, if live migration is supported.
	// * `LIVE_MIGRATE` - Run maintenance using a live migration.
	// * `REBOOT` - Run maintenance using a reboot.
	PreferredMaintenanceAction UpdateInstanceDetailsPreferredMaintenanceActionEnum `mandatory:"false" json:"preferredMaintenanceAction,omitempty"`

	// The parameter acts as a fail-safe to prevent unwanted downtime when updating a running instance.
	// The default is ALLOW_DOWNTIME.
	// * `ALLOW_DOWNTIME` - Compute might reboot the instance while updating the instance if a reboot is required.
	// * `AVOID_DOWNTIME` - If the instance is in running state, Compute tries to update the instance without rebooting
	//                   it. If the instance requires a reboot to be updated, an error is returned and the instance
	//                   is not updated. If the instance is stopped, it is updated and remains in the stopped state.
	UpdateOperationConstraint UpdateInstanceDetailsUpdateOperationConstraintEnum `mandatory:"false" json:"updateOperationConstraint,omitempty"`

	InstanceOptions *InstanceOptions `mandatory:"false" json:"instanceOptions"`

	// A fault domain is a grouping of hardware and infrastructure within an availability domain.
	// Each availability domain contains three fault domains. Fault domains let you distribute your
	// instances so that they are not on the same physical hardware within a single availability domain.
	// A hardware failure or Compute hardware maintenance that affects one fault domain does not affect
	// instances in other fault domains.
	// To get a list of fault domains, use the
	// ListFaultDomains operation in the
	// Identity and Access Management Service API.
	// Example: `FAULT-DOMAIN-1`
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	LaunchOptions *UpdateLaunchOptions `mandatory:"false" json:"launchOptions"`

	AvailabilityConfig *UpdateInstanceAvailabilityConfigDetails `mandatory:"false" json:"availabilityConfig"`

	// For a VM instance, resets the scheduled time that the instance will be reboot migrated for
	// infrastructure maintenance, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// If the instance hasn't been rebooted after this date, Oracle reboots the instance within 24 hours of the time
	// and date that maintenance is due.
	// To get the maximum possible date that a maintenance reboot can be extended,
	// use GetInstanceMaintenanceReboot.
	// Regardless of how the instance is stopped, this flag is reset to empty as soon as the instance reaches the
	// Stopped state.
	// To reboot migrate a bare metal instance, use the InstanceAction operation.
	// For more information, see
	// Infrastructure Maintenance (https://docs.cloud.oracle.com/iaas/Content/Compute/References/infrastructure-maintenance.htm).
	// Example: `2018-05-25T21:10:29.600Z`
	TimeMaintenanceRebootDue *common.SDKTime `mandatory:"false" json:"timeMaintenanceRebootDue"`
}

func (m UpdateInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInstanceDetailsPreferredMaintenanceActionEnum(string(m.PreferredMaintenanceAction)); !ok && m.PreferredMaintenanceAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreferredMaintenanceAction: %s. Supported values are: %s.", m.PreferredMaintenanceAction, strings.Join(GetUpdateInstanceDetailsPreferredMaintenanceActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateInstanceDetailsUpdateOperationConstraintEnum(string(m.UpdateOperationConstraint)); !ok && m.UpdateOperationConstraint != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateOperationConstraint: %s. Supported values are: %s.", m.UpdateOperationConstraint, strings.Join(GetUpdateInstanceDetailsUpdateOperationConstraintEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateInstanceDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CapacityReservationId      *string                                             `json:"capacityReservationId"`
		DefinedTags                map[string]map[string]interface{}                   `json:"definedTags"`
		DisplayName                *string                                             `json:"displayName"`
		FreeformTags               map[string]string                                   `json:"freeformTags"`
		AgentConfig                *UpdateInstanceAgentConfigDetails                   `json:"agentConfig"`
		Metadata                   map[string]string                                   `json:"metadata"`
		ExtendedMetadata           map[string]interface{}                              `json:"extendedMetadata"`
		Shape                      *string                                             `json:"shape"`
		ShapeConfig                *UpdateInstanceShapeConfigDetails                   `json:"shapeConfig"`
		SourceDetails              updateinstancesourcedetails                         `json:"sourceDetails"`
		PreferredMaintenanceAction UpdateInstanceDetailsPreferredMaintenanceActionEnum `json:"preferredMaintenanceAction"`
		UpdateOperationConstraint  UpdateInstanceDetailsUpdateOperationConstraintEnum  `json:"updateOperationConstraint"`
		InstanceOptions            *InstanceOptions                                    `json:"instanceOptions"`
		FaultDomain                *string                                             `json:"faultDomain"`
		LaunchOptions              *UpdateLaunchOptions                                `json:"launchOptions"`
		AvailabilityConfig         *UpdateInstanceAvailabilityConfigDetails            `json:"availabilityConfig"`
		TimeMaintenanceRebootDue   *common.SDKTime                                     `json:"timeMaintenanceRebootDue"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CapacityReservationId = model.CapacityReservationId

	m.DefinedTags = model.DefinedTags

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.AgentConfig = model.AgentConfig

	m.Metadata = model.Metadata

	m.ExtendedMetadata = model.ExtendedMetadata

	m.Shape = model.Shape

	m.ShapeConfig = model.ShapeConfig

	nn, e = model.SourceDetails.UnmarshalPolymorphicJSON(model.SourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SourceDetails = nn.(UpdateInstanceSourceDetails)
	} else {
		m.SourceDetails = nil
	}

	m.PreferredMaintenanceAction = model.PreferredMaintenanceAction

	m.UpdateOperationConstraint = model.UpdateOperationConstraint

	m.InstanceOptions = model.InstanceOptions

	m.FaultDomain = model.FaultDomain

	m.LaunchOptions = model.LaunchOptions

	m.AvailabilityConfig = model.AvailabilityConfig

	m.TimeMaintenanceRebootDue = model.TimeMaintenanceRebootDue

	return
}

// UpdateInstanceDetailsPreferredMaintenanceActionEnum Enum with underlying type: string
type UpdateInstanceDetailsPreferredMaintenanceActionEnum string

// Set of constants representing the allowable values for UpdateInstanceDetailsPreferredMaintenanceActionEnum
const (
	UpdateInstanceDetailsPreferredMaintenanceActionLiveMigrate UpdateInstanceDetailsPreferredMaintenanceActionEnum = "LIVE_MIGRATE"
	UpdateInstanceDetailsPreferredMaintenanceActionReboot      UpdateInstanceDetailsPreferredMaintenanceActionEnum = "REBOOT"
)

var mappingUpdateInstanceDetailsPreferredMaintenanceActionEnum = map[string]UpdateInstanceDetailsPreferredMaintenanceActionEnum{
	"LIVE_MIGRATE": UpdateInstanceDetailsPreferredMaintenanceActionLiveMigrate,
	"REBOOT":       UpdateInstanceDetailsPreferredMaintenanceActionReboot,
}

var mappingUpdateInstanceDetailsPreferredMaintenanceActionEnumLowerCase = map[string]UpdateInstanceDetailsPreferredMaintenanceActionEnum{
	"live_migrate": UpdateInstanceDetailsPreferredMaintenanceActionLiveMigrate,
	"reboot":       UpdateInstanceDetailsPreferredMaintenanceActionReboot,
}

// GetUpdateInstanceDetailsPreferredMaintenanceActionEnumValues Enumerates the set of values for UpdateInstanceDetailsPreferredMaintenanceActionEnum
func GetUpdateInstanceDetailsPreferredMaintenanceActionEnumValues() []UpdateInstanceDetailsPreferredMaintenanceActionEnum {
	values := make([]UpdateInstanceDetailsPreferredMaintenanceActionEnum, 0)
	for _, v := range mappingUpdateInstanceDetailsPreferredMaintenanceActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInstanceDetailsPreferredMaintenanceActionEnumStringValues Enumerates the set of values in String for UpdateInstanceDetailsPreferredMaintenanceActionEnum
func GetUpdateInstanceDetailsPreferredMaintenanceActionEnumStringValues() []string {
	return []string{
		"LIVE_MIGRATE",
		"REBOOT",
	}
}

// GetMappingUpdateInstanceDetailsPreferredMaintenanceActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInstanceDetailsPreferredMaintenanceActionEnum(val string) (UpdateInstanceDetailsPreferredMaintenanceActionEnum, bool) {
	enum, ok := mappingUpdateInstanceDetailsPreferredMaintenanceActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateInstanceDetailsUpdateOperationConstraintEnum Enum with underlying type: string
type UpdateInstanceDetailsUpdateOperationConstraintEnum string

// Set of constants representing the allowable values for UpdateInstanceDetailsUpdateOperationConstraintEnum
const (
	UpdateInstanceDetailsUpdateOperationConstraintAllowDowntime UpdateInstanceDetailsUpdateOperationConstraintEnum = "ALLOW_DOWNTIME"
	UpdateInstanceDetailsUpdateOperationConstraintAvoidDowntime UpdateInstanceDetailsUpdateOperationConstraintEnum = "AVOID_DOWNTIME"
)

var mappingUpdateInstanceDetailsUpdateOperationConstraintEnum = map[string]UpdateInstanceDetailsUpdateOperationConstraintEnum{
	"ALLOW_DOWNTIME": UpdateInstanceDetailsUpdateOperationConstraintAllowDowntime,
	"AVOID_DOWNTIME": UpdateInstanceDetailsUpdateOperationConstraintAvoidDowntime,
}

var mappingUpdateInstanceDetailsUpdateOperationConstraintEnumLowerCase = map[string]UpdateInstanceDetailsUpdateOperationConstraintEnum{
	"allow_downtime": UpdateInstanceDetailsUpdateOperationConstraintAllowDowntime,
	"avoid_downtime": UpdateInstanceDetailsUpdateOperationConstraintAvoidDowntime,
}

// GetUpdateInstanceDetailsUpdateOperationConstraintEnumValues Enumerates the set of values for UpdateInstanceDetailsUpdateOperationConstraintEnum
func GetUpdateInstanceDetailsUpdateOperationConstraintEnumValues() []UpdateInstanceDetailsUpdateOperationConstraintEnum {
	values := make([]UpdateInstanceDetailsUpdateOperationConstraintEnum, 0)
	for _, v := range mappingUpdateInstanceDetailsUpdateOperationConstraintEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInstanceDetailsUpdateOperationConstraintEnumStringValues Enumerates the set of values in String for UpdateInstanceDetailsUpdateOperationConstraintEnum
func GetUpdateInstanceDetailsUpdateOperationConstraintEnumStringValues() []string {
	return []string{
		"ALLOW_DOWNTIME",
		"AVOID_DOWNTIME",
	}
}

// GetMappingUpdateInstanceDetailsUpdateOperationConstraintEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInstanceDetailsUpdateOperationConstraintEnum(val string) (UpdateInstanceDetailsUpdateOperationConstraintEnum, bool) {
	enum, ok := mappingUpdateInstanceDetailsUpdateOperationConstraintEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
