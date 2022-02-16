// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SuppressionSummary The full information representing a suppression.
type SuppressionSummary struct {

	// The OCID for the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The email address of the suppression.
	EmailAddress *string `mandatory:"true" json:"emailAddress"`

	// The unique OCID of the suppression.
	Id *string `mandatory:"true" json:"id"`

	// The reason that the email address was suppressed.
	Reason SuppressionSummaryReasonEnum `mandatory:"false" json:"reason,omitempty"`

	// The date and time a recipient's email address was added to the
	// suppression list, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m SuppressionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SuppressionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSuppressionSummaryReasonEnum(string(m.Reason)); !ok && m.Reason != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Reason: %s. Supported values are: %s.", m.Reason, strings.Join(GetSuppressionSummaryReasonEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SuppressionSummaryReasonEnum Enum with underlying type: string
type SuppressionSummaryReasonEnum string

// Set of constants representing the allowable values for SuppressionSummaryReasonEnum
const (
	SuppressionSummaryReasonUnknown     SuppressionSummaryReasonEnum = "UNKNOWN"
	SuppressionSummaryReasonHardbounce  SuppressionSummaryReasonEnum = "HARDBOUNCE"
	SuppressionSummaryReasonComplaint   SuppressionSummaryReasonEnum = "COMPLAINT"
	SuppressionSummaryReasonManual      SuppressionSummaryReasonEnum = "MANUAL"
	SuppressionSummaryReasonSoftbounce  SuppressionSummaryReasonEnum = "SOFTBOUNCE"
	SuppressionSummaryReasonUnsubscribe SuppressionSummaryReasonEnum = "UNSUBSCRIBE"
)

var mappingSuppressionSummaryReasonEnum = map[string]SuppressionSummaryReasonEnum{
	"UNKNOWN":     SuppressionSummaryReasonUnknown,
	"HARDBOUNCE":  SuppressionSummaryReasonHardbounce,
	"COMPLAINT":   SuppressionSummaryReasonComplaint,
	"MANUAL":      SuppressionSummaryReasonManual,
	"SOFTBOUNCE":  SuppressionSummaryReasonSoftbounce,
	"UNSUBSCRIBE": SuppressionSummaryReasonUnsubscribe,
}

// GetSuppressionSummaryReasonEnumValues Enumerates the set of values for SuppressionSummaryReasonEnum
func GetSuppressionSummaryReasonEnumValues() []SuppressionSummaryReasonEnum {
	values := make([]SuppressionSummaryReasonEnum, 0)
	for _, v := range mappingSuppressionSummaryReasonEnum {
		values = append(values, v)
	}
	return values
}

// GetSuppressionSummaryReasonEnumStringValues Enumerates the set of values in String for SuppressionSummaryReasonEnum
func GetSuppressionSummaryReasonEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"HARDBOUNCE",
		"COMPLAINT",
		"MANUAL",
		"SOFTBOUNCE",
		"UNSUBSCRIBE",
	}
}

// GetMappingSuppressionSummaryReasonEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSuppressionSummaryReasonEnum(val string) (SuppressionSummaryReasonEnum, bool) {
	mappingSuppressionSummaryReasonEnumIgnoreCase := make(map[string]SuppressionSummaryReasonEnum)
	for k, v := range mappingSuppressionSummaryReasonEnum {
		mappingSuppressionSummaryReasonEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSuppressionSummaryReasonEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
