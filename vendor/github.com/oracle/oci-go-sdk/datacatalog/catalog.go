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

// Catalog A data catalog enables you to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
type Catalog struct {

	// OCID of the data catalog instance.
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Data catalog identifier, which can be renamed.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the data catalog was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the data catalog was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The REST front endpoint URL to the data catalog instance.
	ServiceApiUrl *string `mandatory:"false" json:"serviceApiUrl"`

	// The console front endpoint URL to the data catalog instance.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The number of data objects added to the data catalog.
	// Please see the data catalog documentation for further information on how this is calculated.
	NumberOfObjects *int `mandatory:"false" json:"numberOfObjects"`

	// The current state of the data catalog resource.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail.
	// For example, it can be used to provide actionable information for a resource in 'Failed' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The list of private reverse connection endpoints attached to the catalog
	AttachedCatalogPrivateEndpoints []string `mandatory:"false" json:"attachedCatalogPrivateEndpoints"`
}

func (m Catalog) String() string {
	return common.PointerString(m)
}
