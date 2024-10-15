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

// UpdateGroupToRolesMappingDetails Defines the IDP Groups to GoldenGate roles mapping. This field is used only for IAM deployment and does not have any impact on non-IAM deployments.
// For IAM deployment, when user does not specify this mapping, then it has null value and default mapping is used.
// User belonging to each group can only perform the actions according to the role the respective group is mapped to.
type UpdateGroupToRolesMappingDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role securityGroup.
	// It grants administration of security related objects and invoke security related service requests. This role has full privileges.
	SecurityGroupId *string `mandatory:"false" json:"securityGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role administratorGroup.
	// It grants full access to the user, including the ability to alter general, non-security related operational parameters
	// and profiles of the server.
	AdministratorGroupId *string `mandatory:"false" json:"administratorGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role operatorGroup.
	// It allows users to perform only operational actions, like starting and stopping resources.
	// Operators cannot alter the operational parameters or profiles of the MA server.
	OperatorGroupId *string `mandatory:"false" json:"operatorGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the IDP group which will be mapped to goldengate role userGroup.
	// It allows information-only service requests, which do not alter or affect the operation of either the MA.
	// Examples of query and read-only information include performance metric information and resource status and monitoring information
	UserGroupId *string `mandatory:"false" json:"userGroupId"`
}

func (m UpdateGroupToRolesMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateGroupToRolesMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
