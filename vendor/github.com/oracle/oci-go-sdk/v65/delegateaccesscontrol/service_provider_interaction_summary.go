// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceProviderInteractionSummary Summary of customer and operator conversation.
type ServiceProviderInteractionSummary struct {

	// The unique identifier of the message within the scope of the associated access request.
	MessageIdentifier *string `mandatory:"false" json:"messageIdentifier"`

	// ID of the customer or operator who is part of this conversation. For operator, this field is null.
	UserId *string `mandatory:"false" json:"userId"`

	// Name of the customer or operator who is part of this conversation. For operator, the name is "Operator".
	UserName *string `mandatory:"false" json:"userName"`

	// The information exchanged between the customer and the operator.
	Message *string `mandatory:"false" json:"message"`

	// Indicates whether the user is a customer or an operator.
	UserType ServiceProviderInteractionSummaryUserTypeEnum `mandatory:"false" json:"userType,omitempty"`

	// Time when the conversation happened in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format, e.g. '2020-05-22T21:10:29.600Z'.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m ServiceProviderInteractionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceProviderInteractionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingServiceProviderInteractionSummaryUserTypeEnum(string(m.UserType)); !ok && m.UserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UserType: %s. Supported values are: %s.", m.UserType, strings.Join(GetServiceProviderInteractionSummaryUserTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceProviderInteractionSummaryUserTypeEnum Enum with underlying type: string
type ServiceProviderInteractionSummaryUserTypeEnum string

// Set of constants representing the allowable values for ServiceProviderInteractionSummaryUserTypeEnum
const (
	ServiceProviderInteractionSummaryUserTypeCustomer ServiceProviderInteractionSummaryUserTypeEnum = "CUSTOMER"
	ServiceProviderInteractionSummaryUserTypeOperator ServiceProviderInteractionSummaryUserTypeEnum = "OPERATOR"
)

var mappingServiceProviderInteractionSummaryUserTypeEnum = map[string]ServiceProviderInteractionSummaryUserTypeEnum{
	"CUSTOMER": ServiceProviderInteractionSummaryUserTypeCustomer,
	"OPERATOR": ServiceProviderInteractionSummaryUserTypeOperator,
}

var mappingServiceProviderInteractionSummaryUserTypeEnumLowerCase = map[string]ServiceProviderInteractionSummaryUserTypeEnum{
	"customer": ServiceProviderInteractionSummaryUserTypeCustomer,
	"operator": ServiceProviderInteractionSummaryUserTypeOperator,
}

// GetServiceProviderInteractionSummaryUserTypeEnumValues Enumerates the set of values for ServiceProviderInteractionSummaryUserTypeEnum
func GetServiceProviderInteractionSummaryUserTypeEnumValues() []ServiceProviderInteractionSummaryUserTypeEnum {
	values := make([]ServiceProviderInteractionSummaryUserTypeEnum, 0)
	for _, v := range mappingServiceProviderInteractionSummaryUserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceProviderInteractionSummaryUserTypeEnumStringValues Enumerates the set of values in String for ServiceProviderInteractionSummaryUserTypeEnum
func GetServiceProviderInteractionSummaryUserTypeEnumStringValues() []string {
	return []string{
		"CUSTOMER",
		"OPERATOR",
	}
}

// GetMappingServiceProviderInteractionSummaryUserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceProviderInteractionSummaryUserTypeEnum(val string) (ServiceProviderInteractionSummaryUserTypeEnum, bool) {
	enum, ok := mappingServiceProviderInteractionSummaryUserTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
