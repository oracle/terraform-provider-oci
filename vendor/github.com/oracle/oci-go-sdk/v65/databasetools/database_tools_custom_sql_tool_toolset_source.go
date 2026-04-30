// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsCustomSqlToolToolsetSource The SQL source. Can contain multiple statements with a mix of queries, DML, DCL, DLL and scripts.
type DatabaseToolsCustomSqlToolToolsetSource struct {

	// The sql toolset sources type. INLINE is the only possible value.
	Type DatabaseToolsCustomSqlToolToolsetSourceTypeEnum `mandatory:"true" json:"type"`

	// The SQL source. Can contain multiple statements with a mix of queries, DML, DCL, DLL and scripts.
	Value *string `mandatory:"true" json:"value"`
}

func (m DatabaseToolsCustomSqlToolToolsetSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseToolsCustomSqlToolToolsetSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolsCustomSqlToolToolsetSourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDatabaseToolsCustomSqlToolToolsetSourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
