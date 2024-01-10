// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CreateJavaDownloadTokenDetails The attributes to create a new JavaDownloadToken.
type CreateJavaDownloadTokenDetails struct {

	// User provided display name of the JavaDownloadToken.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// User provided description of the JavaDownloadToken.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the tenancy scoped to the JavaDownloadToken.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Expiry time of the token.
	TimeExpires *common.SDKTime `mandatory:"true" json:"timeExpires"`

	// The Java version associated with the token.
	JavaVersion *string `mandatory:"true" json:"javaVersion"`

	// The license type(s) associated with the JavaDownloadToken.
	LicenseType []LicenseTypeEnum `mandatory:"true" json:"licenseType"`

	// The token default attribute.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`. (See Managing Tags and Tag Namespaces (https://docs.cloud.oracle.com/Content/Tagging/Concepts/understandingfreeformtags.htm).)
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`. (See Understanding Free-form Tags (https://docs.cloud.oracle.com/Content/Tagging/Tasks/managingtagsandtagnamespaces.htm)).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJavaDownloadTokenDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateJavaDownloadTokenDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
