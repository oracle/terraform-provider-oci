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

// ReplaceObjectMetadataDetails To replace Objects User metadata we specify the new metadata in the body.
type ReplaceObjectMetadataDetails struct {

	// Arbitrary string keys-values pair for the user-defined metadata for the object.
	// Keys must be in "opc-meta-*" format. Avoid entering confidential information.
	// The size of user-defined metadata is measured by taking the sum of the number of bytes in the UTF-8 encoding
	// of each key and value. The maximum metadata size is 2975 bytes.
	Metadata map[string]string `mandatory:"true" json:"metadata"`
}

func (m ReplaceObjectMetadataDetails) String() string {
	return common.PointerString(m)
}
