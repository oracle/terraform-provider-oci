package main

import (
	"fmt"

	"github.com/MustWin/baremetal-sdk-go"
)

type resourceProvider interface {
	GetOk(string) (interface{}, bool)
}

func getCoreOptionsFromResourceData(resource resourceProvider, keys ...string) (opts []baremetal.CoreOptions) {
	opts = []baremetal.CoreOptions{}

	for _, key := range keys {
		if val, ok := resource.GetOk(key); ok {

			if len(opts) == 0 {
				opts = append(opts, baremetal.CoreOptions{})
			}

			switch key {
			case "availability_domain":
				opts[0].AvailabilityDomain = val.(string)
			case "image_id":
				opts[0].ImageID = val.(string)
			case "instance_id":
				opts[0].InstanceID = val.(string)
			case "vnic_id":
				opts[0].VnicID = val.(string)
			case "page":
				opts[0].Page = val.(string)
			case "limit":
				opts[0].Limit = val.(uint64)
			default:
				panic(fmt.Sprintf("Unknown key '%s' supplied for CoreOptions", key))
			}
		}
	}

	return
}
