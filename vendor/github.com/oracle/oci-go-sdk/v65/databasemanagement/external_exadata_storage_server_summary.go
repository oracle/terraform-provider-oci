// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ExternalExadataStorageServerSummary The Exadata storage server of the Exadata infrastructure.
type ExternalExadataStorageServerSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	Id *string `mandatory:"true" json:"id"`

	// The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The version of the Exadata resource.
	Version *string `mandatory:"false" json:"version"`

	// The internal ID of the Exadata resource.
	InternalId *string `mandatory:"false" json:"internalId"`

	// The status of the Exadata resource.
	Status *string `mandatory:"false" json:"status"`

	// The timestamp of the creation of the Exadata resource.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The timestamp of the last update of the Exadata resource.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The details of the lifecycle state of the Exadata resource.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The additional details of the resource defined in `{"key": "value"}` format.
	// Example: `{"bar-key": "value"}`
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// The make model of the Exadata storage server.
	MakeModel *string `mandatory:"false" json:"makeModel"`

	// The IP address of the Exadata storage server.
	IpAddress *string `mandatory:"false" json:"ipAddress"`

	// The CPU count of the Exadata storage server.
	CpuCount *float32 `mandatory:"false" json:"cpuCount"`

	// The Exadata storage server memory size in GB.
	MemoryGB *float64 `mandatory:"false" json:"memoryGB"`

	// The maximum hard disk IO operations per second of the Exadata storage server.
	MaxHardDiskIOPS *int `mandatory:"false" json:"maxHardDiskIOPS"`

	// The maximum hard disk IO throughput in MB/s of the Exadata storage server.
	MaxHardDiskThroughput *int `mandatory:"false" json:"maxHardDiskThroughput"`

	// The maximum flash disk IO operations per second of the Exadata storage server.
	MaxFlashDiskIOPS *int `mandatory:"false" json:"maxFlashDiskIOPS"`

	// The maximum flash disk IO throughput in MB/s of the Exadata storage server.
	MaxFlashDiskThroughput *int `mandatory:"false" json:"maxFlashDiskThroughput"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connector.
	ConnectorId *string `mandatory:"false" json:"connectorId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ExternalExadataStorageServerSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExternalExadataStorageServerSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalExadataStorageServerSummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalExadataStorageServerSummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalExadataStorageServerSummary) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m ExternalExadataStorageServerSummary) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m ExternalExadataStorageServerSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExternalExadataStorageServerSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExternalExadataStorageServerSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m ExternalExadataStorageServerSummary) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExternalExadataStorageServerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataStorageServerSummary) ValidateEnumValue() (bool, error) {
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
func (m ExternalExadataStorageServerSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataStorageServerSummary ExternalExadataStorageServerSummary
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExternalExadataStorageServerSummary
	}{
		"STORAGE_SERVER_SUMMARY",
		(MarshalTypeExternalExadataStorageServerSummary)(m),
	}

	return json.Marshal(&s)
}
