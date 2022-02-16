// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InstanceAgentCommandSourceViaObjectStorageUriDetails The source of the command when imported from an Object Storage URL.
type InstanceAgentCommandSourceViaObjectStorageUriDetails struct {

	// The Object Storage URL or pre-authenticated request (PAR) for the command.
	SourceUri *string `mandatory:"true" json:"sourceUri"`
}

func (m InstanceAgentCommandSourceViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentCommandSourceViaObjectStorageUriDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandSourceViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandSourceViaObjectStorageUriDetails InstanceAgentCommandSourceViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceAgentCommandSourceViaObjectStorageUriDetails
	}{
		"OBJECT_STORAGE_URI",
		(MarshalTypeInstanceAgentCommandSourceViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
