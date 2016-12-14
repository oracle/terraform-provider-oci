package baremetal

import (
	"errors"
	"net/http"
	"reflect"
	"time"
)

type ObjectSummary struct {
	Name        string    `json:"name"`
	Size        uint64    `json:"size"`
	MD5         string    `json:"md5"`
	TimeCreated time.Time `json:"timeCreated"`
}

type ListObjects struct {
	RequestableResource
	ClientRequestableResource
	Objects       []ObjectSummary `json:"objects"`
	Prefixes      []string        `json:"prefixes"`
	NextStartWith string          `json:"nextStartWith"`
}

// ListObjects lists objects
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/ListObjects/ListObjects
func (c *Client) ListObjects(namespace Namespace, bucket string, opts *ListObjectsOptions) (objects *ListObjects, e error) {
	details := &requestDetails{
		ids: urlParts{
			resourceNamespaces,
			namespace,
			resourceBuckets,
			bucket,
			resourceObjects,
		},
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	objects = &ListObjects{}
	e = response.unmarshal(objects)
	return
}

// Objects are the items stored in ObjectStorage

type Object struct {
	HeadObject
	Size    uint64
	TraceID string
	Body    []byte
}

func (g *Object) Unmarshal(b []byte, toBeFilled interface{}) error {
	rv := reflect.ValueOf(toBeFilled)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("Value passed to unmarshal is not a pointer")
	}
	if po, ok := toBeFilled.(*Object); ok {
		po.Body = b
	} else {
		return errors.New("Value passed in was not an Object")
	}
	return nil
}

type GetObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestableResource
	Range string
}

func (opt *GetObjectOptions) Header() http.Header {
	header := http.Header{}
	if opt != nil && opt.IfMatch != "" {
		header.Set(headerIfMatch, opt.IfMatch)
	}
	if opt != nil && opt.IfNoneMatch != "" {
		header.Set(headerIfNoneMatch, opt.IfMatch)
	}
	if opt != nil && opt.Range != "" {
		header.Set(headerRange, opt.Range)
	}
	return header
}

// GetObject fetches an object from object storage
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/GetObject
func (c *Client) GetObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *GetObjectOptions,
) (object *Object, e error) {
	details := &requestDetails{
		ids: urlParts{
			resourceNamespaces,
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.getRequest(details); e != nil {
		return
	}

	object = &Object{}
	e = response.unmarshal(object)
	object.Namespace = namespace
	object.Bucket = bucketName
	object.ID = objectName
	return
}

type DeleteObject struct {
	RequestableResource
	ClientRequestableResource
	LastModifiedResourceContainer
}

// DeleteObject deletes an object from object storage
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/DeleteObject
func (c *Client) DeleteObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *DeleteObjectOptions,
) (object *DeleteObject, e error) {

	var required interface{}
	details := &requestDetails{
		ids: urlParts{
			resourceNamespaces,
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.request(http.MethodDelete, details); e != nil {
		return
	}

	object = &DeleteObject{}
	e = response.unmarshal(object)
	return
}

type HeadObject struct {
	ETaggedResource
	ClientRequestableResource
	LastModifiedResourceContainer
	ContentResource
	MetadataResource
	ID string
	Bucket string
	Namespace Namespace
}

type HeadObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestableResource
}

// HeadObject fetches the user defined metadata for an object
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/HeadObject
func (c *Client) HeadObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	opts *HeadObjectOptions,
) (headObject *HeadObject, e error) {

	var required interface{}
	details := &requestDetails{
		ids: urlParts{
			resourceNamespaces,
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.request(http.MethodHead, details); e != nil {
		return
	}

	headObject = &HeadObject{}
	e = response.unmarshal(headObject)
	return
}


// PutObject updates an object in object storage
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/objectstorage/20160918/methods/PutObject
func (c *Client) PutObject(
	namespace Namespace,
	bucketName string,
	objectName string,
	content []byte,
	opts *PutObjectOptions,
) (object *Object, e error) {

	details := &requestDetails{
		ids: urlParts{
			resourceNamespaces,
			namespace,
			resourceBuckets,
			bucketName,
			resourceObjects,
			objectName,
		},
		optional: opts,
		required: &Object{
			Body: content,
		},
	}

	var response *requestResponse
	if response, e = c.objectStorageApi.request(http.MethodPut, details); e != nil {
		return
	}

	object = &Object{}
	e = response.unmarshal(object)
	return
}
