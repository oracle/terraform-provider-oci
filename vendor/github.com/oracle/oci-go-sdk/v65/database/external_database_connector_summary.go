// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalDatabaseConnectorSummary An Oracle Cloud Infrastructure resource used to connect to an external Oracle Database.
// This resource stores the database connection string, user credentials, and related details that allow you to
// manage your external database using the Oracle Cloud Infrastructure Console and API interfaces.
type ExternalDatabaseConnectorSummary interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The user-friendly name for the
	// CreateExternalDatabaseConnectorDetails.
	// The name does not have to be unique.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	GetId() *string

	// The current lifecycle state of the external database connector resource.
	GetLifecycleState() ExternalDatabaseConnectorLifecycleStateEnum

	// The date and time the external connector was created.
	GetTimeCreated() *common.SDKTime

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external database resource.
	GetExternalDatabaseId() *string

	// The status of connectivity to the external database.
	GetConnectionStatus() *string

	// The date and time the `connectionStatus` of this external connector was last updated.
	GetTimeConnectionStatusLastUpdated() *common.SDKTime

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}

	// Additional information about the current lifecycle state.
	GetLifecycleDetails() *string
}

type externaldatabaseconnectorsummary struct {
	JsonData                        []byte
	CompartmentId                   *string                                     `mandatory:"true" json:"compartmentId"`
	DisplayName                     *string                                     `mandatory:"true" json:"displayName"`
	Id                              *string                                     `mandatory:"true" json:"id"`
	LifecycleState                  ExternalDatabaseConnectorLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	TimeCreated                     *common.SDKTime                             `mandatory:"true" json:"timeCreated"`
	ExternalDatabaseId              *string                                     `mandatory:"true" json:"externalDatabaseId"`
	ConnectionStatus                *string                                     `mandatory:"true" json:"connectionStatus"`
	TimeConnectionStatusLastUpdated *common.SDKTime                             `mandatory:"true" json:"timeConnectionStatusLastUpdated"`
	FreeformTags                    map[string]string                           `mandatory:"false" json:"freeformTags"`
	DefinedTags                     map[string]map[string]interface{}           `mandatory:"false" json:"definedTags"`
	LifecycleDetails                *string                                     `mandatory:"false" json:"lifecycleDetails"`
	ConnectorType                   string                                      `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *externaldatabaseconnectorsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternaldatabaseconnectorsummary externaldatabaseconnectorsummary
	s := struct {
		Model Unmarshalerexternaldatabaseconnectorsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DisplayName = s.Model.DisplayName
	m.Id = s.Model.Id
	m.LifecycleState = s.Model.LifecycleState
	m.TimeCreated = s.Model.TimeCreated
	m.ExternalDatabaseId = s.Model.ExternalDatabaseId
	m.ConnectionStatus = s.Model.ConnectionStatus
	m.TimeConnectionStatusLastUpdated = s.Model.TimeConnectionStatusLastUpdated
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externaldatabaseconnectorsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := ExternalMacsConnectorSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m externaldatabaseconnectorsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDisplayName returns DisplayName
func (m externaldatabaseconnectorsummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetId returns Id
func (m externaldatabaseconnectorsummary) GetId() *string {
	return m.Id
}

//GetLifecycleState returns LifecycleState
func (m externaldatabaseconnectorsummary) GetLifecycleState() ExternalDatabaseConnectorLifecycleStateEnum {
	return m.LifecycleState
}

//GetTimeCreated returns TimeCreated
func (m externaldatabaseconnectorsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetExternalDatabaseId returns ExternalDatabaseId
func (m externaldatabaseconnectorsummary) GetExternalDatabaseId() *string {
	return m.ExternalDatabaseId
}

//GetConnectionStatus returns ConnectionStatus
func (m externaldatabaseconnectorsummary) GetConnectionStatus() *string {
	return m.ConnectionStatus
}

//GetTimeConnectionStatusLastUpdated returns TimeConnectionStatusLastUpdated
func (m externaldatabaseconnectorsummary) GetTimeConnectionStatusLastUpdated() *common.SDKTime {
	return m.TimeConnectionStatusLastUpdated
}

//GetFreeformTags returns FreeformTags
func (m externaldatabaseconnectorsummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m externaldatabaseconnectorsummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetLifecycleDetails returns LifecycleDetails
func (m externaldatabaseconnectorsummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m externaldatabaseconnectorsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externaldatabaseconnectorsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalDatabaseConnectorLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalDatabaseConnectorLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalDatabaseConnectorSummaryConnectorTypeEnum Enum with underlying type: string
type ExternalDatabaseConnectorSummaryConnectorTypeEnum string

// Set of constants representing the allowable values for ExternalDatabaseConnectorSummaryConnectorTypeEnum
const (
	ExternalDatabaseConnectorSummaryConnectorTypeMacs ExternalDatabaseConnectorSummaryConnectorTypeEnum = "MACS"
)

var mappingExternalDatabaseConnectorSummaryConnectorTypeEnum = map[string]ExternalDatabaseConnectorSummaryConnectorTypeEnum{
	"MACS": ExternalDatabaseConnectorSummaryConnectorTypeMacs,
}

var mappingExternalDatabaseConnectorSummaryConnectorTypeEnumLowerCase = map[string]ExternalDatabaseConnectorSummaryConnectorTypeEnum{
	"macs": ExternalDatabaseConnectorSummaryConnectorTypeMacs,
}

// GetExternalDatabaseConnectorSummaryConnectorTypeEnumValues Enumerates the set of values for ExternalDatabaseConnectorSummaryConnectorTypeEnum
func GetExternalDatabaseConnectorSummaryConnectorTypeEnumValues() []ExternalDatabaseConnectorSummaryConnectorTypeEnum {
	values := make([]ExternalDatabaseConnectorSummaryConnectorTypeEnum, 0)
	for _, v := range mappingExternalDatabaseConnectorSummaryConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalDatabaseConnectorSummaryConnectorTypeEnumStringValues Enumerates the set of values in String for ExternalDatabaseConnectorSummaryConnectorTypeEnum
func GetExternalDatabaseConnectorSummaryConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingExternalDatabaseConnectorSummaryConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalDatabaseConnectorSummaryConnectorTypeEnum(val string) (ExternalDatabaseConnectorSummaryConnectorTypeEnum, bool) {
	enum, ok := mappingExternalDatabaseConnectorSummaryConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
