// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Limits APIs
//
// APIs that interact with the resource limits of a specific resource type.
//

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceSummary A specific OCI service supported by resource limits.
type ServiceSummary struct {

	// The service name. Use this when calling other APIs.
	Name *string `mandatory:"false" json:"name"`

	// The friendly service name.
	Description *string `mandatory:"false" json:"description"`

	// An array of subscription types supported by the service. e,g The type of subscription, such as 'SAAS', 'ERP', 'CRM'.
	SupportedSubscriptions []string `mandatory:"false" json:"supportedSubscriptions"`
}

func (m ServiceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
