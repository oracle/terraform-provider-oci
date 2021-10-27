// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// BulkEditResource The representation of BulkEditResource
type BulkEditResource struct {

	// The unique OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The type of resource. See BulkEditResourceTypes.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// Additional information that identifies the resource for bulk editing of tags. This information is provided in the resource's API documentation.
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m BulkEditResource) String() string {
	return common.PointerString(m)
}
