// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GgcsSourceDetail GGCS source configuration for creating or updating existing extracts.
type GgcsSourceDetail struct {

	// Ggcs source artifact id.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// Action to be done over the user. Allowed values are "CREATE" or "UPDATE".
	Action WorkflowActionEnum `mandatory:"true" json:"action"`

	// Boolean value that determines source operations should start or not.
	ShouldStartSourceOperations *bool `mandatory:"true" json:"shouldStartSourceOperations"`

	// Target uri for the GoldenGate deployment where distribution path needs to be configured.
	TargetUri *string `mandatory:"false" json:"targetUri"`

	// Name of assigned connection for the source.
	TargetConnectionName *string `mandatory:"false" json:"targetConnectionName"`
}

func (m GgcsSourceDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GgcsSourceDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkflowActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetWorkflowActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
