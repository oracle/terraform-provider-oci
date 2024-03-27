// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HeatWaveNode The information about an individual HeatWave node.
type HeatWaveNode struct {

	// The ID associated with the HeatWave node.
	Id *string `mandatory:"true" json:"id"`

	// The status of the HeatWave node. Indicates whether the status of the node is UP, DOWN, or UNKNOWN at the current time.
	Status HeatWaveNodeStatusEnum `mandatory:"true" json:"status"`

	// The date and time the node was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`
}

func (m HeatWaveNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeatWaveNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHeatWaveNodeStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetHeatWaveNodeStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
