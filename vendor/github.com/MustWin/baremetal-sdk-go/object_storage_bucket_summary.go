package baremetal

import (
	"time"
)

// BucketSummary is the list representation of a bucket
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/BucketSummary/

type BucketSummary struct {
	Namespace     Namespace `json:"namespace"`
	Name          string    `json:"name"`
	CompartmentID string    `json:"compartmentId"`
	CreatedBy     string    `json:"createdBy"`
	TimeCreated   time.Time `json:"timeCreated"`
	ETag          string    `json:"etag"`
}

type ListBuckets struct {
	ResourceContainer
	ClientRequestableResource
	BucketSummaries []BucketSummary
}

func (ref *ListBuckets) GetList() interface{} {
	return &ref.BucketSummaries
}

// ListBuckets returns BucketSummaries for all the buckets in a namespace
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/BucketSummary/ListBuckets
func (c *Client) ListBuckets(
	compartmentID string,
	namespaceName Namespace,
	opts *ListBucketsOptions,
) (buckets *ListBuckets, e error) {

	required := listOCIDRequirement{}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		ids:      urlParts{resourceNamespaces, namespaceName, resourceBuckets},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	buckets = &ListBuckets{}
	e = response.unmarshal(buckets)
	return
}
