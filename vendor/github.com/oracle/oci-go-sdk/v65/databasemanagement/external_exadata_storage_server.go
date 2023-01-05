// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalExadataStorageServer The Exadata storage server details.
type ExternalExadataStorageServer struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	Id *string `mandatory:"true" json:"id"`

	// The name of the resource. English letters and "-", "." only.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The version of the resource.
	Version *string `mandatory:"false" json:"version"`

	// The internal ID.
	InternalId *string `mandatory:"false" json:"internalId"`

	// The status of the entity.
	Status *string `mandatory:"false" json:"status"`

	// The timestamp of the creation.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The timestamp of the last update.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The details of the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The additional details of the resource defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of Exadata infrastructure system.
	ExadataInfrastructureId *string `mandatory:"false" json:"exadataInfrastructureId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of Exadata storage grid.
	StorageGridId *string `mandatory:"false" json:"storageGridId"`

	// The make model of the storage server.
	MakeModel *string `mandatory:"false" json:"makeModel"`

	// The IP address of the storage server.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// CPU count of the storage server
	CpuCount *float32 `mandatory:"false" json:"cpuCount"`

	// Storage server memory size in GB
	MemoryGB *float64 `mandatory:"false" json:"memoryGB"`

	// Maximum hard disk IO operations per second of the storage server
	MaxHardDiskIOPS *int `mandatory:"false" json:"maxHardDiskIOPS"`

	// Maximum hard disk IO throughput in MB/s of the storage server
	MaxHardDiskThroughput *int `mandatory:"false" json:"maxHardDiskThroughput"`

	// Maximum flash disk IO operations per second of the storage server
	MaxFlashDiskIOPS *int `mandatory:"false" json:"maxFlashDiskIOPS"`

	// Maximum flash disk IO throughput in MB/s of the storage server
	MaxFlashDiskThroughput *int `mandatory:"false" json:"maxFlashDiskThroughput"`

	Connector *ExternalExadataStorageConnectorSummary `mandatory:"false" json:"connector"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m ExternalExadataStorageServer) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m ExternalExadataStorageServer) GetDisplayName() *string {
	return m.DisplayName
}

//GetVersion returns Version
func (m ExternalExadataStorageServer) GetVersion() *string {
	return m.Version
}

//GetInternalId returns InternalId
func (m ExternalExadataStorageServer) GetInternalId() *string {
	return m.InternalId
}

//GetStatus returns Status
func (m ExternalExadataStorageServer) GetStatus() *string {
	return m.Status
}

//GetLifecycleState returns LifecycleState
func (m ExternalExadataStorageServer) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m ExternalExadataStorageServer) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m ExternalExadataStorageServer) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleDetails returns LifecycleDetails
func (m ExternalExadataStorageServer) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetAdditionalDetails returns AdditionalDetails
func (m ExternalExadataStorageServer) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExternalExadataStorageServer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataStorageServer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDbmResourceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbmResourceLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalExadataStorageServer) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataStorageServer ExternalExadataStorageServer
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExternalExadataStorageServer
	}{
		"STORAGE_SERVER",
		(MarshalTypeExternalExadataStorageServer)(m),
	}

	return json.Marshal(&s)
}
