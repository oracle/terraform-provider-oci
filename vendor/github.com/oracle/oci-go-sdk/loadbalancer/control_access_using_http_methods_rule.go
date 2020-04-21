// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ControlAccessUsingHttpMethodsRule An object that represents the action of returning a specified response code when the requested HTTP method is not in
// the list of allowed methods for the listener. The load balancer does not forward a disallowed request to the back end
// servers. The default response code is `405 Method Not Allowed`.
// If you set the response code to `405` or leave it blank, the system adds an "allow" response header that contains a
// list of the allowed methods for the listener. If you set the response code to anything other than `405` (or blank),
// the system does not add the "allow" response header with a list of allowed methods.
// This rule applies only to HTTP listeners. No more than one `ControlAccessUsingHttpMethodsRule` object can be present in
// a given listener.
type ControlAccessUsingHttpMethodsRule struct {

	// The list of HTTP methods allowed for this listener.
	// By default, you can specify only the standard HTTP methods defined in the
	// HTTP Method Registry (http://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also
	// see a list of supported standard HTTP methods in the Load Balancing service documentation at
	// Managing Rule Sets (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingrulesets.htm).
	// Your backend application must be able to handle the methods specified in this list.
	// The list of HTTP methods is extensible. If you need to configure custom HTTP methods, contact
	// My Oracle Support (http://support.oracle.com/) to remove the restriction for your tenancy.
	// Example: ["GET", "PUT", "POST", "PROPFIND"]
	AllowedMethods []string `mandatory:"true" json:"allowedMethods"`

	// The HTTP status code to return when the requested HTTP method is not in the list of allowed methods.
	// The associated status line returned with the code is mapped from the standard HTTP specification. The
	// default value is `405 (Method Not Allowed)`.
	// Example: 403
	StatusCode *int `mandatory:"false" json:"statusCode"`
}

func (m ControlAccessUsingHttpMethodsRule) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ControlAccessUsingHttpMethodsRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeControlAccessUsingHttpMethodsRule ControlAccessUsingHttpMethodsRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeControlAccessUsingHttpMethodsRule
	}{
		"CONTROL_ACCESS_USING_HTTP_METHODS",
		(MarshalTypeControlAccessUsingHttpMethodsRule)(m),
	}

	return json.Marshal(&s)
}
