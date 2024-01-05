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

// AuthConfig Authentication configuration for Generic REST invocation.
type AuthConfig interface {

	// Generated key that can be used in API calls to identify this object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference
}

type authconfig struct {
	JsonData     []byte
	Key          *string          `mandatory:"false" json:"key"`
	ModelVersion *string          `mandatory:"false" json:"modelVersion"`
	ParentRef    *ParentReference `mandatory:"false" json:"parentRef"`
	ModelType    string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *authconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerauthconfig authconfig
	s := struct {
		Model Unmarshalerauthconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *authconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "OCI_RESOURCE_AUTH_CONFIG":
		mm := ResourcePrincipalAuthConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AuthConfig: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m authconfig) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m authconfig) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m authconfig) GetParentRef() *ParentReference {
	return m.ParentRef
}

func (m authconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m authconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AuthConfigModelTypeEnum Enum with underlying type: string
type AuthConfigModelTypeEnum string

// Set of constants representing the allowable values for AuthConfigModelTypeEnum
const (
	AuthConfigModelTypeOciResourceAuthConfig AuthConfigModelTypeEnum = "OCI_RESOURCE_AUTH_CONFIG"
)

var mappingAuthConfigModelTypeEnum = map[string]AuthConfigModelTypeEnum{
	"OCI_RESOURCE_AUTH_CONFIG": AuthConfigModelTypeOciResourceAuthConfig,
}

var mappingAuthConfigModelTypeEnumLowerCase = map[string]AuthConfigModelTypeEnum{
	"oci_resource_auth_config": AuthConfigModelTypeOciResourceAuthConfig,
}

// GetAuthConfigModelTypeEnumValues Enumerates the set of values for AuthConfigModelTypeEnum
func GetAuthConfigModelTypeEnumValues() []AuthConfigModelTypeEnum {
	values := make([]AuthConfigModelTypeEnum, 0)
	for _, v := range mappingAuthConfigModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAuthConfigModelTypeEnumStringValues Enumerates the set of values in String for AuthConfigModelTypeEnum
func GetAuthConfigModelTypeEnumStringValues() []string {
	return []string{
		"OCI_RESOURCE_AUTH_CONFIG",
	}
}

// GetMappingAuthConfigModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAuthConfigModelTypeEnum(val string) (AuthConfigModelTypeEnum, bool) {
	enum, ok := mappingAuthConfigModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
