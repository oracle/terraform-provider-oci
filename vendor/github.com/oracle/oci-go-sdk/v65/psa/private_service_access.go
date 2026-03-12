// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage Private Service Access (PSA) endpoints. PSA endpoints are used to create private access between resources in a VCN or on-premises and services in Oracle services network. For important details about how PSA endpoints work, see Access to Oracle Services: Private Service Access Endpoints (https://docs.oracle.com/iaas/Content/Network/Concepts/private-service-access.htm).
//

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateServiceAccess Private Service Access (PSA) endpoints are a new way to create private accesss to Oracle services. For important details about how PSA endpoints work, see Access to Oracle Services: Private Service Access Endpoints (https://docs.oracle.com/iaas/Content/Network/Concepts/private-service-access.htm).
type PrivateServiceAccess struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the
	// Private Service Access endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Private Service Access endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN that the private
	// service access endpoint belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the Private Service Access endpoint belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// An OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Private Service Access endpoint's VNIC, which
	// resides in the Private Service Access endpoint's VCN.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// The Private Service Access endpoint's current lifecycle state.
	LifecycleState PrivateServiceAccessLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A unique identifier for the service for which the Private Service Access endpoint was created.
	ServiceId *string `mandatory:"true" json:"serviceId"`

	// The Private Service Access endpoint's FQDN, which can be used to access the associated service.
	// Example: `xyz.oraclecloud.com`
	Fqdns []string `mandatory:"true" json:"fqdns"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// This optional field will indicate to assign IPv6 address to the Private Service Access endpoint when it is created in a dual stack (both IPv4 and IPv6) subnet.
	IsAssignDualstackIpv6 *bool `mandatory:"false" json:"isAssignDualstackIpv6"`

	// A description of this Private Service Access endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the Private Service Access endpoint was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Private Service Access endpoint was last updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A list of the OCIDs of the network security groups that the Private Service Access endpoint's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The Private Service Access endpoint's IPv6 FQDN, which can be used to access the associated service.
	// Example: `xyz.oraclecloud.com`
	FqdnsV6 []string `mandatory:"false" json:"fqdnsV6"`

	// The private IPv4 address (in the VCN) that represents the access point for the
	// associated service.
	Ipv4Ip *string `mandatory:"false" json:"ipv4Ip"`

	// The private IPv6 address (in the VCN) that represents the access point for the
	// associated service. (Optional field)
	Ipv6Ip *string `mandatory:"false" json:"ipv6Ip"`
}

func (m PrivateServiceAccess) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateServiceAccess) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateServiceAccessLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateServiceAccessLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateServiceAccessLifecycleStateEnum Enum with underlying type: string
type PrivateServiceAccessLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateServiceAccessLifecycleStateEnum
const (
	PrivateServiceAccessLifecycleStateCreating PrivateServiceAccessLifecycleStateEnum = "CREATING"
	PrivateServiceAccessLifecycleStateUpdating PrivateServiceAccessLifecycleStateEnum = "UPDATING"
	PrivateServiceAccessLifecycleStateActive   PrivateServiceAccessLifecycleStateEnum = "ACTIVE"
	PrivateServiceAccessLifecycleStateDeleting PrivateServiceAccessLifecycleStateEnum = "DELETING"
	PrivateServiceAccessLifecycleStateDeleted  PrivateServiceAccessLifecycleStateEnum = "DELETED"
	PrivateServiceAccessLifecycleStateFailed   PrivateServiceAccessLifecycleStateEnum = "FAILED"
)

var mappingPrivateServiceAccessLifecycleStateEnum = map[string]PrivateServiceAccessLifecycleStateEnum{
	"CREATING": PrivateServiceAccessLifecycleStateCreating,
	"UPDATING": PrivateServiceAccessLifecycleStateUpdating,
	"ACTIVE":   PrivateServiceAccessLifecycleStateActive,
	"DELETING": PrivateServiceAccessLifecycleStateDeleting,
	"DELETED":  PrivateServiceAccessLifecycleStateDeleted,
	"FAILED":   PrivateServiceAccessLifecycleStateFailed,
}

var mappingPrivateServiceAccessLifecycleStateEnumLowerCase = map[string]PrivateServiceAccessLifecycleStateEnum{
	"creating": PrivateServiceAccessLifecycleStateCreating,
	"updating": PrivateServiceAccessLifecycleStateUpdating,
	"active":   PrivateServiceAccessLifecycleStateActive,
	"deleting": PrivateServiceAccessLifecycleStateDeleting,
	"deleted":  PrivateServiceAccessLifecycleStateDeleted,
	"failed":   PrivateServiceAccessLifecycleStateFailed,
}

// GetPrivateServiceAccessLifecycleStateEnumValues Enumerates the set of values for PrivateServiceAccessLifecycleStateEnum
func GetPrivateServiceAccessLifecycleStateEnumValues() []PrivateServiceAccessLifecycleStateEnum {
	values := make([]PrivateServiceAccessLifecycleStateEnum, 0)
	for _, v := range mappingPrivateServiceAccessLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateServiceAccessLifecycleStateEnumStringValues Enumerates the set of values in String for PrivateServiceAccessLifecycleStateEnum
func GetPrivateServiceAccessLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPrivateServiceAccessLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateServiceAccessLifecycleStateEnum(val string) (PrivateServiceAccessLifecycleStateEnum, bool) {
	enum, ok := mappingPrivateServiceAccessLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
