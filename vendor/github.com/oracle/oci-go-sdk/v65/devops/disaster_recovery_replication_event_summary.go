// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.oracle.com/iaas/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DisasterRecoveryReplicationEventSummary Summary of the Repository Replication Events.
type DisasterRecoveryReplicationEventSummary struct {

	// Unique identifier representing the Event ID
	ReplicationEventId *string `mandatory:"true" json:"replicationEventId"`

	// Unique identifier that is immutable on creation.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Type of the resource
	ResourceType DisasterRecoveryReplicationEventSummaryResourceTypeEnum `mandatory:"false" json:"resourceType,omitempty"`

	// Specifies type of Action performed on the resource.
	ActionType DisasterRecoveryReplicationEventSummaryActionTypeEnum `mandatory:"false" json:"actionType,omitempty"`

	// Time the event was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Additional Data related to the event (i.e. compartmentId)
	EventData *string `mandatory:"false" json:"eventData"`
}

func (m DisasterRecoveryReplicationEventSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryReplicationEventSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDisasterRecoveryReplicationEventSummaryResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetDisasterRecoveryReplicationEventSummaryResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDisasterRecoveryReplicationEventSummaryActionTypeEnum(string(m.ActionType)); !ok && m.ActionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ActionType: %s. Supported values are: %s.", m.ActionType, strings.Join(GetDisasterRecoveryReplicationEventSummaryActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisasterRecoveryReplicationEventSummaryResourceTypeEnum Enum with underlying type: string
type DisasterRecoveryReplicationEventSummaryResourceTypeEnum string

// Set of constants representing the allowable values for DisasterRecoveryReplicationEventSummaryResourceTypeEnum
const (
	DisasterRecoveryReplicationEventSummaryResourceTypeRepository         DisasterRecoveryReplicationEventSummaryResourceTypeEnum = "REPOSITORY"
	DisasterRecoveryReplicationEventSummaryResourceTypePullRequest        DisasterRecoveryReplicationEventSummaryResourceTypeEnum = "PULL_REQUEST"
	DisasterRecoveryReplicationEventSummaryResourceTypeRepositorySettings DisasterRecoveryReplicationEventSummaryResourceTypeEnum = "REPOSITORY_SETTINGS"
	DisasterRecoveryReplicationEventSummaryResourceTypeProjectSettings    DisasterRecoveryReplicationEventSummaryResourceTypeEnum = "PROJECT_SETTINGS"
)

var mappingDisasterRecoveryReplicationEventSummaryResourceTypeEnum = map[string]DisasterRecoveryReplicationEventSummaryResourceTypeEnum{
	"REPOSITORY":          DisasterRecoveryReplicationEventSummaryResourceTypeRepository,
	"PULL_REQUEST":        DisasterRecoveryReplicationEventSummaryResourceTypePullRequest,
	"REPOSITORY_SETTINGS": DisasterRecoveryReplicationEventSummaryResourceTypeRepositorySettings,
	"PROJECT_SETTINGS":    DisasterRecoveryReplicationEventSummaryResourceTypeProjectSettings,
}

var mappingDisasterRecoveryReplicationEventSummaryResourceTypeEnumLowerCase = map[string]DisasterRecoveryReplicationEventSummaryResourceTypeEnum{
	"repository":          DisasterRecoveryReplicationEventSummaryResourceTypeRepository,
	"pull_request":        DisasterRecoveryReplicationEventSummaryResourceTypePullRequest,
	"repository_settings": DisasterRecoveryReplicationEventSummaryResourceTypeRepositorySettings,
	"project_settings":    DisasterRecoveryReplicationEventSummaryResourceTypeProjectSettings,
}

// GetDisasterRecoveryReplicationEventSummaryResourceTypeEnumValues Enumerates the set of values for DisasterRecoveryReplicationEventSummaryResourceTypeEnum
func GetDisasterRecoveryReplicationEventSummaryResourceTypeEnumValues() []DisasterRecoveryReplicationEventSummaryResourceTypeEnum {
	values := make([]DisasterRecoveryReplicationEventSummaryResourceTypeEnum, 0)
	for _, v := range mappingDisasterRecoveryReplicationEventSummaryResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryReplicationEventSummaryResourceTypeEnumStringValues Enumerates the set of values in String for DisasterRecoveryReplicationEventSummaryResourceTypeEnum
func GetDisasterRecoveryReplicationEventSummaryResourceTypeEnumStringValues() []string {
	return []string{
		"REPOSITORY",
		"PULL_REQUEST",
		"REPOSITORY_SETTINGS",
		"PROJECT_SETTINGS",
	}
}

// GetMappingDisasterRecoveryReplicationEventSummaryResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryReplicationEventSummaryResourceTypeEnum(val string) (DisasterRecoveryReplicationEventSummaryResourceTypeEnum, bool) {
	enum, ok := mappingDisasterRecoveryReplicationEventSummaryResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DisasterRecoveryReplicationEventSummaryActionTypeEnum Enum with underlying type: string
type DisasterRecoveryReplicationEventSummaryActionTypeEnum string

// Set of constants representing the allowable values for DisasterRecoveryReplicationEventSummaryActionTypeEnum
const (
	DisasterRecoveryReplicationEventSummaryActionTypeDelete          DisasterRecoveryReplicationEventSummaryActionTypeEnum = "DELETE"
	DisasterRecoveryReplicationEventSummaryActionTypeGitPush         DisasterRecoveryReplicationEventSummaryActionTypeEnum = "GIT_PUSH"
	DisasterRecoveryReplicationEventSummaryActionTypeMerge           DisasterRecoveryReplicationEventSummaryActionTypeEnum = "MERGE"
	DisasterRecoveryReplicationEventSummaryActionTypeCreate          DisasterRecoveryReplicationEventSummaryActionTypeEnum = "CREATE"
	DisasterRecoveryReplicationEventSummaryActionTypeUpdate          DisasterRecoveryReplicationEventSummaryActionTypeEnum = "UPDATE"
	DisasterRecoveryReplicationEventSummaryActionTypeRename          DisasterRecoveryReplicationEventSummaryActionTypeEnum = "RENAME"
	DisasterRecoveryReplicationEventSummaryActionTypeMoveCompartment DisasterRecoveryReplicationEventSummaryActionTypeEnum = "MOVE_COMPARTMENT"
)

var mappingDisasterRecoveryReplicationEventSummaryActionTypeEnum = map[string]DisasterRecoveryReplicationEventSummaryActionTypeEnum{
	"DELETE":           DisasterRecoveryReplicationEventSummaryActionTypeDelete,
	"GIT_PUSH":         DisasterRecoveryReplicationEventSummaryActionTypeGitPush,
	"MERGE":            DisasterRecoveryReplicationEventSummaryActionTypeMerge,
	"CREATE":           DisasterRecoveryReplicationEventSummaryActionTypeCreate,
	"UPDATE":           DisasterRecoveryReplicationEventSummaryActionTypeUpdate,
	"RENAME":           DisasterRecoveryReplicationEventSummaryActionTypeRename,
	"MOVE_COMPARTMENT": DisasterRecoveryReplicationEventSummaryActionTypeMoveCompartment,
}

var mappingDisasterRecoveryReplicationEventSummaryActionTypeEnumLowerCase = map[string]DisasterRecoveryReplicationEventSummaryActionTypeEnum{
	"delete":           DisasterRecoveryReplicationEventSummaryActionTypeDelete,
	"git_push":         DisasterRecoveryReplicationEventSummaryActionTypeGitPush,
	"merge":            DisasterRecoveryReplicationEventSummaryActionTypeMerge,
	"create":           DisasterRecoveryReplicationEventSummaryActionTypeCreate,
	"update":           DisasterRecoveryReplicationEventSummaryActionTypeUpdate,
	"rename":           DisasterRecoveryReplicationEventSummaryActionTypeRename,
	"move_compartment": DisasterRecoveryReplicationEventSummaryActionTypeMoveCompartment,
}

// GetDisasterRecoveryReplicationEventSummaryActionTypeEnumValues Enumerates the set of values for DisasterRecoveryReplicationEventSummaryActionTypeEnum
func GetDisasterRecoveryReplicationEventSummaryActionTypeEnumValues() []DisasterRecoveryReplicationEventSummaryActionTypeEnum {
	values := make([]DisasterRecoveryReplicationEventSummaryActionTypeEnum, 0)
	for _, v := range mappingDisasterRecoveryReplicationEventSummaryActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryReplicationEventSummaryActionTypeEnumStringValues Enumerates the set of values in String for DisasterRecoveryReplicationEventSummaryActionTypeEnum
func GetDisasterRecoveryReplicationEventSummaryActionTypeEnumStringValues() []string {
	return []string{
		"DELETE",
		"GIT_PUSH",
		"MERGE",
		"CREATE",
		"UPDATE",
		"RENAME",
		"MOVE_COMPARTMENT",
	}
}

// GetMappingDisasterRecoveryReplicationEventSummaryActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryReplicationEventSummaryActionTypeEnum(val string) (DisasterRecoveryReplicationEventSummaryActionTypeEnum, bool) {
	enum, ok := mappingDisasterRecoveryReplicationEventSummaryActionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
