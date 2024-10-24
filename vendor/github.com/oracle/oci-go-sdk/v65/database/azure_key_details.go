// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AzureKeyDetails Details for Azure encryption key.
type AzureKeyDetails struct {

	// Azure vault URI
	VaultUri *string `mandatory:"true" json:"vaultUri"`

	// Azure key name
	KeyName *string `mandatory:"true" json:"keyName"`
}

func (m AzureKeyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AzureKeyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AzureKeyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAzureKeyDetails AzureKeyDetails
	s := struct {
		DiscriminatorParam string `json:"provider"`
		MarshalTypeAzureKeyDetails
	}{
		"AZURE",
		(MarshalTypeAzureKeyDetails)(m),
	}

	return json.Marshal(&s)
}
