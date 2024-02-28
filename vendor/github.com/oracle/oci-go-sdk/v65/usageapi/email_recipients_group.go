// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the chosen dimension. The Usage API is used by the Cost Analysis and Carbon Emissions Analysis tools in the Console. See Cost Analysis Overview (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm) and Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailRecipientsGroup The recipients group to receive usage statement email.
type EmailRecipientsGroup struct {

	// The usage statement email recipients group OCID.
	Id *string `mandatory:"true" json:"id"`

	// The customer tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of recipient will receive the usage statement email.
	RecipientsList []EmailRecipient `mandatory:"true" json:"recipientsList"`

	// The email recipient group lifecycle state.
	LifecycleState EmailRecipientsGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m EmailRecipientsGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailRecipientsGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailRecipientsGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailRecipientsGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailRecipientsGroupLifecycleStateEnum Enum with underlying type: string
type EmailRecipientsGroupLifecycleStateEnum string

// Set of constants representing the allowable values for EmailRecipientsGroupLifecycleStateEnum
const (
	EmailRecipientsGroupLifecycleStateActive   EmailRecipientsGroupLifecycleStateEnum = "ACTIVE"
	EmailRecipientsGroupLifecycleStateInactive EmailRecipientsGroupLifecycleStateEnum = "INACTIVE"
)

var mappingEmailRecipientsGroupLifecycleStateEnum = map[string]EmailRecipientsGroupLifecycleStateEnum{
	"ACTIVE":   EmailRecipientsGroupLifecycleStateActive,
	"INACTIVE": EmailRecipientsGroupLifecycleStateInactive,
}

var mappingEmailRecipientsGroupLifecycleStateEnumLowerCase = map[string]EmailRecipientsGroupLifecycleStateEnum{
	"active":   EmailRecipientsGroupLifecycleStateActive,
	"inactive": EmailRecipientsGroupLifecycleStateInactive,
}

// GetEmailRecipientsGroupLifecycleStateEnumValues Enumerates the set of values for EmailRecipientsGroupLifecycleStateEnum
func GetEmailRecipientsGroupLifecycleStateEnumValues() []EmailRecipientsGroupLifecycleStateEnum {
	values := make([]EmailRecipientsGroupLifecycleStateEnum, 0)
	for _, v := range mappingEmailRecipientsGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailRecipientsGroupLifecycleStateEnumStringValues Enumerates the set of values in String for EmailRecipientsGroupLifecycleStateEnum
func GetEmailRecipientsGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingEmailRecipientsGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailRecipientsGroupLifecycleStateEnum(val string) (EmailRecipientsGroupLifecycleStateEnum, bool) {
	enum, ok := mappingEmailRecipientsGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
