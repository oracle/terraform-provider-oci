// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TlsCertificate Resource representing the location of the TLS certificate.
type TlsCertificate interface {
}

type tlscertificate struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *tlscertificate) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertlscertificate tlscertificate
	s := struct {
		Model Unmarshalertlscertificate
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tlscertificate) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "OCI_CERTIFICATES":
		mm := OciTlsCertificate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOCAL_FILE":
		mm := LocalFileTlsCertificate{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m tlscertificate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tlscertificate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TlsCertificateTypeEnum Enum with underlying type: string
type TlsCertificateTypeEnum string

// Set of constants representing the allowable values for TlsCertificateTypeEnum
const (
	TlsCertificateTypeOciCertificates TlsCertificateTypeEnum = "OCI_CERTIFICATES"
	TlsCertificateTypeLocalFile       TlsCertificateTypeEnum = "LOCAL_FILE"
)

var mappingTlsCertificateTypeEnum = map[string]TlsCertificateTypeEnum{
	"OCI_CERTIFICATES": TlsCertificateTypeOciCertificates,
	"LOCAL_FILE":       TlsCertificateTypeLocalFile,
}

var mappingTlsCertificateTypeEnumLowerCase = map[string]TlsCertificateTypeEnum{
	"oci_certificates": TlsCertificateTypeOciCertificates,
	"local_file":       TlsCertificateTypeLocalFile,
}

// GetTlsCertificateTypeEnumValues Enumerates the set of values for TlsCertificateTypeEnum
func GetTlsCertificateTypeEnumValues() []TlsCertificateTypeEnum {
	values := make([]TlsCertificateTypeEnum, 0)
	for _, v := range mappingTlsCertificateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTlsCertificateTypeEnumStringValues Enumerates the set of values in String for TlsCertificateTypeEnum
func GetTlsCertificateTypeEnumStringValues() []string {
	return []string{
		"OCI_CERTIFICATES",
		"LOCAL_FILE",
	}
}

// GetMappingTlsCertificateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTlsCertificateTypeEnum(val string) (TlsCertificateTypeEnum, bool) {
	enum, ok := mappingTlsCertificateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
