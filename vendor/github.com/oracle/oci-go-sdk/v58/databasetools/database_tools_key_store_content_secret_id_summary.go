// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DatabaseToolsKeyStoreContentSecretIdSummary The key store content.
type DatabaseToolsKeyStoreContentSecretIdSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the secret containing the key store.
	SecretId *string `mandatory:"false" json:"secretId"`
}

func (m DatabaseToolsKeyStoreContentSecretIdSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsKeyStoreContentSecretIdSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseToolsKeyStoreContentSecretIdSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseToolsKeyStoreContentSecretIdSummary DatabaseToolsKeyStoreContentSecretIdSummary
	s := struct {
		DiscriminatorParam string `json:"valueType"`
		MarshalTypeDatabaseToolsKeyStoreContentSecretIdSummary
	}{
		"SECRETID",
		(MarshalTypeDatabaseToolsKeyStoreContentSecretIdSummary)(m),
	}

	return json.Marshal(&s)
}
