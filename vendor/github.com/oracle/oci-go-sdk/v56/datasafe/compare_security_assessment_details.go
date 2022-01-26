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

// CompareSecurityAssessmentDetails Details specifying the security assessment used for comparison.
type CompareSecurityAssessmentDetails struct {

	// The OCID of the security assessment. In this case a security assessment can be another security assessment, a latest assessment or a baseline.
	ComparisonSecurityAssessmentId *string `mandatory:"true" json:"comparisonSecurityAssessmentId"`
}

func (m CompareSecurityAssessmentDetails) String() string {
	return common.PointerString(m)
}
