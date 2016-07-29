package baremtlclient

import (
	"fmt"
	"net/url"
)

func buildIdentityURL(resource resourceName, query *url.Values, ids ...string) string {
	urlStr := fmt.Sprintf("%s/%s/%s", identityServiceAPI, identityServiceAPIVersion, resource)

	for _, id := range ids {
		urlStr += "/"
		urlStr += id
	}

	if len(ids) == 0 && query == nil {
		urlStr += "/"
	}

	u, _ := url.Parse(urlStr)
	if query != nil {
		q := u.Query()
		for key, vals := range *query {
			for _, val := range vals {
				q.Add(key, val)
			}
		}

		u.RawQuery += q.Encode()
	}

	return u.String()
}
