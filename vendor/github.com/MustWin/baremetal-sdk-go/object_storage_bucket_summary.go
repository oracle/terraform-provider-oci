// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"time"
)

// BucketSummary is the list representation of a bucket
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/BucketSummary/

type BucketSummary struct {
	Namespace     Namespace `json:"namespace"`
	Name          string    `json:"name"`
	CompartmentID string    `json:"compartmentId"`
	CreatedBy     string    `json:"createdBy"`
	TimeCreated   time.Time `json:"timeCreated"`
	ETag          string    `json:"etag"`
}

type ListBuckets struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	OPCClientRequestIDUnmarshaller
	BucketSummaries []BucketSummary
}

func (ref *ListBuckets) GetList() interface{} {
	return &ref.BucketSummaries
}

// ListBuckets returns BucketSummaries for all the buckets in a namespace
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/BucketSummary/ListBuckets
func (c *Client) ListBuckets(
	compartmentID string,
	namespaceName Namespace,
	opts *ListBucketsOptions,
) (buckets *ListBuckets, e error) {

	required := listOCIDRequirement{}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		ids:      urlParts{namespaceName, resourceBuckets},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	buckets = &ListBuckets{}
	e = resp.unmarshal(buckets)
	return
}
