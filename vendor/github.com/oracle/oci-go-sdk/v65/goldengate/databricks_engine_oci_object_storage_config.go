// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabricksEngineOciObjectStorageConfig Configuration for an OCI Object Storage Databricks Engine artifact location.
type DatabricksEngineOciObjectStorageConfig interface {
}

type databricksengineociobjectstorageconfig struct {
	JsonData           []byte
	AuthenticationType string `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *databricksengineociobjectstorageconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabricksengineociobjectstorageconfig databricksengineociobjectstorageconfig
	s := struct {
		Model Unmarshalerdatabricksengineociobjectstorageconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databricksengineociobjectstorageconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "CUSTOMER_SECRET_KEY":
		mm := DatabricksEngineSecretKeyOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_AUTHENTICATED_REQUEST":
		mm := DatabricksEngineParOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER_PRINCIPAL":
		mm := DatabricksEngineUserPrincipalOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "RESOURCE_PRINCIPAL":
		mm := DatabricksEngineResourcePrincipalOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DatabricksEngineOciObjectStorageConfig: %s.", m.AuthenticationType)
		return *m, nil
	}
}

func (m databricksengineociobjectstorageconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databricksengineociobjectstorageconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum Enum with underlying type: string
type DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum string

// Set of constants representing the allowable values for DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum
const (
	DatabricksEngineOciObjectStorageConfigAuthenticationTypeUserPrincipal           DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum = "USER_PRINCIPAL"
	DatabricksEngineOciObjectStorageConfigAuthenticationTypeResourcePrincipal       DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum = "RESOURCE_PRINCIPAL"
	DatabricksEngineOciObjectStorageConfigAuthenticationTypeCustomerSecretKey       DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum = "CUSTOMER_SECRET_KEY"
	DatabricksEngineOciObjectStorageConfigAuthenticationTypePreAuthenticatedRequest DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum = "PRE_AUTHENTICATED_REQUEST"
)

var mappingDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum = map[string]DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum{
	"USER_PRINCIPAL":            DatabricksEngineOciObjectStorageConfigAuthenticationTypeUserPrincipal,
	"RESOURCE_PRINCIPAL":        DatabricksEngineOciObjectStorageConfigAuthenticationTypeResourcePrincipal,
	"CUSTOMER_SECRET_KEY":       DatabricksEngineOciObjectStorageConfigAuthenticationTypeCustomerSecretKey,
	"PRE_AUTHENTICATED_REQUEST": DatabricksEngineOciObjectStorageConfigAuthenticationTypePreAuthenticatedRequest,
}

var mappingDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnumLowerCase = map[string]DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum{
	"user_principal":            DatabricksEngineOciObjectStorageConfigAuthenticationTypeUserPrincipal,
	"resource_principal":        DatabricksEngineOciObjectStorageConfigAuthenticationTypeResourcePrincipal,
	"customer_secret_key":       DatabricksEngineOciObjectStorageConfigAuthenticationTypeCustomerSecretKey,
	"pre_authenticated_request": DatabricksEngineOciObjectStorageConfigAuthenticationTypePreAuthenticatedRequest,
}

// GetDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnumValues Enumerates the set of values for DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum
func GetDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnumValues() []DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum {
	values := make([]DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum, 0)
	for _, v := range mappingDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnumStringValues Enumerates the set of values in String for DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum
func GetDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnumStringValues() []string {
	return []string{
		"USER_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
		"CUSTOMER_SECRET_KEY",
		"PRE_AUTHENTICATED_REQUEST",
	}
}

// GetMappingDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum(val string) (DatabricksEngineOciObjectStorageConfigAuthenticationTypeEnum, bool) {
	enum, ok := mappingDatabricksEngineOciObjectStorageConfigAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
