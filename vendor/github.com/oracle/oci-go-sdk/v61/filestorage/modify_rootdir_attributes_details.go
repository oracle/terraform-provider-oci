// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// ModifyRootdirAttributesDetails Details for changing the ownership of file system root directory.
type ModifyRootdirAttributesDetails struct {

	// The user id for root directory ownership. However if not specified then it does not change.
	Uid *int `mandatory:"false" json:"uid"`

	// The group id for root directory ownership. However if not specified then it does not change.
	Gid *int `mandatory:"false" json:"gid"`

	// Only standard mode bits are supported like 644, 755 etc. However if not specified then it does not change.
	Mode *int `mandatory:"false" json:"mode"`
}

func (m ModifyRootdirAttributesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModifyRootdirAttributesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
