// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Common set of Object Storage and Archive Storage APIs for managing buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// BulkCopyObjectsItem The parameters required by Object Storage to process a request to bulk copy objects to another bucket.
type BulkCopyObjectsItem struct {

	// Unique number associated with an individual bulk copy item, that serve as a unique identifier of the copy
	// item. This number should be in the range of [1, batchCount]
	ItemNumber *int `mandatory:"true" json:"itemNumber"`

	// The name of the object to be copied.
	SourceObjectName *string `mandatory:"true" json:"sourceObjectName"`

	// The name of the destination object resulting from the copy operation. Avoid entering confidential information.
	DestinationObjectName *string `mandatory:"true" json:"destinationObjectName"`

	// The entity tag (ETag) to match against that of the source object. Used to confirm that the source object
	// with a given name is the version of that object storing a specified ETag.
	SourceObjectIfMatchETag *string `mandatory:"false" json:"sourceObjectIfMatchETag"`

	// The entity tag (ETag) to match against that of the destination object (an object intended to be overwritten).
	// Used to confirm that the destination object stored under a given name is the version of that object
	// storing a specified entity tag.
	DestinationObjectIfMatchETag *string `mandatory:"false" json:"destinationObjectIfMatchETag"`

	// The entity tag (ETag) to avoid matching. The only valid value is '*', which indicates that the request should fail
	// if the object already exists in the destination bucket.
	DestinationObjectIfNoneMatchETag *string `mandatory:"false" json:"destinationObjectIfNoneMatchETag"`

	// Arbitrary string keys and values for the user-defined metadata for the object. Avoid entering confidential
	// information. Metadata key-value pairs entered in this field are assigned to the destination object. If you
	// enter no metadata values, the destination object will inherit any existing metadata values associated with
	// the source object.
	DestinationObjectMetadata map[string]string `mandatory:"false" json:"destinationObjectMetadata"`
}

func (m BulkCopyObjectsItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkCopyObjectsItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
