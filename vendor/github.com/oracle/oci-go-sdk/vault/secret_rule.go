// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// SecretRule A rule that you can apply to a secret to enforce certain conditions on the secret's usage and management.
type SecretRule interface {
}

type secretrule struct {
	JsonData []byte
	RuleType string `json:"ruleType"`
}

// UnmarshalJSON unmarshals json
func (m *secretrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersecretrule secretrule
	s := struct {
		Model Unmarshalersecretrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.RuleType = s.Model.RuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *secretrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RuleType {
	case "SECRET_EXPIRY_RULE":
		mm := SecretExpiryRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SECRET_REUSE_RULE":
		mm := SecretReuseRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m secretrule) String() string {
	return common.PointerString(m)
}

// SecretRuleRuleTypeEnum Enum with underlying type: string
type SecretRuleRuleTypeEnum string

// Set of constants representing the allowable values for SecretRuleRuleTypeEnum
const (
	SecretRuleRuleTypeExpiryRule SecretRuleRuleTypeEnum = "SECRET_EXPIRY_RULE"
	SecretRuleRuleTypeReuseRule  SecretRuleRuleTypeEnum = "SECRET_REUSE_RULE"
)

var mappingSecretRuleRuleType = map[string]SecretRuleRuleTypeEnum{
	"SECRET_EXPIRY_RULE": SecretRuleRuleTypeExpiryRule,
	"SECRET_REUSE_RULE":  SecretRuleRuleTypeReuseRule,
}

// GetSecretRuleRuleTypeEnumValues Enumerates the set of values for SecretRuleRuleTypeEnum
func GetSecretRuleRuleTypeEnumValues() []SecretRuleRuleTypeEnum {
	values := make([]SecretRuleRuleTypeEnum, 0)
	for _, v := range mappingSecretRuleRuleType {
		values = append(values, v)
	}
	return values
}
