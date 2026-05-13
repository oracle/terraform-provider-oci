// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigDefault Manage configuration related to database API gateway config resources.
type DatabaseToolsDatabaseApiGatewayConfigDefault struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools database API gateway config.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools database API gateway config.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time the Database Tools database API gateway config was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Database Tools database API gateway config was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The RESTful service definition location.
	MetadataSource DatabaseApiGatewayConfigMetadataSourceEnum `mandatory:"true" json:"metadataSource"`

	// The current state of the Database Tools database API gateway config.
	LifecycleState DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetDisplayName() *string {
	return m.DisplayName
}

// GetMetadataSource returns MetadataSource
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetMetadataSource() DatabaseApiGatewayConfigMetadataSourceEnum {
	return m.MetadataSource
}

// GetLifecycleState returns LifecycleState
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetLifecycleState() DatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetDefinedTags returns DefinedTags
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetSystemTags returns SystemTags
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetLocks returns Locks
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) GetLocks() []ResourceLock {
	return m.Locks
}

func (m DatabaseToolsDatabaseApiGatewayConfigDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseApiGatewayConfigMetadataSourceEnum(string(m.MetadataSource)); !ok && m.MetadataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetadataSource: %s. Supported values are: %s.", m.MetadataSource, strings.Join(GetDatabaseApiGatewayConfigMetadataSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseToolsDatabaseApiGatewayConfigLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsDatabaseApiGatewayConfigDefault) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsDatabaseApiGatewayConfigDefault DatabaseToolsDatabaseApiGatewayConfigDefault
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsDatabaseApiGatewayConfigDefault
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsDatabaseApiGatewayConfigDefault)(m),
	}

	return json.Marshal(&s)
}
