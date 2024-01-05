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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NetworkDetails Network details for the database system.
type NetworkDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer subnet associated with the database system.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// Private IP in customer subnet. The value is optional.
	// If the IP is not provided, the IP will be chosen from the available IP addresses from the specified subnet.
	PrimaryDbEndpointPrivateIp *string `mandatory:"false" json:"primaryDbEndpointPrivateIp"`

	// List of customer Network Security Group OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with the database system.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m NetworkDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NetworkDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
