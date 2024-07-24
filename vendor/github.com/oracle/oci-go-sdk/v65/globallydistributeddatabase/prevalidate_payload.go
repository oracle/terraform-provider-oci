// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrevalidatePayload Payload required to run prevalidation operation for create sharded database or patch sharded database, based on operation selected.
type PrevalidatePayload interface {
}

type prevalidatepayload struct {
	JsonData  []byte
	Operation string `json:"operation"`
}

// UnmarshalJSON unmarshals json
func (m *prevalidatepayload) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerprevalidatepayload prevalidatepayload
	s := struct {
		Model Unmarshalerprevalidatepayload
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Operation = s.Model.Operation

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *prevalidatepayload) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Operation {
	case "PATCH":
		mm := PrevalidatePatchPayload{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CREATE":
		mm := PrevalidateCreatePayload{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PrevalidatePayload: %s.", m.Operation)
		return *m, nil
	}
}

func (m prevalidatepayload) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m prevalidatepayload) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrevalidatePayloadOperationEnum Enum with underlying type: string
type PrevalidatePayloadOperationEnum string

// Set of constants representing the allowable values for PrevalidatePayloadOperationEnum
const (
	PrevalidatePayloadOperationCreate PrevalidatePayloadOperationEnum = "CREATE"
	PrevalidatePayloadOperationPatch  PrevalidatePayloadOperationEnum = "PATCH"
)

var mappingPrevalidatePayloadOperationEnum = map[string]PrevalidatePayloadOperationEnum{
	"CREATE": PrevalidatePayloadOperationCreate,
	"PATCH":  PrevalidatePayloadOperationPatch,
}

var mappingPrevalidatePayloadOperationEnumLowerCase = map[string]PrevalidatePayloadOperationEnum{
	"create": PrevalidatePayloadOperationCreate,
	"patch":  PrevalidatePayloadOperationPatch,
}

// GetPrevalidatePayloadOperationEnumValues Enumerates the set of values for PrevalidatePayloadOperationEnum
func GetPrevalidatePayloadOperationEnumValues() []PrevalidatePayloadOperationEnum {
	values := make([]PrevalidatePayloadOperationEnum, 0)
	for _, v := range mappingPrevalidatePayloadOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetPrevalidatePayloadOperationEnumStringValues Enumerates the set of values in String for PrevalidatePayloadOperationEnum
func GetPrevalidatePayloadOperationEnumStringValues() []string {
	return []string{
		"CREATE",
		"PATCH",
	}
}

// GetMappingPrevalidatePayloadOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrevalidatePayloadOperationEnum(val string) (PrevalidatePayloadOperationEnum, bool) {
	enum, ok := mappingPrevalidatePayloadOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
