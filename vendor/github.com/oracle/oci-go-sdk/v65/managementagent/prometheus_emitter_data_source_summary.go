// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// PrometheusEmitterDataSourceSummary A Prometheus emitter data source summary.
type PrometheusEmitterDataSourceSummary struct {

	// ID for DataSource.
	Id *string `mandatory:"true" json:"id"`

	// Unique name of the dataSource.
	DataSourceName *string `mandatory:"true" json:"dataSourceName"`
}

//GetId returns Id
func (m PrometheusEmitterDataSourceSummary) GetId() *string {
	return m.Id
}

//GetDataSourceName returns DataSourceName
func (m PrometheusEmitterDataSourceSummary) GetDataSourceName() *string {
	return m.DataSourceName
}

func (m PrometheusEmitterDataSourceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrometheusEmitterDataSourceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PrometheusEmitterDataSourceSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrometheusEmitterDataSourceSummary PrometheusEmitterDataSourceSummary
	s := struct {
		DiscriminatorParam string `json:"dataSourceType"`
		MarshalTypePrometheusEmitterDataSourceSummary
	}{
		"PROMETHEUS_EMITTER",
		(MarshalTypePrometheusEmitterDataSourceSummary)(m),
	}

	return json.Marshal(&s)
}
