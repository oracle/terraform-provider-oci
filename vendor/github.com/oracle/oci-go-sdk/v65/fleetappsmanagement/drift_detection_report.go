// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DriftDetectionReport The List Of Deployed Resources whose values have drifted from initial deployment
type DriftDetectionReport struct {

	// Has the deployed resource drifted from its initial state
	IsDrifted *bool `mandatory:"true" json:"isDrifted"`

	// Drift report file name stored in object storage
	DriftDetectionReportName *string `mandatory:"true" json:"driftDetectionReportName"`

	// Drift report file name stored in object storage
	DriftDetectionReportFileLink *string `mandatory:"true" json:"driftDetectionReportFileLink"`

	// The date and time when the Drift detection report was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeLastDriftReport *common.SDKTime `mandatory:"true" json:"timeLastDriftReport"`

	// A description of the provision.
	ProvisionDescription *string `mandatory:"true" json:"provisionDescription"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableCurrentUserId *string `mandatory:"true" json:"tfVariableCurrentUserId"`

	// An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
	TfVariableCompartmentId *string `mandatory:"true" json:"tfVariableCompartmentId"`
}

func (m DriftDetectionReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DriftDetectionReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
