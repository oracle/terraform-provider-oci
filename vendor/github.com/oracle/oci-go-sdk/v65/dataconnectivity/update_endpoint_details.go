// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateEndpointDetails The information to be updated.
type UpdateEndpointDetails struct {

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists only for cross-compatibility.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Data Connectivity Management Registry description
	Description *string `mandatory:"false" json:"description"`

	// The Data Connectivity Management registry display name; registries can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Update endpoint size for reverse connection capacity.
	EndpointSize *int `mandatory:"false" json:"endpointSize"`

	// The list of NSGs to which the Private Endpoint VNIC must be added.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// List of DNS zones to be used by the data assets.
	// Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com
	DnsZones []string `mandatory:"false" json:"dnsZones"`
}

func (m UpdateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
