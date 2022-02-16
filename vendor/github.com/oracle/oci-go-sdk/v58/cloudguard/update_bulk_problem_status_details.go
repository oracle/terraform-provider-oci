// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBulkProblemStatusDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProblemLifecycleDetailEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetProblemLifecycleDetailEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
