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

// ObjectStorageUpdateDetail Details of Object Storage instance to be updated.
type ObjectStorageUpdateDetail struct {

	// Instance id of the existing Object Storage instance to be updated.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// Mentions which storage tier to use for the bucket,Allowed values are "STANDARD" / "ARCHIVE"
	ObjectVersioning ObjectVersioningEnum `mandatory:"false" json:"objectVersioning,omitempty"`

	// It sets the auto-tiering status on the bucket.Allowed values are "DISABLED" / "INFREQUENTACCESS"
	AutoTiering AutoTieringEnum `mandatory:"false" json:"autoTiering,omitempty"`
}

func (m ObjectStorageUpdateDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ObjectStorageUpdateDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingObjectVersioningEnum(string(m.ObjectVersioning)); !ok && m.ObjectVersioning != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectVersioning: %s. Supported values are: %s.", m.ObjectVersioning, strings.Join(GetObjectVersioningEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutoTieringEnum(string(m.AutoTiering)); !ok && m.AutoTiering != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoTiering: %s. Supported values are: %s.", m.AutoTiering, strings.Join(GetAutoTieringEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
