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

// DisassociateMonitoredResourcesDetails The information required to create new monitored resource association.
type DisassociateMonitoredResourcesDetails struct {

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Association type between source and destination resources.
	AssociationType *string `mandatory:"false" json:"associationType"`

	// Source Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SourceResourceId *string `mandatory:"false" json:"sourceResourceId"`

	// Destination Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DestinationResourceId *string `mandatory:"false" json:"destinationResourceId"`
}

func (m DisassociateMonitoredResourcesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisassociateMonitoredResourcesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
