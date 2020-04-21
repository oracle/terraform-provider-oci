// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
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

// AgreementSummaryAuthorEnum Enum with underlying type: string
type AgreementSummaryAuthorEnum string

// Set of constants representing the allowable values for AgreementSummaryAuthorEnum
const (
	AgreementSummaryAuthorOracle  AgreementSummaryAuthorEnum = "ORACLE"
	AgreementSummaryAuthorPartner AgreementSummaryAuthorEnum = "PARTNER"
	AgreementSummaryAuthorPii     AgreementSummaryAuthorEnum = "PII"
)

var mappingAgreementSummaryAuthor = map[string]AgreementSummaryAuthorEnum{
	"ORACLE":  AgreementSummaryAuthorOracle,
	"PARTNER": AgreementSummaryAuthorPartner,
	"PII":     AgreementSummaryAuthorPii,
}

// GetAgreementSummaryAuthorEnumValues Enumerates the set of values for AgreementSummaryAuthorEnum
func GetAgreementSummaryAuthorEnumValues() []AgreementSummaryAuthorEnum {
	values := make([]AgreementSummaryAuthorEnum, 0)
	for _, v := range mappingAgreementSummaryAuthor {
		values = append(values, v)
	}
	return values
}
