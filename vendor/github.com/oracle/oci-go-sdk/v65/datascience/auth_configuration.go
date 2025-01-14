// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuthConfiguration AuthN/Z configuration for online prediction
type AuthConfiguration interface {
}

type authconfiguration struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *authconfiguration) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthconfiguration authconfiguration
	s := struct {
		Model Unmarshalerauthconfiguration
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authconfiguration) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "IDCS":
		mm := IdcsAuthConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IDCS_CUSTOM_SERVICE":
		mm := IdcsCustomServiceAuthConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IAM":
		mm := IamAuthConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for AuthConfiguration: %s.", m.Type)
		return *m, nil
	}
}

func (m authconfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m authconfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthConfigurationTypeEnum Enum with underlying type: string
type AuthConfigurationTypeEnum string

// Set of constants representing the allowable values for AuthConfigurationTypeEnum
const (
	AuthConfigurationTypeIdcs              AuthConfigurationTypeEnum = "IDCS"
	AuthConfigurationTypeIam               AuthConfigurationTypeEnum = "IAM"
	AuthConfigurationTypeIdcsCustomService AuthConfigurationTypeEnum = "IDCS_CUSTOM_SERVICE"
)

var mappingAuthConfigurationTypeEnum = map[string]AuthConfigurationTypeEnum{
	"IDCS":                AuthConfigurationTypeIdcs,
	"IAM":                 AuthConfigurationTypeIam,
	"IDCS_CUSTOM_SERVICE": AuthConfigurationTypeIdcsCustomService,
}

var mappingAuthConfigurationTypeEnumLowerCase = map[string]AuthConfigurationTypeEnum{
	"idcs":                AuthConfigurationTypeIdcs,
	"iam":                 AuthConfigurationTypeIam,
	"idcs_custom_service": AuthConfigurationTypeIdcsCustomService,
}

// GetAuthConfigurationTypeEnumValues Enumerates the set of values for AuthConfigurationTypeEnum
func GetAuthConfigurationTypeEnumValues() []AuthConfigurationTypeEnum {
	values := make([]AuthConfigurationTypeEnum, 0)
	for _, v := range mappingAuthConfigurationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthConfigurationTypeEnumStringValues Enumerates the set of values in String for AuthConfigurationTypeEnum
func GetAuthConfigurationTypeEnumStringValues() []string {
	return []string{
		"IDCS",
		"IAM",
		"IDCS_CUSTOM_SERVICE",
	}
}

// GetMappingAuthConfigurationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthConfigurationTypeEnum(val string) (AuthConfigurationTypeEnum, bool) {
	enum, ok := mappingAuthConfigurationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
