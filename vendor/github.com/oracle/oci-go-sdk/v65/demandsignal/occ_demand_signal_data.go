// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccDemandSignalData The Data Object For Demand Signal.
type OccDemandSignalData struct {

	// The name of the resource for the data.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The units of the data.
	Units *string `mandatory:"true" json:"units"`

	// The values of forecast.
	Values []OccDemandSignalValue `mandatory:"true" json:"values"`
}

func (m OccDemandSignalData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccDemandSignalData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
