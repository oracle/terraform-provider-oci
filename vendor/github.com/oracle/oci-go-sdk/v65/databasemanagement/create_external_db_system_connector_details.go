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

// CreateExternalDbSystemConnectorDetails The details required to create an external DB system connector.
type CreateExternalDbSystemConnectorDetails interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system.
	GetExternalDbSystemId() *string

	// The user-friendly name for the external connector. The name does not have to be unique.
	GetDisplayName() *string
}

type createexternaldbsystemconnectordetails struct {
	JsonData           []byte
	DisplayName        *string `mandatory:"false" json:"displayName"`
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`
	ConnectorType      string  `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *createexternaldbsystemconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateexternaldbsystemconnectordetails createexternaldbsystemconnectordetails
	s := struct {
		Model Unmarshalercreateexternaldbsystemconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExternalDbSystemId = s.Model.ExternalDbSystemId
	m.DisplayName = s.Model.DisplayName
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createexternaldbsystemconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := CreateExternalDbSystemMacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateExternalDbSystemConnectorDetails: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createexternaldbsystemconnectordetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetExternalDbSystemId returns ExternalDbSystemId
func (m createexternaldbsystemconnectordetails) GetExternalDbSystemId() *string {
	return m.ExternalDbSystemId
}

func (m createexternaldbsystemconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createexternaldbsystemconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateExternalDbSystemConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type CreateExternalDbSystemConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for CreateExternalDbSystemConnectorDetailsConnectorTypeEnum
const (
	CreateExternalDbSystemConnectorDetailsConnectorTypeMacs CreateExternalDbSystemConnectorDetailsConnectorTypeEnum = "MACS"
)

var mappingCreateExternalDbSystemConnectorDetailsConnectorTypeEnum = map[string]CreateExternalDbSystemConnectorDetailsConnectorTypeEnum{
	"MACS": CreateExternalDbSystemConnectorDetailsConnectorTypeMacs,
}

var mappingCreateExternalDbSystemConnectorDetailsConnectorTypeEnumLowerCase = map[string]CreateExternalDbSystemConnectorDetailsConnectorTypeEnum{
	"macs": CreateExternalDbSystemConnectorDetailsConnectorTypeMacs,
}

// GetCreateExternalDbSystemConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for CreateExternalDbSystemConnectorDetailsConnectorTypeEnum
func GetCreateExternalDbSystemConnectorDetailsConnectorTypeEnumValues() []CreateExternalDbSystemConnectorDetailsConnectorTypeEnum {
	values := make([]CreateExternalDbSystemConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingCreateExternalDbSystemConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExternalDbSystemConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for CreateExternalDbSystemConnectorDetailsConnectorTypeEnum
func GetCreateExternalDbSystemConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingCreateExternalDbSystemConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExternalDbSystemConnectorDetailsConnectorTypeEnum(val string) (CreateExternalDbSystemConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingCreateExternalDbSystemConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
