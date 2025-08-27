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

// UpdateCloudDbSystemConnectorDetails The details required to update a cloud DB system connector.
type UpdateCloudDbSystemConnectorDetails interface {
}

type updateclouddbsystemconnectordetails struct {
	JsonData      []byte
	ConnectorType string `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *updateclouddbsystemconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateclouddbsystemconnectordetails updateclouddbsystemconnectordetails
	s := struct {
		Model Unmarshalerupdateclouddbsystemconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateclouddbsystemconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := UpdateCloudDbSystemMacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateCloudDbSystemConnectorDetails: %s.", m.ConnectorType)
		return *m, nil
	}
}

func (m updateclouddbsystemconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateclouddbsystemconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum
const (
	UpdateCloudDbSystemConnectorDetailsConnectorTypeMacs UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum = "MACS"
)

var mappingUpdateCloudDbSystemConnectorDetailsConnectorTypeEnum = map[string]UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum{
	"MACS": UpdateCloudDbSystemConnectorDetailsConnectorTypeMacs,
}

var mappingUpdateCloudDbSystemConnectorDetailsConnectorTypeEnumLowerCase = map[string]UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum{
	"macs": UpdateCloudDbSystemConnectorDetailsConnectorTypeMacs,
}

// GetUpdateCloudDbSystemConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum
func GetUpdateCloudDbSystemConnectorDetailsConnectorTypeEnumValues() []UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum {
	values := make([]UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingUpdateCloudDbSystemConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCloudDbSystemConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum
func GetUpdateCloudDbSystemConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingUpdateCloudDbSystemConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCloudDbSystemConnectorDetailsConnectorTypeEnum(val string) (UpdateCloudDbSystemConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingUpdateCloudDbSystemConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
