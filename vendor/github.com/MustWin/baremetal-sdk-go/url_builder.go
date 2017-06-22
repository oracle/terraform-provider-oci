// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
)

type urlBuilderFn func(string, string, resourceName, url.Values, ...interface{}) string

var urlTemplateReg = regexp.MustCompile(`^.*%s+.*%s+.*$`)

func baseUrlHelper(urlTemplate string, service string, region string) string {
	if !urlTemplateReg.MatchString(urlTemplate) {
		panic("Invalid urlTemplate:" + urlTemplate)
	}

	return fmt.Sprintf(urlTemplate, service, region)
}

func buildCoreURL(urlTemplate string, region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := baseUrlHelper(urlTemplate, coreServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, coreServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildIdentityURL(urlTemplate string, region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := baseUrlHelper(urlTemplate, identityServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, identityServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildDatabaseURL(urlTemplate string, region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := baseUrlHelper(urlTemplate, databaseServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, databaseServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildObjectStorageURL(urlTemplate string, region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := baseUrlHelper(urlTemplate, objectStorageServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s", baseUrl, resourceNamespaces)
	if resource == resourcePAR {
		return buildURL(urlStr, make(map[string][]string), ids...)
	}
	return buildURL(urlStr, query, ids...)
}

func buildLoadBalancerURL(urlTemplate string, region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := baseUrlHelper(urlTemplate, loadBalancerServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, loadBalancerServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildURL(urlStr string, query url.Values, ids ...interface{}) string {
	const separator = "/"
	for _, id := range ids {
		var strVal string

		switch id := id.(type) {
		default:
			panic("Unsupported type")
		case bool:
			strVal = strconv.FormatBool(id)
		case uint64:
			strVal = strconv.FormatUint(id, 10)
		case string:
			strVal = id
		case resourceName:
			strVal = string(id)
		case Namespace:
			strVal = string(id)
		}

		if strVal != separator {
			urlStr += separator
		}

		urlStr += strVal
	}

	u, _ := url.Parse(urlStr)
	if query != nil {
		q := u.Query()
		for key, vals := range query {
			for _, val := range vals {
				q.Add(key, val)
			}
		}

		u.RawQuery += q.Encode()
	}

	return u.String()

}
