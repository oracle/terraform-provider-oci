// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsKeyStoreContentSecretIdMySqlSummary The key store content.
type DatabaseToolsKeyStoreContentSecretIdMySqlSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store.
	SecretId *string `mandatory:"false" json:"secretId"`
}

func (m DatabaseToolsKeyStoreContentSecretIdMySqlSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsKeyStoreContentSecretIdMySqlSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsKeyStoreContentSecretIdMySqlSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsKeyStoreContentSecretIdMySqlSummary DatabaseToolsKeyStoreContentSecretIdMySqlSummary
	s := struct {
		DiscriminatorParam string `json:"valueType"`
		MarshalTypeDatabaseToolsKeyStoreContentSecretIdMySqlSummary
	}{
		"SECRETID",
		(MarshalTypeDatabaseToolsKeyStoreContentSecretIdMySqlSummary)(m),
	}

	return json.Marshal(&s)
}
