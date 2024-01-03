// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package operationsinsights

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataObjectCustomColumnUnit Unit details of a data object column of CUSTOM unit category.
type DataObjectCustomColumnUnit struct {

	// Display name of the column's unit.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Custom column unit.
	Unit *string `mandatory:"false" json:"unit"`
}

// GetDisplayName returns DisplayName
func (m DataObjectCustomColumnUnit) GetDisplayName() *string {
	return m.DisplayName
}

func (m DataObjectCustomColumnUnit) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectCustomColumnUnit) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataObjectCustomColumnUnit) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataObjectCustomColumnUnit DataObjectCustomColumnUnit
	s := struct {
		DiscriminatorParam string `json:"unitCategory"`
		MarshalTypeDataObjectCustomColumnUnit
	}{
		"CUSTOM",
		(MarshalTypeDataObjectCustomColumnUnit)(m),
	}

	return json.Marshal(&s)
}
