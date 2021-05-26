// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v41/common"
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

// CaCertificateCertificateTypeEnum Enum with underlying type: string
type CaCertificateCertificateTypeEnum string

// Set of constants representing the allowable values for CaCertificateCertificateTypeEnum
const (
	CaCertificateCertificateTypePem CaCertificateCertificateTypeEnum = "PEM"
)

var mappingCaCertificateCertificateType = map[string]CaCertificateCertificateTypeEnum{
	"PEM": CaCertificateCertificateTypePem,
}

// GetCaCertificateCertificateTypeEnumValues Enumerates the set of values for CaCertificateCertificateTypeEnum
func GetCaCertificateCertificateTypeEnumValues() []CaCertificateCertificateTypeEnum {
	values := make([]CaCertificateCertificateTypeEnum, 0)
	for _, v := range mappingCaCertificateCertificateType {
		values = append(values, v)
	}
	return values
}
