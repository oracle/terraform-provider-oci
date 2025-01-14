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

// DefaultStopPipelineDetails Attribute details for a default pipeline stop.
type DefaultStopPipelineDetails struct {
}

func (m DefaultStopPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DefaultStopPipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DefaultStopPipelineDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultStopPipelineDetails DefaultStopPipelineDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDefaultStopPipelineDetails
	}{
		"DEFAULT",
		(MarshalTypeDefaultStopPipelineDetails)(m),
	}

	return json.Marshal(&s)
}
