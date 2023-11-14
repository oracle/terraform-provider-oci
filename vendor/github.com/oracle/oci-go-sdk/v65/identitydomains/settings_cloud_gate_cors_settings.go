// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SettingsCloudGateCorsSettings A complex attribute that specifies the Cloud Gate cross origin resource sharing settings.
// **Added In:** 2011192329
// **SCIM++ Properties:**
//   - caseExact: false
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: false
//   - returned: default
//   - type: complex
//   - uniqueness: none
type SettingsCloudGateCorsSettings struct {

	// Allow Null Origin (CORS) for this tenant.
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CloudGateCorsAllowNullOrigin *bool `mandatory:"false" json:"cloudGateCorsAllowNullOrigin"`

	// Enable Cloud Gate Cross-Origin Resource Sharing (CORS) for this tenant.
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	CloudGateCorsEnabled *bool `mandatory:"false" json:"cloudGateCorsEnabled"`

	// Cloud Gate Allowed Cross-Origin Resource Sharing (CORS) Origins for this tenant.
	// **Added In:** 2011192329
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CloudGateCorsAllowedOrigins []string `mandatory:"false" json:"cloudGateCorsAllowedOrigins"`

	// Maximum number of seconds a CORS Pre-flight Response may be cached by client.
	// **Added In:** 2205182039
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	CloudGateCorsMaxAge *int `mandatory:"false" json:"cloudGateCorsMaxAge"`

	// List of Response Headers Cloud Gate is allowed to expose in the CORS Response Header: Access-Control-Expose-Headers.
	// **Added In:** 2205182039
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CloudGateCorsExposedHeaders []string `mandatory:"false" json:"cloudGateCorsExposedHeaders"`
}

func (m SettingsCloudGateCorsSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SettingsCloudGateCorsSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
