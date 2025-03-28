// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateEndpointSummary A Data Flow private endpoint object used in bulk listings.
type PrivateEndpointSummary struct {

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// A user-friendly name. It does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// An array of DNS zone names.
	// Example: `[ "app.examplecorp.com", "app.examplecorp2.com" ]`
	DnsZones []string `mandatory:"true" json:"dnsZones"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// The OCID of a private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this private endpoint.
	LifecycleState PrivateEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"true" json:"ownerPrincipalId"`

	// The OCID of a subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The maximum number of hosts to be accessed through the private endpoint. This value is used
	// to calculate the relevant CIDR block and should be a multiple of 256.  If the value is not a
	// multiple of 256, it is rounded up to the next multiple of 256. For example, 300 is rounded up
	// to 512.
	MaxHostCount *int `mandatory:"false" json:"maxHostCount"`

	// An array of network security group OCIDs.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The username of the user who created the resource.  If the username of the owner does not exist,
	// `null` will be returned and the caller should refer to the ownerPrincipalId value instead.
	OwnerUserName *string `mandatory:"false" json:"ownerUserName"`

	// An array of fqdn/port pairs used to create private endpoint. Each object is a simple key-value pair with FQDN as key and port number as value.
	// [ { fqdn: "scan1.oracle.com", port: "1521"}, { fqdn: "scan2.oracle.com", port: "1521" } ]
	ScanDetails []Scan `mandatory:"false" json:"scanDetails"`
}

func (m PrivateEndpointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
