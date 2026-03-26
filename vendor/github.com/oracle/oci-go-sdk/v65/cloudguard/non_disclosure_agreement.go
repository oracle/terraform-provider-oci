// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NonDisclosureAgreement A non-disclosure agreement that describes terms of use for a particular compliance document.
type NonDisclosureAgreement struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the non-disclosure agreement, which is assigned
	// when you create the non-disclosure agreement as an Oracle Cloud Infrastructure resource and is immutable.
	Id *string `mandatory:"true" json:"id"`

	// The ID of the compliance document associated with the non-disclosure agreement.
	DocumentId *string `mandatory:"true" json:"documentId"`

	// The OCID of the principal that called CreateNonDisclosureAgreement (https://docs.oracle.com/iaas/api/#/en/compliancedocs/release/NonDisclosureAgreement/CreateNonDisclosureAgreement).
	UserId *string `mandatory:"true" json:"userId"`

	// The OCID of the compartment that contains the non-disclosure agreement.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The actual terms of the non-disclosure agreement between the customer and Oracle.
	AgreementContent *string `mandatory:"true" json:"agreementContent"`

	// The date and time the non-disclosure agreement was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m NonDisclosureAgreement) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NonDisclosureAgreement) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
