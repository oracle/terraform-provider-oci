// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"fmt"
	"net/url"
	"strconv"
)

type urlBuilderFn func(string, resourceName, url.Values, ...interface{}) string

func buildCoreURL(region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := fmt.Sprintf(coreServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, coreServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildIdentityURL(region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := fmt.Sprintf(identityServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, identityServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildDatabaseURL(region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := fmt.Sprintf(databaseServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, databaseServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildObjectStorageURL(region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := fmt.Sprintf(objectStorageServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s", baseUrl, resourceNamespaces)
	return buildURL(urlStr, query, ids...)
}

func buildLoadBalancerURL(region string, resource resourceName, query url.Values, ids ...interface{}) string {
	baseUrl := fmt.Sprintf(loadBalancerServiceAPI, region)
	urlStr := fmt.Sprintf("%s/%s/%s", baseUrl, loadBalancerServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildURL(urlStr string, query url.Values, ids ...interface{}) string {
	const seperator = "/"
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

		if strVal != seperator {
			urlStr += seperator
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
