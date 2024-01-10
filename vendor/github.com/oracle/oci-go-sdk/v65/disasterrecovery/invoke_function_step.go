// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InvokeFunctionStep Invoke Oracle function step details.
type InvokeFunctionStep struct {

	// The OCID of function to be invoked.
	// Example: `ocid1.fnfunc.oc1..uniqueID`
	FunctionId *string `mandatory:"true" json:"functionId"`

	// The region in which the function is deployed.
	// Example: `us-ashburn-1`
	FunctionRegion *string `mandatory:"true" json:"functionRegion"`

	// The request body for the function.
	// Example: `{ "FnParam1", "FnParam2" }`
	RequestBody *string `mandatory:"false" json:"requestBody"`
}

func (m InvokeFunctionStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InvokeFunctionStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InvokeFunctionStep) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInvokeFunctionStep InvokeFunctionStep
	s := struct {
		DiscriminatorParam string `json:"stepType"`
		MarshalTypeInvokeFunctionStep
	}{
		"INVOKE_FUNCTION",
		(MarshalTypeInvokeFunctionStep)(m),
	}

	return json.Marshal(&s)
}
