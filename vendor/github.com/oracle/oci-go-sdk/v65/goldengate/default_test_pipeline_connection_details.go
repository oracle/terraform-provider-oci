// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DefaultTestPipelineConnectionDetails Additional attribute with which to test the pipeline's connection. The connectionId must be one of the pipeline's assigned connections.
type DefaultTestPipelineConnectionDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	ConnectionId *string `mandatory:"true" json:"connectionId"`
}

func (m DefaultTestPipelineConnectionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultTestPipelineConnectionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefaultTestPipelineConnectionDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultTestPipelineConnectionDetails DefaultTestPipelineConnectionDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDefaultTestPipelineConnectionDetails
	}{
		"DEFAULT",
		(MarshalTypeDefaultTestPipelineConnectionDetails)(m),
	}

	return json.Marshal(&s)
}
