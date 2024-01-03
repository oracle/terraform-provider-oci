// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PullRequest Pull Request containing the diff between a review branch and a destination branch
type PullRequest struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// PullRequest title, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// DevOps Repository Identifier tells which repository this pull request belongs to
	RepositoryId *string `mandatory:"true" json:"repositoryId"`

	// The source branch which contains the changes to be reviewed. Example: "feature/JIRA-123"
	SourceBranch *string `mandatory:"true" json:"sourceBranch"`

	// The destination branch against which the changes are to be reviewed. Example: "main".
	DestinationBranch *string `mandatory:"true" json:"destinationBranch"`

	// The time the PullRequest was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the PullRequest.
	LifecycleState PullRequestLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The status of the Pull Request.
	LifecycleDetails PullRequestLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	CreatedBy *PrincipalDetails `mandatory:"true" json:"createdBy"`

	// The total number of comments on the pull request.
	TotalComments *int `mandatory:"true" json:"totalComments"`

	// The total number of reviewers on the pull request.
	TotalReviewers *int `mandatory:"true" json:"totalReviewers"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Details of the pull request. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The source branch commit ID when the Pull Request status was last changed to merged or closed
	SourceCommitIdAtTermination *string `mandatory:"false" json:"sourceCommitIdAtTermination"`

	// The merge base commit ID when the Pull Request status was last changed to merged or closed
	MergeBaseCommitIdAtTermination *string `mandatory:"false" json:"mergeBaseCommitIdAtTermination"`

	// The time the PullRequest was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// List of Reviewers.
	Reviewers []Reviewer `mandatory:"false" json:"reviewers"`

	MergeChecks *MergeCheckCollection `mandatory:"false" json:"mergeChecks"`

	MergedBy *PrincipalDetails `mandatory:"false" json:"mergedBy"`

	// The OCID of the forked repository that will act as the source of the changes to be included in the pull request against the parent repository.
	SourceRepositoryId *string `mandatory:"false" json:"sourceRepositoryId"`
}

func (m PullRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PullRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPullRequestLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPullRequestLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPullRequestLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetPullRequestLifecycleDetailsEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PullRequestLifecycleStateEnum Enum with underlying type: string
type PullRequestLifecycleStateEnum string

// Set of constants representing the allowable values for PullRequestLifecycleStateEnum
const (
	PullRequestLifecycleStateCreating PullRequestLifecycleStateEnum = "CREATING"
	PullRequestLifecycleStateUpdating PullRequestLifecycleStateEnum = "UPDATING"
	PullRequestLifecycleStateActive   PullRequestLifecycleStateEnum = "ACTIVE"
	PullRequestLifecycleStateDeleting PullRequestLifecycleStateEnum = "DELETING"
	PullRequestLifecycleStateDeleted  PullRequestLifecycleStateEnum = "DELETED"
	PullRequestLifecycleStateFailed   PullRequestLifecycleStateEnum = "FAILED"
)

var mappingPullRequestLifecycleStateEnum = map[string]PullRequestLifecycleStateEnum{
	"CREATING": PullRequestLifecycleStateCreating,
	"UPDATING": PullRequestLifecycleStateUpdating,
	"ACTIVE":   PullRequestLifecycleStateActive,
	"DELETING": PullRequestLifecycleStateDeleting,
	"DELETED":  PullRequestLifecycleStateDeleted,
	"FAILED":   PullRequestLifecycleStateFailed,
}

var mappingPullRequestLifecycleStateEnumLowerCase = map[string]PullRequestLifecycleStateEnum{
	"creating": PullRequestLifecycleStateCreating,
	"updating": PullRequestLifecycleStateUpdating,
	"active":   PullRequestLifecycleStateActive,
	"deleting": PullRequestLifecycleStateDeleting,
	"deleted":  PullRequestLifecycleStateDeleted,
	"failed":   PullRequestLifecycleStateFailed,
}

// GetPullRequestLifecycleStateEnumValues Enumerates the set of values for PullRequestLifecycleStateEnum
func GetPullRequestLifecycleStateEnumValues() []PullRequestLifecycleStateEnum {
	values := make([]PullRequestLifecycleStateEnum, 0)
	for _, v := range mappingPullRequestLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPullRequestLifecycleStateEnumStringValues Enumerates the set of values in String for PullRequestLifecycleStateEnum
func GetPullRequestLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPullRequestLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPullRequestLifecycleStateEnum(val string) (PullRequestLifecycleStateEnum, bool) {
	enum, ok := mappingPullRequestLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PullRequestLifecycleDetailsEnum Enum with underlying type: string
type PullRequestLifecycleDetailsEnum string

// Set of constants representing the allowable values for PullRequestLifecycleDetailsEnum
const (
	PullRequestLifecycleDetailsOpen     PullRequestLifecycleDetailsEnum = "OPEN"
	PullRequestLifecycleDetailsConflict PullRequestLifecycleDetailsEnum = "CONFLICT"
	PullRequestLifecycleDetailsClosed   PullRequestLifecycleDetailsEnum = "CLOSED"
	PullRequestLifecycleDetailsMerging  PullRequestLifecycleDetailsEnum = "MERGING"
	PullRequestLifecycleDetailsMerged   PullRequestLifecycleDetailsEnum = "MERGED"
)

var mappingPullRequestLifecycleDetailsEnum = map[string]PullRequestLifecycleDetailsEnum{
	"OPEN":     PullRequestLifecycleDetailsOpen,
	"CONFLICT": PullRequestLifecycleDetailsConflict,
	"CLOSED":   PullRequestLifecycleDetailsClosed,
	"MERGING":  PullRequestLifecycleDetailsMerging,
	"MERGED":   PullRequestLifecycleDetailsMerged,
}

var mappingPullRequestLifecycleDetailsEnumLowerCase = map[string]PullRequestLifecycleDetailsEnum{
	"open":     PullRequestLifecycleDetailsOpen,
	"conflict": PullRequestLifecycleDetailsConflict,
	"closed":   PullRequestLifecycleDetailsClosed,
	"merging":  PullRequestLifecycleDetailsMerging,
	"merged":   PullRequestLifecycleDetailsMerged,
}

// GetPullRequestLifecycleDetailsEnumValues Enumerates the set of values for PullRequestLifecycleDetailsEnum
func GetPullRequestLifecycleDetailsEnumValues() []PullRequestLifecycleDetailsEnum {
	values := make([]PullRequestLifecycleDetailsEnum, 0)
	for _, v := range mappingPullRequestLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetPullRequestLifecycleDetailsEnumStringValues Enumerates the set of values in String for PullRequestLifecycleDetailsEnum
func GetPullRequestLifecycleDetailsEnumStringValues() []string {
	return []string{
		"OPEN",
		"CONFLICT",
		"CLOSED",
		"MERGING",
		"MERGED",
	}
}

// GetMappingPullRequestLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPullRequestLifecycleDetailsEnum(val string) (PullRequestLifecycleDetailsEnum, bool) {
	enum, ok := mappingPullRequestLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
