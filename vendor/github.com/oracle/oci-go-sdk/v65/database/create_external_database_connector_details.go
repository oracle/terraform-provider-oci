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

// CreateExternalDatabaseConnectorDetails Details for creating an external database connector resource.
type CreateExternalDatabaseConnectorDetails interface {

	// The user-friendly name for the
	// CreateExternalDatabaseConnectorDetails.
	// The name does not have to be unique.
	GetDisplayName() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external database resource.
	GetExternalDatabaseId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}
}

type createexternaldatabaseconnectordetails struct {
	JsonData           []byte
	DisplayName        *string                           `mandatory:"true" json:"displayName"`
	ExternalDatabaseId *string                           `mandatory:"true" json:"externalDatabaseId"`
	FreeformTags       map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags        map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	ConnectorType      string                            `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *createexternaldatabaseconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateexternaldatabaseconnectordetails createexternaldatabaseconnectordetails
	s := struct {
		Model Unmarshalercreateexternaldatabaseconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.ExternalDatabaseId = s.Model.ExternalDatabaseId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createexternaldatabaseconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := CreateExternalMacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m createexternaldatabaseconnectordetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetExternalDatabaseId returns ExternalDatabaseId
func (m createexternaldatabaseconnectordetails) GetExternalDatabaseId() *string {
	return m.ExternalDatabaseId
}

//GetFreeformTags returns FreeformTags
func (m createexternaldatabaseconnectordetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createexternaldatabaseconnectordetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createexternaldatabaseconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createexternaldatabaseconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateExternalDatabaseConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type CreateExternalDatabaseConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for CreateExternalDatabaseConnectorDetailsConnectorTypeEnum
const (
	CreateExternalDatabaseConnectorDetailsConnectorTypeMacs CreateExternalDatabaseConnectorDetailsConnectorTypeEnum = "MACS"
)

var mappingCreateExternalDatabaseConnectorDetailsConnectorTypeEnum = map[string]CreateExternalDatabaseConnectorDetailsConnectorTypeEnum{
	"MACS": CreateExternalDatabaseConnectorDetailsConnectorTypeMacs,
}

var mappingCreateExternalDatabaseConnectorDetailsConnectorTypeEnumLowerCase = map[string]CreateExternalDatabaseConnectorDetailsConnectorTypeEnum{
	"macs": CreateExternalDatabaseConnectorDetailsConnectorTypeMacs,
}

// GetCreateExternalDatabaseConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for CreateExternalDatabaseConnectorDetailsConnectorTypeEnum
func GetCreateExternalDatabaseConnectorDetailsConnectorTypeEnumValues() []CreateExternalDatabaseConnectorDetailsConnectorTypeEnum {
	values := make([]CreateExternalDatabaseConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingCreateExternalDatabaseConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateExternalDatabaseConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for CreateExternalDatabaseConnectorDetailsConnectorTypeEnum
func GetCreateExternalDatabaseConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingCreateExternalDatabaseConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateExternalDatabaseConnectorDetailsConnectorTypeEnum(val string) (CreateExternalDatabaseConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingCreateExternalDatabaseConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
