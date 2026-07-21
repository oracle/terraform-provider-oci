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

// BdsCapacityReservationConfigurationSummary Summary of a configuration between a BDS cluster and a BDS capacity reservation.
type BdsCapacityReservationConfigurationSummary struct {

	// The OCID of the BDS capacity reservation configuration.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the BDS cluster associated with the BDS capacity reservation.
	BdsInstanceId *string `mandatory:"true" json:"bdsInstanceId"`

	// The OCID of the BDS capacity reservation associated with the BDS cluster.
	BdsCapacityReservationId *string `mandatory:"true" json:"bdsCapacityReservationId"`

	// The display name of the BDS capacity reservation configuration.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The lifecycle state of the BDS capacity reservation configuration.
	LifecycleState BdsCapacityReservationConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the BDS capacity reservation configuration was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the BDS capacity reservation configuration was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BdsCapacityReservationConfigurationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BdsCapacityReservationConfigurationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBdsCapacityReservationConfigurationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBdsCapacityReservationConfigurationLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
