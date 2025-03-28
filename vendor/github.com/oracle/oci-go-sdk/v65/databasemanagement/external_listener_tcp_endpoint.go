// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalListenerTcpEndpoint A `TCP`-based protocol address.
type ExternalListenerTcpEndpoint struct {

	// The host name or IP address.
	Host *string `mandatory:"true" json:"host"`

	// The port number.
	Port *int `mandatory:"true" json:"port"`

	// The list of services registered with the listener.
	Services []string `mandatory:"false" json:"services"`
}

// GetServices returns Services
func (m ExternalListenerTcpEndpoint) GetServices() []string {
	return m.Services
}

func (m ExternalListenerTcpEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalListenerTcpEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalListenerTcpEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalListenerTcpEndpoint ExternalListenerTcpEndpoint
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeExternalListenerTcpEndpoint
	}{
		"TCP",
		(MarshalTypeExternalListenerTcpEndpoint)(m),
	}

	return json.Marshal(&s)
}
