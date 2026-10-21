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

// SparkOciObjectStorageConfig Configuration for an OCI Object Storage Spark artifact location.
type SparkOciObjectStorageConfig interface {
}

type sparkociobjectstorageconfig struct {
	JsonData           []byte
	AuthenticationType string `json:"authenticationType"`
}

// UnmarshalJSON unmarshals json
func (m *sparkociobjectstorageconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersparkociobjectstorageconfig sparkociobjectstorageconfig
	s := struct {
		Model Unmarshalersparkociobjectstorageconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AuthenticationType = s.Model.AuthenticationType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *sparkociobjectstorageconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AuthenticationType {
	case "RESOURCE_PRINCIPAL":
		mm := SparkResourcePrincipalOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PRE_AUTHENTICATED_REQUEST":
		mm := SparkParOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "USER_PRINCIPAL":
		mm := SparkUserPrincipalOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOMER_SECRET_KEY":
		mm := SparkSecretKeyOciObjectStorageConfig{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SparkOciObjectStorageConfig: %s.", m.AuthenticationType)
		return *m, nil
	}
}

func (m sparkociobjectstorageconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m sparkociobjectstorageconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SparkOciObjectStorageConfigAuthenticationTypeEnum Enum with underlying type: string
type SparkOciObjectStorageConfigAuthenticationTypeEnum string

// Set of constants representing the allowable values for SparkOciObjectStorageConfigAuthenticationTypeEnum
const (
	SparkOciObjectStorageConfigAuthenticationTypeUserPrincipal           SparkOciObjectStorageConfigAuthenticationTypeEnum = "USER_PRINCIPAL"
	SparkOciObjectStorageConfigAuthenticationTypeResourcePrincipal       SparkOciObjectStorageConfigAuthenticationTypeEnum = "RESOURCE_PRINCIPAL"
	SparkOciObjectStorageConfigAuthenticationTypeCustomerSecretKey       SparkOciObjectStorageConfigAuthenticationTypeEnum = "CUSTOMER_SECRET_KEY"
	SparkOciObjectStorageConfigAuthenticationTypePreAuthenticatedRequest SparkOciObjectStorageConfigAuthenticationTypeEnum = "PRE_AUTHENTICATED_REQUEST"
)

var mappingSparkOciObjectStorageConfigAuthenticationTypeEnum = map[string]SparkOciObjectStorageConfigAuthenticationTypeEnum{
	"USER_PRINCIPAL":            SparkOciObjectStorageConfigAuthenticationTypeUserPrincipal,
	"RESOURCE_PRINCIPAL":        SparkOciObjectStorageConfigAuthenticationTypeResourcePrincipal,
	"CUSTOMER_SECRET_KEY":       SparkOciObjectStorageConfigAuthenticationTypeCustomerSecretKey,
	"PRE_AUTHENTICATED_REQUEST": SparkOciObjectStorageConfigAuthenticationTypePreAuthenticatedRequest,
}

var mappingSparkOciObjectStorageConfigAuthenticationTypeEnumLowerCase = map[string]SparkOciObjectStorageConfigAuthenticationTypeEnum{
	"user_principal":            SparkOciObjectStorageConfigAuthenticationTypeUserPrincipal,
	"resource_principal":        SparkOciObjectStorageConfigAuthenticationTypeResourcePrincipal,
	"customer_secret_key":       SparkOciObjectStorageConfigAuthenticationTypeCustomerSecretKey,
	"pre_authenticated_request": SparkOciObjectStorageConfigAuthenticationTypePreAuthenticatedRequest,
}

// GetSparkOciObjectStorageConfigAuthenticationTypeEnumValues Enumerates the set of values for SparkOciObjectStorageConfigAuthenticationTypeEnum
func GetSparkOciObjectStorageConfigAuthenticationTypeEnumValues() []SparkOciObjectStorageConfigAuthenticationTypeEnum {
	values := make([]SparkOciObjectStorageConfigAuthenticationTypeEnum, 0)
	for _, v := range mappingSparkOciObjectStorageConfigAuthenticationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSparkOciObjectStorageConfigAuthenticationTypeEnumStringValues Enumerates the set of values in String for SparkOciObjectStorageConfigAuthenticationTypeEnum
func GetSparkOciObjectStorageConfigAuthenticationTypeEnumStringValues() []string {
	return []string{
		"USER_PRINCIPAL",
		"RESOURCE_PRINCIPAL",
		"CUSTOMER_SECRET_KEY",
		"PRE_AUTHENTICATED_REQUEST",
	}
}

// GetMappingSparkOciObjectStorageConfigAuthenticationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSparkOciObjectStorageConfigAuthenticationTypeEnum(val string) (SparkOciObjectStorageConfigAuthenticationTypeEnum, bool) {
	enum, ok := mappingSparkOciObjectStorageConfigAuthenticationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
