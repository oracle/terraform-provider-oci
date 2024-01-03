// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RestoreVaultFromObjectStoreDetails The details of the backup location from which you want to restore the Vault.
type RestoreVaultFromObjectStoreDetails struct {
	BackupLocation BackupLocation `mandatory:"false" json:"backupLocation"`
}

func (m RestoreVaultFromObjectStoreDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestoreVaultFromObjectStoreDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RestoreVaultFromObjectStoreDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		BackupLocation backuplocation `json:"backupLocation"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.BackupLocation.UnmarshalPolymorphicJSON(model.BackupLocation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.BackupLocation = nn.(BackupLocation)
	} else {
		m.BackupLocation = nil
	}

	return
}
