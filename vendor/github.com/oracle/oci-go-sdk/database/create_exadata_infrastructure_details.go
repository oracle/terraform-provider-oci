// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateExadataInfrastructureDetails Request to create Exadata infrastructure.
type CreateExadataInfrastructureDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Exadata infrastructure. The name does not need to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"true" json:"shape"`

	// The time zone of the Exadata infrastructure. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"true" json:"timeZone"`

	// The IP address for the first control plane server.
	CloudControlPlaneServer1 *string `mandatory:"true" json:"cloudControlPlaneServer1"`

	// The IP address for the second control plane server.
	CloudControlPlaneServer2 *string `mandatory:"true" json:"cloudControlPlaneServer2"`

	// The netmask for the control plane network.
	Netmask *string `mandatory:"true" json:"netmask"`

	// The gateway for the control plane network.
	Gateway *string `mandatory:"true" json:"gateway"`

	// The CIDR block for the Exadata administration network.
	AdminNetworkCIDR *string `mandatory:"true" json:"adminNetworkCIDR"`

	// The CIDR block for the Exadata InfiniBand interconnect.
	InfiniBandNetworkCIDR *string `mandatory:"true" json:"infiniBandNetworkCIDR"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServer []string `mandatory:"true" json:"dnsServer"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServer []string `mandatory:"true" json:"ntpServer"`

	// The corporate network proxy for access to the control plane network. Oracle recommends using an HTTPS proxy when possible
	// for enhanced security.
	CorporateProxy *string `mandatory:"false" json:"corporateProxy"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}
