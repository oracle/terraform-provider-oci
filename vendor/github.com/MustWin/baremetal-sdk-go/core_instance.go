package baremetal

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Instance contains information about a compute host.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#Instance
type Instance struct {
	AvailabilityDomain string            `json:"availabilityDomain"`
	CompartmentID      string            `json:"compartmentId"`
	DisplayName        string            `json:"displayName"`
	ID                 string            `json:"id"`
	Image              string            `json:"image"`
	Metadata           map[string]string `json:"metadata"`
	Region             string            `json:"region"`
	Shape              string            `json:"shape"`
	State              string            `json:"state"`
	TimeCreated        Time              `json:"timeCreated"`
	ETag               string            `json:"etag,omitempty"`
	OPCRequestID       string            `json:"opc-request-id,omitempty"`
}

// InstanceList contains a list of instances.
type InstanceList struct {
	// NextPage if present there are more instances available to be requested.
	// Assign the NextPage value to Options.Page in a subsequent ListInstances
	// request.
	NextPage string
	// RequestID Oracle assigned identifier for the request
	RequestID string
	Instances []Instance
}

type LaunchInstanceRequest struct {
	AvailabilityDomain string            `json:"availabilityDomain"`
	CompartmentID      string            `json:"compartmentId"`
	DisplayName        string            `json:"displayName,omitempty"`
	Image              string            `json:"image"`
	Metadata           map[string]string `json:"metadata"`
	Shape              string            `json:"shape"`
	SubnetID           string            `json:"subnetId"`
}

// LaunchInstance initializes and starts a compute instance. Display name is
// set in the opts parameter.  See Oracle documentation for more information
// on other arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#launchInstance
func (c *Client) LaunchInstance(
	availabilityDomain,
	compartmentID,
	image,
	shape,
	subnetID string,
	metadata map[string]string, opts ...Options) (inst *Instance, e error) {

	var displayName string
	if len(opts) > 0 {
		displayName = opts[0].DisplayName
	}

	requestBody := LaunchInstanceRequest{
		AvailabilityDomain: availabilityDomain,
		CompartmentID:      compartmentID,
		DisplayName:        displayName,
		Image:              image,
		Metadata:           metadata,
		Shape:              shape,
		SubnetID:           subnetID,
	}

	req := &sdkRequestOptions{
		body:    requestBody,
		name:    resourceInstances,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, req); e != nil {
		return
	}

	inst = &Instance{}

	if e = json.Unmarshal(response.body, inst); e != nil {
		return
	}

	inst.ETag = response.header.Get(headerETag)
	inst.OPCRequestID = response.header.Get(headerOPCRequestID)

	return
}

// GetInstance retrieves a compute instance with instanceID
func (c *Client) GetInstance(instanceID string) (inst *Instance, e error) {
	req := &sdkRequestOptions{
		name: resourceInstances,
		ids:  urlParts{instanceID},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	inst = &Instance{}

	if e = json.Unmarshal(response.body, inst); e != nil {
		return
	}

	inst.OPCRequestID = response.header.Get(headerOPCRequestID)
	inst.ETag = response.header.Get(headerETag)

	return

}

// UpdateInstance can be used to change the display name of a compute instance
// by assigning the new name to Options.DisplayName
func (c *Client) UpdateInstance(instanceID string, opts ...Options) (inst *Instance, e error) {
	var displayName string

	if len(opts) > 0 {
		displayName = opts[0].DisplayName
	}

	requestBody := struct {
		DisplayName string `json:"displayName,omitempty"`
	}{
		DisplayName: displayName,
	}

	req := &sdkRequestOptions{
		name:    resourceInstances,
		body:    requestBody,
		ids:     urlParts{instanceID},
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, req); e != nil {
		return
	}

	inst = &Instance{}

	if e = json.Unmarshal(response.body, inst); e != nil {
		return
	}

	inst.OPCRequestID = response.header.Get(headerOPCRequestID)
	inst.ETag = response.header.Get(headerETag)

	return
}

// TerminateInstance terminates the compute instance with an ID matching
// instanceID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#terminateInstance
func (c *Client) TerminateInstance(instanceID string, opts ...Options) (e error) {
	req := &sdkRequestOptions{
		name:    resourceInstances,
		ids:     urlParts{instanceID},
		options: opts,
	}

	return c.coreApi.deleteRequest(req)
}

// ListInstances returns a list of compute instances hosted in a compartment. AvailabilityDomain
// may be included in Options to further refine results.
func (c *Client) ListInstances(compartmentID string, opts ...Options) (insts *InstanceList, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceInstances,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	insts = &InstanceList{
		NextPage:  resp.header.Get(headerOPCNextPage),
		RequestID: resp.header.Get(headerOPCRequestID),
	}

	if e = json.Unmarshal(resp.body, &insts.Instances); e != nil {
		return
	}

	return

}

// InstanceAction starts, stops, or resets a compute instance identified by
// instanceID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#instanceAction
func (c *Client) InstanceAction(instanceID string, action instanceActions, opts ...Options) (inst *Instance, e error) {

	reqOpts := &sdkRequestOptions{
		name:    resourceInstances,
		options: opts,
		ids:     urlParts{instanceID},
		query:   url.Values{},
	}

	reqOpts.query.Set(queryAction, string(action))

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	inst = &Instance{}

	if e = json.Unmarshal(response.body, inst); e != nil {
		return
	}

	inst.OPCRequestID = response.header.Get(headerOPCRequestID)
	inst.ETag = response.header.Get(headerETag)

	return

}
