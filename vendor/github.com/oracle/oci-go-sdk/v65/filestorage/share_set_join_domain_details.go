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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShareSetJoinDomainDetails Details for the mount target share set to join a domain controller.
type ShareSetJoinDomainDetails struct {

	// The user is the account name of a sufficiently powerful administrator
	// account. The default domain is the left most domain of the
	// customer-provided DNS name.
	// This user name can be specified in several different ways:
	//   * As an ordinary name (e.g. joe).
	//   * As a name at some domain (e.g. sally@some.domain.com).
	//   * As a name under some domain (e.g. some.domain.com\administrator).
	User *string `mandatory:"true" json:"user"`

	// The credential password of the user to join the domain.
	Password *string `mandatory:"true" json:"password"`
}

func (m ShareSetJoinDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShareSetJoinDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
