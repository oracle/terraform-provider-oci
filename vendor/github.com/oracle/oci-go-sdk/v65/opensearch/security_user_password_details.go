// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityUserPasswordDetails Base password-source details for the OpenSearch security master user.
type SecurityUserPasswordDetails interface {
}

type securityuserpassworddetails struct {
	JsonData   []byte
	SourceType string `json:"sourceType"`
}

// UnmarshalJSON unmarshals json
func (m *securityuserpassworddetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecurityuserpassworddetails securityuserpassworddetails
	s := struct {
		Model Unmarshalersecurityuserpassworddetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.SourceType = s.Model.SourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *securityuserpassworddetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SourceType {
	case "OCI_VAULT_SECRET":
		mm := VaultSecretSecurityUserPasswordDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PLAIN_PASSWORD":
		mm := PlainPasswordSecurityUserPasswordDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SecurityUserPasswordDetails: %s.", m.SourceType)
		return *m, nil
	}
}

func (m securityuserpassworddetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m securityuserpassworddetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityUserPasswordDetailsSourceTypeEnum Enum with underlying type: string
type SecurityUserPasswordDetailsSourceTypeEnum string

// Set of constants representing the allowable values for SecurityUserPasswordDetailsSourceTypeEnum
const (
	SecurityUserPasswordDetailsSourceTypePlainPassword  SecurityUserPasswordDetailsSourceTypeEnum = "PLAIN_PASSWORD"
	SecurityUserPasswordDetailsSourceTypeOciVaultSecret SecurityUserPasswordDetailsSourceTypeEnum = "OCI_VAULT_SECRET"
)

var mappingSecurityUserPasswordDetailsSourceTypeEnum = map[string]SecurityUserPasswordDetailsSourceTypeEnum{
	"PLAIN_PASSWORD":   SecurityUserPasswordDetailsSourceTypePlainPassword,
	"OCI_VAULT_SECRET": SecurityUserPasswordDetailsSourceTypeOciVaultSecret,
}

var mappingSecurityUserPasswordDetailsSourceTypeEnumLowerCase = map[string]SecurityUserPasswordDetailsSourceTypeEnum{
	"plain_password":   SecurityUserPasswordDetailsSourceTypePlainPassword,
	"oci_vault_secret": SecurityUserPasswordDetailsSourceTypeOciVaultSecret,
}

// GetSecurityUserPasswordDetailsSourceTypeEnumValues Enumerates the set of values for SecurityUserPasswordDetailsSourceTypeEnum
func GetSecurityUserPasswordDetailsSourceTypeEnumValues() []SecurityUserPasswordDetailsSourceTypeEnum {
	values := make([]SecurityUserPasswordDetailsSourceTypeEnum, 0)
	for _, v := range mappingSecurityUserPasswordDetailsSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityUserPasswordDetailsSourceTypeEnumStringValues Enumerates the set of values in String for SecurityUserPasswordDetailsSourceTypeEnum
func GetSecurityUserPasswordDetailsSourceTypeEnumStringValues() []string {
	return []string{
		"PLAIN_PASSWORD",
		"OCI_VAULT_SECRET",
	}
}

// GetMappingSecurityUserPasswordDetailsSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityUserPasswordDetailsSourceTypeEnum(val string) (SecurityUserPasswordDetailsSourceTypeEnum, bool) {
	enum, ok := mappingSecurityUserPasswordDetailsSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
