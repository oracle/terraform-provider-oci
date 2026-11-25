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

// OracleManagedKeyManagementSystemDetails Type to be selected to reset Autonomous VM Cluster key management system to Oracle Managed Keys.
type OracleManagedKeyManagementSystemDetails struct {

	// If true, ACDs within this Cluster cannot use a different key management system than what is configured in AVM Cluster.
	IsExclusive *bool `mandatory:"false" json:"isExclusive"`
}

// GetIsExclusive returns IsExclusive
func (m OracleManagedKeyManagementSystemDetails) GetIsExclusive() *bool {
	return m.IsExclusive
}

func (m OracleManagedKeyManagementSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OracleManagedKeyManagementSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OracleManagedKeyManagementSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOracleManagedKeyManagementSystemDetails OracleManagedKeyManagementSystemDetails
	s := struct {
		DiscriminatorParam string `json:"keyManagementSystem"`
		MarshalTypeOracleManagedKeyManagementSystemDetails
	}{
		"ORACLE_MANAGED_KEYS",
		(MarshalTypeOracleManagedKeyManagementSystemDetails)(m),
	}

	return json.Marshal(&s)
}
