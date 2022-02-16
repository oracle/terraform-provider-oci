// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CredentialDetails User credential details to connect to the database. This is supplied via the External Database Service.
type CredentialDetails interface {

	// Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
	GetCredentialSourceName() *string
}

type credentialdetails struct {
	JsonData             []byte
	CredentialSourceName *string `mandatory:"true" json:"credentialSourceName"`
	CredentialType       string  `json:"credentialType"`
}

// UnmarshalJSON unmarshals json
func (m *credentialdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercredentialdetails credentialdetails
	s := struct {
		Model Unmarshalercredentialdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CredentialSourceName = s.Model.CredentialSourceName
	m.CredentialType = s.Model.CredentialType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *credentialdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CredentialType {
	case "CREDENTIALS_BY_SOURCE":
		mm := CredentialsBySource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCredentialSourceName returns CredentialSourceName
func (m credentialdetails) GetCredentialSourceName() *string {
	return m.CredentialSourceName
}

func (m credentialdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m credentialdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CredentialDetailsCredentialTypeEnum Enum with underlying type: string
type CredentialDetailsCredentialTypeEnum string

// Set of constants representing the allowable values for CredentialDetailsCredentialTypeEnum
const (
	CredentialDetailsCredentialTypeCredentialsBySource CredentialDetailsCredentialTypeEnum = "CREDENTIALS_BY_SOURCE"
)

var mappingCredentialDetailsCredentialTypeEnum = map[string]CredentialDetailsCredentialTypeEnum{
	"CREDENTIALS_BY_SOURCE": CredentialDetailsCredentialTypeCredentialsBySource,
}

// GetCredentialDetailsCredentialTypeEnumValues Enumerates the set of values for CredentialDetailsCredentialTypeEnum
func GetCredentialDetailsCredentialTypeEnumValues() []CredentialDetailsCredentialTypeEnum {
	values := make([]CredentialDetailsCredentialTypeEnum, 0)
	for _, v := range mappingCredentialDetailsCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCredentialDetailsCredentialTypeEnumStringValues Enumerates the set of values in String for CredentialDetailsCredentialTypeEnum
func GetCredentialDetailsCredentialTypeEnumStringValues() []string {
	return []string{
		"CREDENTIALS_BY_SOURCE",
	}
}

// GetMappingCredentialDetailsCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCredentialDetailsCredentialTypeEnum(val string) (CredentialDetailsCredentialTypeEnum, bool) {
	mappingCredentialDetailsCredentialTypeEnumIgnoreCase := make(map[string]CredentialDetailsCredentialTypeEnum)
	for k, v := range mappingCredentialDetailsCredentialTypeEnum {
		mappingCredentialDetailsCredentialTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCredentialDetailsCredentialTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
