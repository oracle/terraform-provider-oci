// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CertificateRule A rule that you can apply to a certificate to enforce certain conditions on the certificate's usage and management.
type CertificateRule interface {
}

type certificaterule struct {
	JsonData []byte
	RuleType string `json:"ruleType"`
}

// UnmarshalJSON unmarshals json
func (m *certificaterule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercertificaterule certificaterule
	s := struct {
		Model Unmarshalercertificaterule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuleType = s.Model.RuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *certificaterule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuleType {
	case "CERTIFICATE_RENEWAL_RULE":
		mm := CertificateRenewalRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m certificaterule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m certificaterule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CertificateRuleRuleTypeEnum Enum with underlying type: string
type CertificateRuleRuleTypeEnum string

// Set of constants representing the allowable values for CertificateRuleRuleTypeEnum
const (
	CertificateRuleRuleTypeCertificateRenewalRule CertificateRuleRuleTypeEnum = "CERTIFICATE_RENEWAL_RULE"
)

var mappingCertificateRuleRuleTypeEnum = map[string]CertificateRuleRuleTypeEnum{
	"CERTIFICATE_RENEWAL_RULE": CertificateRuleRuleTypeCertificateRenewalRule,
}

// GetCertificateRuleRuleTypeEnumValues Enumerates the set of values for CertificateRuleRuleTypeEnum
func GetCertificateRuleRuleTypeEnumValues() []CertificateRuleRuleTypeEnum {
	values := make([]CertificateRuleRuleTypeEnum, 0)
	for _, v := range mappingCertificateRuleRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateRuleRuleTypeEnumStringValues Enumerates the set of values in String for CertificateRuleRuleTypeEnum
func GetCertificateRuleRuleTypeEnumStringValues() []string {
	return []string{
		"CERTIFICATE_RENEWAL_RULE",
	}
}

// GetMappingCertificateRuleRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateRuleRuleTypeEnum(val string) (CertificateRuleRuleTypeEnum, bool) {
	mappingCertificateRuleRuleTypeEnumIgnoreCase := make(map[string]CertificateRuleRuleTypeEnum)
	for k, v := range mappingCertificateRuleRuleTypeEnum {
		mappingCertificateRuleRuleTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCertificateRuleRuleTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
