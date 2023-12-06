// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccInfrastructureSummary Summary information about a Compute Cloud@Customer infrastructure.
type CccInfrastructureSummary struct {

	// The Compute Cloud@Customer infrastructure
	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This cannot be changed once created.
	Id *string `mandatory:"true" json:"id"`

	// The name that will be used to display the Compute Cloud@Customer infrastructure
	// in the Oracle Cloud Infrastructure console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with
	// the infrastructure.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network subnet that is
	// used to communicate with Compute Cloud@Customer infrastructure.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Compute Cloud@Customer infrastructure creation date and time. An RFC3339 formatted
	// datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the Compute Cloud@Customer infrastructure.
	LifecycleState CccInfrastructureLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Compute Cloud@Customer infrastructure short name. This is generated at
	// the time the resource is created and cannot be changed. The short name can be
	// used when communicating with Oracle Service and may be used during the configuration
	// of the data center network.
	ShortName *string `mandatory:"false" json:"shortName"`

	// The current connection state of the infrastructure.
	ConnectionState CccInfrastructureConnectionStateEnum `mandatory:"false" json:"connectionState,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CccInfrastructureSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccInfrastructureSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccInfrastructureLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCccInfrastructureLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCccInfrastructureConnectionStateEnum(string(m.ConnectionState)); !ok && m.ConnectionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionState: %s. Supported values are: %s.", m.ConnectionState, strings.Join(GetCccInfrastructureConnectionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
