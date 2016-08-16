package baremetal

import (
	"fmt"
	"net/url"
	"strconv"
)

type urlBuilderFn func(resourceName, url.Values, ...interface{}) string

func buildCoreURL(resource resourceName, query url.Values, ids ...interface{}) string {
	urlStr := fmt.Sprintf("%s/%s/%s", coreServiceAPI, coreServiceAPIVersion, resource)
	return buildURL(urlStr, query, ids...)
}

func buildIdentityURL(resource resourceName, query url.Values, ids ...interface{}) string {
	urlStr := fmt.Sprintf("%s/%s/%s", identityServiceAPI, identityServiceAPIVersion, resource)
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
