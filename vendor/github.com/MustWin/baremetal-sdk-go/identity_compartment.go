package baremetal

import (
	"net/http"
	"time"
)

type Compartment struct {
	ETaggedResource
	CompartmentID  string    `json:"compartmentId"`
	Description    string    `json:"description"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	Name           string    `json:"name"`
	State          string    `json:"lifecycleState"`
	TimeCreated    time.Time `json:"timeCreated"`
}

type ListCompartments struct {
	ResourceContainer
	Compartments []Compartment
}

func (l *ListCompartments) GetList() interface{} {
	return &l.Compartments
}

// CreateCompartment create a new compartment.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/CreateCompartment
func (c *Client) CreateCompartment(name, desc string, opts *RetryTokenOptions) (res *Compartment, e error) {
	required := identityCreationRequirement{
		CompartmentID: c.authInfo.tenancyOCID,
		Description:   desc,
		Name:          name,
	}

	details := &requestDetails{
		name:     resourceCompartments,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &Compartment{}
	e = response.unmarshal(res)
	return
}

// GetCompartment returns the compartment identified by compartmentID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/GetCompartment
func (c *Client) GetCompartment(id string) (res *Compartment, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceCompartments,
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &Compartment{}
	e = response.unmarshal(res)
	return
}

// UpdateCompartment updates the description of a compartment.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/UpdateCompartment
func (c *Client) UpdateCompartment(id string, opts *UpdateIdentityOptions) (res *Compartment, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceCompartments,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &Compartment{}
	e = response.unmarshal(res)
	return
}

// ListCompartments returns a list of compartments. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/ListCompartments
func (c *Client) ListCompartments(opts *ListOptions) (resources *ListCompartments, e error) {
	details := &requestDetails{
		name:     resourceCompartments,
		optional: opts,
		required: ocidRequirement{c.authInfo.tenancyOCID},
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListCompartments{}
	e = getResp.unmarshal(resources)
	return
}
