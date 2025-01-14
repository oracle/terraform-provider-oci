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

// TemplateAnalyticDimensions The scope of analytics data.
type TemplateAnalyticDimensions struct {

	// The OCID of the security assessment of type TEMPLATE.
	TemplateAssessmentId *string `mandatory:"false" json:"templateAssessmentId"`

	// The OCID of the security assessment of type TEMPLATE_BASELINE.
	TemplateBaselineAssessmentId *string `mandatory:"false" json:"templateBaselineAssessmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Indicates whether or not the template security assessment is applied to a target group.
	// If the value is false, it means the template security assessment is applied to a individual target.
	IsGroup *bool `mandatory:"false" json:"isGroup"`
}

func (m TemplateAnalyticDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateAnalyticDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
