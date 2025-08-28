// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EncryptDataDetails Encrypt data details.
type EncryptDataDetails struct {

	// Select whether to use Oracle-managed key (SYSTEM) or your own key (BYOK).
	KeyGenerationType KeyGenerationTypeEnum `mandatory:"true" json:"keyGenerationType"`

	// The OCID of the key to use.
	KeyId *string `mandatory:"false" json:"keyId"`
}

func (m EncryptDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EncryptDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingKeyGenerationTypeEnum(string(m.KeyGenerationType)); !ok && m.KeyGenerationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyGenerationType: %s. Supported values are: %s.", m.KeyGenerationType, strings.Join(GetKeyGenerationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
