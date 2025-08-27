// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateScheduledQueryDetails Object that contains the details about the scheduled query to be created.
type CreateScheduledQueryDetails struct {

	// Name of the scheduled query.
	ScheduledQueryName *string `mandatory:"false" json:"scheduledQueryName"`

	// Type of the scheduled query.
	ScheduledQueryProcessingType ScheduledQueryProcessingTypeEnum `mandatory:"false" json:"scheduledQueryProcessingType,omitempty"`

	// Scheduled query to be run.
	ScheduledQueryText *string `mandatory:"false" json:"scheduledQueryText"`

	// Schedule for the scheduled query.
	ScheduledQuerySchedule *string `mandatory:"false" json:"scheduledQuerySchedule"`

	// Description for the scheduled query.
	ScheduledQueryDescription *string `mandatory:"false" json:"scheduledQueryDescription"`

	// Maximum runtime for the scheduled query in seconds.
	ScheduledQueryMaximumRuntimeInSeconds *int64 `mandatory:"false" json:"scheduledQueryMaximumRuntimeInSeconds"`

	// Retention period for the scheduled query in milliseconds.
	ScheduledQueryRetentionPeriodInMs *int64 `mandatory:"false" json:"scheduledQueryRetentionPeriodInMs"`

	// Processing sub type of the scheduled query.
	ScheduledQueryProcessingSubType ScheduledQueryProcessingSubTypeEnum `mandatory:"false" json:"scheduledQueryProcessingSubType,omitempty"`

	ScheduledQueryProcessingConfiguration *ScheduledQueryProcessingConfig `mandatory:"false" json:"scheduledQueryProcessingConfiguration"`

	// Retention criteria for the scheduled query.
	ScheduledQueryRetentionCriteria ScheduledQueryRetentionCriteriaEnum `mandatory:"false" json:"scheduledQueryRetentionCriteria,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateScheduledQueryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateScheduledQueryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScheduledQueryProcessingTypeEnum(string(m.ScheduledQueryProcessingType)); !ok && m.ScheduledQueryProcessingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledQueryProcessingType: %s. Supported values are: %s.", m.ScheduledQueryProcessingType, strings.Join(GetScheduledQueryProcessingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledQueryProcessingSubTypeEnum(string(m.ScheduledQueryProcessingSubType)); !ok && m.ScheduledQueryProcessingSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledQueryProcessingSubType: %s. Supported values are: %s.", m.ScheduledQueryProcessingSubType, strings.Join(GetScheduledQueryProcessingSubTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScheduledQueryRetentionCriteriaEnum(string(m.ScheduledQueryRetentionCriteria)); !ok && m.ScheduledQueryRetentionCriteria != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledQueryRetentionCriteria: %s. Supported values are: %s.", m.ScheduledQueryRetentionCriteria, strings.Join(GetScheduledQueryRetentionCriteriaEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
