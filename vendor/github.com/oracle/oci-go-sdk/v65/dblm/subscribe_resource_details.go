// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscribeResourceDetails SubscriberDatabasesDetails.
type SubscribeResourceDetails struct {

	// Compartment Id of the subscriber Databases
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of subscriber databases.
	Items []SubscriberDatabaseDetails `mandatory:"true" json:"items"`

	// Unique identifier of SoftwareImage
	SoftwareImageId *string `mandatory:"true" json:"softwareImageId"`

	// Intermediate user to be used for patching, created and maintained by customers. This user requires sudo access to switch as Oracle home owner and root user
	PatchUser *string `mandatory:"false" json:"patchUser"`

	// Path to sudo binary (executable) file
	SudoFilePath *string `mandatory:"false" json:"sudoFilePath"`
}

func (m SubscribeResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscribeResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
