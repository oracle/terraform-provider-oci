// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LogAnalyticsAssociation LogAnalyticsAssociation
type LogAnalyticsAssociation struct {

	// The failure message.
	FailureMessage *string `mandatory:"false" json:"failureMessage"`

	// The agent unique identifier.
	AgentId *string `mandatory:"false" json:"agentId"`

	// The last attempt date.
	TimeLastAttempted *common.SDKTime `mandatory:"false" json:"timeLastAttempted"`

	// The number of times the association will be attempted
	// before failing.
	RetryCount *int64 `mandatory:"false" json:"retryCount"`

	// The source name.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The source display name.
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// The source type internal name.
	SourceTypeName *string `mandatory:"false" json:"sourceTypeName"`

	// The lifecycle status.  Valid values are ACCEPTED, IN_PROGRESS, SUCCEEDED
	// or FAILED.
	LifeCycleState LogAnalyticsAssociationLifeCycleStateEnum `mandatory:"false" json:"lifeCycleState,omitempty"`

	// The entity unique identifier.
	EntityId *string `mandatory:"false" json:"entityId"`

	// The entity name.
	EntityName *string `mandatory:"false" json:"entityName"`

	// The entity type internal name.
	EntityTypeName *string `mandatory:"false" json:"entityTypeName"`

	// The host name.
	Host *string `mandatory:"false" json:"host"`

	// The name of the entity which contains the agent.
	AgentEntityName *string `mandatory:"false" json:"agentEntityName"`

	// The entity type display name.
	EntityTypeDisplayName *string `mandatory:"false" json:"entityTypeDisplayName"`

	// The log group unique identifier.
	LogGroupId *string `mandatory:"false" json:"logGroupId"`

	// The log group name.
	LogGroupName *string `mandatory:"false" json:"logGroupName"`

	// The log group compartment.
	LogGroupCompartment *string `mandatory:"false" json:"logGroupCompartment"`
}

func (m LogAnalyticsAssociation) String() string {
	return common.PointerString(m)
}

// LogAnalyticsAssociationLifeCycleStateEnum Enum with underlying type: string
type LogAnalyticsAssociationLifeCycleStateEnum string

// Set of constants representing the allowable values for LogAnalyticsAssociationLifeCycleStateEnum
const (
	LogAnalyticsAssociationLifeCycleStateAccepted   LogAnalyticsAssociationLifeCycleStateEnum = "ACCEPTED"
	LogAnalyticsAssociationLifeCycleStateInProgress LogAnalyticsAssociationLifeCycleStateEnum = "IN_PROGRESS"
	LogAnalyticsAssociationLifeCycleStateSucceeded  LogAnalyticsAssociationLifeCycleStateEnum = "SUCCEEDED"
	LogAnalyticsAssociationLifeCycleStateFailed     LogAnalyticsAssociationLifeCycleStateEnum = "FAILED"
)

var mappingLogAnalyticsAssociationLifeCycleState = map[string]LogAnalyticsAssociationLifeCycleStateEnum{
	"ACCEPTED":    LogAnalyticsAssociationLifeCycleStateAccepted,
	"IN_PROGRESS": LogAnalyticsAssociationLifeCycleStateInProgress,
	"SUCCEEDED":   LogAnalyticsAssociationLifeCycleStateSucceeded,
	"FAILED":      LogAnalyticsAssociationLifeCycleStateFailed,
}

// GetLogAnalyticsAssociationLifeCycleStateEnumValues Enumerates the set of values for LogAnalyticsAssociationLifeCycleStateEnum
func GetLogAnalyticsAssociationLifeCycleStateEnumValues() []LogAnalyticsAssociationLifeCycleStateEnum {
	values := make([]LogAnalyticsAssociationLifeCycleStateEnum, 0)
	for _, v := range mappingLogAnalyticsAssociationLifeCycleState {
		values = append(values, v)
	}
	return values
}
