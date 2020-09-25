// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v25/common"
)

// PostObjectMetadataResponse Updated object information after user-metadata update.
type PostObjectMetadataResponse struct {

	// The new entity tag (ETag) for the object.
	ETag *string `mandatory:"true" json:"ETag"`

	// The time the object was modified, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`
}

func (m PostObjectMetadataResponse) String() string {
	return common.PointerString(m)
}
