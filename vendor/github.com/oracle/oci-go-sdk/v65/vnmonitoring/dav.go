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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Dav A Direct Attached Vnic.
type Dav struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Direct Attached Vnic.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Direct Attached Vnic's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the Direct Attached Vnic.
	LifecycleState DavLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Index of NIC for Direct Attached Vnic.
	NicIndex *int `mandatory:"true" json:"nicIndex"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// The MAC address for the DAV. This will be a newly allocated MAC address
	// and not the one used by the instance.
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// The substrate IP of DAV and primary VNIC attached to the instance.
	// This field will be null in case the DAV is not attached.
	SubstrateIp *string `mandatory:"false" json:"substrateIp"`

	// The allocated slot id for the Dav.
	SlotId *int `mandatory:"false" json:"slotId"`

	// The VLAN Tag assigned to Direct Attached Vnic.
	VlanTag *int `mandatory:"false" json:"vlanTag"`

	// The MAC address of the Virtual Router.
	VirtualRouterMac *string `mandatory:"false" json:"virtualRouterMac"`

	// Substrate IP address of the remote endpoint.
	RemoteEndpointSubstrateIp *string `mandatory:"false" json:"remoteEndpointSubstrateIp"`

	// List of VCNx Attachments to a DRG.
	VcnxAttachmentIds []string `mandatory:"false" json:"vcnxAttachmentIds"`

	// The label type for Direct Attached Vnic. This is used to determine the
	// label forwarding to be used by the Direct Attached Vnic.
	LabelType DavLabelTypeEnum `mandatory:"false" json:"labelType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Dav) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Dav) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDavLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDavLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDavLabelTypeEnum(string(m.LabelType)); !ok && m.LabelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LabelType: %s. Supported values are: %s.", m.LabelType, strings.Join(GetDavLabelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DavLifecycleStateEnum Enum with underlying type: string
type DavLifecycleStateEnum string

// Set of constants representing the allowable values for DavLifecycleStateEnum
const (
	DavLifecycleStateProvisioning DavLifecycleStateEnum = "PROVISIONING"
	DavLifecycleStateUpdating     DavLifecycleStateEnum = "UPDATING"
	DavLifecycleStateAvailable    DavLifecycleStateEnum = "AVAILABLE"
	DavLifecycleStateTerminating  DavLifecycleStateEnum = "TERMINATING"
	DavLifecycleStateTerminated   DavLifecycleStateEnum = "TERMINATED"
)

var mappingDavLifecycleStateEnum = map[string]DavLifecycleStateEnum{
	"PROVISIONING": DavLifecycleStateProvisioning,
	"UPDATING":     DavLifecycleStateUpdating,
	"AVAILABLE":    DavLifecycleStateAvailable,
	"TERMINATING":  DavLifecycleStateTerminating,
	"TERMINATED":   DavLifecycleStateTerminated,
}

var mappingDavLifecycleStateEnumLowerCase = map[string]DavLifecycleStateEnum{
	"provisioning": DavLifecycleStateProvisioning,
	"updating":     DavLifecycleStateUpdating,
	"available":    DavLifecycleStateAvailable,
	"terminating":  DavLifecycleStateTerminating,
	"terminated":   DavLifecycleStateTerminated,
}

// GetDavLifecycleStateEnumValues Enumerates the set of values for DavLifecycleStateEnum
func GetDavLifecycleStateEnumValues() []DavLifecycleStateEnum {
	values := make([]DavLifecycleStateEnum, 0)
	for _, v := range mappingDavLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDavLifecycleStateEnumStringValues Enumerates the set of values in String for DavLifecycleStateEnum
func GetDavLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"UPDATING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingDavLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDavLifecycleStateEnum(val string) (DavLifecycleStateEnum, bool) {
	enum, ok := mappingDavLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DavLabelTypeEnum Enum with underlying type: string
type DavLabelTypeEnum string

// Set of constants representing the allowable values for DavLabelTypeEnum
const (
	DavLabelTypeMpls DavLabelTypeEnum = "MPLS"
)

var mappingDavLabelTypeEnum = map[string]DavLabelTypeEnum{
	"MPLS": DavLabelTypeMpls,
}

var mappingDavLabelTypeEnumLowerCase = map[string]DavLabelTypeEnum{
	"mpls": DavLabelTypeMpls,
}

// GetDavLabelTypeEnumValues Enumerates the set of values for DavLabelTypeEnum
func GetDavLabelTypeEnumValues() []DavLabelTypeEnum {
	values := make([]DavLabelTypeEnum, 0)
	for _, v := range mappingDavLabelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDavLabelTypeEnumStringValues Enumerates the set of values in String for DavLabelTypeEnum
func GetDavLabelTypeEnumStringValues() []string {
	return []string{
		"MPLS",
	}
}

// GetMappingDavLabelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDavLabelTypeEnum(val string) (DavLabelTypeEnum, bool) {
	enum, ok := mappingDavLabelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
