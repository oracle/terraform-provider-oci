// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAllPackagesOnManagedInstancesInCompartmentDetails The details about the package types to be updated.
type UpdateAllPackagesOnManagedInstancesInCompartmentDetails struct {

	// The compartment being targeted by this operation.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of updates to be applied.
	UpdateTypes []UpdateTypesEnum `mandatory:"false" json:"updateTypes,omitempty"`

	WorkRequestDetails *WorkRequestDetails `mandatory:"false" json:"workRequestDetails"`
}

func (m UpdateAllPackagesOnManagedInstancesInCompartmentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAllPackagesOnManagedInstancesInCompartmentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.UpdateTypes {
		if _, ok := GetMappingUpdateTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateTypes: %s. Supported values are: %s.", val, strings.Join(GetUpdateTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
