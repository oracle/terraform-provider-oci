// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SecurityAssessmentComparisonPerTarget The results of the comparison between two security assessment resources.
type SecurityAssessmentComparisonPerTarget struct {

	// The OCID of the target that is used as a baseline in this comparison.
	BaselineTargetId *string `mandatory:"false" json:"baselineTargetId"`

	// The OCID of the target to be compared against the baseline target.
	CurrentTargetId *string `mandatory:"false" json:"currentTargetId"`

	// A comparison between findings belonging to Auditing category.
	Auditing []Diffs `mandatory:"false" json:"auditing"`

	// A comparison between findings belonging to Authorization Control category.
	AuthorizationControl []Diffs `mandatory:"false" json:"authorizationControl"`

	// Comparison between findings belonging to Data Encryption category.
	DataEncryption []Diffs `mandatory:"false" json:"dataEncryption"`

	// Comparison between findings belonging to Database Configuration category.
	DbConfiguration []Diffs `mandatory:"false" json:"dbConfiguration"`

	// Comparison between findings belonging to Fine-Grained Access Control category.
	FineGrainedAccessControl []Diffs `mandatory:"false" json:"fineGrainedAccessControl"`

	// Comparison between findings belonging to Privileges and Roles category.
	PrivilegesAndRoles []Diffs `mandatory:"false" json:"privilegesAndRoles"`

	// Comparison between findings belonging to User Accounts category.
	UserAccounts []Diffs `mandatory:"false" json:"userAccounts"`
}

func (m SecurityAssessmentComparisonPerTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAssessmentComparisonPerTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
