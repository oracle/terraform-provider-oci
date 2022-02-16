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

// ResumeJobDetails Parameters to specify to resume a Migration Job.
type ResumeJobDetails struct {

	// Name of a migration phase. The Job will wait after executing this
	// phase until Resume Job endpoint is called again.
	WaitAfter OdmsJobPhasesEnum `mandatory:"false" json:"waitAfter,omitempty"`
}

func (m ResumeJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResumeJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOdmsJobPhasesEnum(string(m.WaitAfter)); !ok && m.WaitAfter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WaitAfter: %s. Supported values are: %s.", m.WaitAfter, strings.Join(GetOdmsJobPhasesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
