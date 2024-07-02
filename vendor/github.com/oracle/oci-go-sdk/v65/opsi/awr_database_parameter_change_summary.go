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

// AwrDatabaseParameterChangeSummary A summary of the changes made to a single AWR database parameter.
type AwrDatabaseParameterChangeSummary struct {

	// The ID of the snapshot with the parameter value changed. The snapshot identifier is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /awrHubs/{awrHubId}/awrDatabaseSnapshots
	SnapshotIdentifier *int `mandatory:"true" json:"snapshotIdentifier"`

	// The start time of the interval.
	TimeBegin *common.SDKTime `mandatory:"false" json:"timeBegin"`

	// The end time of the interval.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The database instance number.
	InstanceNumber *int `mandatory:"false" json:"instanceNumber"`

	// The previous value of the database parameter.
	PreviousValue *string `mandatory:"false" json:"previousValue"`

	// The current value of the database parameter.
	Value *string `mandatory:"false" json:"value"`

	// Indicates whether the parameter has been modified after instance startup:
	//  - MODIFIED - Parameter has been modified with ALTER SESSION
	//  - SYSTEM_MOD - Parameter has been modified with ALTER SYSTEM (which causes all the currently logged in sessions values to be modified)
	//  - FALSE - Parameter has not been modified after instance startup
	ValueModified *string `mandatory:"false" json:"valueModified"`

	// Indicates whether the parameter value in the end snapshot is the default.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m AwrDatabaseParameterChangeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseParameterChangeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
