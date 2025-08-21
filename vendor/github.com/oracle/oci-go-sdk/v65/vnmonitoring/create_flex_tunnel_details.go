// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateFlexTunnelDetails These details can be included in a request to create a flex tunnel.
type CreateFlexTunnelDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the flex tunnel.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG for loopback attachment.
	DrgId *string `mandatory:"true" json:"drgId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the transport attachment.
	TransportAttachmentId *string `mandatory:"true" json:"transportAttachmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.
	DrgRouteTableId *string `mandatory:"true" json:"drgRouteTableId"`

	FlexTunnelConfiguration CreateFlexTunnelConfigurationDetails `mandatory:"true" json:"flexTunnelConfiguration"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateFlexTunnelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFlexTunnelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateFlexTunnelDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                              `json:"displayName"`
		FreeformTags            map[string]string                    `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{}    `json:"definedTags"`
		CompartmentId           *string                              `json:"compartmentId"`
		DrgId                   *string                              `json:"drgId"`
		TransportAttachmentId   *string                              `json:"transportAttachmentId"`
		DrgRouteTableId         *string                              `json:"drgRouteTableId"`
		FlexTunnelConfiguration createflextunnelconfigurationdetails `json:"flexTunnelConfiguration"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.DrgId = model.DrgId

	m.TransportAttachmentId = model.TransportAttachmentId

	m.DrgRouteTableId = model.DrgRouteTableId

	nn, e = model.FlexTunnelConfiguration.UnmarshalPolymorphicJSON(model.FlexTunnelConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.FlexTunnelConfiguration = nn.(CreateFlexTunnelConfigurationDetails)
	} else {
		m.FlexTunnelConfiguration = nil
	}

	return
}
