// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming Service API
//
// The API for the Streaming Service.
//

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CustomEncryptionKeyDetails The OCID of the custom encryption key to be used or deleted if currently being used.
type CustomEncryptionKeyDetails struct {

	// Custom Encryption Key (Master Key) ocid.
	KmsKeyId *string `mandatory:"true" json:"kmsKeyId"`
}

func (m CustomEncryptionKeyDetails) String() string {
	return common.PointerString(m)
}
