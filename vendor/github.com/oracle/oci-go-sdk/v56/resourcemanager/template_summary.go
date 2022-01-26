// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TemplateSummary Summary information for a template.
type TemplateSummary struct {

	// Unique identifier of the specified template.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this template.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable display name for the template.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Brief description of the template.
	Description *string `mandatory:"false" json:"description"`

	// whether the template will work for free tier tenancy.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The date and time at which the template was created.
	// Format is defined by RFC3339.
	// Example: `2020-11-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current lifecycle state of the template.
	// Allowable values:
	// - ACTIVE
	LifecycleState TemplateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TemplateSummary) String() string {
	return common.PointerString(m)
}
