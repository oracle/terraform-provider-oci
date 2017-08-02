// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Instance contains information about a compute host.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/
type Instance struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain string                 `json:"availabilityDomain"`
	CompartmentID      string                 `json:"compartmentId"`
	DisplayName        string                 `json:"displayName"`
	ID                 string                 `json:"id"`
	ImageID            string                 `json:"imageId"`
	Metadata           map[string]string      `json:"metadata"`
	ExtendedMetadata   map[string]interface{} `json:"extendedMetadata"`
	Region             string                 `json:"region"`
	Shape              string                 `json:"shape"`
	State              string                 `json:"lifecycleState"`
	TimeCreated        Time                   `json:"timeCreated"`
	IpxeScript         string                 `json:"ipxeScript"`
}

// InstanceCredentials contains first run windows instance credentials
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InstanceCredentials/GetWindowsInstanceInitialCredentials
type InstanceCredentials struct {
	Username string
	Password string
}

// ListInstances contains a list of instances.
type ListInstances struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Instances []Instance
}

func (l *ListInstances) GetList() interface{} {
	return &l.Instances
}

// LaunchInstance initializes and starts a compute instance. Display name is
// set in the opts parameter.  See Oracle documentation for more information
// on other arguments.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/LaunchInstance
func (c *Client) LaunchInstance(
	availabilityDomain,
	compartmentID,
	image,
	shape,
	subnetID string,
	opts *LaunchInstanceOptions) (inst *Instance, e error) {

	required := struct {
		ocidRequirement
		AvailabilityDomain string `header:"-" json:"availabilityDomain" url:"-"`
		ImageID            string `header:"-" json:"imageId" url:"-"`
		Shape              string `header:"-" json:"shape" url:"-"`
		SubnetID           string `header:"-" json:"subnetId,omitempty" url:"-"`
	}{
		AvailabilityDomain: availabilityDomain,
		ImageID:            image,
		Shape:              shape,
		SubnetID:           subnetID,
	}
	required.CompartmentID = compartmentID

	req := &requestDetails{
		name:     resourceInstances,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(req); e != nil {
		return
	}

	inst = &Instance{}
	e = resp.unmarshal(inst)
	return
}

// GetInstance retrieves a compute instance with instanceID
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/GetInstance
func (c *Client) GetInstance(id string) (inst *Instance, e error) {
	details := &requestDetails{
		name: resourceInstances,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	inst = &Instance{}
	e = resp.unmarshal(inst)
	return
}

// UpdateInstance can be used to change the display name of a compute instance
// by assigning the new name to Options.DisplayName
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/UpdateInstance
func (c *Client) UpdateInstance(id string, opts *UpdateOptions) (inst *Instance, e error) {
	details := &requestDetails{
		name:     resourceInstances,
		ids:      urlParts{id},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	inst = &Instance{}
	e = resp.unmarshal(inst)
	return
}

// TerminateInstance terminates the compute instance with an ID matching
// instanceID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/TerminateInstance
func (c *Client) TerminateInstance(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceInstances,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListInstances returns a list of compute instances hosted in a compartment. AvailabilityDomain
// may be included in Options to further refine results.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/ListInstances
func (c *Client) ListInstances(compartmentID string, opts *ListInstancesOptions) (insts *ListInstances, e error) {
	details := &requestDetails{
		name:     resourceInstances,
		required: listOCIDRequirement{CompartmentID: compartmentID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	insts = &ListInstances{}
	e = resp.unmarshal(insts)
	return
}

// InstanceAction starts, stops, or resets a compute instance identified by
// instanceID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Instance/InstanceAction
func (c *Client) InstanceAction(id string, action InstanceActions, opts *HeaderOptions) (inst *Instance, e error) {
	required := struct {
		Action string `header:"-" json:"-" url:"action"`
	}{
		Action: string(action),
	}

	details := &requestDetails{
		name:     resourceInstances,
		ids:      urlParts{id},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	inst = &Instance{}
	e = resp.unmarshal(inst)
	return
}

// Returns the initial credentials for a Windows instance.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InstanceCredentials/
func (c *Client) GetWindowsInstanceInitialCredentials(instanceId string) (creds *InstanceCredentials, e error) {
	details := &requestDetails{
		name: resourceInstances,
		ids:  urlParts{instanceId, "initialCredentials"},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	creds = &InstanceCredentials{}
	e = resp.unmarshal(creds)
	return
}
