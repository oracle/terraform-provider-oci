// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMacsManagedCloudDatabaseInsightDetails The information about database to be analyzed.
type CreateMacsManagedCloudDatabaseInsightDetails struct {

	// Compartment Identifier of database
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Management Agent
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	ConnectionDetails *ConnectionDetails `mandatory:"true" json:"connectionDetails"`

	ConnectionCredentialDetails CredentialDetails `mandatory:"true" json:"connectionCredentialDetails"`

	// OCI database resource type
	DatabaseResourceType *string `mandatory:"true" json:"databaseResourceType"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Database Deployment Type (EXACS will be supported in the future)
	DeploymentType CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum `mandatory:"true" json:"deploymentType"`
}

// GetCompartmentId returns CompartmentId
func (m CreateMacsManagedCloudDatabaseInsightDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateMacsManagedCloudDatabaseInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateMacsManagedCloudDatabaseInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateMacsManagedCloudDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMacsManagedCloudDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateMacsManagedCloudDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMacsManagedCloudDatabaseInsightDetails CreateMacsManagedCloudDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreateMacsManagedCloudDatabaseInsightDetails
	}{
		"MACS_MANAGED_CLOUD_DATABASE",
		(MarshalTypeCreateMacsManagedCloudDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateMacsManagedCloudDatabaseInsightDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags                map[string]string                                              `json:"freeformTags"`
		DefinedTags                 map[string]map[string]interface{}                              `json:"definedTags"`
		SystemTags                  map[string]map[string]interface{}                              `json:"systemTags"`
		CompartmentId               *string                                                        `json:"compartmentId"`
		DatabaseId                  *string                                                        `json:"databaseId"`
		ManagementAgentId           *string                                                        `json:"managementAgentId"`
		ConnectionDetails           *ConnectionDetails                                             `json:"connectionDetails"`
		ConnectionCredentialDetails credentialdetails                                              `json:"connectionCredentialDetails"`
		DatabaseResourceType        *string                                                        `json:"databaseResourceType"`
		DeploymentType              CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum `json:"deploymentType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	m.DatabaseId = model.DatabaseId

	m.ManagementAgentId = model.ManagementAgentId

	m.ConnectionDetails = model.ConnectionDetails

	nn, e = model.ConnectionCredentialDetails.UnmarshalPolymorphicJSON(model.ConnectionCredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ConnectionCredentialDetails = nn.(CredentialDetails)
	} else {
		m.ConnectionCredentialDetails = nil
	}

	m.DatabaseResourceType = model.DatabaseResourceType

	m.DeploymentType = model.DeploymentType

	return
}

// CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum Enum with underlying type: string
type CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum string

// Set of constants representing the allowable values for CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum
const (
	CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeVirtualMachine CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum = "VIRTUAL_MACHINE"
	CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeBareMetal      CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum = "BARE_METAL"
	CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeExacc          CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum = "EXACC"
	CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeExacs          CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum = "EXACS"
)

var mappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum = map[string]CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum{
	"VIRTUAL_MACHINE": CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeVirtualMachine,
	"BARE_METAL":      CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeBareMetal,
	"EXACC":           CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeExacc,
	"EXACS":           CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeExacs,
}

var mappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumLowerCase = map[string]CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum{
	"virtual_machine": CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeVirtualMachine,
	"bare_metal":      CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeBareMetal,
	"exacc":           CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeExacc,
	"exacs":           CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeExacs,
}

// GetCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumValues Enumerates the set of values for CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum
func GetCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumValues() []CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum {
	values := make([]CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum, 0)
	for _, v := range mappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumStringValues Enumerates the set of values in String for CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum
func GetCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_MACHINE",
		"BARE_METAL",
		"EXACC",
		"EXACS",
	}
}

// GetMappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum(val string) (CreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnum, bool) {
	enum, ok := mappingCreateMacsManagedCloudDatabaseInsightDetailsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
