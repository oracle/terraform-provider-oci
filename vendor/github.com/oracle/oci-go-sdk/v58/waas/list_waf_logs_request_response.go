// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListWafLogsRequest wrapper for the ListWafLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/waas/ListWafLogs.go.html to see an example of how to use ListWafLogsRequest.
type ListWafLogsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WAAS policy.
	WaasPolicyId *string `mandatory:"true" contributesTo:"path" name:"waasPolicyId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated call. If unspecified, defaults to `20`.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous paginated call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter that matches log entries where the observed event occurred on or after a date and time specified in RFC 3339 format. If unspecified, defaults to two hours before receipt of the request.
	TimeObservedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeObservedGreaterThanOrEqualTo"`

	// A filter that matches log entries where the observed event occurred before a date and time, specified in RFC 3339 format.
	TimeObservedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeObservedLessThan"`

	// A full text search for logs.
	TextContains *string `mandatory:"false" contributesTo:"query" name:"textContains"`

	// Filters logs by access rule key.
	AccessRuleKey []string `contributesTo:"query" name:"accessRuleKey" collectionFormat:"multi"`

	// Filters logs by Web Application Firewall action.
	Action []ListWafLogsActionEnum `contributesTo:"query" name:"action" omitEmpty:"true" collectionFormat:"multi"`

	// Filters logs by client IP address.
	ClientAddress []string `contributesTo:"query" name:"clientAddress" collectionFormat:"multi"`

	// Filters logs by country code. Country codes are in ISO 3166-1 alpha-2 format. For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	CountryCode []string `contributesTo:"query" name:"countryCode" collectionFormat:"multi"`

	// Filter logs by country name.
	CountryName []string `contributesTo:"query" name:"countryName" collectionFormat:"multi"`

	// Filter logs by device fingerprint.
	Fingerprint []string `contributesTo:"query" name:"fingerprint" collectionFormat:"multi"`

	// Filter logs by HTTP method.
	HttpMethod []ListWafLogsHttpMethodEnum `contributesTo:"query" name:"httpMethod" omitEmpty:"true" collectionFormat:"multi"`

	// Filter logs by incident key.
	IncidentKey []string `contributesTo:"query" name:"incidentKey" collectionFormat:"multi"`

	// Filter by log type. For more information about WAF logs, see Logs (https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/logs.htm).
	LogType []ListWafLogsLogTypeEnum `contributesTo:"query" name:"logType" omitEmpty:"true" collectionFormat:"multi"`

	// Filter by origin IP address.
	OriginAddress []string `contributesTo:"query" name:"originAddress" collectionFormat:"multi"`

	// Filter by referrer.
	Referrer []string `contributesTo:"query" name:"referrer" collectionFormat:"multi"`

	// Filter by request URL.
	RequestUrl []string `contributesTo:"query" name:"requestUrl" collectionFormat:"multi"`

	// Filter by response code.
	ResponseCode []int `contributesTo:"query" name:"responseCode" collectionFormat:"multi"`

	// Filter by threat feed key.
	ThreatFeedKey []string `contributesTo:"query" name:"threatFeedKey" collectionFormat:"multi"`

	// Filter by user agent.
	UserAgent []string `contributesTo:"query" name:"userAgent" collectionFormat:"multi"`

	// Filter by protection rule key.
	ProtectionRuleKey []string `contributesTo:"query" name:"protectionRuleKey" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWafLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWafLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWafLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWafLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWafLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Action {
		if _, ok := GetMappingListWafLogsActionEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", val, strings.Join(GetListWafLogsActionEnumStringValues(), ",")))
		}
	}

	for _, val := range request.HttpMethod {
		if _, ok := GetMappingListWafLogsHttpMethodEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HttpMethod: %s. Supported values are: %s.", val, strings.Join(GetListWafLogsHttpMethodEnumStringValues(), ",")))
		}
	}

	for _, val := range request.LogType {
		if _, ok := GetMappingListWafLogsLogTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LogType: %s. Supported values are: %s.", val, strings.Join(GetListWafLogsLogTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWafLogsResponse wrapper for the ListWafLogs operation
type ListWafLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WafLog instances
	Items []WafLog `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results may remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWafLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWafLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWafLogsActionEnum Enum with underlying type: string
type ListWafLogsActionEnum string

// Set of constants representing the allowable values for ListWafLogsActionEnum
const (
	ListWafLogsActionBlock      ListWafLogsActionEnum = "BLOCK"
	ListWafLogsActionDetect     ListWafLogsActionEnum = "DETECT"
	ListWafLogsActionBypass     ListWafLogsActionEnum = "BYPASS"
	ListWafLogsActionLog        ListWafLogsActionEnum = "LOG"
	ListWafLogsActionRedirected ListWafLogsActionEnum = "REDIRECTED"
)

var mappingListWafLogsActionEnum = map[string]ListWafLogsActionEnum{
	"BLOCK":      ListWafLogsActionBlock,
	"DETECT":     ListWafLogsActionDetect,
	"BYPASS":     ListWafLogsActionBypass,
	"LOG":        ListWafLogsActionLog,
	"REDIRECTED": ListWafLogsActionRedirected,
}

// GetListWafLogsActionEnumValues Enumerates the set of values for ListWafLogsActionEnum
func GetListWafLogsActionEnumValues() []ListWafLogsActionEnum {
	values := make([]ListWafLogsActionEnum, 0)
	for _, v := range mappingListWafLogsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetListWafLogsActionEnumStringValues Enumerates the set of values in String for ListWafLogsActionEnum
func GetListWafLogsActionEnumStringValues() []string {
	return []string{
		"BLOCK",
		"DETECT",
		"BYPASS",
		"LOG",
		"REDIRECTED",
	}
}

// GetMappingListWafLogsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWafLogsActionEnum(val string) (ListWafLogsActionEnum, bool) {
	mappingListWafLogsActionEnumIgnoreCase := make(map[string]ListWafLogsActionEnum)
	for k, v := range mappingListWafLogsActionEnum {
		mappingListWafLogsActionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWafLogsActionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWafLogsHttpMethodEnum Enum with underlying type: string
type ListWafLogsHttpMethodEnum string

// Set of constants representing the allowable values for ListWafLogsHttpMethodEnum
const (
	ListWafLogsHttpMethodOptions ListWafLogsHttpMethodEnum = "OPTIONS"
	ListWafLogsHttpMethodGet     ListWafLogsHttpMethodEnum = "GET"
	ListWafLogsHttpMethodHead    ListWafLogsHttpMethodEnum = "HEAD"
	ListWafLogsHttpMethodPost    ListWafLogsHttpMethodEnum = "POST"
	ListWafLogsHttpMethodPut     ListWafLogsHttpMethodEnum = "PUT"
	ListWafLogsHttpMethodDelete  ListWafLogsHttpMethodEnum = "DELETE"
	ListWafLogsHttpMethodTrace   ListWafLogsHttpMethodEnum = "TRACE"
	ListWafLogsHttpMethodConnect ListWafLogsHttpMethodEnum = "CONNECT"
)

var mappingListWafLogsHttpMethodEnum = map[string]ListWafLogsHttpMethodEnum{
	"OPTIONS": ListWafLogsHttpMethodOptions,
	"GET":     ListWafLogsHttpMethodGet,
	"HEAD":    ListWafLogsHttpMethodHead,
	"POST":    ListWafLogsHttpMethodPost,
	"PUT":     ListWafLogsHttpMethodPut,
	"DELETE":  ListWafLogsHttpMethodDelete,
	"TRACE":   ListWafLogsHttpMethodTrace,
	"CONNECT": ListWafLogsHttpMethodConnect,
}

// GetListWafLogsHttpMethodEnumValues Enumerates the set of values for ListWafLogsHttpMethodEnum
func GetListWafLogsHttpMethodEnumValues() []ListWafLogsHttpMethodEnum {
	values := make([]ListWafLogsHttpMethodEnum, 0)
	for _, v := range mappingListWafLogsHttpMethodEnum {
		values = append(values, v)
	}
	return values
}

// GetListWafLogsHttpMethodEnumStringValues Enumerates the set of values in String for ListWafLogsHttpMethodEnum
func GetListWafLogsHttpMethodEnumStringValues() []string {
	return []string{
		"OPTIONS",
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"TRACE",
		"CONNECT",
	}
}

// GetMappingListWafLogsHttpMethodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWafLogsHttpMethodEnum(val string) (ListWafLogsHttpMethodEnum, bool) {
	mappingListWafLogsHttpMethodEnumIgnoreCase := make(map[string]ListWafLogsHttpMethodEnum)
	for k, v := range mappingListWafLogsHttpMethodEnum {
		mappingListWafLogsHttpMethodEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWafLogsHttpMethodEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWafLogsLogTypeEnum Enum with underlying type: string
type ListWafLogsLogTypeEnum string

// Set of constants representing the allowable values for ListWafLogsLogTypeEnum
const (
	ListWafLogsLogTypeAccess                     ListWafLogsLogTypeEnum = "ACCESS"
	ListWafLogsLogTypeProtectionRules            ListWafLogsLogTypeEnum = "PROTECTION_RULES"
	ListWafLogsLogTypeJsChallenge                ListWafLogsLogTypeEnum = "JS_CHALLENGE"
	ListWafLogsLogTypeCaptcha                    ListWafLogsLogTypeEnum = "CAPTCHA"
	ListWafLogsLogTypeAccessRules                ListWafLogsLogTypeEnum = "ACCESS_RULES"
	ListWafLogsLogTypeThreatFeeds                ListWafLogsLogTypeEnum = "THREAT_FEEDS"
	ListWafLogsLogTypeHumanInteractionChallenge  ListWafLogsLogTypeEnum = "HUMAN_INTERACTION_CHALLENGE"
	ListWafLogsLogTypeDeviceFingerprintChallenge ListWafLogsLogTypeEnum = "DEVICE_FINGERPRINT_CHALLENGE"
	ListWafLogsLogTypeAddressRateLimiting        ListWafLogsLogTypeEnum = "ADDRESS_RATE_LIMITING"
)

var mappingListWafLogsLogTypeEnum = map[string]ListWafLogsLogTypeEnum{
	"ACCESS":                       ListWafLogsLogTypeAccess,
	"PROTECTION_RULES":             ListWafLogsLogTypeProtectionRules,
	"JS_CHALLENGE":                 ListWafLogsLogTypeJsChallenge,
	"CAPTCHA":                      ListWafLogsLogTypeCaptcha,
	"ACCESS_RULES":                 ListWafLogsLogTypeAccessRules,
	"THREAT_FEEDS":                 ListWafLogsLogTypeThreatFeeds,
	"HUMAN_INTERACTION_CHALLENGE":  ListWafLogsLogTypeHumanInteractionChallenge,
	"DEVICE_FINGERPRINT_CHALLENGE": ListWafLogsLogTypeDeviceFingerprintChallenge,
	"ADDRESS_RATE_LIMITING":        ListWafLogsLogTypeAddressRateLimiting,
}

// GetListWafLogsLogTypeEnumValues Enumerates the set of values for ListWafLogsLogTypeEnum
func GetListWafLogsLogTypeEnumValues() []ListWafLogsLogTypeEnum {
	values := make([]ListWafLogsLogTypeEnum, 0)
	for _, v := range mappingListWafLogsLogTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWafLogsLogTypeEnumStringValues Enumerates the set of values in String for ListWafLogsLogTypeEnum
func GetListWafLogsLogTypeEnumStringValues() []string {
	return []string{
		"ACCESS",
		"PROTECTION_RULES",
		"JS_CHALLENGE",
		"CAPTCHA",
		"ACCESS_RULES",
		"THREAT_FEEDS",
		"HUMAN_INTERACTION_CHALLENGE",
		"DEVICE_FINGERPRINT_CHALLENGE",
		"ADDRESS_RATE_LIMITING",
	}
}

// GetMappingListWafLogsLogTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWafLogsLogTypeEnum(val string) (ListWafLogsLogTypeEnum, bool) {
	mappingListWafLogsLogTypeEnumIgnoreCase := make(map[string]ListWafLogsLogTypeEnum)
	for k, v := range mappingListWafLogsLogTypeEnum {
		mappingListWafLogsLogTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWafLogsLogTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
