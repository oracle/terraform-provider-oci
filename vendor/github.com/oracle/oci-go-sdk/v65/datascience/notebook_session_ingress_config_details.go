// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NotebookSessionIngressConfigDetails Notebook Session Ingress configuration details.
type NotebookSessionIngressConfigDetails struct {

	// This is a collection of key-value pairs where the key represents a URI path prefix, and the value specifies the corresponding port.
	// The port must be exposed by the Notebook container.
	// For example, a key-value pair like,
	//   Key: /api/custom-path , Value: 9999
	// If a customer attempts to access /api/custom-path/{xyz}, the traffic will be routed to port 9999.
	// Constraints:
	//   Key: Must be a valid URI path and cannot exceed 128 characters in length.
	//   Value: Must be within the valid port range.
	// Maximum accepted key-value pairs is 5.
	PortMappings []PortMapping `mandatory:"false" json:"portMappings"`
}

func (m NotebookSessionIngressConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NotebookSessionIngressConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
