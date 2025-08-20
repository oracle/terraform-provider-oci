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

// CloudListenerIpcEndpoint An `IPC`-based protocol address.
type CloudListenerIpcEndpoint struct {

	// The unique name of the service.
	Key *string `mandatory:"true" json:"key"`

	// The list of services registered with the listener.
	Services []string `mandatory:"false" json:"services"`
}

// GetServices returns Services
func (m CloudListenerIpcEndpoint) GetServices() []string {
	return m.Services
}

func (m CloudListenerIpcEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudListenerIpcEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudListenerIpcEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudListenerIpcEndpoint CloudListenerIpcEndpoint
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeCloudListenerIpcEndpoint
	}{
		"IPC",
		(MarshalTypeCloudListenerIpcEndpoint)(m),
	}

	return json.Marshal(&s)
}
