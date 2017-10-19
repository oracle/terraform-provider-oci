// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "time"

// JSON bodies are unmarshalled by the json package. The unmarshallers in this
// file are primarily responsible for unmarshalling data returned via the
// response header.  That said, any response data that cannot be unmarshalled
// using json should be configured here.

type OPCRequestIDUnmarshallable interface {
	SetRequestID(id string)
}

type OPCRequestIDUnmarshaller struct {
	RequestID string `json:"RequestID,omitempty" url:"-"`
}

func (r *OPCRequestIDUnmarshaller) SetRequestID(id string) {
	r.RequestID = id
}

type NextPageUnmarshallable interface {
	SetNextPage(np string)
}

type NextPageUnmarshaller struct {
	NextPage string
}

func (r *NextPageUnmarshaller) SetNextPage(np string) {
	r.NextPage = np
}

type ETagUnmarshallable interface {
	SetETag(etag string)
}

type ETagUnmarshaller struct {
	ETag string
}

func (r *ETagUnmarshaller) SetETag(etag string) {
	r.ETag = etag
}

type OPCClientRequestIDUnmarshallable interface {
	SetClientRequestID(id string)
}

type OPCClientRequestIDUnmarshaller struct {
	ClientRequestID string
}

func (r *OPCClientRequestIDUnmarshaller) SetClientRequestID(id string) {
	r.ClientRequestID = id
}

type OPCWorkRequestIDUnmarshallable interface {
	SetWorkRequestID(id string)
}

type OPCWorkRequestIDUnmarshaller struct {
	WorkRequestID string `json:"WorkRequestID,omitempty" url:"-"`
}

func (r *OPCWorkRequestIDUnmarshaller) SetWorkRequestID(id string) {
	r.WorkRequestID = id
}

type LastModifiedUnmarshallable interface {
	SetLastModified(time.Time)
}

type LastModifiedUnmarshaller struct {
	LastModified time.Time
}

func (r *LastModifiedUnmarshaller) SetLastModified(time time.Time) {
	r.LastModified = time
}

type MetadataUnmarshallable interface {
	GetMetadata() map[string]string
	SetMetadata(map[string]string)
}

type MetadataUnmarshaller struct {
	Metadata map[string]string `json:"-" url:"-" header:"-"`
}

func (mr *MetadataUnmarshaller) GetMetadata() map[string]string {
	return mr.Metadata
}

func (mr *MetadataUnmarshaller) SetMetadata(md map[string]string) {
	mr.Metadata = md
}

type ContentUnmarshallable interface {
	MetadataUnmarshallable
	SetContentLength(uint64)
	SetContentMD5(string)
	SetContentType(string)
	SetContentLanguage(string)
	SetContentEncoding(string)
}

type ContentUnmarshaller struct {
	ContentLength   uint64
	ContentRange    string
	ContentMD5      string
	ContentType     string
	ContentLanguage string
	ContentEncoding string
}

func (ref *ContentUnmarshaller) SetContentLength(l uint64) {
	ref.ContentLength = l
}

func (ref *ContentUnmarshaller) SetContentRange(r string) {
	ref.ContentRange = r
}

func (ref *ContentUnmarshaller) SetContentMD5(md5 string) {
	ref.ContentMD5 = md5
}

func (ref *ContentUnmarshaller) SetContentType(t string) {
	ref.ContentType = t
}

func (ref *ContentUnmarshaller) SetContentLanguage(l string) {
	ref.ContentLanguage = l
}

func (ref *ContentUnmarshaller) SetContentEncoding(l string) {
	ref.ContentEncoding = l
}

type BodyUnmarshallable interface {
	SetBody([]byte, interface{}) error
}
