// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// APM Availability Monitoring API
//
// Use the APM Availability Monitoring API to query Scripts, Monitors, Dedicated Vantage Points and On-Premise Vantage Points resources. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RequestAuthenticationDetails Details for request HTTP authentication.
type RequestAuthenticationDetails struct {

	// Request HTTP OAuth scheme.
	OauthScheme OAuthSchemesEnum `mandatory:"false" json:"oauthScheme,omitempty"`

	// User name for authentication.
	AuthUserName *string `mandatory:"false" json:"authUserName"`

	// User password for authentication.
	AuthUserPassword *string `mandatory:"false" json:"authUserPassword"`

	// Authentication token.
	AuthToken *string `mandatory:"false" json:"authToken"`

	// URL to get authentication token.
	AuthUrl *string `mandatory:"false" json:"authUrl"`

	// List of authentication headers. Example: `[{"headerName": "content-type", "headerValue":"json"}]`
	AuthHeaders []Header `mandatory:"false" json:"authHeaders"`

	// Request method.
	AuthRequestMethod RequestMethodsEnum `mandatory:"false" json:"authRequestMethod,omitempty"`

	// Request post body.
	AuthRequestPostBody *string `mandatory:"false" json:"authRequestPostBody"`
}

func (m RequestAuthenticationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RequestAuthenticationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOAuthSchemesEnum(string(m.OauthScheme)); !ok && m.OauthScheme != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OauthScheme: %s. Supported values are: %s.", m.OauthScheme, strings.Join(GetOAuthSchemesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestMethodsEnum(string(m.AuthRequestMethod)); !ok && m.AuthRequestMethod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuthRequestMethod: %s. Supported values are: %s.", m.AuthRequestMethod, strings.Join(GetRequestMethodsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
