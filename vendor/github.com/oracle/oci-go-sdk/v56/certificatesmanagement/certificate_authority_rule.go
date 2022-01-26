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
	"github.com/oracle/oci-go-sdk/v56/common"
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
		return *m, nil
	}
}

func (m certificateauthorityrule) String() string {
	return common.PointerString(m)
}

// CertificateAuthorityRuleRuleTypeEnum Enum with underlying type: string
type CertificateAuthorityRuleRuleTypeEnum string

// Set of constants representing the allowable values for CertificateAuthorityRuleRuleTypeEnum
const (
	CertificateAuthorityRuleRuleTypeCertificateAuthorityIssuanceExpiryRule CertificateAuthorityRuleRuleTypeEnum = "CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE"
)

var mappingCertificateAuthorityRuleRuleType = map[string]CertificateAuthorityRuleRuleTypeEnum{
	"CERTIFICATE_AUTHORITY_ISSUANCE_EXPIRY_RULE": CertificateAuthorityRuleRuleTypeCertificateAuthorityIssuanceExpiryRule,
}

// GetCertificateAuthorityRuleRuleTypeEnumValues Enumerates the set of values for CertificateAuthorityRuleRuleTypeEnum
func GetCertificateAuthorityRuleRuleTypeEnumValues() []CertificateAuthorityRuleRuleTypeEnum {
	values := make([]CertificateAuthorityRuleRuleTypeEnum, 0)
	for _, v := range mappingCertificateAuthorityRuleRuleType {
		values = append(values, v)
	}
	return values
}
