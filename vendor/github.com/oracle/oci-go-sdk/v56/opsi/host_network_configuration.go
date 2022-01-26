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

// HostNetworkConfiguration Network Configuration metric for the host
type HostNetworkConfiguration struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Name of the network interface
	InterfaceName *string `mandatory:"true" json:"interfaceName"`

	// IP address (IPv4 or IPv6) of the network interface
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// MAC address of the network interface. MAC address is a 12-digit hexadecimal number separated by colons or dashes or dots. Following formats are accepted: MM:MM:MM:SS:SS:SS, MM-MM-MM-SS-SS-SS, MM.MM.MM.SS.SS.SS, MMM:MMM:SSS:SSS, MMM-MMM-SSS-SSS, MMM.MMM.SSS.SSS, MMMM:MMSS:SSSS, MMMM-MMSS-SSSS, MMMM.MMSS.SSSS
	MacAddress *string `mandatory:"false" json:"macAddress"`
}

//GetTimeCollected returns TimeCollected
func (m HostNetworkConfiguration) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostNetworkConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostNetworkConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostNetworkConfiguration HostNetworkConfiguration
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostNetworkConfiguration
	}{
		"HOST_NETWORK_CONFIGURATION",
		(MarshalTypeHostNetworkConfiguration)(m),
	}

	return json.Marshal(&s)
}
