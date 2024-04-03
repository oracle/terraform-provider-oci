// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GetWorkflowOrReplicationTargetDetails Response for getting the existence of workflow, or the replication target if workflow completes
type GetWorkflowOrReplicationTargetDetails struct {

	// The type of replication target workflow
	WorkflowType GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum `mandatory:"true" json:"workflowType"`

	// An indicator to tell if the replication target workflow is active
	IsWorkflowActive *bool `mandatory:"true" json:"isWorkflowActive"`

	ReplicationTarget *ReplicationTarget `mandatory:"false" json:"replicationTarget"`
}

func (m GetWorkflowOrReplicationTargetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GetWorkflowOrReplicationTargetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum(string(m.WorkflowType)); !ok && m.WorkflowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkflowType: %s. Supported values are: %s.", m.WorkflowType, strings.Join(GetGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum Enum with underlying type: string
type GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum string

// Set of constants representing the allowable values for GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum
const (
	GetWorkflowOrReplicationTargetDetailsWorkflowTypeCreate GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum = "CREATE"
	GetWorkflowOrReplicationTargetDetailsWorkflowTypeUpdate GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum = "UPDATE"
)

var mappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum = map[string]GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum{
	"CREATE": GetWorkflowOrReplicationTargetDetailsWorkflowTypeCreate,
	"UPDATE": GetWorkflowOrReplicationTargetDetailsWorkflowTypeUpdate,
}

var mappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumLowerCase = map[string]GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum{
	"create": GetWorkflowOrReplicationTargetDetailsWorkflowTypeCreate,
	"update": GetWorkflowOrReplicationTargetDetailsWorkflowTypeUpdate,
}

// GetGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumValues Enumerates the set of values for GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum
func GetGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumValues() []GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum {
	values := make([]GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum, 0)
	for _, v := range mappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumStringValues Enumerates the set of values in String for GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum
func GetGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumStringValues() []string {
	return []string{
		"CREATE",
		"UPDATE",
	}
}

// GetMappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum(val string) (GetWorkflowOrReplicationTargetDetailsWorkflowTypeEnum, bool) {
	enum, ok := mappingGetWorkflowOrReplicationTargetDetailsWorkflowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
