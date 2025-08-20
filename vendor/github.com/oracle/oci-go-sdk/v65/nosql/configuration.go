// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Configuration A tenancy or service-level configuration. The
// service may of the standard `MULTI_TENANCY` type, or of the
// `HOSTED` environment type.
type Configuration interface {
}

type configuration struct {
	JsonData    []byte
	Environment string `json:"environment"`
}

// UnmarshalJSON unmarshals json
func (m *configuration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfiguration configuration
	s := struct {
		Model Unmarshalerconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Environment = s.Model.Environment

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configuration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Environment {
	case "MULTI_TENANCY":
		mm := MultiTenancyConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOSTED":
		mm := HostedConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Configuration: %s.", m.Environment)
		return *m, nil
	}
}

func (m configuration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configuration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigurationEnvironmentEnum Enum with underlying type: string
type ConfigurationEnvironmentEnum string

// Set of constants representing the allowable values for ConfigurationEnvironmentEnum
const (
	ConfigurationEnvironmentMultiTenancy ConfigurationEnvironmentEnum = "MULTI_TENANCY"
	ConfigurationEnvironmentHosted       ConfigurationEnvironmentEnum = "HOSTED"
)

var mappingConfigurationEnvironmentEnum = map[string]ConfigurationEnvironmentEnum{
	"MULTI_TENANCY": ConfigurationEnvironmentMultiTenancy,
	"HOSTED":        ConfigurationEnvironmentHosted,
}

var mappingConfigurationEnvironmentEnumLowerCase = map[string]ConfigurationEnvironmentEnum{
	"multi_tenancy": ConfigurationEnvironmentMultiTenancy,
	"hosted":        ConfigurationEnvironmentHosted,
}

// GetConfigurationEnvironmentEnumValues Enumerates the set of values for ConfigurationEnvironmentEnum
func GetConfigurationEnvironmentEnumValues() []ConfigurationEnvironmentEnum {
	values := make([]ConfigurationEnvironmentEnum, 0)
	for _, v := range mappingConfigurationEnvironmentEnum {
		values = append(values, v)
	}
	return values
}

// GetConfigurationEnvironmentEnumStringValues Enumerates the set of values in String for ConfigurationEnvironmentEnum
func GetConfigurationEnvironmentEnumStringValues() []string {
	return []string{
		"MULTI_TENANCY",
		"HOSTED",
	}
}

// GetMappingConfigurationEnvironmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConfigurationEnvironmentEnum(val string) (ConfigurationEnvironmentEnum, bool) {
	enum, ok := mappingConfigurationEnvironmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
