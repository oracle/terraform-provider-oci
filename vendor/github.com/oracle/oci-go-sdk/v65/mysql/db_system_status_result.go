// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// DbSystemStatusResult The status collected from the DB System.
type DbSystemStatusResult struct {

	// Specifies if the Read/Write endpoint of the DB System can accept client connections.
	CanAcceptClientConnections *bool `mandatory:"false" json:"canAcceptClientConnections"`

	// Specifies if the DB System is in offline mode based on the value of the MySQL system variable
	// offline_mode (https://dev.mysql.com/doc/en/server-system-variables.html#sysvar_offline_mode).
	// If True, access is allowed only to users with specific privileges. If False, access is allowed for all MySQL
	// user accounts.
	IsInOfflineMode *bool `mandatory:"false" json:"isInOfflineMode"`

	// Specifies if connected users can run write queries on the DB System.
	IsWritable *bool `mandatory:"false" json:"isWritable"`

	// Specifies if the single MySQL instance in a standalone DB System or all MySQL instances in a highly available
	// DB System (excluding read replicas) are healthy.
	AreAllMysqlInstancesHealthy *bool `mandatory:"false" json:"areAllMysqlInstancesHealthy"`

	// Specifies if there is any MySQL instance (excluding read replicas) whose storage capacity is below storage reserve
	//  (https://docs.oracle.com/en-us/iaas/mysql-database/doc/health-monitor.html#GUID-C6CE25C7-B728-4C80-B548-A76B42005C83.html).
	IsStorageFull *bool `mandatory:"false" json:"isStorageFull"`

	// The GTID set on the DB System (either GTID_EXECUTED or GTID_AVAILABLE) as specified by the gtidSetType
	// parameter in the request.
	GtidSet *string `mandatory:"false" json:"gtidSet"`

	// Specifies if the GTID set in the gtidSetToApply parameter (if provided in the request) is fully applied
	// on the DB System. If gtidSetToApply is not provided in the request or is "", this field will be empty.
	IsGtidSetApplied *bool `mandatory:"false" json:"isGtidSetApplied"`

	// A list showing the status of the channels attached to the DB System for the requested channel IDs.
	ChannelsStatus []ChannelStatusResult `mandatory:"false" json:"channelsStatus"`
}

func (m DbSystemStatusResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemStatusResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
