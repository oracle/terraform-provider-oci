// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutomatedMountDetails Used for creating NFS Auto Mount backup destinations for autonomous on ExaCC.
type AutomatedMountDetails struct {

	// IP addresses for NFS Auto mount.
	NfsServer []string `mandatory:"true" json:"nfsServer"`

	// Specifies the directory on which to mount the file system
	NfsServerExport *string `mandatory:"true" json:"nfsServerExport"`
}

func (m AutomatedMountDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutomatedMountDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutomatedMountDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutomatedMountDetails AutomatedMountDetails
	s := struct {
		DiscriminatorParam string `json:"mountType"`
		MarshalTypeAutomatedMountDetails
	}{
		"AUTOMATED_MOUNT",
		(MarshalTypeAutomatedMountDetails)(m),
	}

	return json.Marshal(&s)
}
