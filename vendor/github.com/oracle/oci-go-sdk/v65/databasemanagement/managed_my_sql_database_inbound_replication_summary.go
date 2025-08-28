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

// ManagedMySqlDatabaseInboundReplicationSummary Inbound replication information of a MySQL replica.
type ManagedMySqlDatabaseInboundReplicationSummary struct {

	// The host name or IP address of the source this replica is connected to.
	SourceHost *string `mandatory:"true" json:"sourceHost"`

	// The port used to connect to the source.
	SourcePort *int `mandatory:"true" json:"sourcePort"`

	// The Universally Unique Identifier (UUID) value from the source server.
	SourceUuid *string `mandatory:"true" json:"sourceUuid"`

	// The current status of fetch operations.
	FetchStatus *string `mandatory:"false" json:"fetchStatus"`

	// The current status of apply operations.
	ApplyStatus *string `mandatory:"false" json:"applyStatus"`

	// The desired number of seconds that the replica must lag the source.
	DesiredDelaySeconds *int64 `mandatory:"false" json:"desiredDelaySeconds"`

	// If the replica is waiting for the desired delay seconds to pass since the source applied an event, this field contains the number of delay seconds remaining.
	RemainingDelaySeconds *int64 `mandatory:"false" json:"remainingDelaySeconds"`

	// The name of the replication channel.
	ChannelName *string `mandatory:"false" json:"channelName"`

	// The server ID value from the source server.
	SourceServerId *int64 `mandatory:"false" json:"sourceServerId"`

	// Indicates whether the channel assigns global transaction identifiers (GTIDs) to anonymous replicated transactions. OFF means no GTIDs are assigned. LOCAL means a GTID is assigned that includes this replica's own universally unique identifier (UUID). A UUID as value indicates that a GTID is assigned, which includes that manually set UUID value.
	GtidAssignment *string `mandatory:"false" json:"gtidAssignment"`

	// A list of MySqlReplicationApplierFilter records.
	ApplierFilters []MySqlReplicationApplierFilter `mandatory:"false" json:"applierFilters"`

	// The number of seconds the replica is behind the source server.
	SecondsBehindSource *int64 `mandatory:"false" json:"secondsBehindSource"`

	// The set of global transaction IDs corresponding to all transactions received by this replica from the source server. Empty if GTIDs are not in use.
	RetrievedGtidSet *string `mandatory:"false" json:"retrievedGtidSet"`

	// The total size in bytes of all the existing relay log files pertaining to this channel.
	RelayLogStorageSpaceUsed *int64 `mandatory:"false" json:"relayLogStorageSpaceUsed"`

	// The number of transactions received by this replica from the source server.
	TransactionsReceived *int64 `mandatory:"false" json:"transactionsReceived"`

	// The time in seconds that the current transaction took between being committed on the source and being applied on the replica.
	ApplyDelay *float64 `mandatory:"false" json:"applyDelay"`

	// The number of workers currently busy applying transactions from the source server.
	BusyWorkers *int `mandatory:"false" json:"busyWorkers"`

	FetchError *MySqlFetchError `mandatory:"false" json:"fetchError"`

	ApplyError *MySqlApplyError `mandatory:"false" json:"applyError"`
}

func (m ManagedMySqlDatabaseInboundReplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedMySqlDatabaseInboundReplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
