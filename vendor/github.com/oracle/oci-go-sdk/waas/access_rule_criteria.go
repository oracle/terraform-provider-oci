// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AccessRuleCriteria The representation of AccessRuleCriteria
type AccessRuleCriteria struct {

	// The criteria the access rule uses to determine if action should be taken on a request.
	// - **URL_IS:** Matches if the concatenation of request URL path and query is identical to the contents of the `value` field.
	// - **URL_IS_NOT:** Matches if the concatenation of request URL path and query is not identical to the contents of the `value` field.
	// - **URL_STARTS_WITH:** Matches if the concatenation of request URL path and query starts with the contents of the `value` field.
	// - **URL_PART_ENDS_WITH:** Matches if the concatenation of request URL path and query ends with the contents of the `value` field.
	// - **URL_PART_CONTAINS:** Matches if the concatenation of request URL path and query contains the contents of the `value` field.
	// - **URL_REGEX:** Matches if the request is described by the regular expression in the `value` field.
	// - **IP_IS:** Matches if the request originates from an IP address in the `value` field.
	// - **IP_IS_NOT:** Matches if the request does not originate from an IP address in the `value` field.
	// - **HTTP_HEADER_CONTAINS:** The HTTP_HEADER_CONTAINS criteria is defined using a compound value separated by a colon: a header field name and a header field value. `host:test.example.com` is an example of a criteria value where `host` is the header field name and `test.example.com` is the header field value. A request matches when the header field name is a case insensitive match and the header field value is a case insensitive, substring match.
	// *Example:* With a criteria value of `host:test.example.com`, where `host` is the name of the field and `test.example.com` is the value of the host field, a request with the header values, `Host: www.test.example.com` will match, where as a request with header values of `host: www.example.com` or `host: test.sub.example.com` will not match.
	// - **IP_IN_LIST:** Matches if the request originates from one of IP addresses contained in the referenced address list. The `value` in this case is OCID of the address list.
	// - **IP_NOT_IN_LIST:** Matches if the request does not originate from any IP address contained in the referenced address list. The `value` field in this case is OCID of the address list.
	// - **HTTP_METHOD_IS:** Matches if the request method corresponds to the `value` field. The list of available methods: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	// - **HTTP_METHOD_IS_NOT:** Matches if the request method does not correspond to the `value` field. The list of available methods: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	// - **COUNTRY_IS:** Matches if the request originates from a country in the `value` field. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	// - **COUNTRY_IS_NOT:** Matches if the request does not originate from a country in the `value` field. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	// - **USER_AGENT_IS:** Matches if the requesting user agent is identical to the contents of the `value` field. Example: `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
	// - **USER_AGENT_IS_NOT:** Matches if the requesting user agent is not identical to the contents of the `value` field. Example: `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
	Condition AccessRuleCriteriaConditionEnum `mandatory:"true" json:"condition"`

	// The criteria value.
	Value *string `mandatory:"true" json:"value"`
}

func (m AccessRuleCriteria) String() string {
	return common.PointerString(m)
}

// AccessRuleCriteriaConditionEnum Enum with underlying type: string
type AccessRuleCriteriaConditionEnum string

// Set of constants representing the allowable values for AccessRuleCriteriaConditionEnum
const (
	AccessRuleCriteriaConditionUrlIs              AccessRuleCriteriaConditionEnum = "URL_IS"
	AccessRuleCriteriaConditionUrlIsNot           AccessRuleCriteriaConditionEnum = "URL_IS_NOT"
	AccessRuleCriteriaConditionUrlStartsWith      AccessRuleCriteriaConditionEnum = "URL_STARTS_WITH"
	AccessRuleCriteriaConditionUrlPartEndsWith    AccessRuleCriteriaConditionEnum = "URL_PART_ENDS_WITH"
	AccessRuleCriteriaConditionUrlPartContains    AccessRuleCriteriaConditionEnum = "URL_PART_CONTAINS"
	AccessRuleCriteriaConditionUrlRegex           AccessRuleCriteriaConditionEnum = "URL_REGEX"
	AccessRuleCriteriaConditionIpIs               AccessRuleCriteriaConditionEnum = "IP_IS"
	AccessRuleCriteriaConditionIpIsNot            AccessRuleCriteriaConditionEnum = "IP_IS_NOT"
	AccessRuleCriteriaConditionHttpHeaderContains AccessRuleCriteriaConditionEnum = "HTTP_HEADER_CONTAINS"
	AccessRuleCriteriaConditionIpInList           AccessRuleCriteriaConditionEnum = "IP_IN_LIST"
	AccessRuleCriteriaConditionIpNotInList        AccessRuleCriteriaConditionEnum = "IP_NOT_IN_LIST"
	AccessRuleCriteriaConditionHttpMethodIs       AccessRuleCriteriaConditionEnum = "HTTP_METHOD_IS"
	AccessRuleCriteriaConditionHttpMethodIsNot    AccessRuleCriteriaConditionEnum = "HTTP_METHOD_IS_NOT"
	AccessRuleCriteriaConditionCountryIs          AccessRuleCriteriaConditionEnum = "COUNTRY_IS"
	AccessRuleCriteriaConditionCountryIsNot       AccessRuleCriteriaConditionEnum = "COUNTRY_IS_NOT"
	AccessRuleCriteriaConditionUserAgentIs        AccessRuleCriteriaConditionEnum = "USER_AGENT_IS"
	AccessRuleCriteriaConditionUserAgentIsNot     AccessRuleCriteriaConditionEnum = "USER_AGENT_IS_NOT"
)

var mappingAccessRuleCriteriaCondition = map[string]AccessRuleCriteriaConditionEnum{
	"URL_IS":               AccessRuleCriteriaConditionUrlIs,
	"URL_IS_NOT":           AccessRuleCriteriaConditionUrlIsNot,
	"URL_STARTS_WITH":      AccessRuleCriteriaConditionUrlStartsWith,
	"URL_PART_ENDS_WITH":   AccessRuleCriteriaConditionUrlPartEndsWith,
	"URL_PART_CONTAINS":    AccessRuleCriteriaConditionUrlPartContains,
	"URL_REGEX":            AccessRuleCriteriaConditionUrlRegex,
	"IP_IS":                AccessRuleCriteriaConditionIpIs,
	"IP_IS_NOT":            AccessRuleCriteriaConditionIpIsNot,
	"HTTP_HEADER_CONTAINS": AccessRuleCriteriaConditionHttpHeaderContains,
	"IP_IN_LIST":           AccessRuleCriteriaConditionIpInList,
	"IP_NOT_IN_LIST":       AccessRuleCriteriaConditionIpNotInList,
	"HTTP_METHOD_IS":       AccessRuleCriteriaConditionHttpMethodIs,
	"HTTP_METHOD_IS_NOT":   AccessRuleCriteriaConditionHttpMethodIsNot,
	"COUNTRY_IS":           AccessRuleCriteriaConditionCountryIs,
	"COUNTRY_IS_NOT":       AccessRuleCriteriaConditionCountryIsNot,
	"USER_AGENT_IS":        AccessRuleCriteriaConditionUserAgentIs,
	"USER_AGENT_IS_NOT":    AccessRuleCriteriaConditionUserAgentIsNot,
}

// GetAccessRuleCriteriaConditionEnumValues Enumerates the set of values for AccessRuleCriteriaConditionEnum
func GetAccessRuleCriteriaConditionEnumValues() []AccessRuleCriteriaConditionEnum {
	values := make([]AccessRuleCriteriaConditionEnum, 0)
	for _, v := range mappingAccessRuleCriteriaCondition {
		values = append(values, v)
	}
	return values
}
