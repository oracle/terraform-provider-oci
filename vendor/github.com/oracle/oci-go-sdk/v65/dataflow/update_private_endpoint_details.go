// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdatePrivateEndpointDetails The details required to update a private endpoint.
type UpdatePrivateEndpointDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly description. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. It does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An array of DNS zone names.
	// Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]`
	DnsZones []string `mandatory:"false" json:"dnsZones"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The maximum number of hosts to be accessed through the private endpoint. This value is used
	// to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a
	// multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up
	// to 512.
	MaxHostCount *int `mandatory:"false" json:"maxHostCount"`

	// An array of network security group OCIDs.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// An array of fqdn/port pairs used to create private endpoint. Each object is a simple key-value pair with FQDN as key and port number as value.
	// [ { fqdn: "scan1.oracle.com", port: "1521"}, { fqdn: "scan2.oracle.com", port: "1521" } ]
	ScanDetails []Scan `mandatory:"false" json:"scanDetails"`
}

func (m UpdatePrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
