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

// AlertPolicyRule A rule associated with a alert policy.
type AlertPolicyRule struct {

	// The unique key of the alert policy rule.
	Key *string `mandatory:"true" json:"key"`

	// The conditional expression of the alert policy rule which evaluates to boolean value.
	Expression *string `mandatory:"true" json:"expression"`

	// Describes the alert policy rule.
	Description *string `mandatory:"false" json:"description"`

	// The current state of the alert policy rule.
	LifecycleState AlertPolicyRuleLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The display name of the alert policy rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Creation date and time of the alert policy rule, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AlertPolicyRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlertPolicyRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAlertPolicyRuleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlertPolicyRuleLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
