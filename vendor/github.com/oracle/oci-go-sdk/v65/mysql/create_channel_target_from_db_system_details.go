// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateChannelTargetFromDbSystemDetails Parameters detailing how to provision the target endpoint that is a DB System.
type CreateChannelTargetFromDbSystemDetails struct {

	// The OCID of the target DB System.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The case-insensitive name that identifies the replication channel. Channel names
	// must follow the rules defined for MySQL identifiers (https://dev.mysql.com/doc/refman/8.0/en/identifiers.html).
	// The names of non-Deleted Channels must be unique for each DB System.
	ChannelName *string `mandatory:"false" json:"channelName"`

	// The username for the replication applier of the target MySQL DB System.
	ApplierUsername *string `mandatory:"false" json:"applierUsername"`

	// Replication filter rules to be applied at the DB System Channel target.
	Filters []ChannelFilter `mandatory:"false" json:"filters"`

	// Specifies the amount of time, in seconds, that the channel waits before
	// applying a transaction received from the source.
	DelayInSeconds *int `mandatory:"false" json:"delayInSeconds"`

	// Specifies how a replication channel handles the creation and alteration of tables
	// that do not have a primary key. The default value is set to ALLOW.
	TablesWithoutPrimaryKeyHandling ChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum `mandatory:"false" json:"tablesWithoutPrimaryKeyHandling,omitempty"`
}

func (m CreateChannelTargetFromDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateChannelTargetFromDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnum(string(m.TablesWithoutPrimaryKeyHandling)); !ok && m.TablesWithoutPrimaryKeyHandling != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TablesWithoutPrimaryKeyHandling: %s. Supported values are: %s.", m.TablesWithoutPrimaryKeyHandling, strings.Join(GetChannelTargetDbSystemTablesWithoutPrimaryKeyHandlingEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateChannelTargetFromDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateChannelTargetFromDbSystemDetails CreateChannelTargetFromDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"targetType"`
		MarshalTypeCreateChannelTargetFromDbSystemDetails
	}{
		"DBSYSTEM",
		(MarshalTypeCreateChannelTargetFromDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}
