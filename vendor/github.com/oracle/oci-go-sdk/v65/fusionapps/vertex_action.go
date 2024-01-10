// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VertexAction Vertex update action
type VertexAction struct {

	// A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	// Unique identifier of the object that represents the action
	ReferenceKey *string `mandatory:"false" json:"referenceKey"`

	// patch that delivered the vertex update prerequisite
	Artifact *string `mandatory:"false" json:"artifact"`

	// A string that describes whether the change is applied hot or cold
	State ActionStateEnum `mandatory:"false" json:"state,omitempty"`
}

// GetReferenceKey returns ReferenceKey
func (m VertexAction) GetReferenceKey() *string {
	return m.ReferenceKey
}

// GetState returns State
func (m VertexAction) GetState() ActionStateEnum {
	return m.State
}

// GetDescription returns Description
func (m VertexAction) GetDescription() *string {
	return m.Description
}

func (m VertexAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VertexAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingActionStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetActionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m VertexAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVertexAction VertexAction
	s := struct {
		DiscriminatorParam string `json:"actionType"`
		MarshalTypeVertexAction
	}{
		"VERTEX",
		(MarshalTypeVertexAction)(m),
	}

	return json.Marshal(&s)
}
