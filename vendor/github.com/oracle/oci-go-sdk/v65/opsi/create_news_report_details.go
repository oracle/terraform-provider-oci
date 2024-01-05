// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateNewsReportDetails The information about the news report to be created.
type CreateNewsReportDetails struct {

	// The news report name.
	Name *string `mandatory:"true" json:"name"`

	// News report frequency.
	NewsFrequency NewsFrequencyEnum `mandatory:"true" json:"newsFrequency"`

	// The description of the news report.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ONS topic.
	OnsTopicId *string `mandatory:"true" json:"onsTopicId"`

	// Compartment Identifier where the news report will be created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ContentTypes *NewsContentTypes `mandatory:"true" json:"contentTypes"`

	// Language of the news report.
	Locale NewsLocaleEnum `mandatory:"true" json:"locale"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Defines if the news report will be enabled or disabled.
	Status ResourceStatusEnum `mandatory:"false" json:"status,omitempty"`
}

func (m CreateNewsReportDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateNewsReportDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNewsFrequencyEnum(string(m.NewsFrequency)); !ok && m.NewsFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NewsFrequency: %s. Supported values are: %s.", m.NewsFrequency, strings.Join(GetNewsFrequencyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNewsLocaleEnum(string(m.Locale)); !ok && m.Locale != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Locale: %s. Supported values are: %s.", m.Locale, strings.Join(GetNewsLocaleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
