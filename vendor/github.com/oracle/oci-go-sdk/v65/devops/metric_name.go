// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"strings"
)

// MetricNameEnum Enum with underlying type: string
type MetricNameEnum string

// Set of constants representing the allowable values for MetricNameEnum
const (
	MetricNameCommits                              MetricNameEnum = "COMMITS"
	MetricNameLinesAdded                           MetricNameEnum = "LINES_ADDED"
	MetricNameLinesDeleted                         MetricNameEnum = "LINES_DELETED"
	MetricNamePullRequestCreated                   MetricNameEnum = "PULL_REQUEST_CREATED"
	MetricNamePullRequestMerged                    MetricNameEnum = "PULL_REQUEST_MERGED"
	MetricNamePullRequestDeclined                  MetricNameEnum = "PULL_REQUEST_DECLINED"
	MetricNamePullRequestReviewStartDurationInDays MetricNameEnum = "PULL_REQUEST_REVIEW_START_DURATION_IN_DAYS"
	MetricNamePullRequestReviewDurationInDays      MetricNameEnum = "PULL_REQUEST_REVIEW_DURATION_IN_DAYS"
	MetricNamePullRequestApproved                  MetricNameEnum = "PULL_REQUEST_APPROVED"
	MetricNamePullRequestReviewed                  MetricNameEnum = "PULL_REQUEST_REVIEWED"
	MetricNamePullRequestComments                  MetricNameEnum = "PULL_REQUEST_COMMENTS"
)

var mappingMetricNameEnum = map[string]MetricNameEnum{
	"COMMITS":               MetricNameCommits,
	"LINES_ADDED":           MetricNameLinesAdded,
	"LINES_DELETED":         MetricNameLinesDeleted,
	"PULL_REQUEST_CREATED":  MetricNamePullRequestCreated,
	"PULL_REQUEST_MERGED":   MetricNamePullRequestMerged,
	"PULL_REQUEST_DECLINED": MetricNamePullRequestDeclined,
	"PULL_REQUEST_REVIEW_START_DURATION_IN_DAYS": MetricNamePullRequestReviewStartDurationInDays,
	"PULL_REQUEST_REVIEW_DURATION_IN_DAYS":       MetricNamePullRequestReviewDurationInDays,
	"PULL_REQUEST_APPROVED":                      MetricNamePullRequestApproved,
	"PULL_REQUEST_REVIEWED":                      MetricNamePullRequestReviewed,
	"PULL_REQUEST_COMMENTS":                      MetricNamePullRequestComments,
}

var mappingMetricNameEnumLowerCase = map[string]MetricNameEnum{
	"commits":               MetricNameCommits,
	"lines_added":           MetricNameLinesAdded,
	"lines_deleted":         MetricNameLinesDeleted,
	"pull_request_created":  MetricNamePullRequestCreated,
	"pull_request_merged":   MetricNamePullRequestMerged,
	"pull_request_declined": MetricNamePullRequestDeclined,
	"pull_request_review_start_duration_in_days": MetricNamePullRequestReviewStartDurationInDays,
	"pull_request_review_duration_in_days":       MetricNamePullRequestReviewDurationInDays,
	"pull_request_approved":                      MetricNamePullRequestApproved,
	"pull_request_reviewed":                      MetricNamePullRequestReviewed,
	"pull_request_comments":                      MetricNamePullRequestComments,
}

// GetMetricNameEnumValues Enumerates the set of values for MetricNameEnum
func GetMetricNameEnumValues() []MetricNameEnum {
	values := make([]MetricNameEnum, 0)
	for _, v := range mappingMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricNameEnumStringValues Enumerates the set of values in String for MetricNameEnum
func GetMetricNameEnumStringValues() []string {
	return []string{
		"COMMITS",
		"LINES_ADDED",
		"LINES_DELETED",
		"PULL_REQUEST_CREATED",
		"PULL_REQUEST_MERGED",
		"PULL_REQUEST_DECLINED",
		"PULL_REQUEST_REVIEW_START_DURATION_IN_DAYS",
		"PULL_REQUEST_REVIEW_DURATION_IN_DAYS",
		"PULL_REQUEST_APPROVED",
		"PULL_REQUEST_REVIEWED",
		"PULL_REQUEST_COMMENTS",
	}
}

// GetMappingMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricNameEnum(val string) (MetricNameEnum, bool) {
	enum, ok := mappingMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
