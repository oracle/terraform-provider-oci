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

// Agreement The model for an end user license agreement.
type Agreement struct {

	// The unique identifier for the agreement.
	Id *string `mandatory:"true" json:"id"`

	// The content URL of the agreement.
	ContentUrl *string `mandatory:"true" json:"contentUrl"`

	// A time-based signature that can be used to accept an agreement or remove a
	// previously accepted agreement from the list that Marketplace checks before a deployment.
	Signature *string `mandatory:"true" json:"signature"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Who authored the agreement.
	Author AgreementAuthorEnum `mandatory:"false" json:"author,omitempty"`

	// Textual prompt to read and accept the agreement.
	Prompt *string `mandatory:"false" json:"prompt"`
}

func (m Agreement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Agreement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAgreementAuthorEnum(string(m.Author)); !ok && m.Author != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Author: %s. Supported values are: %s.", m.Author, strings.Join(GetAgreementAuthorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AgreementAuthorEnum Enum with underlying type: string
type AgreementAuthorEnum string

// Set of constants representing the allowable values for AgreementAuthorEnum
const (
	AgreementAuthorOracle  AgreementAuthorEnum = "ORACLE"
	AgreementAuthorPartner AgreementAuthorEnum = "PARTNER"
)

var mappingAgreementAuthorEnum = map[string]AgreementAuthorEnum{
	"ORACLE":  AgreementAuthorOracle,
	"PARTNER": AgreementAuthorPartner,
}

// GetAgreementAuthorEnumValues Enumerates the set of values for AgreementAuthorEnum
func GetAgreementAuthorEnumValues() []AgreementAuthorEnum {
	values := make([]AgreementAuthorEnum, 0)
	for _, v := range mappingAgreementAuthorEnum {
		values = append(values, v)
	}
	return values
}

// GetAgreementAuthorEnumStringValues Enumerates the set of values in String for AgreementAuthorEnum
func GetAgreementAuthorEnumStringValues() []string {
	return []string{
		"ORACLE",
		"PARTNER",
	}
}

// GetMappingAgreementAuthorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAgreementAuthorEnum(val string) (AgreementAuthorEnum, bool) {
	mappingAgreementAuthorEnumIgnoreCase := make(map[string]AgreementAuthorEnum)
	for k, v := range mappingAgreementAuthorEnum {
		mappingAgreementAuthorEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAgreementAuthorEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
