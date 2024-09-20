// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateDesktopPoolDetails Provides the details of a request to update the desktop pool.
type UpdateDesktopPoolDetails struct {

	// A user friendly display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user friendly description providing additional information about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The maximum number of desktops permitted in the desktop pool.
	MaximumSize *int `mandatory:"false" json:"maximumSize"`

	// The maximum number of standby desktops available in the desktop pool.
	StandbySize *int `mandatory:"false" json:"standbySize"`

	DevicePolicy *DesktopDevicePolicy `mandatory:"false" json:"devicePolicy"`

	AvailabilityPolicy *DesktopAvailabilityPolicy `mandatory:"false" json:"availabilityPolicy"`

	// Contact information of the desktop pool administrator.
	// Avoid entering confidential information.
	ContactDetails *string `mandatory:"false" json:"contactDetails"`

	// The start time of the desktop pool.
	TimeStartScheduled *common.SDKTime `mandatory:"false" json:"timeStartScheduled"`

	// The stop time of the desktop pool.
	TimeStopScheduled *common.SDKTime `mandatory:"false" json:"timeStopScheduled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	SessionLifecycleActions *UpdateDesktopPoolDesktopSessionLifecycleActions `mandatory:"false" json:"sessionLifecycleActions"`
}

func (m UpdateDesktopPoolDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateDesktopPoolDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
