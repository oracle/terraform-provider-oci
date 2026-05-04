// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary Summary of the Database Tools database API gateway config pool resource.
type DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary struct {

	// A system generated string that uniquely identifies a pool sub resource.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly name. Does not have to be unique, and it’s changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" json:"databaseToolsConnectionId"`

	// The time the resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

// GetKey returns Key
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) GetKey() *string {
	return m.Key
}

// GetDisplayName returns DisplayName
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) GetDisplayName() *string {
	return m.DisplayName
}

// GetDatabaseToolsConnectionId returns DatabaseToolsConnectionId
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) GetDatabaseToolsConnectionId() *string {
	return m.DatabaseToolsConnectionId
}

// GetTimeCreated returns TimeCreated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary DatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary
	}{
		"DEFAULT",
		(MarshalTypeDatabaseToolsDatabaseApiGatewayConfigPoolDefaultSummary)(m),
	}

	return json.Marshal(&s)
}
