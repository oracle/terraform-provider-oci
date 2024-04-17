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

// EmailRecipient The email recipient to receive usage statements for the subscription.
type EmailRecipient struct {

	// the email of the recipient.
	EmailId *string `mandatory:"true" json:"emailId"`

	// The email recipient lifecycle state.
	LifecycleState EmailRecipientLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// the first name of the recipient.
	FirstName *string `mandatory:"false" json:"firstName"`

	// the last name of the recipient.
	LastName *string `mandatory:"false" json:"lastName"`
}

func (m EmailRecipient) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailRecipient) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailRecipientLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailRecipientLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailRecipientLifecycleStateEnum Enum with underlying type: string
type EmailRecipientLifecycleStateEnum string

// Set of constants representing the allowable values for EmailRecipientLifecycleStateEnum
const (
	EmailRecipientLifecycleStateActive   EmailRecipientLifecycleStateEnum = "ACTIVE"
	EmailRecipientLifecycleStateInactive EmailRecipientLifecycleStateEnum = "INACTIVE"
)

var mappingEmailRecipientLifecycleStateEnum = map[string]EmailRecipientLifecycleStateEnum{
	"ACTIVE":   EmailRecipientLifecycleStateActive,
	"INACTIVE": EmailRecipientLifecycleStateInactive,
}

var mappingEmailRecipientLifecycleStateEnumLowerCase = map[string]EmailRecipientLifecycleStateEnum{
	"active":   EmailRecipientLifecycleStateActive,
	"inactive": EmailRecipientLifecycleStateInactive,
}

// GetEmailRecipientLifecycleStateEnumValues Enumerates the set of values for EmailRecipientLifecycleStateEnum
func GetEmailRecipientLifecycleStateEnumValues() []EmailRecipientLifecycleStateEnum {
	values := make([]EmailRecipientLifecycleStateEnum, 0)
	for _, v := range mappingEmailRecipientLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailRecipientLifecycleStateEnumStringValues Enumerates the set of values in String for EmailRecipientLifecycleStateEnum
func GetEmailRecipientLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingEmailRecipientLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailRecipientLifecycleStateEnum(val string) (EmailRecipientLifecycleStateEnum, bool) {
	enum, ok := mappingEmailRecipientLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
