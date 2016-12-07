package core

import "github.com/MustWin/baremetal-sdk-go"

type resourceProvider interface {
	GetOk(string) (interface{}, bool)
}

func setLimitOptions(resource resourceProvider, opts *baremetal.LimitListOptions) {
	if val, ok := resource.GetOk("limit"); ok {
		opts.Limit = uint64(val.(int))
	}
	return
}

func setPageOptions(resource resourceProvider, opts *baremetal.PageListOptions) {
	if val, ok := resource.GetOk("page"); ok {
		opts.Page = val.(string)
	}
	return
}

func setListOptions(resource resourceProvider, opts *baremetal.ListOptions) {
	setLimitOptions(resource, &opts.LimitListOptions)
	setPageOptions(resource, &opts.PageListOptions)
	return
}

func setNextPageOption(nextPage string, opts *baremetal.PageListOptions) (hasNextPage bool) {
	if nextPage == "" {
		hasNextPage = false
	} else {
		hasNextPage = true
		opts.Page = nextPage
	}

	return
}
