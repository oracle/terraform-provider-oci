// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciConsoleSignOnPolicyConsentSearchRequest Clients MAY execute queries without passing parameters on the URL by using the HTTP POST verb combined with the **.search** path extension. The inclusion of **.search** on the end of a valid SCIM endpoint SHALL be used to indicate the HTTP POST verb is intended to be a query operation. To create a new query result set, a SCIM client sends an HTTP POST request to the desired SCIM resource endpoint (ending in **.search**). The body of the POST request MAY include any of the parameters.
type OciConsoleSignOnPolicyConsentSearchRequest struct {

	// The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. Query requests MUST be identified using the following URI: "urn:ietf:params:scim:api:messages:2.0:SearchRequest" REQUIRED.
	Schemas []string `mandatory:"true" json:"schemas"`

	// A multi-valued list of strings indicating the names of resource attributes to return in the response overriding the set of attributes that would be returned by default. Attribute names MUST be in standard attribute notation (Section 3.10 (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.10)) form. See (additional retrieval query parameters (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.9)). OPTIONAL.
	Attributes []string `mandatory:"false" json:"attributes"`

	// A multi-valued list of strings indicating the return type of attribute definition. The specified set of attributes can be fetched by the return type of the attribute. One or more values can be given together to fetch more than one group of attributes. If "attributes" query parameter is also available, union of the two is fetched. Valid values : all, always, never, request, default. Values are case-insensitive. OPTIONAL.
	AttributeSets []AttributeSetsEnum `mandatory:"false" json:"attributeSets,omitempty"`

	// The filter string that is used to request a subset of resources. The filter string MUST be a valid filter expression. See Section 3.4.2.2 (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.4.2.2). OPTIONAL.
	Filter *string `mandatory:"false" json:"filter"`

	// A string that indicates the attribute whose value SHALL be used to order the returned responses. The sortBy attribute MUST be in standard attribute notation (Section 3.10 (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.10)) form. See Sorting section (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.4.2.3). OPTIONAL.
	SortBy *string `mandatory:"false" json:"sortBy"`

	// A string that indicates the order in which the sortBy parameter is applied. Allowed values are "ascending" and "descending". See (Sorting Section (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.4.2.3)). OPTIONAL.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// An integer that indicates the 1-based index of the first query result. See Pagination Section (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.4.2.4). OPTIONAL.
	StartIndex *int `mandatory:"false" json:"startIndex"`

	// An integer that indicates the desired maximum number of query results per page. 1000 is the largest value that you can use. See the Pagination section of the System for Cross-Domain Identity Management Protocol specification for more information. (Section 3.4.2.4 (https://tools.ietf.org/html/draft-ietf-scim-api-19#section-3.4.2.4)). OPTIONAL.
	Count *int `mandatory:"false" json:"count"`
}

func (m OciConsoleSignOnPolicyConsentSearchRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciConsoleSignOnPolicyConsentSearchRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.AttributeSets {
		if _, ok := GetMappingAttributeSetsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AttributeSets: %s. Supported values are: %s.", val, strings.Join(GetAttributeSetsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
