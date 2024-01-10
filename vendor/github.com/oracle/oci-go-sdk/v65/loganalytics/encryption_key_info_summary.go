// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EncryptionKeyInfoSummary This is the summary of an encryption key.
type EncryptionKeyInfoSummary struct {

	// This is the source of the encryption key.
	KeySource EncryptionKeySourceEnum `mandatory:"true" json:"keySource"`

	// This is the key OCID of the encryption key (null if Oracle-managed).
	KeyId *string `mandatory:"true" json:"keyId"`

	// This is the type of data to be encrypted. It can be either active or archival.
	KeyType EncryptionKeyTypeEnum `mandatory:"true" json:"keyType"`
}

func (m EncryptionKeyInfoSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EncryptionKeyInfoSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEncryptionKeySourceEnum(string(m.KeySource)); !ok && m.KeySource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeySource: %s. Supported values are: %s.", m.KeySource, strings.Join(GetEncryptionKeySourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingEncryptionKeyTypeEnum(string(m.KeyType)); !ok && m.KeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyType: %s. Supported values are: %s.", m.KeyType, strings.Join(GetEncryptionKeyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
