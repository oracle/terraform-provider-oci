package baremetal

type Container interface {
	GetList() interface{}
}

type Pageable interface {
	SetNextPage(np string)
}

type Requestable interface {
	SetRequestID(id string)
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
