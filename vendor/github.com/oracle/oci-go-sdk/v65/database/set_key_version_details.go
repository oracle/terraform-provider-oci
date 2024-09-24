// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SetKeyVersionDetails Updating the database key version
type SetKeyVersionDetails interface {
}

type setkeyversiondetails struct {
	JsonData []byte
	Provider string `json:"provider"`
}

// UnmarshalJSON unmarshals json
func (m *setkeyversiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersetkeyversiondetails setkeyversiondetails
	s := struct {
		Model Unmarshalersetkeyversiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Provider = s.Model.Provider

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *setkeyversiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Provider {
	case "OCI":
		mm := OciProviderSetKeyVersionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for SetKeyVersionDetails: %s.", m.Provider)
		return *m, nil
	}
}

func (m setkeyversiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m setkeyversiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SetKeyVersionDetailsProviderEnum Enum with underlying type: string
type SetKeyVersionDetailsProviderEnum string

// Set of constants representing the allowable values for SetKeyVersionDetailsProviderEnum
const (
	SetKeyVersionDetailsProviderOci SetKeyVersionDetailsProviderEnum = "OCI"
)

var mappingSetKeyVersionDetailsProviderEnum = map[string]SetKeyVersionDetailsProviderEnum{
	"OCI": SetKeyVersionDetailsProviderOci,
}

var mappingSetKeyVersionDetailsProviderEnumLowerCase = map[string]SetKeyVersionDetailsProviderEnum{
	"oci": SetKeyVersionDetailsProviderOci,
}

// GetSetKeyVersionDetailsProviderEnumValues Enumerates the set of values for SetKeyVersionDetailsProviderEnum
func GetSetKeyVersionDetailsProviderEnumValues() []SetKeyVersionDetailsProviderEnum {
	values := make([]SetKeyVersionDetailsProviderEnum, 0)
	for _, v := range mappingSetKeyVersionDetailsProviderEnum {
		values = append(values, v)
	}
	return values
}

// GetSetKeyVersionDetailsProviderEnumStringValues Enumerates the set of values in String for SetKeyVersionDetailsProviderEnum
func GetSetKeyVersionDetailsProviderEnumStringValues() []string {
	return []string{
		"OCI",
	}
}

// GetMappingSetKeyVersionDetailsProviderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSetKeyVersionDetailsProviderEnum(val string) (SetKeyVersionDetailsProviderEnum, bool) {
	enum, ok := mappingSetKeyVersionDetailsProviderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
