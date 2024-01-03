// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ApplyJobPlanResolution Deprecated. Use the property `executionPlanStrategy` in `jobOperationDetails` instead.
type ApplyJobPlanResolution struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that specifies the most recently executed plan job.
	PlanJobId *string `mandatory:"false" json:"planJobId"`

	// Specifies whether to use the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the most recently run plan job.
	// `True` if using the latest job OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a plan job that completed successfully.
	IsUseLatestJobId *bool `mandatory:"false" json:"isUseLatestJobId"`

	// Specifies whether to use the configuration directly, without reference to a Plan job.
	// `True` if using the configuration directly. Note that it is not necessary
	// for a Plan job to have run successfully.
	IsAutoApproved *bool `mandatory:"false" json:"isAutoApproved"`
}

func (m ApplyJobPlanResolution) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ApplyJobPlanResolution) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
