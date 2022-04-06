// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// RestartDbSystemDetails DB System restart parameters.
type RestartDbSystemDetails struct {

	// The InnoDB shutdown mode to use, following the option
	// "innodb_fast_shutdown (https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_fast_shutdown)".
	ShutdownType InnoDbShutdownModeEnum `mandatory:"true" json:"shutdownType"`
}

func (m RestartDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestartDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInnoDbShutdownModeEnum(string(m.ShutdownType)); !ok && m.ShutdownType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShutdownType: %s. Supported values are: %s.", m.ShutdownType, strings.Join(GetInnoDbShutdownModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
