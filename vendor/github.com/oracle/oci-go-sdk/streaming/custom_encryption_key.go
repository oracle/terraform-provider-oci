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

// CustomEncryptionKey Custom Encryption Key which will be used for encryption by all the streams in the pool.
type CustomEncryptionKey struct {

	// Custom Encryption Key (Master Key) ocid.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// Life cycle State of the custom key
	KeyState CustomEncryptionKeyKeyStateEnum `mandatory:"false" json:"keyState,omitempty"`
}

func (m CustomEncryptionKey) String() string {
	return common.PointerString(m)
}

// CustomEncryptionKeyKeyStateEnum Enum with underlying type: string
type CustomEncryptionKeyKeyStateEnum string

// Set of constants representing the allowable values for CustomEncryptionKeyKeyStateEnum
const (
	CustomEncryptionKeyKeyStateActive   CustomEncryptionKeyKeyStateEnum = "ACTIVE"
	CustomEncryptionKeyKeyStateCreating CustomEncryptionKeyKeyStateEnum = "CREATING"
	CustomEncryptionKeyKeyStateDeleting CustomEncryptionKeyKeyStateEnum = "DELETING"
	CustomEncryptionKeyKeyStateNone     CustomEncryptionKeyKeyStateEnum = "NONE"
	CustomEncryptionKeyKeyStateFailed   CustomEncryptionKeyKeyStateEnum = "FAILED"
	CustomEncryptionKeyKeyStateUpdating CustomEncryptionKeyKeyStateEnum = "UPDATING"
)

var mappingCustomEncryptionKeyKeyState = map[string]CustomEncryptionKeyKeyStateEnum{
	"ACTIVE":   CustomEncryptionKeyKeyStateActive,
	"CREATING": CustomEncryptionKeyKeyStateCreating,
	"DELETING": CustomEncryptionKeyKeyStateDeleting,
	"NONE":     CustomEncryptionKeyKeyStateNone,
	"FAILED":   CustomEncryptionKeyKeyStateFailed,
	"UPDATING": CustomEncryptionKeyKeyStateUpdating,
}

// GetCustomEncryptionKeyKeyStateEnumValues Enumerates the set of values for CustomEncryptionKeyKeyStateEnum
func GetCustomEncryptionKeyKeyStateEnumValues() []CustomEncryptionKeyKeyStateEnum {
	values := make([]CustomEncryptionKeyKeyStateEnum, 0)
	for _, v := range mappingCustomEncryptionKeyKeyState {
		values = append(values, v)
	}
	return values
}
