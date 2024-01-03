// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateVirtualDeploymentDetails The information to be updated.
type UpdateVirtualDeploymentDetails struct {

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	ServiceDiscovery ServiceDiscoveryConfiguration `mandatory:"false" json:"serviceDiscovery"`

	// The listeners for the virtual deployment.
	Listeners []VirtualDeploymentListener `mandatory:"false" json:"listeners"`

	AccessLogging *AccessLoggingConfiguration `mandatory:"false" json:"accessLogging"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateVirtualDeploymentDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVirtualDeploymentDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateVirtualDeploymentDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		ServiceDiscovery servicediscoveryconfiguration     `json:"serviceDiscovery"`
		Listeners        []VirtualDeploymentListener       `json:"listeners"`
		AccessLogging    *AccessLoggingConfiguration       `json:"accessLogging"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	nn, e = model.ServiceDiscovery.UnmarshalPolymorphicJSON(model.ServiceDiscovery.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ServiceDiscovery = nn.(ServiceDiscoveryConfiguration)
	} else {
		m.ServiceDiscovery = nil
	}

	m.Listeners = make([]VirtualDeploymentListener, len(model.Listeners))
	copy(m.Listeners, model.Listeners)
	m.AccessLogging = model.AccessLogging

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
