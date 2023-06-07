// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GrpcVirtualServiceTrafficRouteRule Rule for routing incoming Virtual Service traffic with GRPC protocol.
type GrpcVirtualServiceTrafficRouteRule struct {

	// The destination of the request.
	Destinations []VirtualDeploymentTrafficRuleTarget `mandatory:"true" json:"destinations"`

	// An object that represents the data to match from the request.
	Metadata []StringMatch `mandatory:"false" json:"metadata"`

	// The fully qualified domain name for the service to match from the request.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// The method name to match from the request. If you specify a methodName, you must also specify a serviceName.
	MethodName *string `mandatory:"false" json:"methodName"`

	RetryConfiguration *GrpcRetryConfiguration `mandatory:"false" json:"retryConfiguration"`
}

//GetDestinations returns Destinations
func (m GrpcVirtualServiceTrafficRouteRule) GetDestinations() []VirtualDeploymentTrafficRuleTarget {
	return m.Destinations
}

func (m GrpcVirtualServiceTrafficRouteRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrpcVirtualServiceTrafficRouteRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GrpcVirtualServiceTrafficRouteRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGrpcVirtualServiceTrafficRouteRule GrpcVirtualServiceTrafficRouteRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeGrpcVirtualServiceTrafficRouteRule
	}{
		"GRPC",
		(MarshalTypeGrpcVirtualServiceTrafficRouteRule)(m),
	}

	return json.Marshal(&s)
}
