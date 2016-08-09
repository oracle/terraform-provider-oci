package baremtlsdk

import (
	"net/url"
	"strconv"
)

type CoreOptions struct {
	AvailabilityDomain string
	Limit              uint64
	Page               string
	ImageID            string
}

func (c *Client) setCoreOptions(query url.Values, options ...CoreOptions) {
	if len(options) > 0 {
		option := options[0]
		if option.AvailabilityDomain != "" {
			query.Set(queryAvailabilityDomain, option.AvailabilityDomain)
		}
		if option.ImageID != "" {
			query.Set(queryImageID, option.ImageID)
		}
		if option.Limit > 0 {
			query.Set(queryLimit, strconv.FormatUint(option.Limit, 10))
		}
		if option.Page != "" {
			query.Set(queryPage, option.Page)
		}
	}
}
