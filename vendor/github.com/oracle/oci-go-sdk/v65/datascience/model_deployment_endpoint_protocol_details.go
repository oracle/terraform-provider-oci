// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ModelDeploymentEndpointProtocolDetails Specifies the protocols to be used by the model deployment.
type ModelDeploymentEndpointProtocolDetails interface {

	// Boolean flag to explicitly opt in or opt out of additional protocol types.
	GetIsProtocolEnabled() *bool
}

type modeldeploymentendpointprotocoldetails struct {
	JsonData          []byte
	IsProtocolEnabled *bool  `mandatory:"false" json:"isProtocolEnabled"`
	Protocol          string `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *modeldeploymentendpointprotocoldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermodeldeploymentendpointprotocoldetails modeldeploymentendpointprotocoldetails
	s := struct {
		Model Unmarshalermodeldeploymentendpointprotocoldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsProtocolEnabled = s.Model.IsProtocolEnabled
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *modeldeploymentendpointprotocoldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "WEB_SOCKET":
		mm := ModelDeploymentWebSocketEndpointProtocolDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ModelDeploymentEndpointProtocolDetails: %s.", m.Protocol)
		return *m, nil
	}
}

// GetIsProtocolEnabled returns IsProtocolEnabled
func (m modeldeploymentendpointprotocoldetails) GetIsProtocolEnabled() *bool {
	return m.IsProtocolEnabled
}

func (m modeldeploymentendpointprotocoldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m modeldeploymentendpointprotocoldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
