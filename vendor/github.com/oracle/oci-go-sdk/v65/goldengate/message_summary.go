// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MessageSummary Deployment message Summary.
type MessageSummary struct {

	// The deployment Message Id.
	Id *string `mandatory:"true" json:"id"`

	// The deployment Message in plain text with optional HTML anchor tags.
	DeploymentMessage *string `mandatory:"true" json:"deploymentMessage"`

	// The deployment Message Status.
	DeploymentMessageStatus MessageSummaryDeploymentMessageStatusEnum `mandatory:"true" json:"deploymentMessageStatus"`
}

func (m MessageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MessageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMessageSummaryDeploymentMessageStatusEnum(string(m.DeploymentMessageStatus)); !ok && m.DeploymentMessageStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentMessageStatus: %s. Supported values are: %s.", m.DeploymentMessageStatus, strings.Join(GetMessageSummaryDeploymentMessageStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MessageSummaryDeploymentMessageStatusEnum Enum with underlying type: string
type MessageSummaryDeploymentMessageStatusEnum string

// Set of constants representing the allowable values for MessageSummaryDeploymentMessageStatusEnum
const (
	MessageSummaryDeploymentMessageStatusInfo    MessageSummaryDeploymentMessageStatusEnum = "INFO"
	MessageSummaryDeploymentMessageStatusWarning MessageSummaryDeploymentMessageStatusEnum = "WARNING"
	MessageSummaryDeploymentMessageStatusError   MessageSummaryDeploymentMessageStatusEnum = "ERROR"
)

var mappingMessageSummaryDeploymentMessageStatusEnum = map[string]MessageSummaryDeploymentMessageStatusEnum{
	"INFO":    MessageSummaryDeploymentMessageStatusInfo,
	"WARNING": MessageSummaryDeploymentMessageStatusWarning,
	"ERROR":   MessageSummaryDeploymentMessageStatusError,
}

var mappingMessageSummaryDeploymentMessageStatusEnumLowerCase = map[string]MessageSummaryDeploymentMessageStatusEnum{
	"info":    MessageSummaryDeploymentMessageStatusInfo,
	"warning": MessageSummaryDeploymentMessageStatusWarning,
	"error":   MessageSummaryDeploymentMessageStatusError,
}

// GetMessageSummaryDeploymentMessageStatusEnumValues Enumerates the set of values for MessageSummaryDeploymentMessageStatusEnum
func GetMessageSummaryDeploymentMessageStatusEnumValues() []MessageSummaryDeploymentMessageStatusEnum {
	values := make([]MessageSummaryDeploymentMessageStatusEnum, 0)
	for _, v := range mappingMessageSummaryDeploymentMessageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMessageSummaryDeploymentMessageStatusEnumStringValues Enumerates the set of values in String for MessageSummaryDeploymentMessageStatusEnum
func GetMessageSummaryDeploymentMessageStatusEnumStringValues() []string {
	return []string{
		"INFO",
		"WARNING",
		"ERROR",
	}
}

// GetMappingMessageSummaryDeploymentMessageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMessageSummaryDeploymentMessageStatusEnum(val string) (MessageSummaryDeploymentMessageStatusEnum, bool) {
	enum, ok := mappingMessageSummaryDeploymentMessageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
