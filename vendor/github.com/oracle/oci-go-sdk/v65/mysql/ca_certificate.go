// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CaCertificate The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
type CaCertificate interface {
}

type cacertificate struct {
	JsonData        []byte
	CertificateType string `json:"certificateType"`
}

// UnmarshalJSON unmarshals json
func (m *cacertificate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercacertificate cacertificate
	s := struct {
		Model Unmarshalercacertificate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CertificateType = s.Model.CertificateType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cacertificate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CertificateType {
	case "PEM":
		mm := PemCaCertificate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m cacertificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cacertificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CaCertificateCertificateTypeEnum Enum with underlying type: string
type CaCertificateCertificateTypeEnum string

// Set of constants representing the allowable values for CaCertificateCertificateTypeEnum
const (
	CaCertificateCertificateTypePem CaCertificateCertificateTypeEnum = "PEM"
)

var mappingCaCertificateCertificateTypeEnum = map[string]CaCertificateCertificateTypeEnum{
	"PEM": CaCertificateCertificateTypePem,
}

var mappingCaCertificateCertificateTypeEnumLowerCase = map[string]CaCertificateCertificateTypeEnum{
	"pem": CaCertificateCertificateTypePem,
}

// GetCaCertificateCertificateTypeEnumValues Enumerates the set of values for CaCertificateCertificateTypeEnum
func GetCaCertificateCertificateTypeEnumValues() []CaCertificateCertificateTypeEnum {
	values := make([]CaCertificateCertificateTypeEnum, 0)
	for _, v := range mappingCaCertificateCertificateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCaCertificateCertificateTypeEnumStringValues Enumerates the set of values in String for CaCertificateCertificateTypeEnum
func GetCaCertificateCertificateTypeEnumStringValues() []string {
	return []string{
		"PEM",
	}
}

// GetMappingCaCertificateCertificateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCaCertificateCertificateTypeEnum(val string) (CaCertificateCertificateTypeEnum, bool) {
	enum, ok := mappingCaCertificateCertificateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
