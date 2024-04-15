// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Language API
//
// OCI Language Service solutions can help enterprise customers integrate AI into their products immediately using our proven,
// pre-trained and custom models or containers, without a need to set up an house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI and ML operations, which shortens the time to market.
//

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PiiEntity PII entity object.
type PiiEntity struct {

	// The number of Unicode code points preceding this entity in the submitted text.
	Offset *int `mandatory:"true" json:"offset"`

	// Length of PII entity text.
	Length *int `mandatory:"true" json:"length"`

	// Entity text like name of person, Organization and so on.
	Text *string `mandatory:"true" json:"text"`

	// Entity type supported
	// PERSON
	// ADDRESS
	// AGE
	// DATE_TIME
	// SSN_OR_TAXPAYER
	// EMAIL
	// PASSPORT_NUMBER_US
	// TELEPHONE_NUMBER
	// DRIVER_ID_US
	// BANK_ACCOUNT_NUMBER
	// BANK_SWIFT
	// BANK_ROUTING
	// CREDIT_DEBIT_NUMBER
	// IP_ADDRESS
	// MAC_ADDRESS
	// COOKIE
	// XSRF_TOKEN
	// AUTH_BASIC
	// AUTH_BEARER
	// JSON_WEB_TOKEN
	// PRIVATE_KEY
	// PUBLIC_KEY
	// OCI_OCID_USER
	// OCI_OCID_TENANCY
	// OCI_SMTP_USERNAME
	// OCI_OCID_REFERENCE
	// OCI_FINGERPRINT
	// OCI_CREDENTIAL
	// OCI_PRE_AUTH_REQUEST
	// OCI_STORAGE_SIGNED_URL
	// OCI_CUSTOMER_SECRET_KEY
	// OCI_ACCESS_KEy
	// MEDICAL_RECORD_NUMBER
	// HEALTH_PLAN_ID
	// URL
	// CERTIFICATE_NUMBER
	// FIN
	// GUIDs
	// VEHICLE_LICENSE_PLATE_US
	// VEHICLE_IDENTIFIER_US
	Type *string `mandatory:"true" json:"type"`

	// Score or confidence for detected PII entity.
	Score *float64 `mandatory:"true" json:"score"`

	// Unique id of the entity.
	Id *string `mandatory:"false" json:"id"`
}

func (m PiiEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PiiEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
