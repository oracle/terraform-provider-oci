// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnnouncementSummary A summary of announcements for Fleet Application Management.
type AnnouncementSummary struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of announcement.
	Type *string `mandatory:"true" json:"type"`

	// Summary of the announcement.
	Summary *string `mandatory:"true" json:"summary"`

	// Announcement start date.
	AnnouncementStart *common.SDKTime `mandatory:"true" json:"announcementStart"`

	// Announcement end date
	AnnouncementEnd *common.SDKTime `mandatory:"true" json:"announcementEnd"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Announcement Details.
	Details *string `mandatory:"false" json:"details"`

	// URL to the announcement.
	Url *string `mandatory:"false" json:"url"`

	// The lifecycle state of the announcement.
	LifecycleState AnnouncementSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m AnnouncementSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnnouncementSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAnnouncementSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAnnouncementSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AnnouncementSummaryLifecycleStateEnum Enum with underlying type: string
type AnnouncementSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AnnouncementSummaryLifecycleStateEnum
const (
	AnnouncementSummaryLifecycleStateActive   AnnouncementSummaryLifecycleStateEnum = "ACTIVE"
	AnnouncementSummaryLifecycleStateInactive AnnouncementSummaryLifecycleStateEnum = "INACTIVE"
	AnnouncementSummaryLifecycleStateDeleted  AnnouncementSummaryLifecycleStateEnum = "DELETED"
	AnnouncementSummaryLifecycleStateFailed   AnnouncementSummaryLifecycleStateEnum = "FAILED"
)

var mappingAnnouncementSummaryLifecycleStateEnum = map[string]AnnouncementSummaryLifecycleStateEnum{
	"ACTIVE":   AnnouncementSummaryLifecycleStateActive,
	"INACTIVE": AnnouncementSummaryLifecycleStateInactive,
	"DELETED":  AnnouncementSummaryLifecycleStateDeleted,
	"FAILED":   AnnouncementSummaryLifecycleStateFailed,
}

var mappingAnnouncementSummaryLifecycleStateEnumLowerCase = map[string]AnnouncementSummaryLifecycleStateEnum{
	"active":   AnnouncementSummaryLifecycleStateActive,
	"inactive": AnnouncementSummaryLifecycleStateInactive,
	"deleted":  AnnouncementSummaryLifecycleStateDeleted,
	"failed":   AnnouncementSummaryLifecycleStateFailed,
}

// GetAnnouncementSummaryLifecycleStateEnumValues Enumerates the set of values for AnnouncementSummaryLifecycleStateEnum
func GetAnnouncementSummaryLifecycleStateEnumValues() []AnnouncementSummaryLifecycleStateEnum {
	values := make([]AnnouncementSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAnnouncementSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAnnouncementSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AnnouncementSummaryLifecycleStateEnum
func GetAnnouncementSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAnnouncementSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnnouncementSummaryLifecycleStateEnum(val string) (AnnouncementSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAnnouncementSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
