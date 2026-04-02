// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PointInTimeDbSystemSourceDetails Details of database system point-in-time recovery.
type PointInTimeDbSystemSourceDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source database system which will be used to perform point-in-time recovery.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The target point-in-time of the source database system that will be restored, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Point-in-time recovery can only performed in granularity of seconds.
	// Example: `2016-08-25T21:10:29Z`
	TimeToRestore *common.SDKTime `mandatory:"true" json:"timeToRestore"`
}

func (m PointInTimeDbSystemSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PointInTimeDbSystemSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PointInTimeDbSystemSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePointInTimeDbSystemSourceDetails PointInTimeDbSystemSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypePointInTimeDbSystemSourceDetails
	}{
		"POINT_IN_TIME_DB_SYSTEM",
		(MarshalTypePointInTimeDbSystemSourceDetails)(m),
	}

	return json.Marshal(&s)
}
