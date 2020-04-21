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

// GlossarySummary Summary of a glossary. A glossary of business terms, such as 'Customer', 'Account', 'Contact', 'Address',
// or 'Product', with definitions, used to provide common meaning across disparate data assets. Business glossaries
// may be hierarchical where some terms may contain child terms to allow them to be used as 'taxonomies'.
// By linking data assets, data entities, and attributes to glossaries and glossary terms, the glossary can act as a
// way of organizing data catalog objects in a hierarchy to make a large number of objects more navigable and easier to
// consume. Objects in the data catalog, such as data assets or data entities, may be linked to any level in the
// glossary, so that the glossary can be used to browse the available data according to the business model of the
// organization.
type GlossarySummary struct {

	// Unique glossary key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// The date and time the glossary was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Detailed description of the glossary.
	Description *string `mandatory:"false" json:"description"`

	// URI to the glossary instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Status of the approval process workflow for this business glossary.
	WorkflowStatus TermWorkflowStatusEnum `mandatory:"false" json:"workflowStatus,omitempty"`

	// State of the Glossary.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m GlossarySummary) String() string {
	return common.PointerString(m)
}
