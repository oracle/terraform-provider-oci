// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAgentReceiverTaskDetails Request details for enabling/disabling the metric receiver on the management agent.
type UpdateAgentReceiverTaskDetails struct {

	// Management Agent Identifier OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AgentId *string `mandatory:"true" json:"agentId"`

	// True to enable the receiver and false to disable the receiver on the agent.
	IsEnable *bool `mandatory:"true" json:"isEnable"`

	ReceiverProperties *AgentReceiverProperties `mandatory:"false" json:"receiverProperties"`

	// Type of the handler.
	HandlerType HandlerTypeEnum `mandatory:"true" json:"handlerType"`
}

func (m UpdateAgentReceiverTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAgentReceiverTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHandlerTypeEnum(string(m.HandlerType)); !ok && m.HandlerType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HandlerType: %s. Supported values are: %s.", m.HandlerType, strings.Join(GetHandlerTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAgentReceiverTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAgentReceiverTaskDetails UpdateAgentReceiverTaskDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateAgentReceiverTaskDetails
	}{
		"UPDATE_AGENT_RECEIVER",
		(MarshalTypeUpdateAgentReceiverTaskDetails)(m),
	}

	return json.Marshal(&s)
}
