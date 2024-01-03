// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateAuthConfigurationDetails AuthN/Z configuration for online prediction
type CreateAuthConfigurationDetails interface {
}

type createauthconfigurationdetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createauthconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateauthconfigurationdetails createauthconfigurationdetails
	s := struct {
		Model Unmarshalercreateauthconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createauthconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "IDCS":
		mm := CreateIdcsAuthConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IAM":
		mm := CreateIamAuthConfigurationCreateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateAuthConfigurationDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m createauthconfigurationdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createauthconfigurationdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAuthConfigurationDetailsTypeEnum Enum with underlying type: string
type CreateAuthConfigurationDetailsTypeEnum string

// Set of constants representing the allowable values for CreateAuthConfigurationDetailsTypeEnum
const (
	CreateAuthConfigurationDetailsTypeIdcs CreateAuthConfigurationDetailsTypeEnum = "IDCS"
	CreateAuthConfigurationDetailsTypeIam  CreateAuthConfigurationDetailsTypeEnum = "IAM"
)

var mappingCreateAuthConfigurationDetailsTypeEnum = map[string]CreateAuthConfigurationDetailsTypeEnum{
	"IDCS": CreateAuthConfigurationDetailsTypeIdcs,
	"IAM":  CreateAuthConfigurationDetailsTypeIam,
}

var mappingCreateAuthConfigurationDetailsTypeEnumLowerCase = map[string]CreateAuthConfigurationDetailsTypeEnum{
	"idcs": CreateAuthConfigurationDetailsTypeIdcs,
	"iam":  CreateAuthConfigurationDetailsTypeIam,
}

// GetCreateAuthConfigurationDetailsTypeEnumValues Enumerates the set of values for CreateAuthConfigurationDetailsTypeEnum
func GetCreateAuthConfigurationDetailsTypeEnumValues() []CreateAuthConfigurationDetailsTypeEnum {
	values := make([]CreateAuthConfigurationDetailsTypeEnum, 0)
	for _, v := range mappingCreateAuthConfigurationDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAuthConfigurationDetailsTypeEnumStringValues Enumerates the set of values in String for CreateAuthConfigurationDetailsTypeEnum
func GetCreateAuthConfigurationDetailsTypeEnumStringValues() []string {
	return []string{
		"IDCS",
		"IAM",
	}
}

// GetMappingCreateAuthConfigurationDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAuthConfigurationDetailsTypeEnum(val string) (CreateAuthConfigurationDetailsTypeEnum, bool) {
	enum, ok := mappingCreateAuthConfigurationDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
