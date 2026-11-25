// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ThalesKeyManagementSystemDetails Details of Thales CAKM to be used by Autonomous VM Cluster for key management.
type ThalesKeyManagementSystemDetails struct {

	// If true, ACDs within this Cluster cannot use a different key management system than what is configured in AVM Cluster.
	IsExclusive *bool `mandatory:"false" json:"isExclusive"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store of Oracle Vault.
	KeyStoreId *string `mandatory:"false" json:"keyStoreId"`

	// Thales HSM Port. Default is 9000.
	ThalesHsmPort *int `mandatory:"false" json:"thalesHsmPort"`
}

// GetIsExclusive returns IsExclusive
func (m ThalesKeyManagementSystemDetails) GetIsExclusive() *bool {
	return m.IsExclusive
}

func (m ThalesKeyManagementSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThalesKeyManagementSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ThalesKeyManagementSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeThalesKeyManagementSystemDetails ThalesKeyManagementSystemDetails
	s := struct {
		DiscriminatorParam string `json:"keyManagementSystem"`
		MarshalTypeThalesKeyManagementSystemDetails
	}{
		"THALES",
		(MarshalTypeThalesKeyManagementSystemDetails)(m),
	}

	return json.Marshal(&s)
}
