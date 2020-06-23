// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/common"
)

// StopDbSystemDetails DB System shutdown parameters.
type StopDbSystemDetails struct {

	// The InnoDB shutdown mode to use, following the option
	// "innodb_fast_shutdown (https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_fast_shutdown)".
	ShutdownType InnoDbShutdownModeEnum `mandatory:"true" json:"shutdownType"`
}

func (m StopDbSystemDetails) String() string {
	return common.PointerString(m)
}
