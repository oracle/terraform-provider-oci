// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateCccInfrastructureDetails The configuration details for creating Compute Cloud@Customer infrastructure.
type CreateCccInfrastructureDetails struct {

	// The name that will be used to display the Compute Cloud@Customer infrastructure
	// in the Oracle Cloud Infrastructure console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with
	// the infrastructure.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Identifier for network subnet that will be used to communicate with Compute Cloud@Customer infrastructure.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A mutable client-meaningful text description of the Compute Cloud@Customer infrastructure.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The current connection state of the Compute Cloud@Customer infrastructure. This value
	// will default to REJECT if the value is not provided. The only valid value at creation
	// time is REJECT.
	ConnectionState CccInfrastructureConnectionStateEnum `mandatory:"false" json:"connectionState,omitempty"`

	// A message describing the current connection state in more detail.
	ConnectionDetails *string `mandatory:"false" json:"connectionDetails"`

	// Schedule used for upgrades. If no schedule is associated with the infrastructure,
	// it can be upgraded at any time.
	CccUpgradeScheduleId *string `mandatory:"false" json:"cccUpgradeScheduleId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateCccInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCccInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCccInfrastructureConnectionStateEnum(string(m.ConnectionState)); !ok && m.ConnectionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionState: %s. Supported values are: %s.", m.ConnectionState, strings.Join(GetCccInfrastructureConnectionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
