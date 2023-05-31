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

// GrpcFaultConfiguration Fault can be used to specify one or more faults to inject while receiving GRPC requests at the listener.
// Faults include aborting the GRPC request from downstream service, and/or delaying proxying of requests.
type GrpcFaultConfiguration struct {
	DelayConfiguration *DelayFaultConfiguration `mandatory:"false" json:"delayConfiguration"`

	AbortConfiguration *GrpcAbortFaultConfiguration `mandatory:"false" json:"abortConfiguration"`
}

func (m GrpcFaultConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrpcFaultConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GrpcFaultConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGrpcFaultConfiguration GrpcFaultConfiguration
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeGrpcFaultConfiguration
	}{
		"GRPC",
		(MarshalTypeGrpcFaultConfiguration)(m),
	}

	return json.Marshal(&s)
}
