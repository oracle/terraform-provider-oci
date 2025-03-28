// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseTargetSummary Details of a Database target member of a Exadata Fleet Update Collection.
// Stored references of the resource documented in
//
//	https://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/Database/
type DatabaseTargetSummary struct {

	// OCID of the target resource in the Exadata Fleet Update Collection.
	Id *string `mandatory:"false" json:"id"`

	// Compartment identifier of the target.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCID of the database home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// OCID of the related VM Cluster or Cloud VM Cluster.
	VmClusterId *string `mandatory:"false" json:"vmClusterId"`

	// OCID of the related Exadata Infrastructure or Cloud Exadata Infrastructure resource.
	InfrastructureId *string `mandatory:"false" json:"infrastructureId"`

	// OCID of the Database sofware image.
	SoftwareImageId *string `mandatory:"false" json:"softwareImageId"`
}

// GetId returns Id
func (m DatabaseTargetSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m DatabaseTargetSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m DatabaseTargetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseTargetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DatabaseTargetSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDatabaseTargetSummary DatabaseTargetSummary
	s := struct {
		DiscriminatorParam string `json:"entityType"`
		MarshalTypeDatabaseTargetSummary
	}{
		"DATABASE",
		(MarshalTypeDatabaseTargetSummary)(m),
	}

	return json.Marshal(&s)
}
