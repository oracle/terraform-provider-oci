// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccAgreementSummary The model for an end user license agreement.
type CccAgreementSummary struct {

	// The unique identifier for the agreement.
	Id *string `mandatory:"true" json:"id"`

	// Compute Cloud@Customer agreement display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The content URL of the agreement.
	ContentUrl *string `mandatory:"true" json:"contentUrl"`

	// Who authored the agreement.
	Author CccAgreementSummaryAuthorEnum `mandatory:"true" json:"author"`

	// Textual prompt to read and accept the agreement.
	Prompt *string `mandatory:"true" json:"prompt"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CccAgreementSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccAgreementSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccAgreementSummaryAuthorEnum(string(m.Author)); !ok && m.Author != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Author: %s. Supported values are: %s.", m.Author, strings.Join(GetCccAgreementSummaryAuthorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccAgreementSummaryAuthorEnum Enum with underlying type: string
type CccAgreementSummaryAuthorEnum string

// Set of constants representing the allowable values for CccAgreementSummaryAuthorEnum
const (
	CccAgreementSummaryAuthorOracle  CccAgreementSummaryAuthorEnum = "ORACLE"
	CccAgreementSummaryAuthorPartner CccAgreementSummaryAuthorEnum = "PARTNER"
	CccAgreementSummaryAuthorPii     CccAgreementSummaryAuthorEnum = "PII"
)

var mappingCccAgreementSummaryAuthorEnum = map[string]CccAgreementSummaryAuthorEnum{
	"ORACLE":  CccAgreementSummaryAuthorOracle,
	"PARTNER": CccAgreementSummaryAuthorPartner,
	"PII":     CccAgreementSummaryAuthorPii,
}

var mappingCccAgreementSummaryAuthorEnumLowerCase = map[string]CccAgreementSummaryAuthorEnum{
	"oracle":  CccAgreementSummaryAuthorOracle,
	"partner": CccAgreementSummaryAuthorPartner,
	"pii":     CccAgreementSummaryAuthorPii,
}

// GetCccAgreementSummaryAuthorEnumValues Enumerates the set of values for CccAgreementSummaryAuthorEnum
func GetCccAgreementSummaryAuthorEnumValues() []CccAgreementSummaryAuthorEnum {
	values := make([]CccAgreementSummaryAuthorEnum, 0)
	for _, v := range mappingCccAgreementSummaryAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetCccAgreementSummaryAuthorEnumStringValues Enumerates the set of values in String for CccAgreementSummaryAuthorEnum
func GetCccAgreementSummaryAuthorEnumStringValues() []string {
	return []string{
		"ORACLE",
		"PARTNER",
		"PII",
	}
}

// GetMappingCccAgreementSummaryAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccAgreementSummaryAuthorEnum(val string) (CccAgreementSummaryAuthorEnum, bool) {
	enum, ok := mappingCccAgreementSummaryAuthorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
