// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// RestDetails REST configuration details.
type RestDetails struct {

	// Select how REST is configured across the DB System instances.
	Configuration RestConfigurationTypeEnum `mandatory:"true" json:"configuration"`

	// The port for REST to listen on. Supported port numbers are 443 and from 1024 to 65535.
	Port *int `mandatory:"false" json:"port"`
}

func (m RestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRestConfigurationTypeEnum(string(m.Configuration)); !ok && m.Configuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Configuration: %s. Supported values are: %s.", m.Configuration, strings.Join(GetRestConfigurationTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
