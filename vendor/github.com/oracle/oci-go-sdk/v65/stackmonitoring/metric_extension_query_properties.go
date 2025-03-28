// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricExtensionQueryProperties Collection method and query properties details of metric extension
type MetricExtensionQueryProperties interface {
}

type metricextensionqueryproperties struct {
	JsonData         []byte
	CollectionMethod string `json:"collectionMethod"`
}

// UnmarshalJSON unmarshals json
func (m *metricextensionqueryproperties) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermetricextensionqueryproperties metricextensionqueryproperties
	s := struct {
		Model Unmarshalermetricextensionqueryproperties
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CollectionMethod = s.Model.CollectionMethod

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *metricextensionqueryproperties) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.CollectionMethod {
	case "OS_COMMAND":
		mm := OsCommandQueryProperties{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL":
		mm := SqlQueryProperties{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JMX":
		mm := JmxQueryProperties{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := HttpQueryProperties{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for MetricExtensionQueryProperties: %s.", m.CollectionMethod)
		return *m, nil
	}
}

func (m metricextensionqueryproperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m metricextensionqueryproperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
