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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Tagslug A resource that has information about the tag and some other important metadata
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Tagslug struct {

	// The date and time the tagSlug was created in target region.
	// Example: `2021-01-04T20:01:29.100Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The type of resource the tag is associated with.
	// Example: "1" for filesystem, "2" for mount target
	Type *string `mandatory:"false" json:"type"`

	// determines if the tagslug is used by a resource or not.
	// Example: "1" for used, "0" for not used
	Used *string `mandatory:"false" json:"used"`

	// The id of the filesystem or mount target if this tag is associated to
	Id1 *string `mandatory:"false" json:"id1"`

	// The id of snapshot if this tag is associated to else 0
	Id2 *string `mandatory:"false" json:"id2"`
}

func (m Tagslug) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Tagslug) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
