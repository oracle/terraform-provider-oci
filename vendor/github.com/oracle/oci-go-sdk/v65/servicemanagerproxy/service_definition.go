// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// Use the Service Manager Proxy API to obtain information about SaaS environments provisioned by Service Manager.
// You can get information such as service types and service environment URLs.
//

package servicemanagerproxy

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ServiceDefinition Details for a service definition.
type ServiceDefinition struct {

	// The service definition type. For example, a service definition type "RGBUOROMS"
	// would be for the service "Oracle Retail Order Management Cloud Service".
	Type *string `mandatory:"true" json:"type"`

	// Display name of the service. For example, "Oracle Retail Order Management Cloud Service".
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Short display name of the service. For example, "Retail Order Management".
	ShortDisplayName *string `mandatory:"true" json:"shortDisplayName"`
}

func (m ServiceDefinition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceDefinition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
