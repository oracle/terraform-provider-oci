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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DbParameters Initialization parameters for a database.
type DbParameters struct {

	// Database instance number.
	InstanceNumber *int `mandatory:"true" json:"instanceNumber"`

	// Database parameter name.
	ParameterName *string `mandatory:"true" json:"parameterName"`

	// Database parameter value.
	ParameterValue *string `mandatory:"true" json:"parameterValue"`

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`

	// AWR snapshot id for the parameter value
	SnapshotId *int `mandatory:"false" json:"snapshotId"`

	// Indicates whether the parameter's value changed in given snapshot or not.
	IsChanged *string `mandatory:"false" json:"isChanged"`

	// Indicates whether this value is the default value or not.
	IsDefault *string `mandatory:"false" json:"isDefault"`
}

// GetTimeCollected returns TimeCollected
func (m DbParameters) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m DbParameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbParameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DbParameters) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDbParameters DbParameters
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeDbParameters
	}{
		"DB_PARAMETERS",
		(MarshalTypeDbParameters)(m),
	}

	return json.Marshal(&s)
}
