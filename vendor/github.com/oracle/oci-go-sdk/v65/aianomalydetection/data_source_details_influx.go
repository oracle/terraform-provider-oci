// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataSourceDetailsInflux Data Source details for influx.
type DataSourceDetailsInflux struct {
	VersionSpecificDetails InfluxDetails `mandatory:"true" json:"versionSpecificDetails"`

	// Username for connection to Influx
	UserName *string `mandatory:"true" json:"userName"`

	// Password Secret Id for the influx connection
	PasswordSecretId *string `mandatory:"true" json:"passwordSecretId"`

	// Measurement name for influx
	MeasurementName *string `mandatory:"true" json:"measurementName"`

	// public IP address and port to influx DB
	Url *string `mandatory:"true" json:"url"`
}

func (m DataSourceDetailsInflux) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataSourceDetailsInflux) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataSourceDetailsInflux) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataSourceDetailsInflux DataSourceDetailsInflux
	s := struct {
		DiscriminatorParam string `json:"dataSourceType"`
		MarshalTypeDataSourceDetailsInflux
	}{
		"INFLUX",
		(MarshalTypeDataSourceDetailsInflux)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DataSourceDetailsInflux) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		VersionSpecificDetails influxdetails `json:"versionSpecificDetails"`
		UserName               *string       `json:"userName"`
		PasswordSecretId       *string       `json:"passwordSecretId"`
		MeasurementName        *string       `json:"measurementName"`
		Url                    *string       `json:"url"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.VersionSpecificDetails.UnmarshalPolymorphicJSON(model.VersionSpecificDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.VersionSpecificDetails = nn.(InfluxDetails)
	} else {
		m.VersionSpecificDetails = nil
	}

	m.UserName = model.UserName

	m.PasswordSecretId = model.PasswordSecretId

	m.MeasurementName = model.MeasurementName

	m.Url = model.Url

	return
}
