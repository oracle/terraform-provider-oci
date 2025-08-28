// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// BdsCapacityReport A report of the host capacity within an availability domain that is available for you to create bds clusters. Host capacity is the physical infrastructure that resources such as compute instances run on. Use the capacity report to determine whether sufficient capacity is available for a shape before you create a bds cluster or change the shape of a bds cluster.
type BdsCapacityReport struct {

	// The OCID for the compartment. This should always be the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Information about the capacity of each requested shape.
	ShapeAvailabilities []CapacityReportShapeAvailability `mandatory:"true" json:"shapeAvailabilities"`

	// The time the report was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m BdsCapacityReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsCapacityReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
