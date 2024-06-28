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

// OpsiDataObjectTypeOpsiDataObjectDetailsInQuery Details applicable for all OPSI data objects of a specific OpsiDataObjectType used in a data object query.
type OpsiDataObjectTypeOpsiDataObjectDetailsInQuery struct {

	// An array of query parameters to be applied, for the OPSI data objects targetted by dataObjectDetailsTarget, before executing the query.
	// Refer to supportedQueryParams of OpsiDataObject for the supported query parameters.
	QueryParams []OpsiDataObjectQueryParam `mandatory:"false" json:"queryParams"`

	// Type of OPSI data object.
	DataObjectType OpsiDataObjectTypeEnum `mandatory:"true" json:"dataObjectType"`
}

// GetQueryParams returns QueryParams
func (m OpsiDataObjectTypeOpsiDataObjectDetailsInQuery) GetQueryParams() []OpsiDataObjectQueryParam {
	return m.QueryParams
}

func (m OpsiDataObjectTypeOpsiDataObjectDetailsInQuery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpsiDataObjectTypeOpsiDataObjectDetailsInQuery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOpsiDataObjectTypeEnum(string(m.DataObjectType)); !ok && m.DataObjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataObjectType: %s. Supported values are: %s.", m.DataObjectType, strings.Join(GetOpsiDataObjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OpsiDataObjectTypeOpsiDataObjectDetailsInQuery) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOpsiDataObjectTypeOpsiDataObjectDetailsInQuery OpsiDataObjectTypeOpsiDataObjectDetailsInQuery
	s := struct {
		DiscriminatorParam string `json:"dataObjectDetailsTarget"`
		MarshalTypeOpsiDataObjectTypeOpsiDataObjectDetailsInQuery
	}{
		"OPSIDATAOBJECTTYPE_OPSIDATAOBJECTS",
		(MarshalTypeOpsiDataObjectTypeOpsiDataObjectDetailsInQuery)(m),
	}

	return json.Marshal(&s)
}
