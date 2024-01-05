// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbSystem Information about a database system.
type DbSystem struct {

	// A unique identifier for the database system. Immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the database system. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the database system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time that the database system was created, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the database system.
	LifecycleState DbSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Type of the database system.
	SystemType DbSystemSystemTypeEnum `mandatory:"true" json:"systemType"`

	// The major and minor versions of the database system software.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The name of the shape for the database instance.
	// Example: `VM.Standard.E4.Flex`
	Shape *string `mandatory:"true" json:"shape"`

	// The total number of OCPUs available to each database instance node.
	InstanceOcpuCount *int `mandatory:"true" json:"instanceOcpuCount"`

	// The total amount of memory available to each database instance node, in gigabytes.
	InstanceMemorySizeInGBs *int `mandatory:"true" json:"instanceMemorySizeInGBs"`

	StorageDetails StorageDetails `mandatory:"true" json:"storageDetails"`

	NetworkDetails *NetworkDetails `mandatory:"true" json:"networkDetails"`

	ManagementPolicy *ManagementPolicy `mandatory:"true" json:"managementPolicy"`

	// A description of the database system.
	Description *string `mandatory:"false" json:"description"`

	// The date and time that the database system was updated, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The database system administrator username.
	AdminUsername *string `mandatory:"false" json:"adminUsername"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the configuration associated with the database system.
	ConfigId *string `mandatory:"false" json:"configId"`

	// Count of instances, or nodes, in the database system.
	InstanceCount *int `mandatory:"false" json:"instanceCount"`

	// The list of instances, or nodes, in the database system.
	Instances []DbInstance `mandatory:"false" json:"instances"`

	Source SourceDetails `mandatory:"false" json:"source"`
}

func (m DbSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemSystemTypeEnum(string(m.SystemType)); !ok && m.SystemType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SystemType: %s. Supported values are: %s.", m.SystemType, strings.Join(GetDbSystemSystemTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *DbSystem) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description             *string                           `json:"description"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		LifecycleDetails        *string                           `json:"lifecycleDetails"`
		AdminUsername           *string                           `json:"adminUsername"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		SystemTags              map[string]map[string]interface{} `json:"systemTags"`
		ConfigId                *string                           `json:"configId"`
		InstanceCount           *int                              `json:"instanceCount"`
		Instances               []DbInstance                      `json:"instances"`
		Source                  sourcedetails                     `json:"source"`
		Id                      *string                           `json:"id"`
		DisplayName             *string                           `json:"displayName"`
		CompartmentId           *string                           `json:"compartmentId"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		LifecycleState          DbSystemLifecycleStateEnum        `json:"lifecycleState"`
		SystemType              DbSystemSystemTypeEnum            `json:"systemType"`
		DbVersion               *string                           `json:"dbVersion"`
		Shape                   *string                           `json:"shape"`
		InstanceOcpuCount       *int                              `json:"instanceOcpuCount"`
		InstanceMemorySizeInGBs *int                              `json:"instanceMemorySizeInGBs"`
		StorageDetails          storagedetails                    `json:"storageDetails"`
		NetworkDetails          *NetworkDetails                   `json:"networkDetails"`
		ManagementPolicy        *ManagementPolicy                 `json:"managementPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleDetails = model.LifecycleDetails

	m.AdminUsername = model.AdminUsername

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.ConfigId = model.ConfigId

	m.InstanceCount = model.InstanceCount

	m.Instances = make([]DbInstance, len(model.Instances))
	copy(m.Instances, model.Instances)
	nn, e = model.Source.UnmarshalPolymorphicJSON(model.Source.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Source = nn.(SourceDetails)
	} else {
		m.Source = nil
	}

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.LifecycleState = model.LifecycleState

	m.SystemType = model.SystemType

	m.DbVersion = model.DbVersion

	m.Shape = model.Shape

	m.InstanceOcpuCount = model.InstanceOcpuCount

	m.InstanceMemorySizeInGBs = model.InstanceMemorySizeInGBs

	nn, e = model.StorageDetails.UnmarshalPolymorphicJSON(model.StorageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.StorageDetails = nn.(StorageDetails)
	} else {
		m.StorageDetails = nil
	}

	m.NetworkDetails = model.NetworkDetails

	m.ManagementPolicy = model.ManagementPolicy

	return
}

// DbSystemLifecycleStateEnum Enum with underlying type: string
type DbSystemLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemLifecycleStateEnum
const (
	DbSystemLifecycleStateCreating       DbSystemLifecycleStateEnum = "CREATING"
	DbSystemLifecycleStateUpdating       DbSystemLifecycleStateEnum = "UPDATING"
	DbSystemLifecycleStateActive         DbSystemLifecycleStateEnum = "ACTIVE"
	DbSystemLifecycleStateInactive       DbSystemLifecycleStateEnum = "INACTIVE"
	DbSystemLifecycleStateDeleting       DbSystemLifecycleStateEnum = "DELETING"
	DbSystemLifecycleStateDeleted        DbSystemLifecycleStateEnum = "DELETED"
	DbSystemLifecycleStateFailed         DbSystemLifecycleStateEnum = "FAILED"
	DbSystemLifecycleStateNeedsAttention DbSystemLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDbSystemLifecycleStateEnum = map[string]DbSystemLifecycleStateEnum{
	"CREATING":        DbSystemLifecycleStateCreating,
	"UPDATING":        DbSystemLifecycleStateUpdating,
	"ACTIVE":          DbSystemLifecycleStateActive,
	"INACTIVE":        DbSystemLifecycleStateInactive,
	"DELETING":        DbSystemLifecycleStateDeleting,
	"DELETED":         DbSystemLifecycleStateDeleted,
	"FAILED":          DbSystemLifecycleStateFailed,
	"NEEDS_ATTENTION": DbSystemLifecycleStateNeedsAttention,
}

var mappingDbSystemLifecycleStateEnumLowerCase = map[string]DbSystemLifecycleStateEnum{
	"creating":        DbSystemLifecycleStateCreating,
	"updating":        DbSystemLifecycleStateUpdating,
	"active":          DbSystemLifecycleStateActive,
	"inactive":        DbSystemLifecycleStateInactive,
	"deleting":        DbSystemLifecycleStateDeleting,
	"deleted":         DbSystemLifecycleStateDeleted,
	"failed":          DbSystemLifecycleStateFailed,
	"needs_attention": DbSystemLifecycleStateNeedsAttention,
}

// GetDbSystemLifecycleStateEnumValues Enumerates the set of values for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumValues() []DbSystemLifecycleStateEnum {
	values := make([]DbSystemLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemLifecycleStateEnum
func GetDbSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDbSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemLifecycleStateEnum(val string) (DbSystemLifecycleStateEnum, bool) {
	enum, ok := mappingDbSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbSystemSystemTypeEnum Enum with underlying type: string
type DbSystemSystemTypeEnum string

// Set of constants representing the allowable values for DbSystemSystemTypeEnum
const (
	DbSystemSystemTypeOciOptimizedStorage DbSystemSystemTypeEnum = "OCI_OPTIMIZED_STORAGE"
)

var mappingDbSystemSystemTypeEnum = map[string]DbSystemSystemTypeEnum{
	"OCI_OPTIMIZED_STORAGE": DbSystemSystemTypeOciOptimizedStorage,
}

var mappingDbSystemSystemTypeEnumLowerCase = map[string]DbSystemSystemTypeEnum{
	"oci_optimized_storage": DbSystemSystemTypeOciOptimizedStorage,
}

// GetDbSystemSystemTypeEnumValues Enumerates the set of values for DbSystemSystemTypeEnum
func GetDbSystemSystemTypeEnumValues() []DbSystemSystemTypeEnum {
	values := make([]DbSystemSystemTypeEnum, 0)
	for _, v := range mappingDbSystemSystemTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemSystemTypeEnumStringValues Enumerates the set of values in String for DbSystemSystemTypeEnum
func GetDbSystemSystemTypeEnumStringValues() []string {
	return []string{
		"OCI_OPTIMIZED_STORAGE",
	}
}

// GetMappingDbSystemSystemTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemSystemTypeEnum(val string) (DbSystemSystemTypeEnum, bool) {
	enum, ok := mappingDbSystemSystemTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
