// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Download API
//
// The APIs for the download engine of the Java Management Service.
//

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// JavaLicenseAcceptanceRecord User acceptance record for a Java license.
type JavaLicenseAcceptanceRecord struct {

	// The unique identifier for the acceptance record.
	Id *string `mandatory:"true" json:"id"`

	// Status of license acceptance.
	LicenseAcceptanceStatus LicenseAcceptanceStatusEnum `mandatory:"true" json:"licenseAcceptanceStatus"`

	// The tenancy OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the user accepting the license.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// License type associated with the acceptance.
	LicenseType LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	CreatedBy *Principal `mandatory:"true" json:"createdBy"`

	// The date and time of license acceptance(formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	LastUpdatedBy *Principal `mandatory:"false" json:"lastUpdatedBy"`

	// The date and time of last update(formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeLastUpdated *common.SDKTime `mandatory:"false" json:"timeLastUpdated"`

	// The current state of the JavaLicenseAcceptanceRecord.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m JavaLicenseAcceptanceRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaLicenseAcceptanceRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLicenseAcceptanceStatusEnum(string(m.LicenseAcceptanceStatus)); !ok && m.LicenseAcceptanceStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseAcceptanceStatus: %s. Supported values are: %s.", m.LicenseAcceptanceStatus, strings.Join(GetLicenseAcceptanceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLicenseTypeEnum(string(m.LicenseType)); !ok && m.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", m.LicenseType, strings.Join(GetLicenseTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
