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

// CreatePeComanagedDatabaseInsightDetails The information about database to be analyzed. Either an opsiPrivateEndpointId or dbmPrivateEndpointId must be specified. If the dbmPrivateEndpointId is specified, a new Operations Insights private endpoint will be created.
type CreatePeComanagedDatabaseInsightDetails struct {

	// Compartment Identifier of database
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	// OCI database resource type
	DatabaseResourceType *string `mandatory:"true" json:"databaseResourceType"`

	// Database service name used for connection requests.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	CredentialDetails CredentialDetails `mandatory:"true" json:"credentialDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
	OpsiPrivateEndpointId *string `mandatory:"false" json:"opsiPrivateEndpointId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint
	DbmPrivateEndpointId *string `mandatory:"false" json:"dbmPrivateEndpointId"`

	ConnectionDetails *PeComanagedDatabaseConnectionDetails `mandatory:"false" json:"connectionDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Database Deployment Type
	DeploymentType CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum `mandatory:"true" json:"deploymentType"`
}

// GetCompartmentId returns CompartmentId
func (m CreatePeComanagedDatabaseInsightDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreatePeComanagedDatabaseInsightDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreatePeComanagedDatabaseInsightDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreatePeComanagedDatabaseInsightDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePeComanagedDatabaseInsightDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum(string(m.DeploymentType)); !ok && m.DeploymentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentType: %s. Supported values are: %s.", m.DeploymentType, strings.Join(GetCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreatePeComanagedDatabaseInsightDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreatePeComanagedDatabaseInsightDetails CreatePeComanagedDatabaseInsightDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreatePeComanagedDatabaseInsightDetails
	}{
		"PE_COMANAGED_DATABASE",
		(MarshalTypeCreatePeComanagedDatabaseInsightDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreatePeComanagedDatabaseInsightDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		FreeformTags          map[string]string                                         `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{}                         `json:"definedTags"`
		OpsiPrivateEndpointId *string                                                   `json:"opsiPrivateEndpointId"`
		DbmPrivateEndpointId  *string                                                   `json:"dbmPrivateEndpointId"`
		ConnectionDetails     *PeComanagedDatabaseConnectionDetails                     `json:"connectionDetails"`
		SystemTags            map[string]map[string]interface{}                         `json:"systemTags"`
		CompartmentId         *string                                                   `json:"compartmentId"`
		DatabaseId            *string                                                   `json:"databaseId"`
		DatabaseResourceType  *string                                                   `json:"databaseResourceType"`
		ServiceName           *string                                                   `json:"serviceName"`
		CredentialDetails     credentialdetails                                         `json:"credentialDetails"`
		DeploymentType        CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum `json:"deploymentType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.OpsiPrivateEndpointId = model.OpsiPrivateEndpointId

	m.DbmPrivateEndpointId = model.DbmPrivateEndpointId

	m.ConnectionDetails = model.ConnectionDetails

	m.SystemTags = model.SystemTags

	m.CompartmentId = model.CompartmentId

	m.DatabaseId = model.DatabaseId

	m.DatabaseResourceType = model.DatabaseResourceType

	m.ServiceName = model.ServiceName

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(CredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.DeploymentType = model.DeploymentType

	return
}

// CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum Enum with underlying type: string
type CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum string

// Set of constants representing the allowable values for CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum
const (
	CreatePeComanagedDatabaseInsightDetailsDeploymentTypeVirtualMachine CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum = "VIRTUAL_MACHINE"
	CreatePeComanagedDatabaseInsightDetailsDeploymentTypeBareMetal      CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum = "BARE_METAL"
	CreatePeComanagedDatabaseInsightDetailsDeploymentTypeExacs          CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum = "EXACS"
)

var mappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum = map[string]CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum{
	"VIRTUAL_MACHINE": CreatePeComanagedDatabaseInsightDetailsDeploymentTypeVirtualMachine,
	"BARE_METAL":      CreatePeComanagedDatabaseInsightDetailsDeploymentTypeBareMetal,
	"EXACS":           CreatePeComanagedDatabaseInsightDetailsDeploymentTypeExacs,
}

var mappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumLowerCase = map[string]CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum{
	"virtual_machine": CreatePeComanagedDatabaseInsightDetailsDeploymentTypeVirtualMachine,
	"bare_metal":      CreatePeComanagedDatabaseInsightDetailsDeploymentTypeBareMetal,
	"exacs":           CreatePeComanagedDatabaseInsightDetailsDeploymentTypeExacs,
}

// GetCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumValues Enumerates the set of values for CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum
func GetCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumValues() []CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum {
	values := make([]CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum, 0)
	for _, v := range mappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumStringValues Enumerates the set of values in String for CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum
func GetCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumStringValues() []string {
	return []string{
		"VIRTUAL_MACHINE",
		"BARE_METAL",
		"EXACS",
	}
}

// GetMappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum(val string) (CreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnum, bool) {
	enum, ok := mappingCreatePeComanagedDatabaseInsightDetailsDeploymentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
