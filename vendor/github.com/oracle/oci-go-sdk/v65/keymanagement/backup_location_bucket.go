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

// BackupLocationBucket Object storage bucket details to upload or download the backup
type BackupLocationBucket struct {
	Namespace *string `mandatory:"true" json:"namespace"`

	BucketName *string `mandatory:"true" json:"bucketName"`

	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m BackupLocationBucket) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupLocationBucket) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BackupLocationBucket) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBackupLocationBucket BackupLocationBucket
	s := struct {
		DiscriminatorParam string `json:"destination"`
		MarshalTypeBackupLocationBucket
	}{
		"BUCKET",
		(MarshalTypeBackupLocationBucket)(m),
	}

	return json.Marshal(&s)
}
