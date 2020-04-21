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

// CreateFaaSActionDetails Create an action that delivers to an Oracle Functions Service endpoint.
type CreateFaaSActionDetails struct {

	// Whether or not this action is currently enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering
	// confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Function hosted by Oracle Functions Service.
	FunctionId *string `mandatory:"false" json:"functionId"`
}

//GetIsEnabled returns IsEnabled
func (m CreateFaaSActionDetails) GetIsEnabled() *bool {
	return m.IsEnabled
}

//GetDescription returns Description
func (m CreateFaaSActionDetails) GetDescription() *string {
	return m.Description
}

func (m CreateFaaSActionDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateFaaSActionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateFaaSActionDetails CreateFaaSActionDetails
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeCreateFaaSActionDetails
	}{
		"FAAS",
		(MarshalTypeCreateFaaSActionDetails)(m),
	}

	return json.Marshal(&s)
}
