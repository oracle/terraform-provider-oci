// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataIntelligences Control Plane API
//
// Use the DataIntelligences Control Plane API to manage dataIntelligences.
//

package dif

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ObjectStorageDetail Object Storage Details required to provision buckets.
type ObjectStorageDetail struct {

	// Id for Object Storage instance to be provisioned.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Mentions whether the object versioning to be enabled or not,Allowed values are "ENABLED" / "DISABLED"/"SUSPENDED"
	ObjectVersioning ObjectVersioningEnum `mandatory:"true" json:"objectVersioning"`

	// Mentions which storage tier to use for the bucket,Allowed values are "STANDARD" / "ARCHIVE"
	StorageTier StorageTierEnum `mandatory:"true" json:"storageTier"`

	// It sets the auto-tiering status on the bucket.Allowed values are "DISABLED" / "INFREQUENTACCESS"
	AutoTiering AutoTieringEnum `mandatory:"false" json:"autoTiering,omitempty"`
}

func (m ObjectStorageDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingObjectVersioningEnum(string(m.ObjectVersioning)); !ok && m.ObjectVersioning != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectVersioning: %s. Supported values are: %s.", m.ObjectVersioning, strings.Join(GetObjectVersioningEnumStringValues(), ",")))
	}
	if _, ok := GetMappingStorageTierEnum(string(m.StorageTier)); !ok && m.StorageTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageTier: %s. Supported values are: %s.", m.StorageTier, strings.Join(GetStorageTierEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutoTieringEnum(string(m.AutoTiering)); !ok && m.AutoTiering != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoTiering: %s. Supported values are: %s.", m.AutoTiering, strings.Join(GetAutoTieringEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
