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

// DbNodeSnapshot A snapshot of the DB node details at the time of the backup.
type DbNodeSnapshot struct {

	// The OCID of the DB node.
	Id *string `mandatory:"true" json:"id"`

	// The name of the DB node.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Intended MySQL version used by the DB node.
	// You can specify or change the intended MySQL version using the create or update shared-storage DB cluster API calls.
	MysqlVersion *string `mandatory:"false" json:"mysqlVersion"`

	// Actual MySQL version used by the DB node.
	// This version is controlled by the service and could be different from the intended MySQL version (mysqlVersion)
	// for the DB node as a side effect of service maintenance events.
	CurrentMysqlVersion *string `mandatory:"false" json:"currentMysqlVersion"`

	// The name of the availability domain the DB node is located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Preference of a DB node as a potential failover target.
	// This is an integer that defines the relative weight/priority
	// of a particular DB node versus another, where lower the number,
	// the higher the preference for that DB node to become a new primary on a failover.
	PromotionTier *int `mandatory:"false" json:"promotionTier"`

	ReadEndpoint *DbNodeReadEndpoint `mandatory:"false" json:"readEndpoint"`
}

func (m DbNodeSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbNodeSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
