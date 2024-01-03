// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicaOverrides By default a read replica inherits the MySQL version, shape, and configuration of the source DB system.
// If you want to override any of these, provide values in the properties, mysqlVersion, shapeName,
// and configurationId. If you set a property value to "", then the value is inherited from its
// source DB system.
type ReplicaOverrides struct {

	// The MySQL version to be used by the read replica.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// The shape to be used by the read replica. The shape determines the resources allocated:
	// CPU cores and memory for VM shapes, CPU cores, memory and storage for non-VM (bare metal) shapes.
	// To get a list of shapes, use the ListShapes operation.
	ShapeName *string `mandatory:"false" json:"shapeName"`

	// The OCID of the Configuration to be used by the read replica.
	ConfigurationId *string `mandatory:"false" json:"configurationId"`
}

func (m ReplicaOverrides) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicaOverrides) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
