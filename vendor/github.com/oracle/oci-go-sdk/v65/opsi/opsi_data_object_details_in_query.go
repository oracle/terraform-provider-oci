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

// OpsiDataObjectDetailsInQuery Details for OPSI data object used in a data object query.
type OpsiDataObjectDetailsInQuery interface {

	// An array of query parameters to be applied, for the OPSI data objects targetted by dataObjectDetailsTarget, before executing the query.
	// Refer to supportedQueryParams of OpsiDataObject for the supported query parameters.
	GetQueryParams() []OpsiDataObjectQueryParam
}

type opsidataobjectdetailsinquery struct {
	JsonData                []byte
	QueryParams             []OpsiDataObjectQueryParam `mandatory:"false" json:"queryParams"`
	DataObjectDetailsTarget string                     `json:"dataObjectDetailsTarget"`
}

// UnmarshalJSON unmarshals json
func (m *opsidataobjectdetailsinquery) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleropsidataobjectdetailsinquery opsidataobjectdetailsinquery
	s := struct {
		Model Unmarshaleropsidataobjectdetailsinquery
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.QueryParams = s.Model.QueryParams
	m.DataObjectDetailsTarget = s.Model.DataObjectDetailsTarget

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *opsidataobjectdetailsinquery) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataObjectDetailsTarget {
	case "INDIVIDUAL_OPSIDATAOBJECT":
		mm := IndividualOpsiDataObjectDetailsInQuery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OPSIDATAOBJECTTYPE_OPSIDATAOBJECTS":
		mm := OpsiDataObjectTypeOpsiDataObjectDetailsInQuery{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for OpsiDataObjectDetailsInQuery: %s.", m.DataObjectDetailsTarget)
		return *m, nil
	}
}

// GetQueryParams returns QueryParams
func (m opsidataobjectdetailsinquery) GetQueryParams() []OpsiDataObjectQueryParam {
	return m.QueryParams
}

func (m opsidataobjectdetailsinquery) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m opsidataobjectdetailsinquery) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
