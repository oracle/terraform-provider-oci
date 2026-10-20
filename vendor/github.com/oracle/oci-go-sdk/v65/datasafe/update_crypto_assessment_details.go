// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateCryptoAssessmentDetails Updates one or more attributes of the specified crypto assessment.
type UpdateCryptoAssessmentDetails struct {

	// The display name of the crypto assessment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Updates the schedule associated with this latest crypto assessment. The schedule uses the format:
	// <version-string>;<version-specific-schedule>
	// For v1, the version-specific schedule format is:
	// <ss> <mm> <hh> <day-of-week> <day-of-month>
	// Specify either day-of-week for weekly schedules or day-of-month for monthly schedules. Do not specify both.
	// For monthly schedules, day-of-month must be between 1 and 28. If the service generates a default monthly schedule for an assessment created on day 29, 30, or 31 of a month, it uses day 28.
	// Schedule updates are supported only for assessments of type LATEST.
	Schedule *string `mandatory:"false" json:"schedule"`

	// Indicates whether scheduled execution is active for this latest crypto assessment.
	// When set to true, the existing schedule is used unless a new schedule is provided.
	// When set to false, the schedule value is retained but scheduled execution is paused.
	// Schedule state updates are supported only for assessments of type LATEST.
	IsAssessmentScheduled *bool `mandatory:"false" json:"isAssessmentScheduled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCryptoAssessmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCryptoAssessmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
