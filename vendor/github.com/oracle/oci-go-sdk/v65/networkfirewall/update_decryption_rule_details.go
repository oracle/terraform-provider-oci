// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateDecryptionRuleDetails Request for updating Decryption Rule used in the firewall policy rules.
// A Decryption Rule is used to define which traffic should be decrypted by the firewall, and how it should do so.
type UpdateDecryptionRuleDetails struct {
	Condition *DecryptionRuleMatchCriteria `mandatory:"true" json:"condition"`

	// Action:
	// * NO_DECRYPT - Matching traffic is not decrypted.
	// * DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`.
	Action DecryptionActionTypeEnum `mandatory:"true" json:"action"`

	// The name of the decryption profile to use.
	DecryptionProfile *string `mandatory:"false" json:"decryptionProfile"`

	// The name of a mapped secret. Its `type` must match that of the specified decryption profile.
	Secret *string `mandatory:"false" json:"secret"`

	Position *RulePosition `mandatory:"false" json:"position"`
}

func (m UpdateDecryptionRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDecryptionRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDecryptionActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDecryptionActionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
