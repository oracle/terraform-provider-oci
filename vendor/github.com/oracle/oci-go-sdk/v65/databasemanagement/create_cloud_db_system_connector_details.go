// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCloudDbSystemConnectorDetails The details required to create a cloud DB system connector.
type CreateCloudDbSystemConnectorDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system.
	GetCloudDbSystemId() *string

	// The user-friendly name for the cloud connector. The name does not have to be unique.
	GetDisplayName() *string
}

type createclouddbsystemconnectordetails struct {
	JsonData        []byte
	DisplayName     *string `mandatory:"false" json:"displayName"`
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`
	ConnectorType   string  `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *createclouddbsystemconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateclouddbsystemconnectordetails createclouddbsystemconnectordetails
	s := struct {
		Model Unmarshalercreateclouddbsystemconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CloudDbSystemId = s.Model.CloudDbSystemId
	m.DisplayName = s.Model.DisplayName
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createclouddbsystemconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := CreateCloudDbSystemMacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateCloudDbSystemConnectorDetails: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m createclouddbsystemconnectordetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetCloudDbSystemId returns CloudDbSystemId
func (m createclouddbsystemconnectordetails) GetCloudDbSystemId() *string {
	return m.CloudDbSystemId
}

func (m createclouddbsystemconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createclouddbsystemconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateCloudDbSystemConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type CreateCloudDbSystemConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for CreateCloudDbSystemConnectorDetailsConnectorTypeEnum
const (
	CreateCloudDbSystemConnectorDetailsConnectorTypeMacs CreateCloudDbSystemConnectorDetailsConnectorTypeEnum = "MACS"
)

var mappingCreateCloudDbSystemConnectorDetailsConnectorTypeEnum = map[string]CreateCloudDbSystemConnectorDetailsConnectorTypeEnum{
	"MACS": CreateCloudDbSystemConnectorDetailsConnectorTypeMacs,
}

var mappingCreateCloudDbSystemConnectorDetailsConnectorTypeEnumLowerCase = map[string]CreateCloudDbSystemConnectorDetailsConnectorTypeEnum{
	"macs": CreateCloudDbSystemConnectorDetailsConnectorTypeMacs,
}

// GetCreateCloudDbSystemConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for CreateCloudDbSystemConnectorDetailsConnectorTypeEnum
func GetCreateCloudDbSystemConnectorDetailsConnectorTypeEnumValues() []CreateCloudDbSystemConnectorDetailsConnectorTypeEnum {
	values := make([]CreateCloudDbSystemConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingCreateCloudDbSystemConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateCloudDbSystemConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for CreateCloudDbSystemConnectorDetailsConnectorTypeEnum
func GetCreateCloudDbSystemConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingCreateCloudDbSystemConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateCloudDbSystemConnectorDetailsConnectorTypeEnum(val string) (CreateCloudDbSystemConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingCreateCloudDbSystemConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
