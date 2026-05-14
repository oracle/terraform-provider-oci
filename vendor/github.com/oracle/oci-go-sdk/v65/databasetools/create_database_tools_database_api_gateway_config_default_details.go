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

// CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails Details for the new Database Tools database API gateway config for the default type.
type CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools database API gateway config.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The RESTful service definition location.
	MetadataSource DatabaseApiGatewayConfigMetadataSourceEnum `mandatory:"true" json:"metadataSource"`
}

// GetCompartmentId returns CompartmentId
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDisplayName returns DisplayName
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetMetadataSource returns MetadataSource
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) GetMetadataSource() DatabaseApiGatewayConfigMetadataSourceEnum {
	return m.MetadataSource
}

// GetDefinedTags returns DefinedTags
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetLocks returns Locks
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) GetLocks() []ResourceLock {
	return m.Locks
}

func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabaseApiGatewayConfigMetadataSourceEnum(string(m.MetadataSource)); !ok && m.MetadataSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetadataSource: %s. Supported values are: %s.", m.MetadataSource, strings.Join(GetDatabaseApiGatewayConfigMetadataSourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails CreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails
	}{
		"DEFAULT",
		(MarshalTypeCreateDatabaseToolsDatabaseApiGatewayConfigDefaultDetails)(m),
	}

	return json.Marshal(&s)
}
