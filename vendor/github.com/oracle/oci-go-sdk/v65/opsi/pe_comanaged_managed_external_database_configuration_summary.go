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

// PeComanagedManagedExternalDatabaseConfigurationSummary Configuration Summary of a Private Endpoint Co-managed External database.
type PeComanagedManagedExternalDatabaseConfigurationSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database insight resource.
	DatabaseInsightId *string `mandatory:"true" json:"databaseInsightId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The database name. The database name is unique within the tenancy.
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// The user-friendly name for the database. The name does not have to be unique.
	DatabaseDisplayName *string `mandatory:"true" json:"databaseDisplayName"`

	// Ops Insights internal representation of the database type.
	DatabaseType *string `mandatory:"true" json:"databaseType"`

	// The version of the database.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// Name of the CDB.Only applies to PDB.
	CdbName *string `mandatory:"true" json:"cdbName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	ParentId *string `mandatory:"true" json:"parentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"true" json:"opsiPrivateEndpointId"`

	// Array of hostname and instance name.
	Instances []HostInstanceMap `mandatory:"true" json:"instances"`

	ExadataDetails *ExadataDetails `mandatory:"true" json:"exadataDetails"`

	// Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
	ProcessorCount *int `mandatory:"false" json:"processorCount"`
}

// GetDatabaseInsightId returns DatabaseInsightId
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetDatabaseInsightId() *string {
	return m.DatabaseInsightId
}

// GetCompartmentId returns CompartmentId
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetDatabaseName returns DatabaseName
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetDatabaseName() *string {
	return m.DatabaseName
}

// GetDatabaseDisplayName returns DatabaseDisplayName
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetDatabaseDisplayName() *string {
	return m.DatabaseDisplayName
}

// GetDatabaseType returns DatabaseType
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetDatabaseType() *string {
	return m.DatabaseType
}

// GetDatabaseVersion returns DatabaseVersion
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetDatabaseVersion() *string {
	return m.DatabaseVersion
}

// GetCdbName returns CdbName
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetCdbName() *string {
	return m.CdbName
}

// GetDefinedTags returns DefinedTags
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetFreeformTags returns FreeformTags
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetProcessorCount returns ProcessorCount
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) GetProcessorCount() *int {
	return m.ProcessorCount
}

func (m PeComanagedManagedExternalDatabaseConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PeComanagedManagedExternalDatabaseConfigurationSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePeComanagedManagedExternalDatabaseConfigurationSummary PeComanagedManagedExternalDatabaseConfigurationSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypePeComanagedManagedExternalDatabaseConfigurationSummary
	}{
		"PE_COMANAGED_DATABASE",
		(MarshalTypePeComanagedManagedExternalDatabaseConfigurationSummary)(m),
	}

	return json.Marshal(&s)
}
