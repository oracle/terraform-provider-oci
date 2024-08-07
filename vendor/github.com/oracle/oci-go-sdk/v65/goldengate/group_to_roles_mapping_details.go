// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GroupToRolesMappingDetails Defines the IDP Groups to GoldenGate roles mapping.
type GroupToRolesMappingDetails struct {

	// Grants administration of security related objects and invoke security related service requests.
	// This role has full privileges.
	SecurityGroup *string `mandatory:"false" json:"securityGroup"`

	// Grants full access to the user, including the ability to alter general, non-security related operational parameters
	// and profiles of the server.
	AdministratorGroup *string `mandatory:"false" json:"administratorGroup"`

	// Allows users to perform only operational actions, like starting and stopping resources.
	// Operators cannot alter the operational parameters or profiles of the MA server.
	OperatorGroup *string `mandatory:"false" json:"operatorGroup"`

	// Allows information-only service requests, which do not alter or affect the operation of either the MA.
	// Examples of query and read-only information include performance metric information and resource status and monitoring information
	UserGroup *string `mandatory:"false" json:"userGroup"`
}

func (m GroupToRolesMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupToRolesMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
