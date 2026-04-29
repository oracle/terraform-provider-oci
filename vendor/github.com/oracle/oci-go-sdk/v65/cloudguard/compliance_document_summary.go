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

// ComplianceDocumentSummary A summary representation of the compliance document.
type ComplianceDocumentSummary struct {

	// A unique identifier for the document that is assigned when you create
	// the document as an Oracle Cloud Infrastructure resource and is immutable.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name or title for the compliance document. You cannot update this value later.
	// Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// The current lifecycle state of the compliance document.
	LifecycleState ComplianceDocumentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The file name of the compliance document.
	DocumentFileName *string `mandatory:"true" json:"documentFileName"`

	// The version number of the compliance document.
	Version *int `mandatory:"true" json:"version"`

	// The type of compliance document.
	Type *string `mandatory:"true" json:"type"`

	// The environment, also known as platform or business pillar, to which the compliance document belongs.
	Platform *string `mandatory:"true" json:"platform"`

	// The date and time the compliance document was last updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The date and time the compliance document was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The saas service name to which compliance document belongs.
	SaasServiceName *string `mandatory:"false" json:"saasServiceName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ComplianceDocumentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComplianceDocumentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingComplianceDocumentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetComplianceDocumentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
