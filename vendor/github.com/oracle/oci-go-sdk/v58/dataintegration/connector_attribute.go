// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ConnectorAttribute Marker class for connector attributes.
type ConnectorAttribute interface {
}

type connectorattribute struct {
	JsonData  []byte
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *connectorattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconnectorattribute connectorattribute
	s := struct {
		Model Unmarshalerconnectorattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *connectorattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "EXTERNAL_STORAGE":
		mm := ExternalStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m connectorattribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m connectorattribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConnectorAttributeModelTypeEnum Enum with underlying type: string
type ConnectorAttributeModelTypeEnum string

// Set of constants representing the allowable values for ConnectorAttributeModelTypeEnum
const (
	ConnectorAttributeModelTypeExternalStorage ConnectorAttributeModelTypeEnum = "EXTERNAL_STORAGE"
)

var mappingConnectorAttributeModelTypeEnum = map[string]ConnectorAttributeModelTypeEnum{
	"EXTERNAL_STORAGE": ConnectorAttributeModelTypeExternalStorage,
}

// GetConnectorAttributeModelTypeEnumValues Enumerates the set of values for ConnectorAttributeModelTypeEnum
func GetConnectorAttributeModelTypeEnumValues() []ConnectorAttributeModelTypeEnum {
	values := make([]ConnectorAttributeModelTypeEnum, 0)
	for _, v := range mappingConnectorAttributeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConnectorAttributeModelTypeEnumStringValues Enumerates the set of values in String for ConnectorAttributeModelTypeEnum
func GetConnectorAttributeModelTypeEnumStringValues() []string {
	return []string{
		"EXTERNAL_STORAGE",
	}
}

// GetMappingConnectorAttributeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConnectorAttributeModelTypeEnum(val string) (ConnectorAttributeModelTypeEnum, bool) {
	mappingConnectorAttributeModelTypeEnumIgnoreCase := make(map[string]ConnectorAttributeModelTypeEnum)
	for k, v := range mappingConnectorAttributeModelTypeEnum {
		mappingConnectorAttributeModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingConnectorAttributeModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
