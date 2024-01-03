// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PostInstallationActionSettings List of available post actions you can execute after the successful Java installation.
type PostInstallationActionSettings struct {

	// The following post JRE installation actions are supported by the field:
	// - Disable TLS 1.0 , TLS 1.1
	DisabledTlsVersions []TlsVersionsEnum `mandatory:"false" json:"disabledTlsVersions"`

	// Restores JDK root certificates with the certificates that are available in the operating system.
	// The following action is supported by the field:
	// - Replace JDK root certificates with a list provided by the operating system.
	ShouldReplaceCertificatesOperatingSystem *bool `mandatory:"false" json:"shouldReplaceCertificatesOperatingSystem"`

	MinimumKeySizeSettings *MinimumKeySizeSettings `mandatory:"false" json:"minimumKeySizeSettings"`

	// Sets FileHandler and ConsoleHandler as handlers in logging.properties file.
	AddLoggingHandler *bool `mandatory:"false" json:"addLoggingHandler"`

	// Sets the logging level in logging.properties file.
	GlobalLoggingLevel GlobalLoggingLevelEnum `mandatory:"false" json:"globalLoggingLevel,omitempty"`

	Proxies *Proxies `mandatory:"false" json:"proxies"`
}

func (m PostInstallationActionSettings) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PostInstallationActionSettings) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGlobalLoggingLevelEnum(string(m.GlobalLoggingLevel)); !ok && m.GlobalLoggingLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GlobalLoggingLevel: %s. Supported values are: %s.", m.GlobalLoggingLevel, strings.Join(GetGlobalLoggingLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
