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

// DesktopPoolSummary Provides information about the desktop pool such as number of active desktops, name, OCID, state, and maximum size.
type DesktopPoolSummary struct {

	// The OCID of the desktop pool.
	Id *string `mandatory:"true" json:"id"`

	// A user friendly display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the desktop pool.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The maximum number of desktops permitted in the desktop pool.
	MaximumSize *int `mandatory:"true" json:"maximumSize"`

	// Contact information of the desktop pool administrator.
	// Avoid entering confidential information.
	ContactDetails *string `mandatory:"true" json:"contactDetails"`

	// The OCID of the compartment which will contain the desktop pool.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The number of active desktops in the desktop pool.
	ActiveDesktops *int `mandatory:"false" json:"activeDesktops"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DesktopPoolSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopPoolSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
