// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityPolicyEntryStateSummary The resource represents the state of a specific entry type deployment on a target.
type SecurityPolicyEntryStateSummary struct {

	// Unique id of the security policy entry state.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the security policy entry associated.
	SecurityPolicyEntryId *string `mandatory:"true" json:"securityPolicyEntryId"`

	// The current deployment status of the security policy deployment and the security policy entry associated.
	DeploymentStatus SecurityPolicyEntryStateDeploymentStatusEnum `mandatory:"true" json:"deploymentStatus"`

	// The OCID of the security policy deployment associated.
	SecurityPolicyDeploymentId *string `mandatory:"false" json:"securityPolicyDeploymentId"`
}

func (m SecurityPolicyEntryStateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityPolicyEntryStateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityPolicyEntryStateDeploymentStatusEnum(string(m.DeploymentStatus)); !ok && m.DeploymentStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeploymentStatus: %s. Supported values are: %s.", m.DeploymentStatus, strings.Join(GetSecurityPolicyEntryStateDeploymentStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
