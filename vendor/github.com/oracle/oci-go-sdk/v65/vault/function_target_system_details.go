// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FunctionTargetSystemDetails Details of the OCI function that vault secret connects to.
type FunctionTargetSystemDetails struct {

	// The unique identifier (OCID) of the OCI Functions that vault secret connects to.
	FunctionId *string `mandatory:"true" json:"functionId"`
}

func (m FunctionTargetSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionTargetSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FunctionTargetSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFunctionTargetSystemDetails FunctionTargetSystemDetails
	s := struct {
		DiscriminatorParam string `json:"targetSystemType"`
		MarshalTypeFunctionTargetSystemDetails
	}{
		"FUNCTION",
		(MarshalTypeFunctionTargetSystemDetails)(m),
	}

	return json.Marshal(&s)
}
