// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Streaming API
//
// Use the Streaming API to produce and consume messages, create streams and stream pools, and manage related items. For more information, see Streaming (https://docs.cloud.oracle.com/Content/Streaming/Concepts/streamingoverview.htm).
//

package streaming

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomEncryptionKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomEncryptionKeyKeyStateEnum(string(m.KeyState)); !ok && m.KeyState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyState: %s. Supported values are: %s.", m.KeyState, strings.Join(GetCustomEncryptionKeyKeyStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingCustomEncryptionKeyKeyStateEnum = map[string]CustomEncryptionKeyKeyStateEnum{
	"ACTIVE":   CustomEncryptionKeyKeyStateActive,
	"CREATING": CustomEncryptionKeyKeyStateCreating,
	"DELETING": CustomEncryptionKeyKeyStateDeleting,
	"NONE":     CustomEncryptionKeyKeyStateNone,
	"FAILED":   CustomEncryptionKeyKeyStateFailed,
	"UPDATING": CustomEncryptionKeyKeyStateUpdating,
}

var mappingCustomEncryptionKeyKeyStateEnumLowerCase = map[string]CustomEncryptionKeyKeyStateEnum{
	"active":   CustomEncryptionKeyKeyStateActive,
	"creating": CustomEncryptionKeyKeyStateCreating,
	"deleting": CustomEncryptionKeyKeyStateDeleting,
	"none":     CustomEncryptionKeyKeyStateNone,
	"failed":   CustomEncryptionKeyKeyStateFailed,
	"updating": CustomEncryptionKeyKeyStateUpdating,
}

// GetCustomEncryptionKeyKeyStateEnumValues Enumerates the set of values for CustomEncryptionKeyKeyStateEnum
func GetCustomEncryptionKeyKeyStateEnumValues() []CustomEncryptionKeyKeyStateEnum {
	values := make([]CustomEncryptionKeyKeyStateEnum, 0)
	for _, v := range mappingCustomEncryptionKeyKeyStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCustomEncryptionKeyKeyStateEnumStringValues Enumerates the set of values in String for CustomEncryptionKeyKeyStateEnum
func GetCustomEncryptionKeyKeyStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"NONE",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingCustomEncryptionKeyKeyStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCustomEncryptionKeyKeyStateEnum(val string) (CustomEncryptionKeyKeyStateEnum, bool) {
	enum, ok := mappingCustomEncryptionKeyKeyStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
