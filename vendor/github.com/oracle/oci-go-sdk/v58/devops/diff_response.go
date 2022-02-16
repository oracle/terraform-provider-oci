// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DiffResponse Response object for obtaining list of changed files.
type DiffResponse struct {

	// List of changes in the difference.
	Changes []DiffResponseEntry `mandatory:"true" json:"changes"`

	// Boolean value to indicate if all changes are included in the response.
	AreAllChangesIncluded *bool `mandatory:"false" json:"areAllChangesIncluded"`

	// Count of each type of change in difference.
	ChangeTypeCount map[string]int `mandatory:"false" json:"changeTypeCount"`

	// The ID of the common commit between source and target.
	CommonCommit *string `mandatory:"false" json:"commonCommit"`

	// The number of commits source is ahead of target by.
	CommitsAheadCount *int `mandatory:"false" json:"commitsAheadCount"`

	// The number of commits source is behind target by.
	CommitsBehindCount *int `mandatory:"false" json:"commitsBehindCount"`

	// The number of lines added in whole difference.
	AddedLinesCount *int `mandatory:"false" json:"addedLinesCount"`

	// The number of lines deleted in whole difference.
	DeletedLinesCount *int `mandatory:"false" json:"deletedLinesCount"`
}

func (m DiffResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiffResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
