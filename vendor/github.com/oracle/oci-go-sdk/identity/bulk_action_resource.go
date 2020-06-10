// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BulkActionResource The bulk action resource entity.
type BulkActionResource struct {

	// The resource identifier.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The resource type.
	EntityType *string `mandatory:"true" json:"entityType"`

	// Additional information that helps to identity the resource for bulk action.
	// DELETE and UPDATE APIs for most resource types only require the resource identifier(ocid).
	// But additional metadata is required for some resource types.
	// This information is provided in the resource's public API document. It is also
	// available through the ListBulkActionResourceTypes API.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m BulkActionResource) String() string {
	return common.PointerString(m)
}
