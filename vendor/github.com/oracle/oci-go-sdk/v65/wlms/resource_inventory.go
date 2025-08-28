// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// WebLogic Management Service API
//
// WebLogic Management Service is an OCI service that enables a unified view and management of WebLogic domains
// in Oracle Cloud Infrastructure. Features include on-demand patching of WebLogic domains, rollback of the
// last applied patch, discovery and management of WebLogic instances on a compute host.
//

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceInventory The inventory of WebLogic domains and managed instances in the selected compartment.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type ResourceInventory struct {

	// The number of WebLogic domains discovered in the selected compartment.
	WlsDomainCount *int `mandatory:"true" json:"wlsDomainCount"`

	// The number of managed instances in the selected compartment.
	ManagedInstanceCount *int `mandatory:"true" json:"managedInstanceCount"`

	// The number of WebLogic domains that are not in the latest patch.
	WlsDomainNotInLatestPatchCount *int `mandatory:"false" json:"wlsDomainNotInLatestPatchCount"`

	// The number of WebLogic domains that have warnings in the patch readiness check.
	WlsDomainWithPatchReadinessWarningCount *int `mandatory:"false" json:"wlsDomainWithPatchReadinessWarningCount"`

	// The number of WebLogic domains that have errors in the patch readiness check.
	WlsDomainWithPatchReadinessErrorCount *int `mandatory:"false" json:"wlsDomainWithPatchReadinessErrorCount"`

	// The number of managed instance that have WebLogic domains with warnings in the patch readiness check.
	ManagedInstanceWithPatchReadinessWarningCount *int `mandatory:"false" json:"managedInstanceWithPatchReadinessWarningCount"`

	// The number of managed instance that have WebLogic domains with errors in the patch readiness check.
	ManagedInstanceWithPatchReadinessErrorCount *int `mandatory:"false" json:"managedInstanceWithPatchReadinessErrorCount"`
}

func (m ResourceInventory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceInventory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
