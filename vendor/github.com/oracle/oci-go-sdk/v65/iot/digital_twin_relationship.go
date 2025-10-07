// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Internet of Things API
//
// Use the Internet of Things (IoT) API to manage IoT domain groups, domains, and digital twin resources including models, adapters, instances, and relationships.
// For more information, see Internet of Things (https://docs.oracle.com/iaas/Content/internet-of-things/home.htm).
//

package iot

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DigitalTwinRelationship A digital twin relationship represents an association between two digital twin instances.
// It can include properties to describe the nature and context of the association within the IoT domain.
// To use any API operations, you must be authorized in an IAM policy. If you are not authorized, contact an administrator.
// If you are an administrator who needs to create policies to grant users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type DigitalTwinRelationship struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IoT domain.
	IotDomainId *string `mandatory:"true" json:"iotDomainId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Its the name of the relationship that links two digital twin instances. Here, it is the relationship name of the source digital twin model.
	ContentPath *string `mandatory:"true" json:"contentPath"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source digital twin instance.
	SourceDigitalTwinInstanceId *string `mandatory:"true" json:"sourceDigitalTwinInstanceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target digital twin instance.
	TargetDigitalTwinInstanceId *string `mandatory:"true" json:"targetDigitalTwinInstanceId"`

	// The current state of the digital twin relationship.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the resource was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// A short description of the resource.
	Description *string `mandatory:"false" json:"description"`

	// The value(s) of the relationship properties defined in the source digital twin model.
	Content map[string]interface{} `mandatory:"false" json:"content"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time when the resource was last updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m DigitalTwinRelationship) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DigitalTwinRelationship) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
