// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DecryptionRule Decryption Rule used in the firewall policy rules.
// A Decryption Rule is used to define which traffic should be decrypted by the firewall, and how it should do so.
type DecryptionRule struct {

	// Name for the decryption rule, must be unique within the policy.
	Name *string `mandatory:"true" json:"name"`

	Condition *DecryptionRuleMatchCriteria `mandatory:"true" json:"condition"`

	// Action:
	// * NO_DECRYPT - Matching traffic is not decrypted.
	// * DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`.
	Action DecryptionRuleActionEnum `mandatory:"true" json:"action"`

	// The name of the decryption profile to use.
	DecryptionProfile *string `mandatory:"false" json:"decryptionProfile"`

	// The name of a mapped secret. Its `type` must match that of the specified decryption profile.
	Secret *string `mandatory:"false" json:"secret"`
}

func (m DecryptionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DecryptionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDecryptionRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDecryptionRuleActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DecryptionRuleActionEnum Enum with underlying type: string
type DecryptionRuleActionEnum string

// Set of constants representing the allowable values for DecryptionRuleActionEnum
const (
	DecryptionRuleActionNoDecrypt DecryptionRuleActionEnum = "NO_DECRYPT"
	DecryptionRuleActionDecrypt   DecryptionRuleActionEnum = "DECRYPT"
)

var mappingDecryptionRuleActionEnum = map[string]DecryptionRuleActionEnum{
	"NO_DECRYPT": DecryptionRuleActionNoDecrypt,
	"DECRYPT":    DecryptionRuleActionDecrypt,
}

var mappingDecryptionRuleActionEnumLowerCase = map[string]DecryptionRuleActionEnum{
	"no_decrypt": DecryptionRuleActionNoDecrypt,
	"decrypt":    DecryptionRuleActionDecrypt,
}

// GetDecryptionRuleActionEnumValues Enumerates the set of values for DecryptionRuleActionEnum
func GetDecryptionRuleActionEnumValues() []DecryptionRuleActionEnum {
	values := make([]DecryptionRuleActionEnum, 0)
	for _, v := range mappingDecryptionRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDecryptionRuleActionEnumStringValues Enumerates the set of values in String for DecryptionRuleActionEnum
func GetDecryptionRuleActionEnumStringValues() []string {
	return []string{
		"NO_DECRYPT",
		"DECRYPT",
	}
}

// GetMappingDecryptionRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDecryptionRuleActionEnum(val string) (DecryptionRuleActionEnum, bool) {
	enum, ok := mappingDecryptionRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
