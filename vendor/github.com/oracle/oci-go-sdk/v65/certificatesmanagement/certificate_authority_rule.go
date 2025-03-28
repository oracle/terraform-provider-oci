// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CertificateAuthorityRule A rule that you can apply to a certificate authority (CA) to enforce certain conditions on its usage and management.
type CertificateAuthorityRule interface {
}

type certificateauthorityrule struct {
	JsonData []byte
	RuleType string `json:"ruleType"`
}

// UnmarshalJSON unmarshals json
func (m *certificateauthorityrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercertificateauthorityrule certificateauthorityrule
	s := struct {
		Model Unmarshalercertificateauthorityrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuleType = s.Model.RuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *certificateauthorityrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuleType {
	case "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE":
		mm := CertificateAuthorityIssuanceExpiryRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CertificateAuthorityRule: %s.", m.RuleType)
		return *m, nil
	}
}

func (m certificateauthorityrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m certificateauthorityrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CertificateAuthorityRuleRuleTypeEnum Enum with underlying type: string
type CertificateAuthorityRuleRuleTypeEnum string

// Set of constants representing the allowable values for CertificateAuthorityRuleRuleTypeEnum
const (
	CertificateAuthorityRuleRuleTypeCertificateAuthorityIssuanceExpiryRule CertificateAuthorityRuleRuleTypeEnum = "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"
)

var mappingCertificateAuthorityRuleRuleTypeEnum = map[string]CertificateAuthorityRuleRuleTypeEnum{
	"CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE": CertificateAuthorityRuleRuleTypeCertificateAuthorityIssuanceExpiryRule,
}

var mappingCertificateAuthorityRuleRuleTypeEnumLowerCase = map[string]CertificateAuthorityRuleRuleTypeEnum{
	"certificate_authority_issuance_expiry_rule": CertificateAuthorityRuleRuleTypeCertificateAuthorityIssuanceExpiryRule,
}

// GetCertificateAuthorityRuleRuleTypeEnumValues Enumerates the set of values for CertificateAuthorityRuleRuleTypeEnum
func GetCertificateAuthorityRuleRuleTypeEnumValues() []CertificateAuthorityRuleRuleTypeEnum {
	values := make([]CertificateAuthorityRuleRuleTypeEnum, 0)
	for _, v := range mappingCertificateAuthorityRuleRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCertificateAuthorityRuleRuleTypeEnumStringValues Enumerates the set of values in String for CertificateAuthorityRuleRuleTypeEnum
func GetCertificateAuthorityRuleRuleTypeEnumStringValues() []string {
	return []string{
		"CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE",
	}
}

// GetMappingCertificateAuthorityRuleRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCertificateAuthorityRuleRuleTypeEnum(val string) (CertificateAuthorityRuleRuleTypeEnum, bool) {
	enum, ok := mappingCertificateAuthorityRuleRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
