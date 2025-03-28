// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AgentEventData Provides additional information for an agent event.
type AgentEventData struct {

	// Type of agent operation.
	OperationType AgentEventDataOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Status of the agent operation.
	Status EventStatusEnum `mandatory:"true" json:"status"`

	AdditionalDetails *WorkRequestEventDataAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m AgentEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgentEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAgentEventDataOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetAgentEventDataOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEventStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetEventStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AgentEventDataOperationTypeEnum Enum with underlying type: string
type AgentEventDataOperationTypeEnum string

// Set of constants representing the allowable values for AgentEventDataOperationTypeEnum
const (
	AgentEventDataOperationTypeListPackages    AgentEventDataOperationTypeEnum = "LIST_PACKAGES"
	AgentEventDataOperationTypeUploadContent   AgentEventDataOperationTypeEnum = "UPLOAD_CONTENT"
	AgentEventDataOperationTypeSyncAgentConfig AgentEventDataOperationTypeEnum = "SYNC_AGENT_CONFIG"
)

var mappingAgentEventDataOperationTypeEnum = map[string]AgentEventDataOperationTypeEnum{
	"LIST_PACKAGES":     AgentEventDataOperationTypeListPackages,
	"UPLOAD_CONTENT":    AgentEventDataOperationTypeUploadContent,
	"SYNC_AGENT_CONFIG": AgentEventDataOperationTypeSyncAgentConfig,
}

var mappingAgentEventDataOperationTypeEnumLowerCase = map[string]AgentEventDataOperationTypeEnum{
	"list_packages":     AgentEventDataOperationTypeListPackages,
	"upload_content":    AgentEventDataOperationTypeUploadContent,
	"sync_agent_config": AgentEventDataOperationTypeSyncAgentConfig,
}

// GetAgentEventDataOperationTypeEnumValues Enumerates the set of values for AgentEventDataOperationTypeEnum
func GetAgentEventDataOperationTypeEnumValues() []AgentEventDataOperationTypeEnum {
	values := make([]AgentEventDataOperationTypeEnum, 0)
	for _, v := range mappingAgentEventDataOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAgentEventDataOperationTypeEnumStringValues Enumerates the set of values in String for AgentEventDataOperationTypeEnum
func GetAgentEventDataOperationTypeEnumStringValues() []string {
	return []string{
		"LIST_PACKAGES",
		"UPLOAD_CONTENT",
		"SYNC_AGENT_CONFIG",
	}
}

// GetMappingAgentEventDataOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgentEventDataOperationTypeEnum(val string) (AgentEventDataOperationTypeEnum, bool) {
	enum, ok := mappingAgentEventDataOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
