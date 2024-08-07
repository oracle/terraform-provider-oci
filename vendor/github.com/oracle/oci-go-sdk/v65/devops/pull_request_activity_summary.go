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

// PullRequestActivitySummary Summary of an activity record in a pull request
type PullRequestActivitySummary interface {

	// activity identifier
	GetId() *string

	GetPrincipal() *PrincipalDetails

	// pullRequest OCID
	GetPullRequestId() *string

	// The time the action was performed. An RFC3339 formatted datetime string
	GetTimeOccurred() *common.SDKTime
}

type pullrequestactivitysummary struct {
	JsonData      []byte
	Id            *string           `mandatory:"true" json:"id"`
	Principal     *PrincipalDetails `mandatory:"true" json:"principal"`
	PullRequestId *string           `mandatory:"true" json:"pullRequestId"`
	TimeOccurred  *common.SDKTime   `mandatory:"true" json:"timeOccurred"`
	ActivityType  string            `json:"activityType"`
}

// UnmarshalJSON unmarshals json
func (m *pullrequestactivitysummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerpullrequestactivitysummary pullrequestactivitysummary
	s := struct {
		Model Unmarshalerpullrequestactivitysummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.Principal = s.Model.Principal
	m.PullRequestId = s.Model.PullRequestId
	m.TimeOccurred = s.Model.TimeOccurred
	m.ActivityType = s.Model.ActivityType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *pullrequestactivitysummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ActivityType {
	case "COMMIT":
		mm := CommitActivitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMMENT":
		mm := CommentActivitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPROVAL":
		mm := ApprovalActivitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REVIEWER":
		mm := ReviewerActivitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LIFECYCLE":
		mm := PullRequestLifecycleActivitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for PullRequestActivitySummary: %s.", m.ActivityType)
		return *m, nil
	}
}

// GetId returns Id
func (m pullrequestactivitysummary) GetId() *string {
	return m.Id
}

// GetPrincipal returns Principal
func (m pullrequestactivitysummary) GetPrincipal() *PrincipalDetails {
	return m.Principal
}

// GetPullRequestId returns PullRequestId
func (m pullrequestactivitysummary) GetPullRequestId() *string {
	return m.PullRequestId
}

// GetTimeOccurred returns TimeOccurred
func (m pullrequestactivitysummary) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

func (m pullrequestactivitysummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m pullrequestactivitysummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PullRequestActivitySummaryActivityTypeEnum Enum with underlying type: string
type PullRequestActivitySummaryActivityTypeEnum string

// Set of constants representing the allowable values for PullRequestActivitySummaryActivityTypeEnum
const (
	PullRequestActivitySummaryActivityTypeLifecycle PullRequestActivitySummaryActivityTypeEnum = "LIFECYCLE"
	PullRequestActivitySummaryActivityTypeApproval  PullRequestActivitySummaryActivityTypeEnum = "APPROVAL"
	PullRequestActivitySummaryActivityTypeCommit    PullRequestActivitySummaryActivityTypeEnum = "COMMIT"
	PullRequestActivitySummaryActivityTypeReviewer  PullRequestActivitySummaryActivityTypeEnum = "REVIEWER"
	PullRequestActivitySummaryActivityTypeComment   PullRequestActivitySummaryActivityTypeEnum = "COMMENT"
)

var mappingPullRequestActivitySummaryActivityTypeEnum = map[string]PullRequestActivitySummaryActivityTypeEnum{
	"LIFECYCLE": PullRequestActivitySummaryActivityTypeLifecycle,
	"APPROVAL":  PullRequestActivitySummaryActivityTypeApproval,
	"COMMIT":    PullRequestActivitySummaryActivityTypeCommit,
	"REVIEWER":  PullRequestActivitySummaryActivityTypeReviewer,
	"COMMENT":   PullRequestActivitySummaryActivityTypeComment,
}

var mappingPullRequestActivitySummaryActivityTypeEnumLowerCase = map[string]PullRequestActivitySummaryActivityTypeEnum{
	"lifecycle": PullRequestActivitySummaryActivityTypeLifecycle,
	"approval":  PullRequestActivitySummaryActivityTypeApproval,
	"commit":    PullRequestActivitySummaryActivityTypeCommit,
	"reviewer":  PullRequestActivitySummaryActivityTypeReviewer,
	"comment":   PullRequestActivitySummaryActivityTypeComment,
}

// GetPullRequestActivitySummaryActivityTypeEnumValues Enumerates the set of values for PullRequestActivitySummaryActivityTypeEnum
func GetPullRequestActivitySummaryActivityTypeEnumValues() []PullRequestActivitySummaryActivityTypeEnum {
	values := make([]PullRequestActivitySummaryActivityTypeEnum, 0)
	for _, v := range mappingPullRequestActivitySummaryActivityTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPullRequestActivitySummaryActivityTypeEnumStringValues Enumerates the set of values in String for PullRequestActivitySummaryActivityTypeEnum
func GetPullRequestActivitySummaryActivityTypeEnumStringValues() []string {
	return []string{
		"LIFECYCLE",
		"APPROVAL",
		"COMMIT",
		"REVIEWER",
		"COMMENT",
	}
}

// GetMappingPullRequestActivitySummaryActivityTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPullRequestActivitySummaryActivityTypeEnum(val string) (PullRequestActivitySummaryActivityTypeEnum, bool) {
	enum, ok := mappingPullRequestActivitySummaryActivityTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
