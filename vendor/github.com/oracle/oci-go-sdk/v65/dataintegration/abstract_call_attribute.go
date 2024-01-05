// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AbstractCallAttribute The abstract write attribute.
type AbstractCallAttribute interface {

	// The fetch size for reading.
	GetFetchSize() *int
}

type abstractcallattribute struct {
	JsonData  []byte
	FetchSize *int   `mandatory:"false" json:"fetchSize"`
	ModelType string `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *abstractcallattribute) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerabstractcallattribute abstractcallattribute
	s := struct {
		Model Unmarshalerabstractcallattribute
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.FetchSize = s.Model.FetchSize
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *abstractcallattribute) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "BIP_CALL_ATTRIBUTE":
		mm := BipCallAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GENERIC_REST_CALL_ATTRIBUTE":
		mm := GenericRestCallAttribute{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AbstractCallAttribute: %s.", m.ModelType)
		return *m, nil
	}
}

// GetFetchSize returns FetchSize
func (m abstractcallattribute) GetFetchSize() *int {
	return m.FetchSize
}

func (m abstractcallattribute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m abstractcallattribute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AbstractCallAttributeModelTypeEnum Enum with underlying type: string
type AbstractCallAttributeModelTypeEnum string

// Set of constants representing the allowable values for AbstractCallAttributeModelTypeEnum
const (
	AbstractCallAttributeModelTypeBipCallAttribute         AbstractCallAttributeModelTypeEnum = "BIP_CALL_ATTRIBUTE"
	AbstractCallAttributeModelTypeGenericRestCallAttribute AbstractCallAttributeModelTypeEnum = "GENERIC_REST_CALL_ATTRIBUTE"
)

var mappingAbstractCallAttributeModelTypeEnum = map[string]AbstractCallAttributeModelTypeEnum{
	"BIP_CALL_ATTRIBUTE":          AbstractCallAttributeModelTypeBipCallAttribute,
	"GENERIC_REST_CALL_ATTRIBUTE": AbstractCallAttributeModelTypeGenericRestCallAttribute,
}

var mappingAbstractCallAttributeModelTypeEnumLowerCase = map[string]AbstractCallAttributeModelTypeEnum{
	"bip_call_attribute":          AbstractCallAttributeModelTypeBipCallAttribute,
	"generic_rest_call_attribute": AbstractCallAttributeModelTypeGenericRestCallAttribute,
}

// GetAbstractCallAttributeModelTypeEnumValues Enumerates the set of values for AbstractCallAttributeModelTypeEnum
func GetAbstractCallAttributeModelTypeEnumValues() []AbstractCallAttributeModelTypeEnum {
	values := make([]AbstractCallAttributeModelTypeEnum, 0)
	for _, v := range mappingAbstractCallAttributeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAbstractCallAttributeModelTypeEnumStringValues Enumerates the set of values in String for AbstractCallAttributeModelTypeEnum
func GetAbstractCallAttributeModelTypeEnumStringValues() []string {
	return []string{
		"BIP_CALL_ATTRIBUTE",
		"GENERIC_REST_CALL_ATTRIBUTE",
	}
}

// GetMappingAbstractCallAttributeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAbstractCallAttributeModelTypeEnum(val string) (AbstractCallAttributeModelTypeEnum, bool) {
	enum, ok := mappingAbstractCallAttributeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
