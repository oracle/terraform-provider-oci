// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateBaselineableMetricDetails Summary for the baseline-able metric
type CreateBaselineableMetricDetails struct {

	// OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// name of the metric
	Name *string `mandatory:"true" json:"name"`

	// metric column name
	Column *string `mandatory:"true" json:"column"`

	// namespace of the metric
	Namespace *string `mandatory:"true" json:"namespace"`

	// Resource group of the metric
	ResourceGroup *string `mandatory:"true" json:"resourceGroup"`
}

func (m CreateBaselineableMetricDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBaselineableMetricDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
