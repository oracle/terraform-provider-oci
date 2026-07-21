// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeCapacityReservations Compute capacity reservation ID mappings by domain. For a multi-AD region, domain1, domain2,
// and domain3 correspond to AD1, AD2, and AD3 respectively. For a single-AD region, domain1,
// domain2, and domain3 correspond to FD1, FD2, and FD3 respectively.
type ComputeCapacityReservations struct {

	// Capacity reservation OCID corresponding to AD1 for a multi-AD region or FD1 for a single-AD region.
	Domain1ReservationId *string `mandatory:"false" json:"domain1ReservationId"`

	// Capacity reservation OCID corresponding to AD2 for a multi-AD region or FD2 for a single-AD region.
	Domain2ReservationId *string `mandatory:"false" json:"domain2ReservationId"`

	// Capacity reservation OCID corresponding to AD3 for a multi-AD region or FD3 for a single-AD region.
	Domain3ReservationId *string `mandatory:"false" json:"domain3ReservationId"`
}

func (m ComputeCapacityReservations) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeCapacityReservations) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
