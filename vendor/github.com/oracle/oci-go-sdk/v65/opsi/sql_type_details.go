// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlTypeDetails SQL details
type SqlTypeDetails struct {

	// SQL identifier
	SqlId *string `mandatory:"true" json:"sqlId"`

	// First 3800 characters of the SQL text
	SqlText *string `mandatory:"true" json:"sqlText"`

	// SQL identifier
	IsSqlTextTruncated *bool `mandatory:"true" json:"isSqlTextTruncated"`

	// SQL command name (such as SELECT, INSERT)
	SqlCommand *string `mandatory:"true" json:"sqlCommand"`
}

func (m SqlTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlTypeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlTypeDetails SqlTypeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSqlTypeDetails
	}{
		"SQL",
		(MarshalTypeSqlTypeDetails)(m),
	}

	return json.Marshal(&s)
}
