// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.cloud.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciOptimizedStorageDetails Storage details of the database system.
type OciOptimizedStorageDetails struct {

	// Specifies if the block volume used for the database system is regional or AD-local.
	// If not specified, it will be set to false.
	// If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified.
	// If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified.
	IsRegionallyDurable *bool `mandatory:"true" json:"isRegionallyDurable"`

	// Specifies the availability domain of AD-local storage.
	// If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified.
	// If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Guaranteed input/output storage requests per second (IOPS) available to the database system.
	Iops *int64 `mandatory:"false" json:"iops"`
}

// GetIsRegionallyDurable returns IsRegionallyDurable
func (m OciOptimizedStorageDetails) GetIsRegionallyDurable() *bool {
	return m.IsRegionallyDurable
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m OciOptimizedStorageDetails) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

func (m OciOptimizedStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciOptimizedStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciOptimizedStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciOptimizedStorageDetails OciOptimizedStorageDetails
	s := struct {
		DiscriminatorParam string `json:"systemType"`
		MarshalTypeOciOptimizedStorageDetails
	}{
		"OCI_OPTIMIZED_STORAGE",
		(MarshalTypeOciOptimizedStorageDetails)(m),
	}

	return json.Marshal(&s)
}
