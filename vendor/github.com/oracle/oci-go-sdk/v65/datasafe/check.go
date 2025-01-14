// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Check The security rule to be evaluated by security assessment to create finding.
type Check struct {

	// A unique identifier for the check.
	Key *string `mandatory:"true" json:"key"`

	// The short title for the check.
	Title *string `mandatory:"true" json:"title"`

	// The explanation of the issue in this check. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
	Remarks *string `mandatory:"false" json:"remarks"`

	// Provides information on whether the check is related to a CIS Oracle Database Benchmark recommendation, STIG rule, GDPR Article/Recital or related to the Oracle Best Practice.
	References *References `mandatory:"false" json:"references"`

	// The category to which the check belongs to.
	Category *string `mandatory:"false" json:"category"`

	// Provides a recommended approach to take to remediate the check reported.
	Oneline *string `mandatory:"false" json:"oneline"`

	// The severity of the check as suggested by Data Safe security assessment. This will be the default severity in the template baseline security assessment.
	SuggestedSeverity *string `mandatory:"false" json:"suggestedSeverity"`
}

func (m Check) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Check) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
