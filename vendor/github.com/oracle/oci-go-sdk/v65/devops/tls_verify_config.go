// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TlsVerifyConfig TLS configuration used by build service to verify TLS connection.
type TlsVerifyConfig interface {
}

type tlsverifyconfig struct {
	JsonData      []byte
	TlsVerifyMode string `json:"tlsVerifyMode"`
}

// UnmarshalJSON unmarshals json
func (m *tlsverifyconfig) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertlsverifyconfig tlsverifyconfig
	s := struct {
		Model Unmarshalertlsverifyconfig
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TlsVerifyMode = s.Model.TlsVerifyMode

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tlsverifyconfig) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TlsVerifyMode {
	case "CA_CERTIFICATE_VERIFY":
		mm := CaCertVerify{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TlsVerifyConfig: %s.", m.TlsVerifyMode)
		return *m, nil
	}
}

func (m tlsverifyconfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tlsverifyconfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TlsVerifyConfigTlsVerifyModeEnum Enum with underlying type: string
type TlsVerifyConfigTlsVerifyModeEnum string

// Set of constants representing the allowable values for TlsVerifyConfigTlsVerifyModeEnum
const (
	TlsVerifyConfigTlsVerifyModeCaCertificateVerify TlsVerifyConfigTlsVerifyModeEnum = "CA_CERTIFICATE_VERIFY"
)

var mappingTlsVerifyConfigTlsVerifyModeEnum = map[string]TlsVerifyConfigTlsVerifyModeEnum{
	"CA_CERTIFICATE_VERIFY": TlsVerifyConfigTlsVerifyModeCaCertificateVerify,
}

var mappingTlsVerifyConfigTlsVerifyModeEnumLowerCase = map[string]TlsVerifyConfigTlsVerifyModeEnum{
	"ca_certificate_verify": TlsVerifyConfigTlsVerifyModeCaCertificateVerify,
}

// GetTlsVerifyConfigTlsVerifyModeEnumValues Enumerates the set of values for TlsVerifyConfigTlsVerifyModeEnum
func GetTlsVerifyConfigTlsVerifyModeEnumValues() []TlsVerifyConfigTlsVerifyModeEnum {
	values := make([]TlsVerifyConfigTlsVerifyModeEnum, 0)
	for _, v := range mappingTlsVerifyConfigTlsVerifyModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTlsVerifyConfigTlsVerifyModeEnumStringValues Enumerates the set of values in String for TlsVerifyConfigTlsVerifyModeEnum
func GetTlsVerifyConfigTlsVerifyModeEnumStringValues() []string {
	return []string{
		"CA_CERTIFICATE_VERIFY",
	}
}

// GetMappingTlsVerifyConfigTlsVerifyModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTlsVerifyConfigTlsVerifyModeEnum(val string) (TlsVerifyConfigTlsVerifyModeEnum, bool) {
	enum, ok := mappingTlsVerifyConfigTlsVerifyModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
