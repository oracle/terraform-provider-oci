// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateUserAssessmentDetails Updates one or more attributes of the specified user assessment.
type UpdateUserAssessmentDetails struct {

	// The description of the user assessment.
	Description *string `mandatory:"false" json:"description"`

	// The display name of the user assessment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The schedule for periodically saving the assessment. This is applicable only for assessments of type save schedule and latest assessment. It updates the existing schedule in a specified format:
	// <version-string>;<version-specific-schedule>
	// Allowed version strings - "v1"
	// v1's version specific schedule -<ss> <mm> <hh> <day-of-week> <day-of-month>
	// Each of the above fields potentially introduce constraints. A workrequest is created only
	// when clock time satisfies all the constraints. Constraints introduced:
	// 1. seconds = <ss> (So, the allowed range for <ss> is [0, 59])
	// 2. minutes = <mm> (So, the allowed range for <mm> is [0, 59])
	// 3. hours = <hh> (So, the allowed range for <hh> is [0, 23])
	// <day-of-week> can be either '*' (without quotes or a number between 1(Monday) and 7(Sunday))
	// 4. No constraint introduced when it is '*'. When not, day of week must equal the given value
	// <day-of-month> can be either '*' (without quotes or a number between 1 and 28)
	// 5. No constraint introduced when it is '*'. When not, day of month must equal the given value
	Schedule *string `mandatory:"false" json:"schedule"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateUserAssessmentDetails) String() string {
	return common.PointerString(m)
}
