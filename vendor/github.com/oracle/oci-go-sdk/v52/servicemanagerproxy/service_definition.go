// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Manager Proxy API
//
// API to manage Service manager proxy.
//

package servicemanagerproxy

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// ServiceDefinition Model for details associated with service
type ServiceDefinition struct {

	// The service definition type.
	Type *string `mandatory:"true" json:"type"`

	// Display name of the service.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Short display name of the service.
	ShortDisplayName *string `mandatory:"true" json:"shortDisplayName"`
}

func (m ServiceDefinition) String() string {
	return common.PointerString(m)
}
