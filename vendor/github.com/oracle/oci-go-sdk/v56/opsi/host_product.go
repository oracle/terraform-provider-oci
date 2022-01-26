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

// HostProduct Product metric for the host
type HostProduct struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Vendor of the product
	Vendor *string `mandatory:"false" json:"vendor"`

	// Name of the product
	Name *string `mandatory:"false" json:"name"`

	// Version of the product
	Version *string `mandatory:"false" json:"version"`
}

//GetTimeCollected returns TimeCollected
func (m HostProduct) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostProduct) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostProduct) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostProduct HostProduct
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostProduct
	}{
		"HOST_PRODUCT",
		(MarshalTypeHostProduct)(m),
	}

	return json.Marshal(&s)
}
