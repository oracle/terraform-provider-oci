// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateAlertPolicyRuleDetails The details used to create a new alert policy rule.
type CreateAlertPolicyRuleDetails struct {

	// The conditional expression of the alert policy rule which evaluates to boolean value.
	Expression *string `mandatory:"true" json:"expression"`

	// Describes the alert policy rule.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the alert policy rule.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateAlertPolicyRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAlertPolicyRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
