// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CircuitBreakingRetryBudget Specifies a limit on concurrent retries in relation to the number of active requests.
type CircuitBreakingRetryBudget struct {

	// Specifies the limit on concurrent retries as a percentage of the sum of active requests and
	// active pending requests.
	// For example,
	// if there are 100 active requests and the budget_percent is set to 25, there may be 25 active retries.
	BudgetPercent *int `mandatory:"false" json:"budgetPercent"`

	// Specifies the minimum retry concurrency allowed for the retry budget.
	// The limit on the number of active retries may never go below this number.
	MinRetryConcurrency *int `mandatory:"false" json:"minRetryConcurrency"`
}

func (m CircuitBreakingRetryBudget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CircuitBreakingRetryBudget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
