// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedMySqlDatabaseGeneralReplicationInformation General information about the replication of a MySQL server.
type ManagedMySqlDatabaseGeneralReplicationInformation struct {

	// This server's ID.
	ServerId *int64 `mandatory:"true" json:"serverId"`

	// This server's Universally Unique Identifier (UUID).
	ServerUuid *string `mandatory:"true" json:"serverUuid"`

	// If the value is ON, the instance is configured as read_only. If the value is SUPER, the instance is configured as super_read_only. If the value is OFF, the instance is neither read_only nor super_read_only.
	ReadOnly MySqlReadOnlyEnum `mandatory:"true" json:"readOnly"`

	// The type of the instance for example, Source, Replica, Primary Group Member, and Secondary Group Member. If the instance is replicating from one or more sources and has one or more replicas, which means, it belongs to a replication chain, the instance type can be Replica/Source.
	InstanceType *string `mandatory:"false" json:"instanceType"`

	// This server's host name.
	HostName *string `mandatory:"false" json:"hostName"`

	// The number of the port on which the server listens for TCP/IP connections.
	Port *int `mandatory:"false" json:"port"`

	// The number of seconds the replica is behind the source. When multiple sources are involved, this is the maximum value across all sources.
	SecondsBehindSourceMax *int64 `mandatory:"false" json:"secondsBehindSourceMax"`

	// A summary of the current status of fetch operations.
	FetchStatusSummary *string `mandatory:"false" json:"fetchStatusSummary"`

	// A summary of the current status of apply operations.
	ApplyStatusSummary *string `mandatory:"false" json:"applyStatusSummary"`

	// Specifies if high availability is enabled on this server.
	IsHighAvailabilityEnabled *bool `mandatory:"false" json:"isHighAvailabilityEnabled"`

	// The state of this server as a group replication member.
	HighAvailabilityMemberState *string `mandatory:"false" json:"highAvailabilityMemberState"`

	// The number of sources this server is replicating from.
	InboundReplicationsCount *int `mandatory:"false" json:"inboundReplicationsCount"`

	// The Global Transaction Identifier (GTID) mode of this server.
	GtidMode *string `mandatory:"false" json:"gtidMode"`

	// The set of global transaction identifiers for transactions that have been executed on this source server.
	ExecutedGtidSet *string `mandatory:"false" json:"executedGtidSet"`

	// The status of binary logging on this server.
	BinaryLogging *string `mandatory:"false" json:"binaryLogging"`

	// The binary logging format used by this server.
	BinaryLogFormat *string `mandatory:"false" json:"binaryLogFormat"`

	// The number of replicas replicating from this server.
	OutboundReplicationsCount *int `mandatory:"false" json:"outboundReplicationsCount"`
}

func (m ManagedMySqlDatabaseGeneralReplicationInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabaseGeneralReplicationInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMySqlReadOnlyEnum(string(m.ReadOnly)); !ok && m.ReadOnly != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReadOnly: %s. Supported values are: %s.", m.ReadOnly, strings.Join(GetMySqlReadOnlyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
