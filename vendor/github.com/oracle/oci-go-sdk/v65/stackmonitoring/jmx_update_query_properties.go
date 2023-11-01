// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// JmxUpdateQueryProperties Query Properties applicable to JMX type of collection method
type JmxUpdateQueryProperties struct {

	// JMX Managed Bean Query or Metric Service Table name
	ManagedBeanQuery *string `mandatory:"false" json:"managedBeanQuery"`

	// List of JMX attributes or Metric Service Table columns separated by semi-colon
	JmxAttributes *string `mandatory:"false" json:"jmxAttributes"`

	// Semi-colon separated list of key properties from Managed Bean ObjectName to be used as key metrics
	IdentityMetric *string `mandatory:"false" json:"identityMetric"`

	// Prefix for an auto generated metric, in case multiple rows with non unique key values are returned
	AutoRowPrefix *string `mandatory:"false" json:"autoRowPrefix"`

	// Indicates if Metric Service is enabled on server domain
	IsMetricServiceEnabled *bool `mandatory:"false" json:"isMetricServiceEnabled"`
}

func (m JmxUpdateQueryProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m JmxUpdateQueryProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m JmxUpdateQueryProperties) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJmxUpdateQueryProperties JmxUpdateQueryProperties
	s := struct {
		DiscriminatorParam string `json:"collectionMethod"`
		MarshalTypeJmxUpdateQueryProperties
	}{
		"JMX",
		(MarshalTypeJmxUpdateQueryProperties)(m),
	}

	return json.Marshal(&s)
}
