package baremetal

import (
	"net/http"
	"time"
)

type CustomUnmarshaler interface {
	Unmarshal([]byte, interface{}) error
}

type Container interface {
	GetList() interface{}
}

type Pageable interface {
	SetNextPage(np string)
}

type Requestable interface {
	SetRequestID(id string)
}

type ClientRequestable interface {
	SetClientRequestID(id string)
}

type ClientRequestOptions struct {
	OPCClientRequestID string
}

func (ref *ClientRequestOptions) Header() http.Header {
	h := http.Header{}
	if ref.OPCClientRequestID != "" {
		h.Add(headerOPCClientRequestID, ref.OPCClientRequestID)
	}
	return h
}

type ETagged interface {
	SetETag(etag string)
}

type RequestableResource struct {
	RequestID string
}

func (r *RequestableResource) SetRequestID(id string) {
	r.RequestID = id
}

type LastModifiedResource interface {
	SetLastModified(time.Time)
}

type MetaDataRequestable interface {
	SetMetadata(map[string]string)
}

type MetadataResource struct {
	Metadata map[string]string
}

func (mr *MetadataResource) SetMetadata(md map[string]string) {
	mr.Metadata = md
}

type ContentRequestable interface {
	MetaDataRequestable
	SetContentLength(uint64)
	SetContentMD5(string)
	SetContentType(string)
	SetContentLanguage(string)
	SetContentEncoding(string)
}

type ContentResource struct {
	ContentLength   uint64
	ContentRange    string
	MD5             string
	ContentType     string
	ContentLanguage string
	ContentEncoding string
}

func (ref *ContentResource) SetContentLength(l uint64) {
	ref.ContentLength = l
}

func (ref *ContentResource) SetContentRange(r string) {
	ref.ContentRange = r
}

func (ref *ContentResource) SetMD5(md5 string) {
	ref.MD5 = md5
}

func (ref *ContentResource) SetContentType(t string) {
	ref.ContentType = t
}

func (ref *ContentResource) SetContentLanguage(l string) {
	ref.ContentLanguage = l
}

func (ref *ContentResource) SetContentEncoding(l string) {
	ref.ContentEncoding = l
}

type ClientRequestableResource struct {
	ClientRequestID string
}

func (r *ClientRequestableResource) SetClientRequestID(id string) {
	r.ClientRequestID = id
}

type LastModifiedResourceContainer struct {
	LastModified time.Time
}

func (r *LastModifiedResourceContainer) SetLastModified(time time.Time) {
	r.LastModified = time
}

type ResourceContainer struct {
	RequestableResource
	NextPage string
}

func (r *ResourceContainer) SetNextPage(np string) {
	r.NextPage = np
}

type ETaggedResource struct {
	RequestableResource
	ETag string
}

func (r *ETaggedResource) SetETag(etag string) {
	r.ETag = etag
}
