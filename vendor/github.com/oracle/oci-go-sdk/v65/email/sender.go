// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Sender The full information representing an approved sender.
type Sender struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Email address of the sender.
	EmailAddress *string `mandatory:"true" json:"emailAddress"`

	// The unique OCID of the sender.
	Id *string `mandatory:"true" json:"id"`

	// Value of the SPF field. For more information about SPF, please see
	// SPF Authentication (https://docs.cloud.oracle.com/Content/Email/Concepts/overview.htm#components).
	IsSpf *bool `mandatory:"false" json:"isSpf"`

	// The sender's current lifecycle state.
	LifecycleState SenderLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the approved sender was added in "YYYY-MM-ddThh:mmZ"
	// format with a Z offset, as defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The email domain used to assert responsibility for emails sent from this sender.
	EmailDomainId *string `mandatory:"false" json:"emailDomainId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Sender) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Sender) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSenderLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSenderLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SenderLifecycleStateEnum Enum with underlying type: string
type SenderLifecycleStateEnum string

// Set of constants representing the allowable values for SenderLifecycleStateEnum
const (
	SenderLifecycleStateCreating SenderLifecycleStateEnum = "CREATING"
	SenderLifecycleStateActive   SenderLifecycleStateEnum = "ACTIVE"
	SenderLifecycleStateDeleting SenderLifecycleStateEnum = "DELETING"
	SenderLifecycleStateDeleted  SenderLifecycleStateEnum = "DELETED"
)

var mappingSenderLifecycleStateEnum = map[string]SenderLifecycleStateEnum{
	"CREATING": SenderLifecycleStateCreating,
	"ACTIVE":   SenderLifecycleStateActive,
	"DELETING": SenderLifecycleStateDeleting,
	"DELETED":  SenderLifecycleStateDeleted,
}

var mappingSenderLifecycleStateEnumLowerCase = map[string]SenderLifecycleStateEnum{
	"creating": SenderLifecycleStateCreating,
	"active":   SenderLifecycleStateActive,
	"deleting": SenderLifecycleStateDeleting,
	"deleted":  SenderLifecycleStateDeleted,
}

// GetSenderLifecycleStateEnumValues Enumerates the set of values for SenderLifecycleStateEnum
func GetSenderLifecycleStateEnumValues() []SenderLifecycleStateEnum {
	values := make([]SenderLifecycleStateEnum, 0)
	for _, v := range mappingSenderLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSenderLifecycleStateEnumStringValues Enumerates the set of values in String for SenderLifecycleStateEnum
func GetSenderLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSenderLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSenderLifecycleStateEnum(val string) (SenderLifecycleStateEnum, bool) {
	enum, ok := mappingSenderLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
