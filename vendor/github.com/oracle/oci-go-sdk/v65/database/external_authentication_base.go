// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalAuthenticationBase The object holds external authentication details for Oracle Autonomous Databases.
type ExternalAuthenticationBase interface {
}

type externalauthenticationbase struct {
	JsonData []byte
	Method   string `json:"method"`
}

// UnmarshalJSON unmarshals json
func (m *externalauthenticationbase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexternalauthenticationbase externalauthenticationbase
	s := struct {
		Model Unmarshalerexternalauthenticationbase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Method = s.Model.Method

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *externalauthenticationbase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Method {
	case "OCI_IAM":
		mm := OciIamAuthenticationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ExternalAuthenticationBase: %s.", m.Method)
		return *m, nil
	}
}

func (m externalauthenticationbase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m externalauthenticationbase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalAuthenticationBaseMethodEnum Enum with underlying type: string
type ExternalAuthenticationBaseMethodEnum string

// Set of constants representing the allowable values for ExternalAuthenticationBaseMethodEnum
const (
	ExternalAuthenticationBaseMethodOciIam ExternalAuthenticationBaseMethodEnum = "OCI_IAM"
)

var mappingExternalAuthenticationBaseMethodEnum = map[string]ExternalAuthenticationBaseMethodEnum{
	"OCI_IAM": ExternalAuthenticationBaseMethodOciIam,
}

var mappingExternalAuthenticationBaseMethodEnumLowerCase = map[string]ExternalAuthenticationBaseMethodEnum{
	"oci_iam": ExternalAuthenticationBaseMethodOciIam,
}

// GetExternalAuthenticationBaseMethodEnumValues Enumerates the set of values for ExternalAuthenticationBaseMethodEnum
func GetExternalAuthenticationBaseMethodEnumValues() []ExternalAuthenticationBaseMethodEnum {
	values := make([]ExternalAuthenticationBaseMethodEnum, 0)
	for _, v := range mappingExternalAuthenticationBaseMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalAuthenticationBaseMethodEnumStringValues Enumerates the set of values in String for ExternalAuthenticationBaseMethodEnum
func GetExternalAuthenticationBaseMethodEnumStringValues() []string {
	return []string{
		"OCI_IAM",
	}
}

// GetMappingExternalAuthenticationBaseMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalAuthenticationBaseMethodEnum(val string) (ExternalAuthenticationBaseMethodEnum, bool) {
	enum, ok := mappingExternalAuthenticationBaseMethodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
