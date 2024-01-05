// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OperatorAccessControl API
//
// Operator Access Control enables you to control the time duration and the actions an Oracle operator can perform on your Exadata Cloud@Customer infrastructure.
// Using logging service, you can view a near real-time audit report of all actions performed by an Oracle operator.
// Use the table of contents and search tool to explore the OperatorAccessControl API.
//

package operatoraccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InteractionSummary Summary of access request customer and operator conversation.
type InteractionSummary struct {

	// The uniqueId of the message.
	Id *string `mandatory:"false" json:"id"`

	// customer or operator id who is part of this conversation.
	UserId *string `mandatory:"false" json:"userId"`

	// customer or operator Name who is part of this conversation.
	UserName *string `mandatory:"false" json:"userName"`

	// contains the information exchanged between operator and customer.
	Message *string `mandatory:"false" json:"message"`

	// Whether the userConversation is an operator or customer.
	UserType *string `mandatory:"false" json:"userType"`

	// Time when the conversation happened in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format. Example: '2020-05-22T21:10:29.600Z'
	TimeOfConversation *common.SDKTime `mandatory:"false" json:"timeOfConversation"`
}

func (m InteractionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InteractionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
