// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateConnectDescriptor Connect Descriptor details. If a Private Endpoint was specified in the Connection, the host entry should be a valid IP address.
type UpdateConnectDescriptor struct {

	// Host or IP address of the connect descriptor.
	Host *string `mandatory:"false" json:"host"`

	// Port of the connect descriptor.
	Port *int `mandatory:"false" json:"port"`

	// Database service name.
	DatabaseServiceName *string `mandatory:"false" json:"databaseServiceName"`

	// Connect String. If specified, this will override the stored connect descriptor details.
	// If a Private Endpoint was specified in the Connection, the host entry should be a valid IP address.
	// Supported formats:
	// Easy connect: <host>:<port>/<db_service_name>
	// Long format: (description= (address=(port=<port>)(host=<host>))(connect_data=(service_name=<db_service_name>)))
	ConnectString *string `mandatory:"false" json:"connectString"`
}

func (m UpdateConnectDescriptor) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateConnectDescriptor) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
