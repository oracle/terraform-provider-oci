// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BaseResourceConfiguration Configuration for a given 'resource'
type BaseResourceConfiguration interface {
}

type baseresourceconfiguration struct {
	JsonData []byte
	Resource string `json:"resource"`
}

// UnmarshalJSON unmarshals json
func (m *baseresourceconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbaseresourceconfiguration baseresourceconfiguration
	s := struct {
		Model Unmarshalerbaseresourceconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Resource = s.Model.Resource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *baseresourceconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Resource {
	case "EXADATA":
		mm := ExadataResourceConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NETWORK":
		mm := NetworkResourceConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "STORAGE":
		mm := StorageResourceConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE":
		mm := ComputeResourceConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BaseResourceConfiguration: %s.", m.Resource)
		return *m, nil
	}
}

func (m baseresourceconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m baseresourceconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseResourceConfigurationResourceEnum Enum with underlying type: string
type BaseResourceConfigurationResourceEnum string

// Set of constants representing the allowable values for BaseResourceConfigurationResourceEnum
const (
	BaseResourceConfigurationResourceCompute BaseResourceConfigurationResourceEnum = "COMPUTE"
	BaseResourceConfigurationResourceExadata BaseResourceConfigurationResourceEnum = "EXADATA"
	BaseResourceConfigurationResourceStorage BaseResourceConfigurationResourceEnum = "STORAGE"
	BaseResourceConfigurationResourceNetwork BaseResourceConfigurationResourceEnum = "NETWORK"
)

var mappingBaseResourceConfigurationResourceEnum = map[string]BaseResourceConfigurationResourceEnum{
	"COMPUTE": BaseResourceConfigurationResourceCompute,
	"EXADATA": BaseResourceConfigurationResourceExadata,
	"STORAGE": BaseResourceConfigurationResourceStorage,
	"NETWORK": BaseResourceConfigurationResourceNetwork,
}

var mappingBaseResourceConfigurationResourceEnumLowerCase = map[string]BaseResourceConfigurationResourceEnum{
	"compute": BaseResourceConfigurationResourceCompute,
	"exadata": BaseResourceConfigurationResourceExadata,
	"storage": BaseResourceConfigurationResourceStorage,
	"network": BaseResourceConfigurationResourceNetwork,
}

// GetBaseResourceConfigurationResourceEnumValues Enumerates the set of values for BaseResourceConfigurationResourceEnum
func GetBaseResourceConfigurationResourceEnumValues() []BaseResourceConfigurationResourceEnum {
	values := make([]BaseResourceConfigurationResourceEnum, 0)
	for _, v := range mappingBaseResourceConfigurationResourceEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseResourceConfigurationResourceEnumStringValues Enumerates the set of values in String for BaseResourceConfigurationResourceEnum
func GetBaseResourceConfigurationResourceEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"EXADATA",
		"STORAGE",
		"NETWORK",
	}
}

// GetMappingBaseResourceConfigurationResourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseResourceConfigurationResourceEnum(val string) (BaseResourceConfigurationResourceEnum, bool) {
	enum, ok := mappingBaseResourceConfigurationResourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
