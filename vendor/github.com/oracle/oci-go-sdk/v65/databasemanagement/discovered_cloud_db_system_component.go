// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DiscoveredCloudDbSystemComponent The details of a cloud DB system component.
type DiscoveredCloudDbSystemComponent interface {

	// The identifier of the discovered DB system component.
	GetComponentId() *string

	// The user-friendly name for the discovered DB system component. The name does not have to be unique.
	GetDisplayName() *string

	// The name of the discovered DB system component.
	GetComponentName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing OCI resource matching the discovered DB system component.
	GetResourceId() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas OCI resource matching the discovered DB system component.
	GetDbaasId() *string

	// Indicates whether the DB system component should be provisioned as an OCI resource or not.
	GetIsSelectedForMonitoring() *bool

	// The state of the discovered DB system component.
	GetStatus() DiscoveredCloudDbSystemComponentStatusEnum

	// The list of associated components.
	GetAssociatedComponents() []AssociatedCloudComponent
}

type discoveredclouddbsystemcomponent struct {
	JsonData                []byte
	ResourceId              *string                                    `mandatory:"false" json:"resourceId"`
	DbaasId                 *string                                    `mandatory:"false" json:"dbaasId"`
	IsSelectedForMonitoring *bool                                      `mandatory:"false" json:"isSelectedForMonitoring"`
	Status                  DiscoveredCloudDbSystemComponentStatusEnum `mandatory:"false" json:"status,omitempty"`
	AssociatedComponents    []AssociatedCloudComponent                 `mandatory:"false" json:"associatedComponents"`
	ComponentId             *string                                    `mandatory:"true" json:"componentId"`
	DisplayName             *string                                    `mandatory:"true" json:"displayName"`
	ComponentName           *string                                    `mandatory:"true" json:"componentName"`
	ComponentType           string                                     `json:"componentType"`
}

// UnmarshalJSON unmarshals json
func (m *discoveredclouddbsystemcomponent) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdiscoveredclouddbsystemcomponent discoveredclouddbsystemcomponent
	s := struct {
		Model Unmarshalerdiscoveredclouddbsystemcomponent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComponentId = s.Model.ComponentId
	m.DisplayName = s.Model.DisplayName
	m.ComponentName = s.Model.ComponentName
	m.ResourceId = s.Model.ResourceId
	m.DbaasId = s.Model.DbaasId
	m.IsSelectedForMonitoring = s.Model.IsSelectedForMonitoring
	m.Status = s.Model.Status
	m.AssociatedComponents = s.Model.AssociatedComponents
	m.ComponentType = s.Model.ComponentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *discoveredclouddbsystemcomponent) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComponentType {
	case "CLUSTER_INSTANCE":
		mm := DiscoveredCloudClusterInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLUGGABLE_DATABASE":
		mm := DiscoveredCloudPluggableDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLUSTER":
		mm := DiscoveredCloudCluster{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASM":
		mm := DiscoveredCloudAsm{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LISTENER":
		mm := DiscoveredCloudListener{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_INSTANCE":
		mm := DiscoveredCloudDbInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE":
		mm := DiscoveredCloudDatabase{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_HOME":
		mm := DiscoveredCloudDbHome{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATABASE_NODE":
		mm := DiscoveredCloudDbNode{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ASM_INSTANCE":
		mm := DiscoveredCloudAsmInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DiscoveredCloudDbSystemComponent: %s.", m.ComponentType)
		return *m, nil
	}
}

// GetResourceId returns ResourceId
func (m discoveredclouddbsystemcomponent) GetResourceId() *string {
	return m.ResourceId
}

// GetDbaasId returns DbaasId
func (m discoveredclouddbsystemcomponent) GetDbaasId() *string {
	return m.DbaasId
}

// GetIsSelectedForMonitoring returns IsSelectedForMonitoring
func (m discoveredclouddbsystemcomponent) GetIsSelectedForMonitoring() *bool {
	return m.IsSelectedForMonitoring
}

// GetStatus returns Status
func (m discoveredclouddbsystemcomponent) GetStatus() DiscoveredCloudDbSystemComponentStatusEnum {
	return m.Status
}

// GetAssociatedComponents returns AssociatedComponents
func (m discoveredclouddbsystemcomponent) GetAssociatedComponents() []AssociatedCloudComponent {
	return m.AssociatedComponents
}

// GetComponentId returns ComponentId
func (m discoveredclouddbsystemcomponent) GetComponentId() *string {
	return m.ComponentId
}

// GetDisplayName returns DisplayName
func (m discoveredclouddbsystemcomponent) GetDisplayName() *string {
	return m.DisplayName
}

// GetComponentName returns ComponentName
func (m discoveredclouddbsystemcomponent) GetComponentName() *string {
	return m.ComponentName
}

func (m discoveredclouddbsystemcomponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m discoveredclouddbsystemcomponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDiscoveredCloudDbSystemComponentStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDiscoveredCloudDbSystemComponentStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveredCloudDbSystemComponentStatusEnum Enum with underlying type: string
type DiscoveredCloudDbSystemComponentStatusEnum string

// Set of constants representing the allowable values for DiscoveredCloudDbSystemComponentStatusEnum
const (
	DiscoveredCloudDbSystemComponentStatusNew                       DiscoveredCloudDbSystemComponentStatusEnum = "NEW"
	DiscoveredCloudDbSystemComponentStatusExisting                  DiscoveredCloudDbSystemComponentStatusEnum = "EXISTING"
	DiscoveredCloudDbSystemComponentStatusExistingBasic             DiscoveredCloudDbSystemComponentStatusEnum = "EXISTING_BASIC"
	DiscoveredCloudDbSystemComponentStatusExistingPe                DiscoveredCloudDbSystemComponentStatusEnum = "EXISTING_PE"
	DiscoveredCloudDbSystemComponentStatusExistingBasicPe           DiscoveredCloudDbSystemComponentStatusEnum = "EXISTING_BASIC_PE"
	DiscoveredCloudDbSystemComponentStatusMarkedForUpgrade          DiscoveredCloudDbSystemComponentStatusEnum = "MARKED_FOR_UPGRADE"
	DiscoveredCloudDbSystemComponentStatusMarkedForMigration        DiscoveredCloudDbSystemComponentStatusEnum = "MARKED_FOR_MIGRATION"
	DiscoveredCloudDbSystemComponentStatusMarkedForUpgradeMigration DiscoveredCloudDbSystemComponentStatusEnum = "MARKED_FOR_UPGRADE_MIGRATION"
	DiscoveredCloudDbSystemComponentStatusMarkedForDeletion         DiscoveredCloudDbSystemComponentStatusEnum = "MARKED_FOR_DELETION"
	DiscoveredCloudDbSystemComponentStatusUnknown                   DiscoveredCloudDbSystemComponentStatusEnum = "UNKNOWN"
)

var mappingDiscoveredCloudDbSystemComponentStatusEnum = map[string]DiscoveredCloudDbSystemComponentStatusEnum{
	"NEW":                          DiscoveredCloudDbSystemComponentStatusNew,
	"EXISTING":                     DiscoveredCloudDbSystemComponentStatusExisting,
	"EXISTING_BASIC":               DiscoveredCloudDbSystemComponentStatusExistingBasic,
	"EXISTING_PE":                  DiscoveredCloudDbSystemComponentStatusExistingPe,
	"EXISTING_BASIC_PE":            DiscoveredCloudDbSystemComponentStatusExistingBasicPe,
	"MARKED_FOR_UPGRADE":           DiscoveredCloudDbSystemComponentStatusMarkedForUpgrade,
	"MARKED_FOR_MIGRATION":         DiscoveredCloudDbSystemComponentStatusMarkedForMigration,
	"MARKED_FOR_UPGRADE_MIGRATION": DiscoveredCloudDbSystemComponentStatusMarkedForUpgradeMigration,
	"MARKED_FOR_DELETION":          DiscoveredCloudDbSystemComponentStatusMarkedForDeletion,
	"UNKNOWN":                      DiscoveredCloudDbSystemComponentStatusUnknown,
}

var mappingDiscoveredCloudDbSystemComponentStatusEnumLowerCase = map[string]DiscoveredCloudDbSystemComponentStatusEnum{
	"new":                          DiscoveredCloudDbSystemComponentStatusNew,
	"existing":                     DiscoveredCloudDbSystemComponentStatusExisting,
	"existing_basic":               DiscoveredCloudDbSystemComponentStatusExistingBasic,
	"existing_pe":                  DiscoveredCloudDbSystemComponentStatusExistingPe,
	"existing_basic_pe":            DiscoveredCloudDbSystemComponentStatusExistingBasicPe,
	"marked_for_upgrade":           DiscoveredCloudDbSystemComponentStatusMarkedForUpgrade,
	"marked_for_migration":         DiscoveredCloudDbSystemComponentStatusMarkedForMigration,
	"marked_for_upgrade_migration": DiscoveredCloudDbSystemComponentStatusMarkedForUpgradeMigration,
	"marked_for_deletion":          DiscoveredCloudDbSystemComponentStatusMarkedForDeletion,
	"unknown":                      DiscoveredCloudDbSystemComponentStatusUnknown,
}

// GetDiscoveredCloudDbSystemComponentStatusEnumValues Enumerates the set of values for DiscoveredCloudDbSystemComponentStatusEnum
func GetDiscoveredCloudDbSystemComponentStatusEnumValues() []DiscoveredCloudDbSystemComponentStatusEnum {
	values := make([]DiscoveredCloudDbSystemComponentStatusEnum, 0)
	for _, v := range mappingDiscoveredCloudDbSystemComponentStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveredCloudDbSystemComponentStatusEnumStringValues Enumerates the set of values in String for DiscoveredCloudDbSystemComponentStatusEnum
func GetDiscoveredCloudDbSystemComponentStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"EXISTING",
		"EXISTING_BASIC",
		"EXISTING_PE",
		"EXISTING_BASIC_PE",
		"MARKED_FOR_UPGRADE",
		"MARKED_FOR_MIGRATION",
		"MARKED_FOR_UPGRADE_MIGRATION",
		"MARKED_FOR_DELETION",
		"UNKNOWN",
	}
}

// GetMappingDiscoveredCloudDbSystemComponentStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveredCloudDbSystemComponentStatusEnum(val string) (DiscoveredCloudDbSystemComponentStatusEnum, bool) {
	enum, ok := mappingDiscoveredCloudDbSystemComponentStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
