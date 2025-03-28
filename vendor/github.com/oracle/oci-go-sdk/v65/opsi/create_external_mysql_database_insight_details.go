// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExternalMysqlDatabaseInsightDetails The information about database to be analyzed.
type CreateExternalMysqlDatabaseInsightDetails struct {

	// Compartment Identifier of database
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The DBM owned database connector OCID (https://docs.oracle.com/iaas/database-management/doc/view-connector-details.html) mapping to the database credentials and connection details.
	DatabaseConnectorId *string `mandatory:"true" json:"databaseConnectorId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetCompartmentId returns CompartmentId
func (m CreateExternalMysqlDatabaseInsightDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateExternalMysqlDatabaseInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateExternalMysqlDatabaseInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateExternalMysqlDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExternalMysqlDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateExternalMysqlDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateExternalMysqlDatabaseInsightDetails CreateExternalMysqlDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreateExternalMysqlDatabaseInsightDetails
	}{
		"EXTERNAL_MYSQL_DATABASE_SYSTEM",
		(MarshalTypeCreateExternalMysqlDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}
