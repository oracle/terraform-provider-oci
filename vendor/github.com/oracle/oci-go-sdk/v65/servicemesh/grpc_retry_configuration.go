// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GrpcRetryConfiguration GRPC retry configuration for virtual service route rule.
type GrpcRetryConfiguration struct {

	// The number of retries to be allowed for a given request.
	MaxRetries *int `mandatory:"true" json:"maxRetries"`

	// The timeout per attempt for a given request, including the initial call and any retries.
	RetryTimeoutInMs *int64 `mandatory:"false" json:"retryTimeoutInMs"`

	// Specifies the conditions under which retry takes place. One or more policies can be specified.
	RetryCritiera []GrpcRetryCritieriaEnum `mandatory:"false" json:"retryCritiera"`

	// Base time interval between retries.
	BackoffBaseDelayInMs *int64 `mandatory:"false" json:"backoffBaseDelayInMs"`

	// Maximum time between retries.
	BackoffMaxDelayInMs *int64 `mandatory:"false" json:"backoffMaxDelayInMs"`
}

func (m GrpcRetryConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GrpcRetryConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
