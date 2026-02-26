// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestoreSoftDeletedObjectDetails The details needed to restore a soft deleted object.
type RestoreSoftDeletedObjectDetails struct {

	// The soft-deleted object to restore.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The version identifier of the soft-deleted object to restore.
	VersionId *string `mandatory:"true" json:"versionId"`

	// The entity tag (ETag) to match against that of the soft-deleted object. Used to confirm that the soft-deleted
	// object with a given name is the version of that object storing a specified ETag.
	ObjectIfMatchETag *string `mandatory:"false" json:"objectIfMatchETag"`

	// The name to assign to the restored object. If not provided, defaults to the original object name.
	NewObjectName *string `mandatory:"false" json:"newObjectName"`

	// The entity tag (ETag) to match against that of the new object (an object intended to be overwritten).
	// Used to confirm that the new object stored under a given name is the version of that object
	// storing a specified entity tag.
	NewObjectIfMatchETag *string `mandatory:"false" json:"newObjectIfMatchETag"`

	// The entity tag (ETag) to avoid matching. The only valid value is '*', which indicates that the request should fail
	// if the object already exists.
	NewObjectIfNoneMatchETag *string `mandatory:"false" json:"newObjectIfNoneMatchETag"`
}

func (m RestoreSoftDeletedObjectDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreSoftDeletedObjectDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
