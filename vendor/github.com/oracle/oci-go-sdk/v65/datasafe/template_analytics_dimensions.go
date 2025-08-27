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

// TemplateAnalyticsDimensions The scope of analytics data.
type TemplateAnalyticsDimensions struct {

	// The OCID of the security assessment of type TEMPLATE.
	TemplateAssessmentId *string `mandatory:"false" json:"templateAssessmentId"`

	// The number of checks inside the template assessment.
	TotalChecks *int `mandatory:"false" json:"totalChecks"`

	// The OCID of the security assessment of type TEMPLATE_BASELINE.
	TemplateBaselineAssessmentId *string `mandatory:"false" json:"templateBaselineAssessmentId"`

	// The OCID of the target database. This field will be in the response if the template was applied on an individual target.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The OCID of the target database group that the group assessment is created for.
	// This field will be in the response if the template was applied on a target group.
	TargetDatabaseGroupId *string `mandatory:"false" json:"targetDatabaseGroupId"`

	// Indicates whether or not the template security assessment is applied to a target group.
	// If the value is false, it means the template security assessment is applied to a individual target.
	IsGroup *bool `mandatory:"false" json:"isGroup"`

	// Indicates whether or not the comparison between the latest assessment and the template baseline assessment is done.
	// If the value is false, it means the comparison is not done yet.
	IsCompared *bool `mandatory:"false" json:"isCompared"`

	// The date and time when the comparison was made upon the template baseline. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeLastCompared *common.SDKTime `mandatory:"false" json:"timeLastCompared"`

	// Indicates whether or not the latest assessment is compliant with the template baseline assessment.
	// If the value is false, it means there is drift in the comparison report and the totalChecksFailed field will have a non-zero value.
	IsCompliant *bool `mandatory:"false" json:"isCompliant"`

	// Indicates how many checks in the template have drifts in the comparison report. This field is only present if isCompliant is false.
	TotalChecksFailed *int `mandatory:"false" json:"totalChecksFailed"`

	// The number of the target(s) inside the target group for which the template baseline assessment was created for.
	// If the isGroup field is false, the value will be 1, representing the single target.
	TotalTargets *int `mandatory:"false" json:"totalTargets"`

	// The number of the target(s) that have drifts in the comparison report.
	// This field is only present if isCompared is true.
	TotalNonCompliantTargets *int `mandatory:"false" json:"totalNonCompliantTargets"`
}

func (m TemplateAnalyticsDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateAnalyticsDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
