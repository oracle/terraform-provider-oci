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

// UpdateConfigurationDetails A tenancy or service-level configuration. The discriminator value
// `UpdateConfigurationDetails.environment` must match the service's
// environment type.
type UpdateConfigurationDetails interface {
}

type updateconfigurationdetails struct {
	JsonData    []byte
	Environment string `json:"environment"`
}

// UnmarshalJSON unmarshals json
func (m *updateconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconfigurationdetails updateconfigurationdetails
	s := struct {
		Model Unmarshalerupdateconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Environment = s.Model.Environment

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Environment {
	case "MULTI_TENANCY":
		mm := UpdateMultiTenancyConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOSTED":
		mm := UpdateHostedConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateConfigurationDetails: %s.", m.Environment)
		return *m, nil
	}
}

func (m updateconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateConfigurationDetailsEnvironmentEnum Enum with underlying type: string
type UpdateConfigurationDetailsEnvironmentEnum string

// Set of constants representing the allowable values for UpdateConfigurationDetailsEnvironmentEnum
const (
	UpdateConfigurationDetailsEnvironmentMultiTenancy UpdateConfigurationDetailsEnvironmentEnum = "MULTI_TENANCY"
	UpdateConfigurationDetailsEnvironmentHosted       UpdateConfigurationDetailsEnvironmentEnum = "HOSTED"
)

var mappingUpdateConfigurationDetailsEnvironmentEnum = map[string]UpdateConfigurationDetailsEnvironmentEnum{
	"MULTI_TENANCY": UpdateConfigurationDetailsEnvironmentMultiTenancy,
	"HOSTED":        UpdateConfigurationDetailsEnvironmentHosted,
}

var mappingUpdateConfigurationDetailsEnvironmentEnumLowerCase = map[string]UpdateConfigurationDetailsEnvironmentEnum{
	"multi_tenancy": UpdateConfigurationDetailsEnvironmentMultiTenancy,
	"hosted":        UpdateConfigurationDetailsEnvironmentHosted,
}

// GetUpdateConfigurationDetailsEnvironmentEnumValues Enumerates the set of values for UpdateConfigurationDetailsEnvironmentEnum
func GetUpdateConfigurationDetailsEnvironmentEnumValues() []UpdateConfigurationDetailsEnvironmentEnum {
	values := make([]UpdateConfigurationDetailsEnvironmentEnum, 0)
	for _, v := range mappingUpdateConfigurationDetailsEnvironmentEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateConfigurationDetailsEnvironmentEnumStringValues Enumerates the set of values in String for UpdateConfigurationDetailsEnvironmentEnum
func GetUpdateConfigurationDetailsEnvironmentEnumStringValues() []string {
	return []string{
		"MULTI_TENANCY",
		"HOSTED",
	}
}

// GetMappingUpdateConfigurationDetailsEnvironmentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateConfigurationDetailsEnvironmentEnum(val string) (UpdateConfigurationDetailsEnvironmentEnum, bool) {
	enum, ok := mappingUpdateConfigurationDetailsEnvironmentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
