// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v45/common"
)

// HostDetails Partial information about a host which includes id, name, type.
type HostDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host.
	Id *string `mandatory:"true" json:"id"`

	// The host name. The host name is unique amongst the hosts managed by the same management agent.
	HostName *string `mandatory:"true" json:"hostName"`

	// Platform type.
	PlatformType HostDetailsPlatformTypeEnum `mandatory:"true" json:"platformType"`

	// The identifier of the agent.
	AgentIdentifier *string `mandatory:"true" json:"agentIdentifier"`

	// The user-friendly name for the host. The name does not have to be unique.
	HostDisplayName *string `mandatory:"false" json:"hostDisplayName"`
}

func (m HostDetails) String() string {
	return common.PointerString(m)
}

// HostDetailsPlatformTypeEnum Enum with underlying type: string
type HostDetailsPlatformTypeEnum string

// Set of constants representing the allowable values for HostDetailsPlatformTypeEnum
const (
	HostDetailsPlatformTypeLinux HostDetailsPlatformTypeEnum = "LINUX"
)

var mappingHostDetailsPlatformType = map[string]HostDetailsPlatformTypeEnum{
	"LINUX": HostDetailsPlatformTypeLinux,
}

// GetHostDetailsPlatformTypeEnumValues Enumerates the set of values for HostDetailsPlatformTypeEnum
func GetHostDetailsPlatformTypeEnumValues() []HostDetailsPlatformTypeEnum {
	values := make([]HostDetailsPlatformTypeEnum, 0)
	for _, v := range mappingHostDetailsPlatformType {
		values = append(values, v)
	}
	return values
}
