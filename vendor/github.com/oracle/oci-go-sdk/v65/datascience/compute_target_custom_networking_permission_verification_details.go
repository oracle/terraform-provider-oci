// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeTargetCustomNetworkingPermissionVerificationDetails Parameters required to validate custom networking permissions for compute target operations.
type ComputeTargetCustomNetworkingPermissionVerificationDetails struct {

	// Customer subnet OCID used for custom egress secondary VNIC operations.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The custom networking operation to verify against the compute target vRP policy model.
	OperationType ComputeTargetCustomNetworkingOperationTypeEnum `mandatory:"true" json:"operationType"`

	// Principal type for this workload/action.
	PrincipalType *string `mandatory:"true" json:"principalType"`

	// Principal id for validation.
	PrincipalId *string `mandatory:"true" json:"principalId"`

	// Tenancy id of the resource.
	ResourceTenancyId *string `mandatory:"true" json:"resourceTenancyId"`

	// Optional VNIC OCID context for detach operations.
	VnicId *string `mandatory:"false" json:"vnicId"`
}

func (m ComputeTargetCustomNetworkingPermissionVerificationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeTargetCustomNetworkingPermissionVerificationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComputeTargetCustomNetworkingOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetComputeTargetCustomNetworkingOperationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
