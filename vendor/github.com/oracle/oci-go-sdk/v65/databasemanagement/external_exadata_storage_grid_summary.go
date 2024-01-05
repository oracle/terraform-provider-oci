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

// ExternalExadataStorageGridSummary The Exadata storage server grid of the Exadata infrastructure.
type ExternalExadataStorageGridSummary struct {

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

	// The number of Exadata storage servers in the Exadata infrastructure.
	ServerCount *float32 `mandatory:"false" json:"serverCount"`

	// The current lifecycle state of the database resource.
	LifecycleState DbmResourceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

// GetId returns Id
func (m ExternalExadataStorageGridSummary) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m ExternalExadataStorageGridSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetVersion returns Version
func (m ExternalExadataStorageGridSummary) GetVersion() *string {
	return m.Version
}

// GetInternalId returns InternalId
func (m ExternalExadataStorageGridSummary) GetInternalId() *string {
	return m.InternalId
}

// GetStatus returns Status
func (m ExternalExadataStorageGridSummary) GetStatus() *string {
	return m.Status
}

// GetLifecycleState returns LifecycleState
func (m ExternalExadataStorageGridSummary) GetLifecycleState() DbmResourceLifecycleStateEnum {
	return m.LifecycleState
}

// GetTimeCreated returns TimeCreated
func (m ExternalExadataStorageGridSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ExternalExadataStorageGridSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m ExternalExadataStorageGridSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetAdditionalDetails returns AdditionalDetails
func (m ExternalExadataStorageGridSummary) GetAdditionalDetails() map[string]string {
	return m.AdditionalDetails
}

func (m ExternalExadataStorageGridSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalExadataStorageGridSummary) ValidateEnumValue() (bool, error) {
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
func (m ExternalExadataStorageGridSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalExadataStorageGridSummary ExternalExadataStorageGridSummary
	s := struct {
		DiscriminatorParam string `json:"resourceType"`
		MarshalTypeExternalExadataStorageGridSummary
	}{
		"STORAGE_GRID_SUMMARY",
		(MarshalTypeExternalExadataStorageGridSummary)(m),
	}

	return json.Marshal(&s)
}
