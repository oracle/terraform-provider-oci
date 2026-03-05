// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateDbSystemSourceFromDbSystemDetails The source DB System identifier (OCID) and region from which the new DB system will be
// cloned by copying its data. Optionally, channel properties can be provided to create a replication
// channel between the newly created DB system and the source DB system.
type CreateDbSystemSourceFromDbSystemDetails struct {

	// The OCID of the DB system to be used as the source for the new DB System.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The region identifier of the source region where the DB system exists, only if it is in a different region.
	// If the source DB system is in the same region, then no region must be specified.
	// For more information, please see Regions and Availability Domains (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm).
	Region *string `mandatory:"false" json:"region"`

	Channel *CreateDbSystemSourceFromDbSystemChannelDetails `mandatory:"false" json:"channel"`
}

func (m CreateDbSystemSourceFromDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDbSystemSourceFromDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDbSystemSourceFromDbSystemDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDbSystemSourceFromDbSystemDetails CreateDbSystemSourceFromDbSystemDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeCreateDbSystemSourceFromDbSystemDetails
	}{
		"DBSYSTEM",
		(MarshalTypeCreateDbSystemSourceFromDbSystemDetails)(m),
	}

	return json.Marshal(&s)
}
