// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateOciFssSecurityDetails Security options for OCI FSS File System.
type CreateOciFssSecurityDetails interface {

	// Determines whether in-transit encryption needs to be enables.
	// Check https://docs.oracle.com/en-us/iaas/Content/File/Tasks/intransitencryption.htm#Using_Intransit_Encryption for more details.
	GetIsEncryptedInTransit() *bool
}

type createocifsssecuritydetails struct {
	JsonData             []byte
	IsEncryptedInTransit *bool  `mandatory:"false" json:"isEncryptedInTransit"`
	Auth                 string `json:"auth"`
}

// UnmarshalJSON unmarshals json
func (m *createocifsssecuritydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateocifsssecuritydetails createocifsssecuritydetails
	s := struct {
		Model Unmarshalercreateocifsssecuritydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsEncryptedInTransit = s.Model.IsEncryptedInTransit
	m.Auth = s.Model.Auth

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createocifsssecuritydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Auth {
	case "SYS":
		mm := CreateOciFssSysSecurityDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateOciFssSecurityDetails: %s.", m.Auth)
		return *m, nil
	}
}

// GetIsEncryptedInTransit returns IsEncryptedInTransit
func (m createocifsssecuritydetails) GetIsEncryptedInTransit() *bool {
	return m.IsEncryptedInTransit
}

func (m createocifsssecuritydetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createocifsssecuritydetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOciFssSecurityDetailsAuthEnum Enum with underlying type: string
type CreateOciFssSecurityDetailsAuthEnum string

// Set of constants representing the allowable values for CreateOciFssSecurityDetailsAuthEnum
const (
	CreateOciFssSecurityDetailsAuthSys CreateOciFssSecurityDetailsAuthEnum = "SYS"
)

var mappingCreateOciFssSecurityDetailsAuthEnum = map[string]CreateOciFssSecurityDetailsAuthEnum{
	"SYS": CreateOciFssSecurityDetailsAuthSys,
}

var mappingCreateOciFssSecurityDetailsAuthEnumLowerCase = map[string]CreateOciFssSecurityDetailsAuthEnum{
	"sys": CreateOciFssSecurityDetailsAuthSys,
}

// GetCreateOciFssSecurityDetailsAuthEnumValues Enumerates the set of values for CreateOciFssSecurityDetailsAuthEnum
func GetCreateOciFssSecurityDetailsAuthEnumValues() []CreateOciFssSecurityDetailsAuthEnum {
	values := make([]CreateOciFssSecurityDetailsAuthEnum, 0)
	for _, v := range mappingCreateOciFssSecurityDetailsAuthEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOciFssSecurityDetailsAuthEnumStringValues Enumerates the set of values in String for CreateOciFssSecurityDetailsAuthEnum
func GetCreateOciFssSecurityDetailsAuthEnumStringValues() []string {
	return []string{
		"SYS",
	}
}

// GetMappingCreateOciFssSecurityDetailsAuthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOciFssSecurityDetailsAuthEnum(val string) (CreateOciFssSecurityDetailsAuthEnum, bool) {
	enum, ok := mappingCreateOciFssSecurityDetailsAuthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
