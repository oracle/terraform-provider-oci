// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage privateServiceAccess.
//

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateServiceAccessSummary A summary of private service access information. This object is returned when listing private service accesses.
type PrivateServiceAccessSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the
	// private service access.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private service access.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN that the private
	// service access belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet that the private service access
	// belongs to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// An OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private service access's VNIC, which
	// resides in the private service access's VCN .
	VnicId *string `mandatory:"true" json:"vnicId"`

	// The private service access's current lifecycle state.
	LifecycleState PrivateServiceAccessLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A unique service identifier for which the private service access was created.
	ServiceId *string `mandatory:"true" json:"serviceId"`

	// The private service access IPv4 FQDNs, which are going to be used to access the service.
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

	// A description of this private service access.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the private service access was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the PrivateServiceAccess was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A list of the OCIDs of the network security groups that the private service access's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private IP address (in the consumer's VCN) that represents the access point for the
	// associated service.
	Ipv4Ip *string `mandatory:"false" json:"ipv4Ip"`
}

func (m PrivateServiceAccessSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateServiceAccessSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateServiceAccessLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateServiceAccessLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
