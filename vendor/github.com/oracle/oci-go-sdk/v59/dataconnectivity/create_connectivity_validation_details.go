// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// CreateConnectivityValidationDetails Input to perform connector validation. If defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
type CreateConnectivityValidationDetails interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Details about the operator.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []OutputPort

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// An array of parameters used in the data flow.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues
}

type createconnectivityvalidationdetails struct {
	JsonData       []byte
	Key            *string          `mandatory:"false" json:"key"`
	ModelVersion   *string          `mandatory:"false" json:"modelVersion"`
	ParentRef      *ParentReference `mandatory:"false" json:"parentRef"`
	Name           *string          `mandatory:"false" json:"name"`
	Description    *string          `mandatory:"false" json:"description"`
	ObjectVersion  *int             `mandatory:"false" json:"objectVersion"`
	InputPorts     []InputPort      `mandatory:"false" json:"inputPorts"`
	OutputPorts    []OutputPort     `mandatory:"false" json:"outputPorts"`
	ObjectStatus   *int             `mandatory:"false" json:"objectStatus"`
	Identifier     *string          `mandatory:"false" json:"identifier"`
	Parameters     []Parameter      `mandatory:"false" json:"parameters"`
	OpConfigValues *ConfigValues    `mandatory:"false" json:"opConfigValues"`
	ModelType      string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createconnectivityvalidationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateconnectivityvalidationdetails createconnectivityvalidationdetails
	s := struct {
		Model Unmarshalercreateconnectivityvalidationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectVersion = s.Model.ObjectVersion
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createconnectivityvalidationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "TARGET_OPERATOR":
		mm := Target{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOURCE_OPERATOR":
		mm := Source{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m createconnectivityvalidationdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createconnectivityvalidationdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m createconnectivityvalidationdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m createconnectivityvalidationdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m createconnectivityvalidationdetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m createconnectivityvalidationdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m createconnectivityvalidationdetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m createconnectivityvalidationdetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m createconnectivityvalidationdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m createconnectivityvalidationdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m createconnectivityvalidationdetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m createconnectivityvalidationdetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m createconnectivityvalidationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createconnectivityvalidationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateConnectivityValidationDetailsModelTypeEnum Enum with underlying type: string
type CreateConnectivityValidationDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateConnectivityValidationDetailsModelTypeEnum
const (
	CreateConnectivityValidationDetailsModelTypeSourceOperator CreateConnectivityValidationDetailsModelTypeEnum = "SOURCE_OPERATOR"
	CreateConnectivityValidationDetailsModelTypeTargetOperator CreateConnectivityValidationDetailsModelTypeEnum = "TARGET_OPERATOR"
)

var mappingCreateConnectivityValidationDetailsModelTypeEnum = map[string]CreateConnectivityValidationDetailsModelTypeEnum{
	"SOURCE_OPERATOR": CreateConnectivityValidationDetailsModelTypeSourceOperator,
	"TARGET_OPERATOR": CreateConnectivityValidationDetailsModelTypeTargetOperator,
}

// GetCreateConnectivityValidationDetailsModelTypeEnumValues Enumerates the set of values for CreateConnectivityValidationDetailsModelTypeEnum
func GetCreateConnectivityValidationDetailsModelTypeEnumValues() []CreateConnectivityValidationDetailsModelTypeEnum {
	values := make([]CreateConnectivityValidationDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateConnectivityValidationDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateConnectivityValidationDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateConnectivityValidationDetailsModelTypeEnum
func GetCreateConnectivityValidationDetailsModelTypeEnumStringValues() []string {
	return []string{
		"SOURCE_OPERATOR",
		"TARGET_OPERATOR",
	}
}
