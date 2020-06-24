// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateCatalogPrivateEndpointDetails Information about the new private endpoint resource
type CreateCatalogPrivateEndpointDetails struct {

	// List of DNS zones to be used by the data assets to be harvested.
	// Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com
	DnsZones []string `mandatory:"true" json:"dnsZones"`

	// The OCID of subnet to which the reverse connection is to be created
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Display name of the private endpoint resource being created.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateCatalogPrivateEndpointDetails) String() string {
	return common.PointerString(m)
}
