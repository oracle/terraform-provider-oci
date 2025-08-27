// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateProvisionDetails The data to create a FamProvision.
type CreateProvisionDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the FamProvision in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item.
	PackageCatalogItemId *string `mandatory:"true" json:"packageCatalogItemId"`

	// A OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item to a file with key/value pairs to set up variables for createStack API.
	ConfigCatalogItemId *string `mandatory:"true" json:"configCatalogItemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" json:"fleetId"`

	// A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableTenancyId *string `mandatory:"true" json:"tfVariableTenancyId"`

	// A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableRegionId *string `mandatory:"true" json:"tfVariableRegionId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of the provision.
	ProvisionDescription *string `mandatory:"false" json:"provisionDescription"`

	// An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableCurrentUserId *string `mandatory:"false" json:"tfVariableCurrentUserId"`

	// An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableCompartmentId *string `mandatory:"false" json:"tfVariableCompartmentId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateProvisionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateProvisionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
