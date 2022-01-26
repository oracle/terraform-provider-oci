// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Apdex An Apdex configuration rule.
// The Apdex Score is computed based on how the response time of a span compares to two predefined threshold values.
// The first threshold defines the maximum response time that can still be considered satisfactory for the end user.
// The second one defines the maximum response time that can be considered tolerable. All times larger than that will
// be considered frustrating for the end user.
// An Apdex configuration rule works by selecting a subset of spans based on a filter expression and applying the
// two threshold comparisons to compute a score for each of the selected spans.
// The rule has a property "isApplyToErrorSpans" that controls whether or not to compute the Apdex for spans that have
// have been marked as errors. If this property is set to true, then error spans will have their Apdex score computed
// the same as non-error ones. If set to false, then computation for error spans will be skipped, and the score will
// be set to "frustrating" regardless of the configured thresholds. The default is "false".
// The property "isEnabled" controls whether an Apdex score is computed for the spans. Can be used to disable Apdex
// scores for certain spans. The default is "true".
// The property "priority" is used to define the importance of the rule when it's part of a rule set.
// Lower values indicate a higher priority. Rules with higher priorities will be evaluated first in the rule set. The
// priority of the rules must be unique within a rule set.
type Apdex struct {

	// The string that defines the Span Filter expression.
	FilterText *string `mandatory:"true" json:"filterText"`

	// The priority controls the order in which multiple rules in a rule set are applied. Lower values indicate higher
	// priorities. Rules with higher priority are applied first, and once a match is found, the rest of the rules are
	// ignored. Rules within the same rule set cannot have the same priority.
	Priority *int `mandatory:"true" json:"priority"`

	// Specifies if the Apdex rule will be computed for spans matching the rule. Can be used to make sure certain
	// spans don't get an Apdex score. The default is "true".
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The maximum response time in milliseconds that will be considered satisfactory for the end user.
	SatisfiedResponseTime *int `mandatory:"false" json:"satisfiedResponseTime"`

	// The maximum response time in milliseconds that will be considered tolerable for the end user. Response
	// times beyond this threshold will be considered frustrating.
	// This value cannot be lower than "satisfiedResponseTime".
	ToleratingResponseTime *int `mandatory:"false" json:"toleratingResponseTime"`

	// If true, the rule will compute the actual Apdex score for spans that have been marked as errors. If false,
	// the rule will always set the Apdex for error spans to frustrating, regardless of the configured thresholds.
	// Default is false.
	IsApplyToErrorSpans *bool `mandatory:"false" json:"isApplyToErrorSpans"`

	// A user-friendly name that provides a short description this rule.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m Apdex) String() string {
	return common.PointerString(m)
}
