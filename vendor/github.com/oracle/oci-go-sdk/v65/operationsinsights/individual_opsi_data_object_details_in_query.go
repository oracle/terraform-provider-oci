// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package operationsinsights

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IndividualOpsiDataObjectDetailsInQuery Details applicable for an individual OPSI data object used in a data object query.
type IndividualOpsiDataObjectDetailsInQuery struct {

	// Unique OPSI data object identifier.
	DataObjectIdentifier *string `mandatory:"true" json:"dataObjectIdentifier"`

	// An array of query parameters to be applied, for the OPSI data objects targetted by dataObjectDetailsTarget, before executing the query.
	// Refer to supportedQueryParams of OpsiDataObject for the supported query parameters.
	QueryParams []OpsiDataObjectQueryParam `mandatory:"false" json:"queryParams"`
}

// GetQueryParams returns QueryParams
func (m IndividualOpsiDataObjectDetailsInQuery) GetQueryParams() []OpsiDataObjectQueryParam {
	return m.QueryParams
}

func (m IndividualOpsiDataObjectDetailsInQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IndividualOpsiDataObjectDetailsInQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m IndividualOpsiDataObjectDetailsInQuery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIndividualOpsiDataObjectDetailsInQuery IndividualOpsiDataObjectDetailsInQuery
	s := struct {
		DiscriminatorParam string `json:"dataObjectDetailsTarget"`
		MarshalTypeIndividualOpsiDataObjectDetailsInQuery
	}{
		"INDIVIDUAL_OPSIDATAOBJECT",
		(MarshalTypeIndividualOpsiDataObjectDetailsInQuery)(m),
	}

	return json.Marshal(&s)
}
