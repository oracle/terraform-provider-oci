// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CommitActivitySummary activity describing commit updates to a pull request
type CommitActivitySummary struct {

	// activity identifier
	Id *string `mandatory:"true" json:"id"`

	Principal *PrincipalDetails `mandatory:"true" json:"principal"`

	// pullRequest OCID
	PullRequestId *string `mandatory:"true" json:"pullRequestId"`

	// The time the action was performed. An RFC3339 formatted datetime string
	TimeOccurred *common.SDKTime `mandatory:"true" json:"timeOccurred"`

	// list of commits added to a PR source branch after git push
	CommitsAdded []RepositoryCommit `mandatory:"true" json:"commitsAdded"`

	// list of commits removed from a PR source branch after git push
	CommitsRemoved []RepositoryCommit `mandatory:"true" json:"commitsRemoved"`
}

// GetId returns Id
func (m CommitActivitySummary) GetId() *string {
	return m.Id
}

// GetPrincipal returns Principal
func (m CommitActivitySummary) GetPrincipal() *PrincipalDetails {
	return m.Principal
}

// GetPullRequestId returns PullRequestId
func (m CommitActivitySummary) GetPullRequestId() *string {
	return m.PullRequestId
}

// GetTimeOccurred returns TimeOccurred
func (m CommitActivitySummary) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

func (m CommitActivitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CommitActivitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CommitActivitySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCommitActivitySummary CommitActivitySummary
	s := struct {
		DiscriminatorParam string `json:"activityType"`
		MarshalTypeCommitActivitySummary
	}{
		"COMMIT",
		(MarshalTypeCommitActivitySummary)(m),
	}

	return json.Marshal(&s)
}
