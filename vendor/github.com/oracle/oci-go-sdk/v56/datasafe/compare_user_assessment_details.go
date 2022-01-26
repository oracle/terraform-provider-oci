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

// CompareUserAssessmentDetails The details of the user assessment comparison.
type CompareUserAssessmentDetails struct {

	// The OCID of the user assessment to be compared. You can compare with another user assessment, a latest assessment, or a baseline.
	ComparisonUserAssessmentId *string `mandatory:"true" json:"comparisonUserAssessmentId"`
}

func (m CompareUserAssessmentDetails) String() string {
	return common.PointerString(m)
}
