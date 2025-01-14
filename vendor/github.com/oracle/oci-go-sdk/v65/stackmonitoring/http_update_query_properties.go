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

// HttpUpdateQueryProperties Query properties applicable to HTTP type of collection method
type HttpUpdateQueryProperties struct {

	// Http(s) end point URL
	Url *string `mandatory:"false" json:"url"`

	ScriptDetails *UpdateHttpScriptFileDetails `mandatory:"false" json:"scriptDetails"`

	// Type of content response given by the http(s) URL
	ResponseContentType HttpResponseContentTypesEnum `mandatory:"false" json:"responseContentType,omitempty"`

	// Supported protocol of resources to be associated with this metric extension. This is optional and defaults to HTTPS, which uses secure connection to the URL
	ProtocolType HttpProtocolTypesEnum `mandatory:"false" json:"protocolType,omitempty"`
}

func (m HttpUpdateQueryProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HttpUpdateQueryProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHttpResponseContentTypesEnum(string(m.ResponseContentType)); !ok && m.ResponseContentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponseContentType: %s. Supported values are: %s.", m.ResponseContentType, strings.Join(GetHttpResponseContentTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHttpProtocolTypesEnum(string(m.ProtocolType)); !ok && m.ProtocolType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtocolType: %s. Supported values are: %s.", m.ProtocolType, strings.Join(GetHttpProtocolTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HttpUpdateQueryProperties) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHttpUpdateQueryProperties HttpUpdateQueryProperties
	s := struct {
		DiscriminatorParam string `json:"collectionMethod"`
		MarshalTypeHttpUpdateQueryProperties
	}{
		"HTTP",
		(MarshalTypeHttpUpdateQueryProperties)(m),
	}

	return json.Marshal(&s)
}
