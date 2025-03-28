// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Security Attribute API
//
// Use the Security Attributes API to manage security attributes and security attribute namespaces. For more information, see the documentation for Security Attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attributes.htm) and Security Attribute Namespaces (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/managing-security-attribute-namespaces.htm).
//

package securityattribute

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityAttributeWorkRequestLogSummary The log entity.
type SecurityAttributeWorkRequestLogSummary struct {

	// A human-readable error string.
	Message *string `mandatory:"true" json:"message"`

	// Date and time the log was written, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`
}

func (m SecurityAttributeWorkRequestLogSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAttributeWorkRequestLogSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
