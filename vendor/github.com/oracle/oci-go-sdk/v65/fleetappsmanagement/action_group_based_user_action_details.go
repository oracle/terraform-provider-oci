// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActionGroupBasedUserActionDetails Details for a user action to be performed on an action group.
type ActionGroupBasedUserActionDetails struct {

	// Unique identifier for the action group.
	ActionGroupId *string `mandatory:"true" json:"actionGroupId"`

	// Action to be Performed.
	Action UserActionDetailsActionEnum `mandatory:"true" json:"action"`
}

// GetAction returns Action
func (m ActionGroupBasedUserActionDetails) GetAction() UserActionDetailsActionEnum {
	return m.Action
}

func (m ActionGroupBasedUserActionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActionGroupBasedUserActionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUserActionDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetUserActionDetailsActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ActionGroupBasedUserActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeActionGroupBasedUserActionDetails ActionGroupBasedUserActionDetails
	s := struct {
		DiscriminatorParam string `json:"level"`
		MarshalTypeActionGroupBasedUserActionDetails
	}{
		"ACTION_GROUP",
		(MarshalTypeActionGroupBasedUserActionDetails)(m),
	}

	return json.Marshal(&s)
}
