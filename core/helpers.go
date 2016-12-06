package core

import "github.com/MustWin/baremetal-sdk-go"

type resourceProvider interface {
	GetOk(string) (interface{}, bool)
}

func setListOptions(resource resourceProvider, opts *baremetal.ListOptions) {
	if val, ok := resource.GetOk("limit"); ok {
		opts.Limit = uint64(val.(int))
	}

	if val, ok := resource.GetOk("page"); ok {
		opts.Page = val.(string)
	}

	return
}

func setNextPageOption(nextPage string, opts *baremetal.ListOptions) (hasNextPage bool) {
	if nextPage == "" {
		hasNextPage = false
	} else {
		hasNextPage = true
		opts.Page = nextPage
	}

	return
}
