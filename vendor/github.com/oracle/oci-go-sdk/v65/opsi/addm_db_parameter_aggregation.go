// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddmDbParameterAggregation Summarizes change history for specific database parameter
type AddmDbParameterAggregation struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database insight.
	Id *string `mandatory:"true" json:"id"`

	// Name of  parameter
	Name *string `mandatory:"true" json:"name"`

	// Indicates whether the parameter's value changed during the selected time range (TRUE) or
	// did not change during the selected time range (FALSE)
	IsChanged *bool `mandatory:"true" json:"isChanged"`

	// Number of database instance
	InstNum *int `mandatory:"false" json:"instNum"`

	// Parameter default value
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// Parameter value when time period began
	BeginValue *string `mandatory:"false" json:"beginValue"`

	// Parameter value when time period ended
	EndValue *string `mandatory:"false" json:"endValue"`

	// Indicates whether the parameter's end value was set to the default value (TRUE) or was
	// specified in the parameter file (FALSE)
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Indicates whether the parameter has been modified after instance starup
	// MODIFIED - Parameter has been modified with ALTER SESSION
	// SYSTEM_MOD - Parameter has been modified with ALTER SYSTEM
	// FALSE - Parameter has not been modified after instance starup
	ValueModified *string `mandatory:"false" json:"valueModified"`

	// Indicates whether the parameter is a high impact parameter (TRUE) or not (FALSE)
	IsHighImpact *bool `mandatory:"false" json:"isHighImpact"`
}

func (m AddmDbParameterAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbParameterAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
