// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociationSummary Summary of an association
type AssociationSummary struct {

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
	LifeCycleState AssociationSummaryLifeCycleStateEnum `mandatory:"false" json:"lifeCycleState,omitempty"`

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

	// A list of association properties.
	AssociationProperties []AssociationProperty `mandatory:"false" json:"associationProperties"`

	// The entity compartment OCID.
	EntityCompartmentId *string `mandatory:"false" json:"entityCompartmentId"`

	// The collection rule OCID.
	CollectionRuleId *string `mandatory:"false" json:"collectionRuleId"`
}

func (m AssociationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAssociationSummaryLifeCycleStateEnum(string(m.LifeCycleState)); !ok && m.LifeCycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifeCycleState: %s. Supported values are: %s.", m.LifeCycleState, strings.Join(GetAssociationSummaryLifeCycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociationSummaryLifeCycleStateEnum Enum with underlying type: string
type AssociationSummaryLifeCycleStateEnum string

// Set of constants representing the allowable values for AssociationSummaryLifeCycleStateEnum
const (
	AssociationSummaryLifeCycleStateAccepted   AssociationSummaryLifeCycleStateEnum = "ACCEPTED"
	AssociationSummaryLifeCycleStateInProgress AssociationSummaryLifeCycleStateEnum = "IN_PROGRESS"
	AssociationSummaryLifeCycleStateSucceeded  AssociationSummaryLifeCycleStateEnum = "SUCCEEDED"
	AssociationSummaryLifeCycleStateFailed     AssociationSummaryLifeCycleStateEnum = "FAILED"
)

var mappingAssociationSummaryLifeCycleStateEnum = map[string]AssociationSummaryLifeCycleStateEnum{
	"ACCEPTED":    AssociationSummaryLifeCycleStateAccepted,
	"IN_PROGRESS": AssociationSummaryLifeCycleStateInProgress,
	"SUCCEEDED":   AssociationSummaryLifeCycleStateSucceeded,
	"FAILED":      AssociationSummaryLifeCycleStateFailed,
}

var mappingAssociationSummaryLifeCycleStateEnumLowerCase = map[string]AssociationSummaryLifeCycleStateEnum{
	"accepted":    AssociationSummaryLifeCycleStateAccepted,
	"in_progress": AssociationSummaryLifeCycleStateInProgress,
	"succeeded":   AssociationSummaryLifeCycleStateSucceeded,
	"failed":      AssociationSummaryLifeCycleStateFailed,
}

// GetAssociationSummaryLifeCycleStateEnumValues Enumerates the set of values for AssociationSummaryLifeCycleStateEnum
func GetAssociationSummaryLifeCycleStateEnumValues() []AssociationSummaryLifeCycleStateEnum {
	values := make([]AssociationSummaryLifeCycleStateEnum, 0)
	for _, v := range mappingAssociationSummaryLifeCycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociationSummaryLifeCycleStateEnumStringValues Enumerates the set of values in String for AssociationSummaryLifeCycleStateEnum
func GetAssociationSummaryLifeCycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingAssociationSummaryLifeCycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociationSummaryLifeCycleStateEnum(val string) (AssociationSummaryLifeCycleStateEnum, bool) {
	enum, ok := mappingAssociationSummaryLifeCycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
