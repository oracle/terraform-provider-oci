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

// OciFssSecurity Security options for OCI FSS File System.
type OciFssSecurity interface {

	// Determines whether in-transit encryption needs to be enables.
	// Check https://docs.oracle.com/en-us/iaas/Content/File/Tasks/intransitencryption.htm#Using_Intransit_Encryption for more details.
	GetIsEncryptedInTransit() *bool
}

type ocifsssecurity struct {
	JsonData             []byte
	IsEncryptedInTransit *bool  `mandatory:"false" json:"isEncryptedInTransit"`
	Auth                 string `json:"auth"`
}

// UnmarshalJSON unmarshals json
func (m *ocifsssecurity) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerocifsssecurity ocifsssecurity
	s := struct {
		Model Unmarshalerocifsssecurity
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
func (m *ocifsssecurity) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Auth {
	case "SYS":
		mm := OciFssSysSecurity{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for OciFssSecurity: %s.", m.Auth)
		return *m, nil
	}
}

// GetIsEncryptedInTransit returns IsEncryptedInTransit
func (m ocifsssecurity) GetIsEncryptedInTransit() *bool {
	return m.IsEncryptedInTransit
}

func (m ocifsssecurity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ocifsssecurity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciFssSecurityAuthEnum Enum with underlying type: string
type OciFssSecurityAuthEnum string

// Set of constants representing the allowable values for OciFssSecurityAuthEnum
const (
	OciFssSecurityAuthSys OciFssSecurityAuthEnum = "SYS"
)

var mappingOciFssSecurityAuthEnum = map[string]OciFssSecurityAuthEnum{
	"SYS": OciFssSecurityAuthSys,
}

var mappingOciFssSecurityAuthEnumLowerCase = map[string]OciFssSecurityAuthEnum{
	"sys": OciFssSecurityAuthSys,
}

// GetOciFssSecurityAuthEnumValues Enumerates the set of values for OciFssSecurityAuthEnum
func GetOciFssSecurityAuthEnumValues() []OciFssSecurityAuthEnum {
	values := make([]OciFssSecurityAuthEnum, 0)
	for _, v := range mappingOciFssSecurityAuthEnum {
		values = append(values, v)
	}
	return values
}

// GetOciFssSecurityAuthEnumStringValues Enumerates the set of values in String for OciFssSecurityAuthEnum
func GetOciFssSecurityAuthEnumStringValues() []string {
	return []string{
		"SYS",
	}
}

// GetMappingOciFssSecurityAuthEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciFssSecurityAuthEnum(val string) (OciFssSecurityAuthEnum, bool) {
	enum, ok := mappingOciFssSecurityAuthEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
