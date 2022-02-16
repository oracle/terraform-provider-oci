// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateOperationsInsightsWarehouseUserDetails The information about a Operations Insights Warehouse User to be created. Input compartmentId MUST be the root compartment.
type CreateOperationsInsightsWarehouseUserDetails struct {

	// OPSI Warehouse OCID
	OperationsInsightsWarehouseId *string `mandatory:"true" json:"operationsInsightsWarehouseId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Username for schema which would have access to AWR Data,  Enterprise Manager Data and Operations Insights OPSI Hub.
	Name *string `mandatory:"true" json:"name"`

	// User provided connection password for the AWR Data,  Enterprise Manager Data and Operations Insights OPSI Hub.
	ConnectionPassword *string `mandatory:"true" json:"connectionPassword"`

	// Indicate whether user has access to AWR data.
	IsAwrDataAccess *bool `mandatory:"true" json:"isAwrDataAccess"`

	// Indicate whether user has access to EM data.
	IsEmDataAccess *bool `mandatory:"false" json:"isEmDataAccess"`

	// Indicate whether user has access to OPSI data.
	IsOpsiDataAccess *bool `mandatory:"false" json:"isOpsiDataAccess"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateOperationsInsightsWarehouseUserDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOperationsInsightsWarehouseUserDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
