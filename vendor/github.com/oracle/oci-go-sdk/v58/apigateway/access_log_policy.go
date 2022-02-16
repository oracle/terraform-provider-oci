// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// AccessLogPolicy Configures the logging policies for the access logs of an API Deployment.
type AccessLogPolicy struct {

	// Enables pushing of access logs to the legacy OCI Object Storage log archival bucket.
	// Oracle recommends using the OCI Logging service to enable, retrieve, and query access logs
	// for an API Deployment. If there is an active log object for the API Deployment and its
	// category is set to 'access' in OCI Logging service, the logs will not be uploaded to the
	// legacy OCI Object Storage log archival bucket.
	// Please note that the functionality to push to the legacy OCI Object Storage log
	// archival bucket has been deprecated and will be removed in the future.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m AccessLogPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AccessLogPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
