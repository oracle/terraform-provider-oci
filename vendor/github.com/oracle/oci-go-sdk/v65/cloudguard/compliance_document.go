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

// ComplianceDocument A compliance document that exists in the tenancy.
type ComplianceDocument struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compliance document, which is assigned
	// when you create the document as an Oracle Cloud Infrastructure resource and is immutable.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name or title for the compliance document. You cannot update this value later.
	// Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The date and time the compliance document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the compliance document.
	LifecycleState ComplianceDocumentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The file name of the compliance document.
	DocumentFileName *string `mandatory:"true" json:"documentFileName"`

	// The version number of the compliance document.
	Version *int `mandatory:"true" json:"version"`

	// The type of compliance document. For definitions of supported types of compliance documents, see Types of Compliance Documents (https://docs.oracle.com/iaas/en-us/iaas/Content/ComplianceDocuments/Concepts/compliancedocsoverview.htm#DocTypes).
	Type *string `mandatory:"true" json:"type"`

	// The information technology infrastructure platform, or set of services, to which the compliance document belongs. A platform
	// can also be described as an environment or a business pillar. For definitions of supported environments, see Types of Environments (https://docs.oracle.com/iaas/en-us/iaas/Content/ComplianceDocuments/Concepts/compliancedocsoverview.htm#EnvironmentTypes).
	Platform *string `mandatory:"true" json:"platform"`

	// The date and time the compliance document was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The saas service name to which compliance document belongs. For other types such as 'OCI' / 'PaaS' this value will be null.
	SaasServiceName *string `mandatory:"false" json:"saasServiceName"`
}

func (m ComplianceDocument) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceDocument) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComplianceDocumentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComplianceDocumentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ComplianceDocumentLifecycleStateEnum Enum with underlying type: string
type ComplianceDocumentLifecycleStateEnum string

// Set of constants representing the allowable values for ComplianceDocumentLifecycleStateEnum
const (
	ComplianceDocumentLifecycleStateActive   ComplianceDocumentLifecycleStateEnum = "ACTIVE"
	ComplianceDocumentLifecycleStateInactive ComplianceDocumentLifecycleStateEnum = "INACTIVE"
)

var mappingComplianceDocumentLifecycleStateEnum = map[string]ComplianceDocumentLifecycleStateEnum{
	"ACTIVE":   ComplianceDocumentLifecycleStateActive,
	"INACTIVE": ComplianceDocumentLifecycleStateInactive,
}

var mappingComplianceDocumentLifecycleStateEnumLowerCase = map[string]ComplianceDocumentLifecycleStateEnum{
	"active":   ComplianceDocumentLifecycleStateActive,
	"inactive": ComplianceDocumentLifecycleStateInactive,
}

// GetComplianceDocumentLifecycleStateEnumValues Enumerates the set of values for ComplianceDocumentLifecycleStateEnum
func GetComplianceDocumentLifecycleStateEnumValues() []ComplianceDocumentLifecycleStateEnum {
	values := make([]ComplianceDocumentLifecycleStateEnum, 0)
	for _, v := range mappingComplianceDocumentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetComplianceDocumentLifecycleStateEnumStringValues Enumerates the set of values in String for ComplianceDocumentLifecycleStateEnum
func GetComplianceDocumentLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingComplianceDocumentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComplianceDocumentLifecycleStateEnum(val string) (ComplianceDocumentLifecycleStateEnum, bool) {
	enum, ok := mappingComplianceDocumentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
