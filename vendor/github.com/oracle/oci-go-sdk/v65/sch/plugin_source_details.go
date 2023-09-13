// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
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

// PluginSourceDetails The pluggable source.
// For configuration instructions, see
// Creating a Service Connector (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/create-service-connector.htm).
type PluginSourceDetails struct {

	// The specific plugin type to be used.
	PluginName *string `mandatory:"true" json:"pluginName"`

	// The config map for a specific plugin type.
	ConfigMap *interface{} `mandatory:"true" json:"configMap"`
}

func (m PluginSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluginSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PluginSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePluginSourceDetails PluginSourceDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypePluginSourceDetails
	}{
		"plugin",
		(MarshalTypePluginSourceDetails)(m),
	}

	return json.Marshal(&s)
}
