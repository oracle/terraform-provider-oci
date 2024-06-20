// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigureShardedDatabaseGsmsDetails Details of the request to configure new global service manager(GSM) instances for the sharded database.
type ConfigureShardedDatabaseGsmsDetails struct {

	// Names of old global service manager(GSM) instances corresponding to which new GSM instances need to be configured.
	OldGsmNames []string `mandatory:"true" json:"oldGsmNames"`

	// Flag to indicate if new global service manager(GSM) instances shall use latest image or re-use image used by existing
	// GSM instances.
	IsLatestGsmImage *bool `mandatory:"true" json:"isLatestGsmImage"`
}

func (m ConfigureShardedDatabaseGsmsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConfigureShardedDatabaseGsmsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
