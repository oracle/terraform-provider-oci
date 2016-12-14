package baremetal

import (
	"net/http"
)

// Bucket stores arbitrary objects on a given key
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/Bucket/

type Bucket struct {
	ETaggedResource
	Namespace     Namespace         `json:"namespace"`
	Name          string            `json:"name"`
	CompartmentID string            `json:"compartmentId"`
	Metadata      map[string]string `json:"metadata"`
	CreatedBy     string            `json:"createdBy"`
	TimeCreated   Time              `json:"timeCreated"`
}

// CreateBucket initializes and creates a storage bucket. Namespace is
// set in the opts parameter. See Oracle documentation for more information
// on other arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/Bucket/CreateBucket
func (c *Client) CreateBucket(
	compartmentID string,
	name string,
	namespaceName Namespace,
	opts *CreateBucketOptions,
) (bckt *Bucket, e error) {

	required := struct {
		ocidRequirement
		Name string `json:"name" url:"-"`
	}{
		Name: name,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		ids:      urlParts{namespaceName, resourceBuckets},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.request(http.MethodPost, details); e != nil {
		return
	}

	bckt = &Bucket{}
	e = response.unmarshal(bckt)
	return
}

// GetBucket gets the current representation of the given bucket in the given namespace.
//
// See: https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/Bucket/GetBucket
func (c *Client) GetBucket(
	bucketName string,
	namespaceName Namespace,
) (bckt *Bucket, e error) {
	details := &requestDetails{
		ids: urlParts{namespaceName, resourceBuckets, bucketName},
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	bckt = &Bucket{}
	e = response.unmarshal(bckt)
	return
}

// UpdateBucket performs a partial (or full) update of a bucket, currently including just the user-defined metadata
//
// See: https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/Bucket/UpdateBucket
func (c *Client) UpdateBucket(
	compartmentID string,
	name string,
	namespaceName Namespace,
	opts *UpdateBucketOptions,
) (bckt *Bucket, e error) {

	required := struct {
		ocidRequirement
		Name string `json:"name" url:"-"`
	}{
		Name: name,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		ids:      urlParts{namespaceName, resourceBuckets, name},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.request(http.MethodPut, details); e != nil {
		return
	}

	bckt = &Bucket{}
	e = response.unmarshal(bckt)
	return
}

// DeleteBucket deletes a bucket if it is already empty. If the bucket is not empty, use DeleteObject first.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/Bucket/DeleteBucket
func (c *Client) DeleteBucket(
	name string,
	namespaceName Namespace,
	opts *IfMatchOptions,
) (e error) {
	required := struct {
		ocidRequirement
		Name string `json:"name" url:"-"`
	}{
		Name: name,
	}

	details := &requestDetails{
		ids:      urlParts{namespaceName, resourceBuckets, name},
		optional: opts,
		required: required,
	}

	return c.objectStorageApi.deleteRequest(details)
}

type HeadBucket struct {
	ETaggedResource
	ClientRequestableResource
}

type HeadBucketOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestableResource
}

// HeadBucket checks that a bucket exists and returns the ETag
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/HeadBucket
func (c *Client) HeadBucket(
	namespace Namespace,
	bucketName string,
	opts *HeadBucketOptions,
) (headBucket *HeadBucket, e error) {

	var required interface{}
	details := &requestDetails{
		ids: urlParts{
			namespace,
			resourceBuckets,
			bucketName,
		},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	headBucket = &HeadBucket{}
	e = response.unmarshal(headBucket)
	return
}
