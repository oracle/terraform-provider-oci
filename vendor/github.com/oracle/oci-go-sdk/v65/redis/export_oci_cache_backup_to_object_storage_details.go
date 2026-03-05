// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExportOciCacheBackupToObjectStorageDetails Parameters to export an OCI Cache Backup’s RDB file(s) to Object Storage.
type ExportOciCacheBackupToObjectStorageDetails struct {

	// The Object Storage namespace name.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The target Object Storage bucket name.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Optional prefix under which the service will place the exported object(s).
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m ExportOciCacheBackupToObjectStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportOciCacheBackupToObjectStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
