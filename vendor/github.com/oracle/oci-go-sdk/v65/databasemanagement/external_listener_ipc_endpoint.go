// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalListenerIpcEndpoint An `IPC`-based protocol address.
type ExternalListenerIpcEndpoint struct {

	// The unique name of the service.
	Key *string `mandatory:"true" json:"key"`

	// The list of services registered with the listener.
	Services []string `mandatory:"false" json:"services"`
}

// GetServices returns Services
func (m ExternalListenerIpcEndpoint) GetServices() []string {
	return m.Services
}

func (m ExternalListenerIpcEndpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalListenerIpcEndpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExternalListenerIpcEndpoint) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalListenerIpcEndpoint ExternalListenerIpcEndpoint
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeExternalListenerIpcEndpoint
	}{
		"IPC",
		(MarshalTypeExternalListenerIpcEndpoint)(m),
	}

	return json.Marshal(&s)
}
