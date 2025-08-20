// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateMonitoringTemplateDetails The information about updating a monitoring template. The monitoring template displayName should be unique in a compartment.
type UpdateMonitoringTemplateDetails struct {

	// A user-friendly name for the monitoring template. It is unique and mutable in nature.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description for the monitoring template. It does not have to be unique, and it's changeable.
	Description *string `mandatory:"false" json:"description"`

	// A list of destinations for alarm notifications. Each destination is represented by the OCID of a related resource
	Destinations []string `mandatory:"false" json:"destinations"`

	// User can create the out of box alarm only for multiple resourceTypes not for individual resource instances and groups for specified compartment.
	IsAlarmsEnabled *bool `mandatory:"false" json:"isAlarmsEnabled"`

	// Whether the alarm notification is enabled or disabled, it will be Enabled by default.
	IsSplitNotificationEnabled *bool `mandatory:"false" json:"isSplitNotificationEnabled"`

	// List of members of this monitoring template.
	Members []MemberReference `mandatory:"false" json:"members"`

	// The frequency for re-submitting alarm notifications, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, PT4H indicates four hours. Minimum- PT1M. Maximum - P30D.
	RepeatNotificationDuration *string `mandatory:"false" json:"repeatNotificationDuration"`

	// The format to use for alarm notifications.
	MessageFormat MessageFormatEnum `mandatory:"false" json:"messageFormat,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateMonitoringTemplateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMonitoringTemplateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMessageFormatEnum(string(m.MessageFormat)); !ok && m.MessageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageFormat: %s. Supported values are: %s.", m.MessageFormat, strings.Join(GetMessageFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
