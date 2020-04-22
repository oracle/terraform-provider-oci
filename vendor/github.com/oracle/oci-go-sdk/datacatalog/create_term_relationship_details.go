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

// CreateTermRelationshipDetails Properties used in term relationship create operations.
type CreateTermRelationshipDetails struct {

	// A user-friendly display name. Is changeable. The combination of 'displayName' and 'parentTermKey'
	// must be unique. Avoid entering confidential information. This is the same as 'relationshipType' for 'termRelationship'.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique id of the related term.
	RelatedTermKey *string `mandatory:"true" json:"relatedTermKey"`

	// Detailed description of the term relationship usually defined at the time of creation.
	Description *string `mandatory:"false" json:"description"`
}

func (m CreateTermRelationshipDetails) String() string {
	return common.PointerString(m)
}
