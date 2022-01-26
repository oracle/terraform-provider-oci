// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Service Key Management API
//
// API for managing and performing operations with keys and vaults. (For the API for managing secrets, see the Vault Service
// Secret Management API. For the API for retrieving secrets, see the Vault Service Secret Retrieval API.)
//

package keymanagement

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// BackupLocationUri PreAuthenticated object storage URI to upload or download the backup
type BackupLocationUri struct {
	Uri *string `mandatory:"true" json:"uri"`
}

func (m BackupLocationUri) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m BackupLocationUri) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBackupLocationUri BackupLocationUri
	s := struct {
		DiscriminatorParam string `json:"destination"`
		MarshalTypeBackupLocationUri
	}{
		"PRE_AUTHENTICATED_REQUEST_URI",
		(MarshalTypeBackupLocationUri)(m),
	}

	return json.Marshal(&s)
}
