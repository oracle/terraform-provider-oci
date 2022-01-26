// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateBulkProblemStatusDetails List of problem ids to be passed in to update the Problem status.
type UpdateBulkProblemStatusDetails struct {

	// Action taken by user
	Status ProblemLifecycleDetailEnum `mandatory:"true" json:"status"`

	// List of ProblemIds to be passed in to update the Problem status.
	ProblemIds []string `mandatory:"true" json:"problemIds"`

	// User defined comment to be passed in to update the problem.
	Comment *string `mandatory:"false" json:"comment"`
}

func (m UpdateBulkProblemStatusDetails) String() string {
	return common.PointerString(m)
}
