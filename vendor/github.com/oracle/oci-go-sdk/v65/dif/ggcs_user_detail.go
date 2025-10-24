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

// GgcsUserDetail Ggcs user credentials to be created or updated.
type GgcsUserDetail struct {

	// username for the user.
	UserName *string `mandatory:"true" json:"userName"`

	// Vault OCID containing password for existing or new user.
	SecretId *string `mandatory:"true" json:"secretId"`

	// Type of GoldenGate user. Allowed values are "OPERATOR".
	UserType *string `mandatory:"true" json:"userType"`

	// Action to be done over the user. Allowed values are "CREATE" or "UPDATE".
	Action WorkflowActionEnum `mandatory:"true" json:"action"`
}

func (m GgcsUserDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GgcsUserDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkflowActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetWorkflowActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
