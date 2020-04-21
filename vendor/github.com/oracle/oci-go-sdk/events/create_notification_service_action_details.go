// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateNotificationServiceActionDetails Create an action that delivers to an Oracle Notification Service topic.
type CreateNotificationServiceActionDetails struct {

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the topic to which messages are delivered.
	TopicId *string `mandatory:"false" json:"topicId"`
}

//GetIsEnabled returns IsEnabled
func (m CreateNotificationServiceActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m CreateNotificationServiceActionDetails) GetDescription() *string {
	return m.Description
}

func (m CreateNotificationServiceActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateNotificationServiceActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateNotificationServiceActionDetails CreateNotificationServiceActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateNotificationServiceActionDetails
	}{
		"ONS",
		(MarshalTypeCreateNotificationServiceActionDetails)(m),
	}

	return json.Marshal(&s)
}
