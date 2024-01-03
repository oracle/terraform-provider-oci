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

// ApprovalActivitySummary activity describing a reviewer's approval decision
type ApprovalActivitySummary struct {

	// activity identifier
	Id *string `mandatory:"true" json:"id"`

	Principal *PrincipalDetails `mandatory:"true" json:"principal"`

	// pullRequest OCID
	PullRequestId *string `mandatory:"true" json:"pullRequestId"`

	// The time the action was performed. An RFC3339 formatted datetime string
	TimeOccurred *common.SDKTime `mandatory:"true" json:"timeOccurred"`

	// The approval status of a reviewer
	Status ApprovalActivitySummaryStatusEnum `mandatory:"true" json:"status"`
}

// GetId returns Id
func (m ApprovalActivitySummary) GetId() *string {
	return m.Id
}

// GetPrincipal returns Principal
func (m ApprovalActivitySummary) GetPrincipal() *PrincipalDetails {
	return m.Principal
}

// GetPullRequestId returns PullRequestId
func (m ApprovalActivitySummary) GetPullRequestId() *string {
	return m.PullRequestId
}

// GetTimeOccurred returns TimeOccurred
func (m ApprovalActivitySummary) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

func (m ApprovalActivitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApprovalActivitySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingApprovalActivitySummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetApprovalActivitySummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ApprovalActivitySummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApprovalActivitySummary ApprovalActivitySummary
	s := struct {
		DiscriminatorParam string `json:"activityType"`
		MarshalTypeApprovalActivitySummary
	}{
		"APPROVAL",
		(MarshalTypeApprovalActivitySummary)(m),
	}

	return json.Marshal(&s)
}

// ApprovalActivitySummaryStatusEnum Enum with underlying type: string
type ApprovalActivitySummaryStatusEnum string

// Set of constants representing the allowable values for ApprovalActivitySummaryStatusEnum
const (
	ApprovalActivitySummaryStatusApproved   ApprovalActivitySummaryStatusEnum = "APPROVED"
	ApprovalActivitySummaryStatusUnapproved ApprovalActivitySummaryStatusEnum = "UNAPPROVED"
)

var mappingApprovalActivitySummaryStatusEnum = map[string]ApprovalActivitySummaryStatusEnum{
	"APPROVED":   ApprovalActivitySummaryStatusApproved,
	"UNAPPROVED": ApprovalActivitySummaryStatusUnapproved,
}

var mappingApprovalActivitySummaryStatusEnumLowerCase = map[string]ApprovalActivitySummaryStatusEnum{
	"approved":   ApprovalActivitySummaryStatusApproved,
	"unapproved": ApprovalActivitySummaryStatusUnapproved,
}

// GetApprovalActivitySummaryStatusEnumValues Enumerates the set of values for ApprovalActivitySummaryStatusEnum
func GetApprovalActivitySummaryStatusEnumValues() []ApprovalActivitySummaryStatusEnum {
	values := make([]ApprovalActivitySummaryStatusEnum, 0)
	for _, v := range mappingApprovalActivitySummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetApprovalActivitySummaryStatusEnumStringValues Enumerates the set of values in String for ApprovalActivitySummaryStatusEnum
func GetApprovalActivitySummaryStatusEnumStringValues() []string {
	return []string{
		"APPROVED",
		"UNAPPROVED",
	}
}

// GetMappingApprovalActivitySummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingApprovalActivitySummaryStatusEnum(val string) (ApprovalActivitySummaryStatusEnum, bool) {
	enum, ok := mappingApprovalActivitySummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
