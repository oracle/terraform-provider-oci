// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkEditTagsResourceType The representation of BulkEditTagsResourceType
type BulkEditTagsResourceType struct {

	// The unique name of the resource type.
	ResourceType *string `mandatory:"true" json:"resourceType"`

	// The metadata keys required to identify the resource.
	// For example, for a bucket, the value of `metadataKeys` will be "namespaceName", "bucketName".
	// This information will match the API documentation.
	// See UpdateBucket (https://docs.cloud.oracle.com/api/#/en/objectstorage/latest/Bucket/UpdateBucket) and
	// DeleteBucket (https://docs.cloud.oracle.com/api/#/en/objectstorage/latest/Bucket/DeleteBucket).
	MetadataKeys []string `mandatory:"false" json:"metadataKeys"`
}

func (m BulkEditTagsResourceType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkEditTagsResourceType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
