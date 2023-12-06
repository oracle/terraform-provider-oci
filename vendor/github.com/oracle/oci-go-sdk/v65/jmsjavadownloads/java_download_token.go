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

// JavaDownloadToken A JavaDownloadToken is a primary resource for the script friendly URLs. The value of this token serves as the authorization token for the download.
type JavaDownloadToken struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the JavaDownloadToken.
	Id *string `mandatory:"true" json:"id"`

	// User provided display name of the JavaDownloadToken.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy scoped to the JavaDownloadToken.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	CreatedBy *Principal `mandatory:"true" json:"createdBy"`

	// User provided description of the JavaDownloadToken.
	Description *string `mandatory:"true" json:"description"`

	// Uniquely generated value for the JavaDownloadToken.
	Value *string `mandatory:"true" json:"value"`

	// The time the JavaDownloadToken was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The expiry time of the JavaDownloadToken. An RFC3339 formatted datetime string.
	TimeExpires *common.SDKTime `mandatory:"true" json:"timeExpires"`

	// The associated Java version of the JavaDownloadToken.
	JavaVersion *string `mandatory:"true" json:"javaVersion"`

	// The current state of the JavaDownloadToken.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	LastUpdatedBy *Principal `mandatory:"false" json:"lastUpdatedBy"`

	// The time the JavaDownloadToken was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The time the JavaDownloadToken was last used for download. An RFC3339 formatted datetime string.
	TimeLastUsed *common.SDKTime `mandatory:"false" json:"timeLastUsed"`

	// The license type(s) associated with the JavaDownloadToken.
	LicenseType []LicenseTypeEnum `mandatory:"false" json:"licenseType"`

	// A flag to indicate if the token is default.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Possible lifecycle substates.
	LifecycleDetails TokenLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

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

func (m JavaDownloadToken) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JavaDownloadToken) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTokenLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetTokenLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
