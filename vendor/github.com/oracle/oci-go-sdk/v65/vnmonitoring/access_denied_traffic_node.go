// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AccessDeniedTrafficNode Defines the configuration of a traffic node to which the user is denied access.
type AccessDeniedTrafficNode struct {
	EgressTraffic *EgressTrafficSpec `mandatory:"false" json:"egressTraffic"`

	NextHopRoutingAction RoutingAction `mandatory:"false" json:"nextHopRoutingAction"`

	EgressSecurityAction SecurityAction `mandatory:"false" json:"egressSecurityAction"`

	IngressSecurityAction SecurityAction `mandatory:"false" json:"ingressSecurityAction"`
}

// GetEgressTraffic returns EgressTraffic
func (m AccessDeniedTrafficNode) GetEgressTraffic() *EgressTrafficSpec {
	return m.EgressTraffic
}

// GetNextHopRoutingAction returns NextHopRoutingAction
func (m AccessDeniedTrafficNode) GetNextHopRoutingAction() RoutingAction {
	return m.NextHopRoutingAction
}

// GetEgressSecurityAction returns EgressSecurityAction
func (m AccessDeniedTrafficNode) GetEgressSecurityAction() SecurityAction {
	return m.EgressSecurityAction
}

// GetIngressSecurityAction returns IngressSecurityAction
func (m AccessDeniedTrafficNode) GetIngressSecurityAction() SecurityAction {
	return m.IngressSecurityAction
}

func (m AccessDeniedTrafficNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessDeniedTrafficNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AccessDeniedTrafficNode) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAccessDeniedTrafficNode AccessDeniedTrafficNode
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAccessDeniedTrafficNode
	}{
		"ACCESS_DENIED",
		(MarshalTypeAccessDeniedTrafficNode)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AccessDeniedTrafficNode) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		EgressTraffic         *EgressTrafficSpec `json:"egressTraffic"`
		NextHopRoutingAction  routingaction      `json:"nextHopRoutingAction"`
		EgressSecurityAction  securityaction     `json:"egressSecurityAction"`
		IngressSecurityAction securityaction     `json:"ingressSecurityAction"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.EgressTraffic = model.EgressTraffic

	nn, e = model.NextHopRoutingAction.UnmarshalPolymorphicJSON(model.NextHopRoutingAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.NextHopRoutingAction = nn.(RoutingAction)
	} else {
		m.NextHopRoutingAction = nil
	}

	nn, e = model.EgressSecurityAction.UnmarshalPolymorphicJSON(model.EgressSecurityAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.EgressSecurityAction = nn.(SecurityAction)
	} else {
		m.EgressSecurityAction = nil
	}

	nn, e = model.IngressSecurityAction.UnmarshalPolymorphicJSON(model.IngressSecurityAction.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.IngressSecurityAction = nn.(SecurityAction)
	} else {
		m.IngressSecurityAction = nil
	}

	return
}
