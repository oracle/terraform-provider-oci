// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// Use the Management Agent API to manage your infrastructure's management agents, including their plugins and install keys.
// For more information, see Management Agent (https://docs.cloud.oracle.com/iaas/management-agents/index.html).
//

package managementagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSourceSummary The information about the dataSources that agent is associated to.
type DataSourceSummary interface {

	// Data source type and name identifier.
	GetKey() *string

	// Unique name of the dataSource.
	GetName() *string
}

type datasourcesummary struct {
	JsonData []byte
	Key      *string `mandatory:"true" json:"key"`
	Name     *string `mandatory:"true" json:"name"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *datasourcesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasourcesummary datasourcesummary
	s := struct {
		Model Unmarshalerdatasourcesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.Name = s.Model.Name
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasourcesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "PROMETHEUS_EMITTER":
		mm := PrometheusEmitterDataSourceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KUBERNETES_CLUSTER":
		mm := KubernetesClusterDataSourceSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for DataSourceSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetKey returns Key
func (m datasourcesummary) GetKey() *string {
	return m.Key
}

// GetName returns Name
func (m datasourcesummary) GetName() *string {
	return m.Name
}

func (m datasourcesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m datasourcesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
