// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KmsKey Information about the state of the service's encryption key management.
// The following properties are read-only and ignored when this object is
// used in UpdateConfiguration: kmsKeyState, timeCreated, timeUpdated.
type KmsKey struct {

	// The OCID of the KMS encryption key assigned to this Hosted
	// Environment. If the Hosted Environment is using an
	// Oracle-managed Key, then the id will be a null string..
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the vault containing the encryption key assigned
	// to this Hosted Environment. If the Hosted Environment is
	// using an Oracle-managed Key, then the kmsVaultId will be a
	// null string.
	KmsVaultId *string `mandatory:"false" json:"kmsVaultId"`

	// The current state of the encryption key assigned to this
	// Hosted Environment. Oracle-managed keys will always report
	// an ACTIVE state.
	KmsKeyState KmsKeyKmsKeyStateEnum `mandatory:"false" json:"kmsKeyState,omitempty"`

	// The timestamp when encryption key was first enabled for this Hosted Environment.
	// RFC3339 formatted.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The timestamp of the last update to the encryption key status. RFC3339 formatted.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m KmsKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KmsKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingKmsKeyKmsKeyStateEnum(string(m.KmsKeyState)); !ok && m.KmsKeyState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KmsKeyState: %s. Supported values are: %s.", m.KmsKeyState, strings.Join(GetKmsKeyKmsKeyStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KmsKeyKmsKeyStateEnum Enum with underlying type: string
type KmsKeyKmsKeyStateEnum string

// Set of constants representing the allowable values for KmsKeyKmsKeyStateEnum
const (
	KmsKeyKmsKeyStateUpdating  KmsKeyKmsKeyStateEnum = "UPDATING"
	KmsKeyKmsKeyStateActive    KmsKeyKmsKeyStateEnum = "ACTIVE"
	KmsKeyKmsKeyStateDeleted   KmsKeyKmsKeyStateEnum = "DELETED"
	KmsKeyKmsKeyStateFailed    KmsKeyKmsKeyStateEnum = "FAILED"
	KmsKeyKmsKeyStateReverting KmsKeyKmsKeyStateEnum = "REVERTING"
	KmsKeyKmsKeyStateDisabled  KmsKeyKmsKeyStateEnum = "DISABLED"
)

var mappingKmsKeyKmsKeyStateEnum = map[string]KmsKeyKmsKeyStateEnum{
	"UPDATING":  KmsKeyKmsKeyStateUpdating,
	"ACTIVE":    KmsKeyKmsKeyStateActive,
	"DELETED":   KmsKeyKmsKeyStateDeleted,
	"FAILED":    KmsKeyKmsKeyStateFailed,
	"REVERTING": KmsKeyKmsKeyStateReverting,
	"DISABLED":  KmsKeyKmsKeyStateDisabled,
}

var mappingKmsKeyKmsKeyStateEnumLowerCase = map[string]KmsKeyKmsKeyStateEnum{
	"updating":  KmsKeyKmsKeyStateUpdating,
	"active":    KmsKeyKmsKeyStateActive,
	"deleted":   KmsKeyKmsKeyStateDeleted,
	"failed":    KmsKeyKmsKeyStateFailed,
	"reverting": KmsKeyKmsKeyStateReverting,
	"disabled":  KmsKeyKmsKeyStateDisabled,
}

// GetKmsKeyKmsKeyStateEnumValues Enumerates the set of values for KmsKeyKmsKeyStateEnum
func GetKmsKeyKmsKeyStateEnumValues() []KmsKeyKmsKeyStateEnum {
	values := make([]KmsKeyKmsKeyStateEnum, 0)
	for _, v := range mappingKmsKeyKmsKeyStateEnum {
		values = append(values, v)
	}
	return values
}

// GetKmsKeyKmsKeyStateEnumStringValues Enumerates the set of values in String for KmsKeyKmsKeyStateEnum
func GetKmsKeyKmsKeyStateEnumStringValues() []string {
	return []string{
		"UPDATING",
		"ACTIVE",
		"DELETED",
		"FAILED",
		"REVERTING",
		"DISABLED",
	}
}

// GetMappingKmsKeyKmsKeyStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKmsKeyKmsKeyStateEnum(val string) (KmsKeyKmsKeyStateEnum, bool) {
	enum, ok := mappingKmsKeyKmsKeyStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
