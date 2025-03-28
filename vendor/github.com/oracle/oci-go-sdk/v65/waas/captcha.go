// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Captcha The settings of the CAPTCHA challenge. If a specific URL should be accessed only by a human, a CAPTCHA challenge can be placed at the URL to protect the web application from bots.
// *Warning:* Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Captcha struct {

	// The unique URL path at which to show the CAPTCHA challenge.
	Url *string `mandatory:"true" json:"url"`

	// The amount of time before the CAPTCHA expires, in seconds. If unspecified, defaults to `300`.
	SessionExpirationInSeconds *int `mandatory:"true" json:"sessionExpirationInSeconds"`

	// The title used when displaying a CAPTCHA challenge. If unspecified, defaults to `Are you human?`
	Title *string `mandatory:"true" json:"title"`

	// The text to show when incorrect CAPTCHA text is entered. If unspecified, defaults to `The CAPTCHA was incorrect. Try again.`
	FailureMessage *string `mandatory:"true" json:"failureMessage"`

	// The text to show on the label of the CAPTCHA challenge submit button. If unspecified, defaults to `Yes, I am human`.
	SubmitLabel *string `mandatory:"true" json:"submitLabel"`

	// The text to show in the header when showing a CAPTCHA challenge. If unspecified, defaults to 'We have detected an increased number of attempts to access this website. To help us keep this site secure, please let us know that you are not a robot by entering the text from the image below.'
	HeaderText *string `mandatory:"false" json:"headerText"`

	// The text to show in the footer when showing a CAPTCHA challenge. If unspecified, defaults to 'Enter the letters and numbers as they are shown in the image above.'
	FooterText *string `mandatory:"false" json:"footerText"`
}

func (m Captcha) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Captcha) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
