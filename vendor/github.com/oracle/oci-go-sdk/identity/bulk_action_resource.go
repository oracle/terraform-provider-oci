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

	// The resource OCID.
	Identifier *string `mandatory:"true" json:"identifier"`

	// The resource-type. To get the list of supported resource-types use
	// ListBulkActionResourceTypes.
	EntityType *string `mandatory:"true" json:"entityType"`

	// Additional information that helps to identity the resource for bulk action.
	// The APIs to delete and move most resource types only require the resource identifier (ocid).
	// But some resource-types require additional identifying information.
	// This information is provided in the resource's public API document. It is also
	// available through the
	// ListBulkActionResourceTypes.
	// **Example**:
	// The APIs to delete or move the `buckets` resource-type require `namespaceName` and `bucketName` to identify the resource, as
	// shown in the APIs, DeleteBucket and
	// UpdateBucket.
	// To add a bucket for bulk actions, specify `namespaceName` and `bucketName` in
	// the metadata property as shown in this example
	//     {
	//       "identifier": "<OCID_of_bucket>"
	//       "entityType": "bucket",
	//       "metadata":
	//       {
	//         "namespaceName": "sampleNamespace",
	//         "bucketName": "sampleBucket"
	//       }
	//     }
	Metadata map[string]string `mandatory:"false" json:"metadata"`
}

func (m BulkActionResource) String() string {
	return common.PointerString(m)
}
