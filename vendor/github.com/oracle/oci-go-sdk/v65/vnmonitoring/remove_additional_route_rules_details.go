// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveAdditionalRouteRulesDetails The configuration details for the remove rules operation. Only one of the three properties should be
// supplied to call the removeAdditionalRouteRules API.
type RemoveAdditionalRouteRulesDetails struct {

	// The list of route rule identifiers used for removing route rules from route table.
	AdditionalRouteRuleIds []string `mandatory:"false" json:"additionalRouteRuleIds"`

	// The list of destinations used for removing route rules from route table. This is only supplied
	// when additionalRouteRuleIds are not provided.
	Destinations []string `mandatory:"false" json:"destinations"`

	// The pairs of <destination, target> used for removing route rules from route table. This is only
	// supplied when additionalRouteRules and destinations are not provided.
	RouteDestinationRouteTargets []RouteDestinationRouteTargetDetails `mandatory:"false" json:"routeDestinationRouteTargets"`
}

func (m RemoveAdditionalRouteRulesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveAdditionalRouteRulesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
