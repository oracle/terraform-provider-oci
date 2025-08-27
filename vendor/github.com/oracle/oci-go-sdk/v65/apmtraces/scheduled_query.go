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

// ScheduledQuery Scheduled Query object.
type ScheduledQuery struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the scheduled query . An OCID is generated
	// when the scheduled query is created.
	Id *string `mandatory:"false" json:"id"`

	// Processing type of the scheduled query.
	ScheduledQueryProcessingType ScheduledQueryProcessingTypeEnum `mandatory:"false" json:"scheduledQueryProcessingType,omitempty"`

	// Name of the scheduled query.
	ScheduledQueryName *string `mandatory:"false" json:"scheduledQueryName"`

	// Scheduled query to be run.
	ScheduledQueryText *string `mandatory:"false" json:"scheduledQueryText"`

	// Description for the scheduled query.
	ScheduledQueryDescription *string `mandatory:"false" json:"scheduledQueryDescription"`

	// Schedule for the scheduled query.
	ScheduledQuerySchedule *string `mandatory:"false" json:"scheduledQuerySchedule"`

	// Maximum runtime for the scheduled query in seconds.
	ScheduledQueryMaximumRuntimeInSeconds *int64 `mandatory:"false" json:"scheduledQueryMaximumRuntimeInSeconds"`

	// Next run for the scheduled query.
	ScheduledQueryNextRunInMs *int64 `mandatory:"false" json:"scheduledQueryNextRunInMs"`

	// Retention period for the scheduled query in milliseconds.
	ScheduledQueryRetentionPeriodInMs *int64 `mandatory:"false" json:"scheduledQueryRetentionPeriodInMs"`

	// Processing sub type of the scheduled query.
	ScheduledQueryProcessingSubType ScheduledQueryProcessingSubTypeEnum `mandatory:"false" json:"scheduledQueryProcessingSubType,omitempty"`

	ScheduledQueryProcessingConfiguration *ScheduledQueryProcessingConfig `mandatory:"false" json:"scheduledQueryProcessingConfiguration"`

	// Retention criteria for the scheduled query.
	ScheduledQueryRetentionCriteria ScheduledQueryRetentionCriteriaEnum `mandatory:"false" json:"scheduledQueryRetentionCriteria,omitempty"`

	// Scheduled query instances.
	ScheduledQueryInstances *string `mandatory:"false" json:"scheduledQueryInstances"`

	// The current lifecycle state of the Scheduled Query.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ScheduledQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledQuery) ValidateEnumValue() (bool, error) {
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
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
