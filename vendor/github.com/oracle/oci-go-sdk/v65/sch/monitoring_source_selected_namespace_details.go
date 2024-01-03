// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoringSourceSelectedNamespaceDetails The namespaces for the compartment-specific list.
type MonitoringSourceSelectedNamespaceDetails struct {

	// The namespaces for the compartment-specific list.
	Namespaces []MonitoringSourceSelectedNamespace `mandatory:"true" json:"namespaces"`
}

func (m MonitoringSourceSelectedNamespaceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoringSourceSelectedNamespaceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m MonitoringSourceSelectedNamespaceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMonitoringSourceSelectedNamespaceDetails MonitoringSourceSelectedNamespaceDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeMonitoringSourceSelectedNamespaceDetails
	}{
		"selected",
		(MarshalTypeMonitoringSourceSelectedNamespaceDetails)(m),
	}

	return json.Marshal(&s)
}
