// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateExadataInfrastructureDetails Updates the Exadata infrastructure. Applies to Exadata Cloud@Customer instances only.
// See UpdateCloudExadataInfrastructureDetails for information on updating Exadata Cloud Service cloud Exadata infrastructure resources.
type UpdateExadataInfrastructureDetails struct {

	// The IP address for the first control plane server.
	CloudControlPlaneServer1 *string `mandatory:"false" json:"cloudControlPlaneServer1"`

	// The IP address for the second control plane server.
	CloudControlPlaneServer2 *string `mandatory:"false" json:"cloudControlPlaneServer2"`

	// The netmask for the control plane network.
	Netmask *string `mandatory:"false" json:"netmask"`

	// The gateway for the control plane network.
	Gateway *string `mandatory:"false" json:"gateway"`

	// The CIDR block for the Exadata administration network.
	AdminNetworkCIDR *string `mandatory:"false" json:"adminNetworkCIDR"`

	// The CIDR block for the Exadata InfiniBand interconnect.
	InfiniBandNetworkCIDR *string `mandatory:"false" json:"infiniBandNetworkCIDR"`

	// The corporate network proxy for access to the control plane network.
	CorporateProxy *string `mandatory:"false" json:"corporateProxy"`

	// The list of contacts for the Exadata infrastructure.
	Contacts []ExadataInfrastructureContact `mandatory:"false" json:"contacts"`

	MaintenanceWindow *MaintenanceWindow `mandatory:"false" json:"maintenanceWindow"`

	// The requested number of additional storage servers for the Exadata infrastructure.
	AdditionalStorageCount *int `mandatory:"false" json:"additionalStorageCount"`

	// The requested number of additional compute servers for the Exadata infrastructure.
	AdditionalComputeCount *int `mandatory:"false" json:"additionalComputeCount"`

	// Oracle Exadata System Model specification. The system model determines the amount of compute or storage
	// server resources available for use. For more information, please see System and Shape Configuration Options
	//  (https://docs.oracle.com/en/engineered-systems/exadata-cloud-at-customer/ecccm/ecc-system-config-options.html#GUID-9E090174-5C57-4EB1-9243-B470F9F10D6B)
	AdditionalComputeSystemModel UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum `mandatory:"false" json:"additionalComputeSystemModel,omitempty"`

	// The list of DNS server IP addresses. Maximum of 3 allowed.
	DnsServer []string `mandatory:"false" json:"dnsServer"`

	// The list of NTP server IP addresses. Maximum of 3 allowed.
	NtpServer []string `mandatory:"false" json:"ntpServer"`

	// The time zone of the Exadata infrastructure. For details, see Exadata Infrastructure Time Zones (https://docs.cloud.oracle.com/Content/Database/References/timezones.htm).
	TimeZone *string `mandatory:"false" json:"timeZone"`

	// Indicates whether cps offline diagnostic report is enabled for this Exadata infrastructure. This will allow a customer to quickly check status themselves and fix problems on their end, saving time and frustration
	// for both Oracle and the customer when they find the CPS in a disconnected state.You can enable offline diagnostic report during Exadata infrastructure provisioning. You can also disable or enable it at any time
	// using the UpdateExadatainfrastructure API.
	IsCpsOfflineReportEnabled *bool `mandatory:"false" json:"isCpsOfflineReportEnabled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateExadataInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateExadataInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum(string(m.AdditionalComputeSystemModel)); !ok && m.AdditionalComputeSystemModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AdditionalComputeSystemModel: %s. Supported values are: %s.", m.AdditionalComputeSystemModel, strings.Join(GetUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum Enum with underlying type: string
type UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum string

// Set of constants representing the allowable values for UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum
const (
	UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX7  UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum = "X7"
	UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX8  UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum = "X8"
	UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX8m UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum = "X8M"
	UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX9m UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum = "X9M"
)

var mappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum = map[string]UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum{
	"X7":  UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX7,
	"X8":  UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX8,
	"X8M": UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX8m,
	"X9M": UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX9m,
}

var mappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumLowerCase = map[string]UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum{
	"x7":  UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX7,
	"x8":  UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX8,
	"x8m": UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX8m,
	"x9m": UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelX9m,
}

// GetUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumValues Enumerates the set of values for UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum
func GetUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumValues() []UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum {
	values := make([]UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum, 0)
	for _, v := range mappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumStringValues Enumerates the set of values in String for UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum
func GetUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumStringValues() []string {
	return []string{
		"X7",
		"X8",
		"X8M",
		"X9M",
	}
}

// GetMappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum(val string) (UpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnum, bool) {
	enum, ok := mappingUpdateExadataInfrastructureDetailsAdditionalComputeSystemModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
