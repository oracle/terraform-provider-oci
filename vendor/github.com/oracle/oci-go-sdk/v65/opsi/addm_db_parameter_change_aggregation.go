// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddmDbParameterChangeAggregation Change record for AWR database parameter
type AddmDbParameterChangeAggregation struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database insight.
	Id *string `mandatory:"true" json:"id"`

	// Begin time of interval which includes change
	TimeBegin *common.SDKTime `mandatory:"true" json:"timeBegin"`

	// End time of interval which includes change
	TimeEnd *common.SDKTime `mandatory:"true" json:"timeEnd"`

	// Instance number
	InstNum *int `mandatory:"true" json:"instNum"`

	// AWR snapshot id which includes the parameter value change
	SnapshotId *int `mandatory:"true" json:"snapshotId"`

	// Previous value
	PreviousValue *string `mandatory:"false" json:"previousValue"`

	// Current value
	Value *string `mandatory:"false" json:"value"`
}

func (m AddmDbParameterChangeAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddmDbParameterChangeAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
