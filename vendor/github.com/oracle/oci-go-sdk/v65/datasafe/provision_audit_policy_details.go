// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ProvisionAuditPolicyDetails Details for audit policy provisioning.
type ProvisionAuditPolicyDetails struct {

	// The audit policy details for provisioning.
	ProvisionAuditConditions []ProvisionAuditConditions `mandatory:"true" json:"provisionAuditConditions"`

	// Option provided to users at the target to indicate whether the Data Safe service account has to be excluded while provisioning the audit policies.
	IsDataSafeServiceAccountExcluded *bool `mandatory:"false" json:"isDataSafeServiceAccountExcluded"`
}

func (m ProvisionAuditPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProvisionAuditPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
