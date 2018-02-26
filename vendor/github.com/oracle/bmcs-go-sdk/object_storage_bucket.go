// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// Bucket stores arbitrary objects on a given key
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/Bucket/

type Bucket struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	Namespace     Namespace         `json:"namespace"`
	Name          string            `json:"name"`
	CompartmentID string            `json:"compartmentId"`
	Metadata      map[string]string `json:"metadata"`
	CreatedBy     string            `json:"createdBy"`
	TimeCreated   Time              `json:"timeCreated"`
	AccessType    BucketAccessType  `json:"publicAccessType"`
}

// CreateBucket initializes and creates a storage bucket. Namespace is
// set in the opts parameter. See Oracle documentation for more information
// on other arguments.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/Bucket/CreateBucket
func (c *Client) CreateBucket(
	compartmentID string,
	name string,
	namespaceName Namespace,
	opts *CreateBucketOptions,
) (bckt *Bucket, e error) {

	required := struct {
		ocidRequirement
		Name string `header:"-" json:"name" url:"-"`
	}{
		Name: name,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		ids:      urlParts{namespaceName, resourceBuckets},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.objectStorageApi.postRequest(details); e != nil {
		return
	}

	bckt = &Bucket{}
	e = resp.unmarshal(bckt)
	return
}

// GetBucket gets the current representation of the given bucket in the given namespace.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/Bucket/GetBucket
func (c *Client) GetBucket(
	bucketName string,
	namespaceName Namespace,
) (bckt *Bucket, e error) {
	details := &requestDetails{
		ids: urlParts{namespaceName, resourceBuckets, bucketName},
	}

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	bckt = &Bucket{}
	e = resp.unmarshal(bckt)
	return
}

// UpdateBucket performs a partial (or full) update of a bucket, currently including just the user-defined metadata
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/Bucket/UpdateBucket
func (c *Client) UpdateBucket(
	compartmentID string,
	name string,
	namespaceName Namespace,
	opts *UpdateBucketOptions,
) (bckt *Bucket, e error) {

	required := struct {
		ocidRequirement
	}{}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		ids:      urlParts{namespaceName, resourceBuckets, name},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.objectStorageApi.postRequest(details); e != nil {
		return
	}

	bckt = &Bucket{}
	e = resp.unmarshal(bckt)
	return
}

// DeleteBucket deletes a bucket if it is already empty. If the bucket is not empty, use DeleteObject first.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/Bucket/DeleteBucket
func (c *Client) DeleteBucket(
	name string,
	namespaceName Namespace,
	opts *IfMatchOptions,
) (e error) {
	required := struct {
		ocidRequirement
		Name string `header:"-" json:"name" url:"-"`
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
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	OPCClientRequestIDUnmarshaller
}

type HeadBucketOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	OPCClientRequestIDUnmarshaller
}

// HeadBucket checks that a bucket exists and returns the ETag
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/methods/HeadBucket
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

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	headBucket = &HeadBucket{}
	e = resp.unmarshal(headBucket)
	return
}
