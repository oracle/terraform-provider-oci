// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// AgreementSummary The model for a summary of an end user license agreement.
type AgreementSummary struct {

	// The unique identifier for the agreement.
	Id *string `mandatory:"false" json:"id"`

	// The content URL of the agreement.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// Who authored the agreement.
	Author AgreementSummaryAuthorEnum `mandatory:"false" json:"author,omitempty"`

	// Textual prompt to read and accept the agreement.
	Prompt *string `mandatory:"false" json:"prompt"`
}

func (m AgreementSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AgreementSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAgreementSummaryAuthorEnum(string(m.Author)); !ok && m.Author != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Author: %s. Supported values are: %s.", m.Author, strings.Join(GetAgreementSummaryAuthorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AgreementSummaryAuthorEnum Enum with underlying type: string
type AgreementSummaryAuthorEnum string

// Set of constants representing the allowable values for AgreementSummaryAuthorEnum
const (
	AgreementSummaryAuthorOracle  AgreementSummaryAuthorEnum = "ORACLE"
	AgreementSummaryAuthorPartner AgreementSummaryAuthorEnum = "PARTNER"
	AgreementSummaryAuthorPii     AgreementSummaryAuthorEnum = "PII"
)

var mappingAgreementSummaryAuthorEnum = map[string]AgreementSummaryAuthorEnum{
	"ORACLE":  AgreementSummaryAuthorOracle,
	"PARTNER": AgreementSummaryAuthorPartner,
	"PII":     AgreementSummaryAuthorPii,
}

// GetAgreementSummaryAuthorEnumValues Enumerates the set of values for AgreementSummaryAuthorEnum
func GetAgreementSummaryAuthorEnumValues() []AgreementSummaryAuthorEnum {
	values := make([]AgreementSummaryAuthorEnum, 0)
	for _, v := range mappingAgreementSummaryAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetAgreementSummaryAuthorEnumStringValues Enumerates the set of values in String for AgreementSummaryAuthorEnum
func GetAgreementSummaryAuthorEnumStringValues() []string {
	return []string{
		"ORACLE",
		"PARTNER",
		"PII",
	}
}

// GetMappingAgreementSummaryAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgreementSummaryAuthorEnum(val string) (AgreementSummaryAuthorEnum, bool) {
	mappingAgreementSummaryAuthorEnumIgnoreCase := make(map[string]AgreementSummaryAuthorEnum)
	for k, v := range mappingAgreementSummaryAuthorEnum {
		mappingAgreementSummaryAuthorEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAgreementSummaryAuthorEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
