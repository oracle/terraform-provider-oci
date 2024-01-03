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

// AssociatedResourceSummary Summary information for a resource associated with a stack or job.
type AssociatedResourceSummary struct {

	// Unique identifier for the resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Name of the resource.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Resource type. For more information about resource types supported for the Oracle Cloud Infrastructure (OCI) provider, see Oracle Cloud Infrastructure Provider (https://registry.terraform.io/providers/oracle/oci/latest/docs).
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Resource attribute values. Each value is represented as a key-value pair.
	// Example: `{"state": "AVAILABLE"}`
	Attributes map[string]string `mandatory:"false" json:"attributes"`

	// The date and time when the stack was created.
	// Format is defined by RFC3339.
	// Example: `2022-07-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Resource region.
	// For information about regions, see Regions and Availability Domains (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
	// Example: `us-phoenix-1`
	Region *string `mandatory:"false" json:"region"`

	// Terraform resource address.
	ResourceAddress *string `mandatory:"false" json:"resourceAddress"`
}

func (m AssociatedResourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedResourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
