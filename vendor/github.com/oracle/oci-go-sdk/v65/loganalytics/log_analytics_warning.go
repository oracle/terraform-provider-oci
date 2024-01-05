// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// LogAnalyticsWarning LogAnalyticsWarning
type LogAnalyticsWarning struct {

	// The unique identifier of the agent associated with the warning
	AgentId *string `mandatory:"false" json:"agentId"`

	// The host containing the agent associated with the warning
	HostName *string `mandatory:"false" json:"hostName"`

	// The display name of the rule which triggered the warning
	RuleDisplayName *string `mandatory:"false" json:"ruleDisplayName"`

	// The name of the source associated with the warning
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The entity compartment ID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The display name of the source associated with the warning
	SourceDisplayName *string `mandatory:"false" json:"sourceDisplayName"`

	// The name of the entity associated with the warning
	EntityName *string `mandatory:"false" json:"entityName"`

	// The time at which the warning was most recently collected
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// The unique identifier of the warning
	WarningId *string `mandatory:"false" json:"warningId"`

	// The date at which the warning was initially triggered
	TimeOfInitialWarning *common.SDKTime `mandatory:"false" json:"timeOfInitialWarning"`

	// A flag indicating if the warning is currently active
	IsActive *bool `mandatory:"false" json:"isActive"`

	// A flag indicating if the warning is currently suppressed
	IsSuppressed *bool `mandatory:"false" json:"isSuppressed"`

	// The most recent date on which the warning was triggered
	TimeOfLatestWarning *common.SDKTime `mandatory:"false" json:"timeOfLatestWarning"`

	// The warning level - either pattern, rule, or source.
	WarningLevel *string `mandatory:"false" json:"warningLevel"`

	// A description of the warning intended for the consumer of the warning.  It will
	// usually detail the cause of the warning, may suggest a remedy, and can contain any
	// other relevant information the consumer might find useful
	WarningMessage *string `mandatory:"false" json:"warningMessage"`

	// The unique identifier of the warning pattern
	PatternId *string `mandatory:"false" json:"patternId"`

	// The text of the pattern used by the warning
	PatternText *string `mandatory:"false" json:"patternText"`

	// The unique identifier of the rule associated with the warning
	RuleId *string `mandatory:"false" json:"ruleId"`

	// The unique identifier of the source associated with the warning
	SourceId *string `mandatory:"false" json:"sourceId"`

	// The user who suppressed the warning, or empty if the warning is not suppressed
	SuppressedBy *string `mandatory:"false" json:"suppressedBy"`

	// The unique identifier of the entity associated with the warning
	EntityId *string `mandatory:"false" json:"entityId"`

	// The type of the entity associated with the warning
	EntityType *string `mandatory:"false" json:"entityType"`

	// The display name of the entity type associated with the warning
	EntityTypeDisplayName *string `mandatory:"false" json:"entityTypeDisplayName"`

	// The display name of the warning type
	TypeDisplayName *string `mandatory:"false" json:"typeDisplayName"`

	// The internal name of the warning
	TypeName *string `mandatory:"false" json:"typeName"`

	// The warning severity
	Severity *int `mandatory:"false" json:"severity"`
}

func (m LogAnalyticsWarning) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsWarning) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
