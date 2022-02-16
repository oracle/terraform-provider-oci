// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// API for the File Storage service. Use this API to manage file systems, mount targets, and snapshots. For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ChangeFileSystemCompartmentDetails Details for changing the compartment.
type ChangeFileSystemCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to move the file system to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeFileSystemCompartmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeFileSystemCompartmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
