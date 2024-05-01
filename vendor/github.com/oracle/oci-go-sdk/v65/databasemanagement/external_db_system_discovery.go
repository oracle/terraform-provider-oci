// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDbSystemDiscovery The details of an external DB system discovery.
type ExternalDbSystemDiscovery struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system discovery.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the DB system. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the management agent
	// used for the external DB system discovery.
	AgentId *string `mandatory:"true" json:"agentId"`

	// The current lifecycle state of the external DB system discovery resource.
	LifecycleState ExternalDbSystemDiscoveryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the external DB system discovery was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the external DB system discovery was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The directory in which Oracle Grid Infrastructure is installed.
	GridHome *string `mandatory:"false" json:"gridHome"`

	// The list of DB system components that were found in the DB system discovery.
	DiscoveredComponents []DiscoveredExternalDbSystemComponent `mandatory:"false" json:"discoveredComponents"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ExternalDbSystemDiscovery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalDbSystemDiscovery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDbSystemDiscoveryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDbSystemDiscoveryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ExternalDbSystemDiscovery) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		GridHome             *string                                     `json:"gridHome"`
		DiscoveredComponents []discoveredexternaldbsystemcomponent       `json:"discoveredComponents"`
		ResourceId           *string                                     `json:"resourceId"`
		LifecycleDetails     *string                                     `json:"lifecycleDetails"`
		FreeformTags         map[string]string                           `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{}           `json:"definedTags"`
		SystemTags           map[string]map[string]interface{}           `json:"systemTags"`
		Id                   *string                                     `json:"id"`
		DisplayName          *string                                     `json:"displayName"`
		CompartmentId        *string                                     `json:"compartmentId"`
		AgentId              *string                                     `json:"agentId"`
		LifecycleState       ExternalDbSystemDiscoveryLifecycleStateEnum `json:"lifecycleState"`
		TimeCreated          *common.SDKTime                             `json:"timeCreated"`
		TimeUpdated          *common.SDKTime                             `json:"timeUpdated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.GridHome = model.GridHome

	m.DiscoveredComponents = make([]DiscoveredExternalDbSystemComponent, len(model.DiscoveredComponents))
	for i, n := range model.DiscoveredComponents {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.DiscoveredComponents[i] = nn.(DiscoveredExternalDbSystemComponent)
		} else {
			m.DiscoveredComponents[i] = nil
		}
	}
	m.ResourceId = model.ResourceId

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.AgentId = model.AgentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	return
}

// ExternalDbSystemDiscoveryLifecycleStateEnum Enum with underlying type: string
type ExternalDbSystemDiscoveryLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalDbSystemDiscoveryLifecycleStateEnum
const (
	ExternalDbSystemDiscoveryLifecycleStateCreating ExternalDbSystemDiscoveryLifecycleStateEnum = "CREATING"
	ExternalDbSystemDiscoveryLifecycleStateActive   ExternalDbSystemDiscoveryLifecycleStateEnum = "ACTIVE"
	ExternalDbSystemDiscoveryLifecycleStateInactive ExternalDbSystemDiscoveryLifecycleStateEnum = "INACTIVE"
	ExternalDbSystemDiscoveryLifecycleStateUpdating ExternalDbSystemDiscoveryLifecycleStateEnum = "UPDATING"
	ExternalDbSystemDiscoveryLifecycleStateDeleting ExternalDbSystemDiscoveryLifecycleStateEnum = "DELETING"
	ExternalDbSystemDiscoveryLifecycleStateDeleted  ExternalDbSystemDiscoveryLifecycleStateEnum = "DELETED"
	ExternalDbSystemDiscoveryLifecycleStateFailed   ExternalDbSystemDiscoveryLifecycleStateEnum = "FAILED"
)

var mappingExternalDbSystemDiscoveryLifecycleStateEnum = map[string]ExternalDbSystemDiscoveryLifecycleStateEnum{
	"CREATING": ExternalDbSystemDiscoveryLifecycleStateCreating,
	"ACTIVE":   ExternalDbSystemDiscoveryLifecycleStateActive,
	"INACTIVE": ExternalDbSystemDiscoveryLifecycleStateInactive,
	"UPDATING": ExternalDbSystemDiscoveryLifecycleStateUpdating,
	"DELETING": ExternalDbSystemDiscoveryLifecycleStateDeleting,
	"DELETED":  ExternalDbSystemDiscoveryLifecycleStateDeleted,
	"FAILED":   ExternalDbSystemDiscoveryLifecycleStateFailed,
}

var mappingExternalDbSystemDiscoveryLifecycleStateEnumLowerCase = map[string]ExternalDbSystemDiscoveryLifecycleStateEnum{
	"creating": ExternalDbSystemDiscoveryLifecycleStateCreating,
	"active":   ExternalDbSystemDiscoveryLifecycleStateActive,
	"inactive": ExternalDbSystemDiscoveryLifecycleStateInactive,
	"updating": ExternalDbSystemDiscoveryLifecycleStateUpdating,
	"deleting": ExternalDbSystemDiscoveryLifecycleStateDeleting,
	"deleted":  ExternalDbSystemDiscoveryLifecycleStateDeleted,
	"failed":   ExternalDbSystemDiscoveryLifecycleStateFailed,
}

// GetExternalDbSystemDiscoveryLifecycleStateEnumValues Enumerates the set of values for ExternalDbSystemDiscoveryLifecycleStateEnum
func GetExternalDbSystemDiscoveryLifecycleStateEnumValues() []ExternalDbSystemDiscoveryLifecycleStateEnum {
	values := make([]ExternalDbSystemDiscoveryLifecycleStateEnum, 0)
	for _, v := range mappingExternalDbSystemDiscoveryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDbSystemDiscoveryLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalDbSystemDiscoveryLifecycleStateEnum
func GetExternalDbSystemDiscoveryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalDbSystemDiscoveryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDbSystemDiscoveryLifecycleStateEnum(val string) (ExternalDbSystemDiscoveryLifecycleStateEnum, bool) {
	enum, ok := mappingExternalDbSystemDiscoveryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
