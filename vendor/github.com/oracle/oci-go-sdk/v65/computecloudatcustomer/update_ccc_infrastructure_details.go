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

// UpdateCccInfrastructureDetails Updates Compute Cloud@Customer infrastructure configuration details.
type UpdateCccInfrastructureDetails struct {

	// The name that will be used to display the Compute Cloud@Customer infrastructure
	// in the Oracle Cloud Infrastructure console. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A mutable client-meaningful text description of the Compute Cloud@Customer infrastructure.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network subnet that is
	// used to communicate with Compute Cloud@Customer infrastructure.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An updated connection state of the Compute Cloud@Customer infrastructure.
	ConnectionState CccInfrastructureConnectionStateEnum `mandatory:"false" json:"connectionState,omitempty"`

	// A message describing the current connection state in more detail.
	ConnectionDetails *string `mandatory:"false" json:"connectionDetails"`

	// Schedule used for upgrades. If no schedule is associated with the infrastructure,
	// it can be updated at any time.
	CccUpgradeScheduleId *string `mandatory:"false" json:"cccUpgradeScheduleId"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateCccInfrastructureDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateCccInfrastructureDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCccInfrastructureConnectionStateEnum(string(m.ConnectionState)); !ok && m.ConnectionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConnectionState: %s. Supported values are: %s.", m.ConnectionState, strings.Join(GetCccInfrastructureConnectionStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
