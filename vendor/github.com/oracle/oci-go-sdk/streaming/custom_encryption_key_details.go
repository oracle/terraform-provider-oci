// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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
