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

// TemplateAssociationAnalyticsDimensions The scope of template association analytics data.
type TemplateAssociationAnalyticsDimensions struct {

	// The OCID of the security assessment of type TEMPLATE.
	TemplateAssessmentId *string `mandatory:"false" json:"templateAssessmentId"`

	// The OCID of the security assessment of type TEMPLATE_BASELINE.
	TemplateBaselineAssessmentId *string `mandatory:"false" json:"templateBaselineAssessmentId"`

	// The OCID of the target database group that the group assessment is created for.
	// This field will be in the response if the template was applied on a target group.
	TargetDatabaseGroupId *string `mandatory:"false" json:"targetDatabaseGroupId"`

	// The OCID of the target database.
	// If the template was applied on a target group, this field will be the OCID of the target members of the target group.
	// If the template was applied on an individual target, this field will contain that targetId.
	TargetId *string `mandatory:"false" json:"targetId"`
}

func (m TemplateAssociationAnalyticsDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateAssociationAnalyticsDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
