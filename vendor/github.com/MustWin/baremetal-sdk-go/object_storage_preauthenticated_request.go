package baremetal

type PreauthenticatedRequest struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	ObjectName  string        `json:"objectName"`
	AccessURI   string        `json:"accessUri"`
	AccessType  PARAccessType `json:"accessType"`
	TimeExpires Time          `json:"timeExpires"`
	TimeCreated Time          `json:"timeCreated"`
}

type PreauthenticatedRequestSummary struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	ObjectName  string        `json:"objectName"`
	AccessType  PARAccessType `json:"accessType"`
	TimeExpires Time          `json:"timeExpires"`
	TimeCreated Time          `json:"timeCreated"`
}

type ListPreauthenticatedRequests struct {
	OPCClientRequestIDUnmarshaller
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	PreauthenticatedRequests []PreauthenticatedRequestSummary
}

func buildPARUrlParts(namespace Namespace, bucketName string, rest ...interface{}) urlParts {
	fixedParts := urlParts{namespace, resourceBuckets, bucketName, resourcePAR}
	parts := fixedParts
	for _, elem := range rest {
		parts = append(parts, elem)
	}
	return parts
}

// CreatePreauthenticatedRequest creates a url that can be used to
// access an object or a bucket with a url. The preauthenticated request is identified by its id
// See Oracle documentation for more information
//
//https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/PreauthenticatedRequest/CreatePreauthenticatedRequest
func (c *Client) CreatePreauthenticatedRequest(
	namespace Namespace,
	bucketName string,
	parDetails *CreatePreauthenticatedRequestDetails,

) (par *PreauthenticatedRequest, e error) {
	details := requestDetails{
		ids:      buildPARUrlParts(namespace, bucketName),
		name:     resourcePAR,
		required: parDetails,
	}

	var res *response
	if res, e = c.objectStorageApi.postRequest(&details); e != nil {
		return
	}
	par = &PreauthenticatedRequest{}
	e = res.unmarshal(&par)
	return
}

// DeletePreauthenticatedRequest deletes a preauthenticated request by its id.
//
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/PreauthenticatedRequest/DeletePreauthenticatedRequest
func (c *Client) DeletePreauthenticatedRequest(namespace Namespace,
	bucketName string,
	parId string,
	options *ClientRequestOptions,
) (e error) {
	details := &requestDetails{
		ids:      buildPARUrlParts(namespace, bucketName, parId),
		optional: options,
	}

	return c.objectStorageApi.deleteRequest(details)
}

// Gets information about a previously created preauthenticated request
//
// https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/PreauthenticatedRequestSummary/GetPreauthenticatedRequest
func (c *Client) GetPreauthenticatedRequest(namespace Namespace,
	bucketName string,
	parId string,
	options *ClientRequestOptions,
) (par *PreauthenticatedRequestSummary, e error) {
	details := requestDetails{
		ids:      buildPARUrlParts(namespace, bucketName, parId),
		optional: options,
	}
	var res *response
	if res, e = c.objectStorageApi.getRequest(&details); e != nil {
		return
	}
	par = &PreauthenticatedRequestSummary{}
	e = res.unmarshal(&par)
	return

}

// Lists preauthenticated request on a give bucket, optionally receives a prefix for filtering by name
//
//https://docs.us-phoenix-1.oraclecloud.com/api/#/en/objectstorage/20160918/PreauthenticatedRequestSummary/ListPreauthenticatedRequests
func (c *Client) ListPreauthenticatedRequests(
	namespace Namespace,
	bucketName string,
	opt *ListPreauthenticatedRequestOptions) (parList *ListPreauthenticatedRequests, e error) {
	details := requestDetails{
		ids:      buildPARUrlParts(namespace, bucketName),
		optional: opt,
	}

	var resp *response
	if resp, e = c.objectStorageApi.getRequest(&details); e != nil {
		return
	}

	parList = &ListPreauthenticatedRequests{}
	e = resp.unmarshal(parList)
	return
}

func (ref *ListPreauthenticatedRequests) GetList() []PreauthenticatedRequestSummary {
	return ref.PreauthenticatedRequests
}
