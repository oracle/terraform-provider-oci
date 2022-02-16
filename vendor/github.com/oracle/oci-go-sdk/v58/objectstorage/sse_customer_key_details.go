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

// SseCustomerKeyDetails Specifies the details of the customer-provided encryption key (SSE-C) associated with an object.
type SseCustomerKeyDetails struct {

	// Specifies the encryption algorithm. The only supported value is "AES256".
	Algorithm SseCustomerKeyDetailsAlgorithmEnum `mandatory:"true" json:"algorithm"`

	// Specifies the base64-encoded 256-bit encryption key to use to encrypt or decrypt the object data.
	Key *string `mandatory:"true" json:"key"`

	// Specifies the base64-encoded SHA256 hash of the encryption key. This value is used to check the integrity
	// of the encryption key.
	KeySha256 *string `mandatory:"true" json:"keySha256"`
}

func (m SseCustomerKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SseCustomerKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSseCustomerKeyDetailsAlgorithmEnum(string(m.Algorithm)); !ok && m.Algorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Algorithm: %s. Supported values are: %s.", m.Algorithm, strings.Join(GetSseCustomerKeyDetailsAlgorithmEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SseCustomerKeyDetailsAlgorithmEnum Enum with underlying type: string
type SseCustomerKeyDetailsAlgorithmEnum string

// Set of constants representing the allowable values for SseCustomerKeyDetailsAlgorithmEnum
const (
	SseCustomerKeyDetailsAlgorithmAes256 SseCustomerKeyDetailsAlgorithmEnum = "AES256"
)

var mappingSseCustomerKeyDetailsAlgorithmEnum = map[string]SseCustomerKeyDetailsAlgorithmEnum{
	"AES256": SseCustomerKeyDetailsAlgorithmAes256,
}

// GetSseCustomerKeyDetailsAlgorithmEnumValues Enumerates the set of values for SseCustomerKeyDetailsAlgorithmEnum
func GetSseCustomerKeyDetailsAlgorithmEnumValues() []SseCustomerKeyDetailsAlgorithmEnum {
	values := make([]SseCustomerKeyDetailsAlgorithmEnum, 0)
	for _, v := range mappingSseCustomerKeyDetailsAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetSseCustomerKeyDetailsAlgorithmEnumStringValues Enumerates the set of values in String for SseCustomerKeyDetailsAlgorithmEnum
func GetSseCustomerKeyDetailsAlgorithmEnumStringValues() []string {
	return []string{
		"AES256",
	}
}

// GetMappingSseCustomerKeyDetailsAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSseCustomerKeyDetailsAlgorithmEnum(val string) (SseCustomerKeyDetailsAlgorithmEnum, bool) {
	mappingSseCustomerKeyDetailsAlgorithmEnumIgnoreCase := make(map[string]SseCustomerKeyDetailsAlgorithmEnum)
	for k, v := range mappingSseCustomerKeyDetailsAlgorithmEnum {
		mappingSseCustomerKeyDetailsAlgorithmEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSseCustomerKeyDetailsAlgorithmEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
