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

// CryptoAssessmentFindingCollection List of crypto deviation findings for the specified crypto assessment.
type CryptoAssessmentFindingCollection struct {
	Summary *CryptoAssessmentFindingSummaryMetrics `mandatory:"true" json:"summary"`

	// Array of crypto deviation findings.
	Items []CryptoAssessmentFindingSummary `mandatory:"true" json:"items"`

	// The OCID of the target database for the specified crypto assessment.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Assessment type recorded in the findings table.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" json:"assessmentType,omitempty"`

	// Database version for the target database associated with the specified crypto assessment.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`
}

func (m CryptoAssessmentFindingCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentFindingCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(m.AssessmentType)); !ok && m.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", m.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
