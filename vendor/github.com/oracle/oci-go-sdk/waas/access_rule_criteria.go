// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AccessRuleCriteria When defined, the parent challenge would be applied only for the requests that matched all the listed conditions.
type AccessRuleCriteria struct {

	// The criteria the access rule and JavaScript Challenge uses to determine if action should be taken on a request.
	// - **URL_IS:** Matches if the concatenation of request URL path and query is identical to the contents of the `value` field. URL must start with a `/`.
	// - **URL_IS_NOT:** Matches if the concatenation of request URL path and query is not identical to the contents of the `value` field. URL must start with a `/`.
	// - **URL_STARTS_WITH:** Matches if the concatenation of request URL path and query starts with the contents of the `value` field. URL must start with a `/`.
	// - **URL_PART_ENDS_WITH:** Matches if the concatenation of request URL path and query ends with the contents of the `value` field.
	// - **URL_PART_CONTAINS:** Matches if the concatenation of request URL path and query contains the contents of the `value` field.
	// - **URL_REGEX:** Matches if the concatenation of request URL path and query is described by the regular expression in the value field. The value must be a valid regular expression recognized by the PCRE library in Nginx (https://www.pcre.org).
	// - **URL_DOES_NOT_MATCH_REGEX:** Matches if the concatenation of request URL path and query is not described by the regular expression in the `value` field. The value must be a valid regular expression recognized by the PCRE library in Nginx (https://www.pcre.org).
	// - **URL_DOES_NOT_START_WITH:** Matches if the concatenation of request URL path and query does not start with the contents of the `value` field.
	// - **URL_PART_DOES_NOT_CONTAIN:** Matches if the concatenation of request URL path and query does not contain the contents of the `value` field.
	// - **URL_PART_DOES_NOT_END_WITH:** Matches if the concatenation of request URL path and query does not end with the contents of the `value` field.
	// - **IP_IS:** Matches if the request originates from one of the IP addresses contained in the defined address list. The `value` in this case is string with one or multiple IPs or CIDR notations separated by new line symbol \n
	// *Example:* "1.1.1.1\n1.1.1.2\n1.2.2.1/30"
	// - **IP_IS_NOT:** Matches if the request does not originate from any of the IP addresses contained in the defined address list. The `value` in this case is string with one or multiple IPs or CIDR notations separated by new line symbol \n
	// *Example:* "1.1.1.1\n1.1.1.2\n1.2.2.1/30"
	// - **IP_IN_LIST:** Matches if the request originates from one of the IP addresses contained in the referenced address list. The `value` in this case is OCID of the address list.
	// - **IP_NOT_IN_LIST:** Matches if the request does not originate from any IP address contained in the referenced address list. The `value` field in this case is OCID of the address list.
	// - **HTTP_HEADER_CONTAINS:** The HTTP_HEADER_CONTAINS criteria is defined using a compound value separated by a colon: a header field name and a header field value. `host:test.example.com` is an example of a criteria value where `host` is the header field name and `test.example.com` is the header field value. A request matches when the header field name is a case insensitive match and the header field value is a case insensitive, substring match.
	// *Example:* With a criteria value of `host:test.example.com`, where `host` is the name of the field and `test.example.com` is the value of the host field, a request with the header values, `Host: www.test.example.com` will match, where as a request with header values of `host: www.example.com` or `host: test.sub.example.com` will not match.
	// - **HTTP_METHOD_IS:** Matches if the request method is identical to one of the values listed in field. The `value` in this case is string with one or multiple HTTP methods separated by new line symbol \n The list of available methods: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	//  *Example:* "GET\nPOST"
	// - **HTTP_METHOD_IS_NOT:** Matches if the request is not identical to any of the contents of the `value` field. The `value` in this case is string with one or multiple HTTP methods separated by new line symbol \n The list of available methods: `GET`, `HEAD`, `POST`, `PUT`, `DELETE`, `CONNECT`, `OPTIONS`, `TRACE`, `PATCH`
	//  *Example:* "GET\nPOST"
	// - **COUNTRY_IS:** Matches if the request originates from one of countries in the `value` field. The `value` in this case is string with one or multiple countries separated by new line symbol \n Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	// *Example:* "AL\nDZ\nAM"
	// - **COUNTRY_IS_NOT:** Matches if the request does not originate from any of countries in the `value` field. The `value` in this case is string with one or multiple countries separated by new line symbol \n Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	// *Example:* "AL\nDZ\nAM"
	// - **USER_AGENT_IS:** Matches if the requesting user agent is identical to the contents of the `value` field.
	// *Example:* `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
	// - **USER_AGENT_IS_NOT:** Matches if the requesting user agent is not identical to the contents of the `value` field.
	// *Example:* `Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:35.0) Gecko/20100101 Firefox/35.0`
	Condition AccessRuleCriteriaConditionEnum `mandatory:"true" json:"condition"`

	// The criteria value.
	Value *string `mandatory:"true" json:"value"`

	// When enabled, the condition will be matched with case-sensitive rules.
	IsCaseSensitive *bool `mandatory:"false" json:"isCaseSensitive"`
}

func (m AccessRuleCriteria) String() string {
	return common.PointerString(m)
}

// AccessRuleCriteriaConditionEnum Enum with underlying type: string
type AccessRuleCriteriaConditionEnum string

// Set of constants representing the allowable values for AccessRuleCriteriaConditionEnum
const (
	AccessRuleCriteriaConditionUrlIs                 AccessRuleCriteriaConditionEnum = "URL_IS"
	AccessRuleCriteriaConditionUrlIsNot              AccessRuleCriteriaConditionEnum = "URL_IS_NOT"
	AccessRuleCriteriaConditionUrlStartsWith         AccessRuleCriteriaConditionEnum = "URL_STARTS_WITH"
	AccessRuleCriteriaConditionUrlPartEndsWith       AccessRuleCriteriaConditionEnum = "URL_PART_ENDS_WITH"
	AccessRuleCriteriaConditionUrlPartContains       AccessRuleCriteriaConditionEnum = "URL_PART_CONTAINS"
	AccessRuleCriteriaConditionUrlRegex              AccessRuleCriteriaConditionEnum = "URL_REGEX"
	AccessRuleCriteriaConditionUrlDoesNotMatchRegex  AccessRuleCriteriaConditionEnum = "URL_DOES_NOT_MATCH_REGEX"
	AccessRuleCriteriaConditionUrlDoesNotStartWith   AccessRuleCriteriaConditionEnum = "URL_DOES_NOT_START_WITH"
	AccessRuleCriteriaConditionUrlPartDoesNotContain AccessRuleCriteriaConditionEnum = "URL_PART_DOES_NOT_CONTAIN"
	AccessRuleCriteriaConditionUrlPartDoesNotEndWith AccessRuleCriteriaConditionEnum = "URL_PART_DOES_NOT_END_WITH"
	AccessRuleCriteriaConditionIpIs                  AccessRuleCriteriaConditionEnum = "IP_IS"
	AccessRuleCriteriaConditionIpIsNot               AccessRuleCriteriaConditionEnum = "IP_IS_NOT"
	AccessRuleCriteriaConditionIpInList              AccessRuleCriteriaConditionEnum = "IP_IN_LIST"
	AccessRuleCriteriaConditionIpNotInList           AccessRuleCriteriaConditionEnum = "IP_NOT_IN_LIST"
	AccessRuleCriteriaConditionHttpHeaderContains    AccessRuleCriteriaConditionEnum = "HTTP_HEADER_CONTAINS"
	AccessRuleCriteriaConditionHttpMethodIs          AccessRuleCriteriaConditionEnum = "HTTP_METHOD_IS"
	AccessRuleCriteriaConditionHttpMethodIsNot       AccessRuleCriteriaConditionEnum = "HTTP_METHOD_IS_NOT"
	AccessRuleCriteriaConditionCountryIs             AccessRuleCriteriaConditionEnum = "COUNTRY_IS"
	AccessRuleCriteriaConditionCountryIsNot          AccessRuleCriteriaConditionEnum = "COUNTRY_IS_NOT"
	AccessRuleCriteriaConditionUserAgentIs           AccessRuleCriteriaConditionEnum = "USER_AGENT_IS"
	AccessRuleCriteriaConditionUserAgentIsNot        AccessRuleCriteriaConditionEnum = "USER_AGENT_IS_NOT"
)

var mappingAccessRuleCriteriaCondition = map[string]AccessRuleCriteriaConditionEnum{
	"URL_IS":                     AccessRuleCriteriaConditionUrlIs,
	"URL_IS_NOT":                 AccessRuleCriteriaConditionUrlIsNot,
	"URL_STARTS_WITH":            AccessRuleCriteriaConditionUrlStartsWith,
	"URL_PART_ENDS_WITH":         AccessRuleCriteriaConditionUrlPartEndsWith,
	"URL_PART_CONTAINS":          AccessRuleCriteriaConditionUrlPartContains,
	"URL_REGEX":                  AccessRuleCriteriaConditionUrlRegex,
	"URL_DOES_NOT_MATCH_REGEX":   AccessRuleCriteriaConditionUrlDoesNotMatchRegex,
	"URL_DOES_NOT_START_WITH":    AccessRuleCriteriaConditionUrlDoesNotStartWith,
	"URL_PART_DOES_NOT_CONTAIN":  AccessRuleCriteriaConditionUrlPartDoesNotContain,
	"URL_PART_DOES_NOT_END_WITH": AccessRuleCriteriaConditionUrlPartDoesNotEndWith,
	"IP_IS":                AccessRuleCriteriaConditionIpIs,
	"IP_IS_NOT":            AccessRuleCriteriaConditionIpIsNot,
	"IP_IN_LIST":           AccessRuleCriteriaConditionIpInList,
	"IP_NOT_IN_LIST":       AccessRuleCriteriaConditionIpNotInList,
	"HTTP_HEADER_CONTAINS": AccessRuleCriteriaConditionHttpHeaderContains,
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
