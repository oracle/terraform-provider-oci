// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// StorageDetails Storage details of the database system.
type StorageDetails interface {

	// Specifies if the block volume used for the database system is regional or AD-local.
	// If not specified, it will be set to false.
	// If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified.
	// If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified.
	GetIsRegionallyDurable() *bool

	// Specifies the availability domain of AD-local storage.
	// If `isRegionallyDurable` is set to true, `availabilityDomain` should not be specified.
	// If `isRegionallyDurable` is set to false, `availabilityDomain` must be specified.
	GetAvailabilityDomain() *string
}

type storagedetails struct {
	JsonData            []byte
	AvailabilityDomain  *string `mandatory:"false" json:"availabilityDomain"`
	IsRegionallyDurable *bool   `mandatory:"true" json:"isRegionallyDurable"`
	SystemType          string  `json:"systemType"`
}

// UnmarshalJSON unmarshals json
func (m *storagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstoragedetails storagedetails
	s := struct {
		Model Unmarshalerstoragedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.IsRegionallyDurable = s.Model.IsRegionallyDurable
	m.AvailabilityDomain = s.Model.AvailabilityDomain
	m.SystemType = s.Model.SystemType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *storagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.SystemType {
	case "OCI_OPTIMIZED_STORAGE":
		mm := OciOptimizedStorageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for StorageDetails: %s.", m.SystemType)
		return *m, nil
	}
}

// GetAvailabilityDomain returns AvailabilityDomain
func (m storagedetails) GetAvailabilityDomain() *string {
	return m.AvailabilityDomain
}

// GetIsRegionallyDurable returns IsRegionallyDurable
func (m storagedetails) GetIsRegionallyDurable() *bool {
	return m.IsRegionallyDurable
}

func (m storagedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m storagedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
