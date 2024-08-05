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

// Reviewer Reviewer information.
type Reviewer struct {

	// Pull Request reviewer id
	PrincipalId *string `mandatory:"true" json:"principalId"`

	// the name of the principal
	PrincipalName *string `mandatory:"false" json:"principalName"`

	// the type of principal
	PrincipalType ReviewerPrincipalTypeEnum `mandatory:"false" json:"principalType,omitempty"`

	// The state of the principal, it can be active or inactive or suppressed for emails
	PrincipalState ReviewerPrincipalStateEnum `mandatory:"false" json:"principalState,omitempty"`

	// The current state of the Review.
	Status ReviewerStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m Reviewer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Reviewer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReviewerPrincipalTypeEnum(string(m.PrincipalType)); !ok && m.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", m.PrincipalType, strings.Join(GetReviewerPrincipalTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReviewerPrincipalStateEnum(string(m.PrincipalState)); !ok && m.PrincipalState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalState: %s. Supported values are: %s.", m.PrincipalState, strings.Join(GetReviewerPrincipalStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReviewerStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetReviewerStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReviewerPrincipalTypeEnum Enum with underlying type: string
type ReviewerPrincipalTypeEnum string

// Set of constants representing the allowable values for ReviewerPrincipalTypeEnum
const (
	ReviewerPrincipalTypeService  ReviewerPrincipalTypeEnum = "SERVICE"
	ReviewerPrincipalTypeUser     ReviewerPrincipalTypeEnum = "USER"
	ReviewerPrincipalTypeInstance ReviewerPrincipalTypeEnum = "INSTANCE"
	ReviewerPrincipalTypeResource ReviewerPrincipalTypeEnum = "RESOURCE"
)

var mappingReviewerPrincipalTypeEnum = map[string]ReviewerPrincipalTypeEnum{
	"SERVICE":  ReviewerPrincipalTypeService,
	"USER":     ReviewerPrincipalTypeUser,
	"INSTANCE": ReviewerPrincipalTypeInstance,
	"RESOURCE": ReviewerPrincipalTypeResource,
}

var mappingReviewerPrincipalTypeEnumLowerCase = map[string]ReviewerPrincipalTypeEnum{
	"service":  ReviewerPrincipalTypeService,
	"user":     ReviewerPrincipalTypeUser,
	"instance": ReviewerPrincipalTypeInstance,
	"resource": ReviewerPrincipalTypeResource,
}

// GetReviewerPrincipalTypeEnumValues Enumerates the set of values for ReviewerPrincipalTypeEnum
func GetReviewerPrincipalTypeEnumValues() []ReviewerPrincipalTypeEnum {
	values := make([]ReviewerPrincipalTypeEnum, 0)
	for _, v := range mappingReviewerPrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReviewerPrincipalTypeEnumStringValues Enumerates the set of values in String for ReviewerPrincipalTypeEnum
func GetReviewerPrincipalTypeEnumStringValues() []string {
	return []string{
		"SERVICE",
		"USER",
		"INSTANCE",
		"RESOURCE",
	}
}

// GetMappingReviewerPrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReviewerPrincipalTypeEnum(val string) (ReviewerPrincipalTypeEnum, bool) {
	enum, ok := mappingReviewerPrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReviewerPrincipalStateEnum Enum with underlying type: string
type ReviewerPrincipalStateEnum string

// Set of constants representing the allowable values for ReviewerPrincipalStateEnum
const (
	ReviewerPrincipalStateActive     ReviewerPrincipalStateEnum = "ACTIVE"
	ReviewerPrincipalStateInactive   ReviewerPrincipalStateEnum = "INACTIVE"
	ReviewerPrincipalStateSuppressed ReviewerPrincipalStateEnum = "SUPPRESSED"
)

var mappingReviewerPrincipalStateEnum = map[string]ReviewerPrincipalStateEnum{
	"ACTIVE":     ReviewerPrincipalStateActive,
	"INACTIVE":   ReviewerPrincipalStateInactive,
	"SUPPRESSED": ReviewerPrincipalStateSuppressed,
}

var mappingReviewerPrincipalStateEnumLowerCase = map[string]ReviewerPrincipalStateEnum{
	"active":     ReviewerPrincipalStateActive,
	"inactive":   ReviewerPrincipalStateInactive,
	"suppressed": ReviewerPrincipalStateSuppressed,
}

// GetReviewerPrincipalStateEnumValues Enumerates the set of values for ReviewerPrincipalStateEnum
func GetReviewerPrincipalStateEnumValues() []ReviewerPrincipalStateEnum {
	values := make([]ReviewerPrincipalStateEnum, 0)
	for _, v := range mappingReviewerPrincipalStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReviewerPrincipalStateEnumStringValues Enumerates the set of values in String for ReviewerPrincipalStateEnum
func GetReviewerPrincipalStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"SUPPRESSED",
	}
}

// GetMappingReviewerPrincipalStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReviewerPrincipalStateEnum(val string) (ReviewerPrincipalStateEnum, bool) {
	enum, ok := mappingReviewerPrincipalStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReviewerStatusEnum Enum with underlying type: string
type ReviewerStatusEnum string

// Set of constants representing the allowable values for ReviewerStatusEnum
const (
	ReviewerStatusApproved      ReviewerStatusEnum = "APPROVED"
	ReviewerStatusReviewPending ReviewerStatusEnum = "REVIEW_PENDING"
)

var mappingReviewerStatusEnum = map[string]ReviewerStatusEnum{
	"APPROVED":       ReviewerStatusApproved,
	"REVIEW_PENDING": ReviewerStatusReviewPending,
}

var mappingReviewerStatusEnumLowerCase = map[string]ReviewerStatusEnum{
	"approved":       ReviewerStatusApproved,
	"review_pending": ReviewerStatusReviewPending,
}

// GetReviewerStatusEnumValues Enumerates the set of values for ReviewerStatusEnum
func GetReviewerStatusEnumValues() []ReviewerStatusEnum {
	values := make([]ReviewerStatusEnum, 0)
	for _, v := range mappingReviewerStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReviewerStatusEnumStringValues Enumerates the set of values in String for ReviewerStatusEnum
func GetReviewerStatusEnumStringValues() []string {
	return []string{
		"APPROVED",
		"REVIEW_PENDING",
	}
}

// GetMappingReviewerStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReviewerStatusEnum(val string) (ReviewerStatusEnum, bool) {
	enum, ok := mappingReviewerStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
