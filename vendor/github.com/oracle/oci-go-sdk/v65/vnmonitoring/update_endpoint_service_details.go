// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateEndpointServiceDetails Information that can be updated for an endpoint service.
type UpdateEndpointServiceDetails struct {

	// A description of the endpoint service. For Oracle services that use the "trusted" mode of the
	// private endpoint service, customers never see this description. Avoid entering confidential
	// information.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Some services want to restrict access to the resources represented by an endpoint service so
	// that only a single private endpoint in the customer VCN has access.
	// For example, the endpoint service might represent a particular service resource (such as a
	// particular database). The service might want to allow access to that particular resource
	// from only a single private endpoint.
	// Defaults to `false`.
	// Example: `true`
	AreMultiplePrivateEndpointsPerVcnAllowed *bool `mandatory:"false" json:"areMultiplePrivateEndpointsPerVcnAllowed"`

	// Reserved for future use.
	IsVcnMetadataEnabled *bool `mandatory:"false" json:"isVcnMetadataEnabled"`

	// List of service IP addresses (in the service VCN) that handle requests to the endpoint service.
	ServiceIps []EndpointServiceIpDetails `mandatory:"false" json:"serviceIps"`

	// The ports on the endpoint service IPs that are open for private endpoint traffic for this
	// endpoint service. If you provide no ports, all open ports on the service IPs are accessible.
	Ports []int `mandatory:"false" json:"ports"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdateEndpointServiceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateEndpointServiceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
