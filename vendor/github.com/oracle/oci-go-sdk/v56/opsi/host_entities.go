// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// HostEntities Database entities running on the host
type HostEntities struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Name of the database entity
	EntityName *string `mandatory:"true" json:"entityName"`

	// Type of the database entity
	EntityType *string `mandatory:"true" json:"entityType"`
}

//GetTimeCollected returns TimeCollected
func (m HostEntities) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostEntities) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostEntities) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostEntities HostEntities
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostEntities
	}{
		"HOST_ENTITIES",
		(MarshalTypeHostEntities)(m),
	}

	return json.Marshal(&s)
}
