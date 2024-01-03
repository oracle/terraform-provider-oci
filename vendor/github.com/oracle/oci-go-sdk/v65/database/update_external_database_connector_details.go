// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateExternalDatabaseConnectorDetails Details for updating an external database connector.
type UpdateExternalDatabaseConnectorDetails interface {

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}

	// The user-friendly name for the
	// CreateExternalDatabaseConnectorDetails.
	// The name does not have to be unique.
	GetDisplayName() *string
}

type updateexternaldatabaseconnectordetails struct {
	JsonData      []byte
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	ConnectorType string                            `json:"connectorType"`
}

// UnmarshalJSON unmarshals json
func (m *updateexternaldatabaseconnectordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateexternaldatabaseconnectordetails updateexternaldatabaseconnectordetails
	s := struct {
		Model Unmarshalerupdateexternaldatabaseconnectordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DisplayName = s.Model.DisplayName
	m.ConnectorType = s.Model.ConnectorType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateexternaldatabaseconnectordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectorType {
	case "MACS":
		mm := UpdateExternalMacsConnectorDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateExternalDatabaseConnectorDetails: %s.", m.ConnectorType)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m updateexternaldatabaseconnectordetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m updateexternaldatabaseconnectordetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetDisplayName returns DisplayName
func (m updateexternaldatabaseconnectordetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m updateexternaldatabaseconnectordetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateexternaldatabaseconnectordetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum Enum with underlying type: string
type UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum string

// Set of constants representing the allowable values for UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum
const (
	UpdateExternalDatabaseConnectorDetailsConnectorTypeMacs UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum = "MACS"
)

var mappingUpdateExternalDatabaseConnectorDetailsConnectorTypeEnum = map[string]UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum{
	"MACS": UpdateExternalDatabaseConnectorDetailsConnectorTypeMacs,
}

var mappingUpdateExternalDatabaseConnectorDetailsConnectorTypeEnumLowerCase = map[string]UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum{
	"macs": UpdateExternalDatabaseConnectorDetailsConnectorTypeMacs,
}

// GetUpdateExternalDatabaseConnectorDetailsConnectorTypeEnumValues Enumerates the set of values for UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum
func GetUpdateExternalDatabaseConnectorDetailsConnectorTypeEnumValues() []UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum {
	values := make([]UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum, 0)
	for _, v := range mappingUpdateExternalDatabaseConnectorDetailsConnectorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExternalDatabaseConnectorDetailsConnectorTypeEnumStringValues Enumerates the set of values in String for UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum
func GetUpdateExternalDatabaseConnectorDetailsConnectorTypeEnumStringValues() []string {
	return []string{
		"MACS",
	}
}

// GetMappingUpdateExternalDatabaseConnectorDetailsConnectorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExternalDatabaseConnectorDetailsConnectorTypeEnum(val string) (UpdateExternalDatabaseConnectorDetailsConnectorTypeEnum, bool) {
	enum, ok := mappingUpdateExternalDatabaseConnectorDetailsConnectorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
