// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InfrastructureServer Database Infrastructure Server details.
type InfrastructureServer struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Data Server of Infrastructure.
	Id *string `mandatory:"true" json:"id"`

	// Database Infrastructure Server name.
	ServerName *string `mandatory:"true" json:"serverName"`

	// Database Infrastructure Server IP address.
	ServerIpAddress *string `mandatory:"true" json:"serverIpAddress"`

	// Database Infrastructure Server ILOM name.
	IlomName *string `mandatory:"true" json:"ilomName"`

	// Database Infrastructure Server ILOM IP address.
	IlomIpAddress *string `mandatory:"true" json:"ilomIpAddress"`

	// The current state of the Database Infrastructure server.
	LifecycleState InfrastructureServerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Number of database virtual machines hosted on the server.
	BaseVmCount *int `mandatory:"false" json:"baseVmCount"`

	// Number of instances hosted on the server.
	InstanceVmCount *int `mandatory:"false" json:"instanceVmCount"`

	ComputeCapacity *ComputeCapacityDetails `mandatory:"false" json:"computeCapacity"`
}

func (m InfrastructureServer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InfrastructureServer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInfrastructureServerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInfrastructureServerLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
