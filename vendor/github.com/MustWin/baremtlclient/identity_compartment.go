package baremtlsdk

// CreateCompartment create a new compartment.
import (
	"encoding/json"
	"net/http"
)

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createCompartment
func (c *Client) CreateCompartment(compartmentName, compartmentDescription string, options ...Options) (compartment *IdentityResource, e error) {
	createRequest := CreateIdentityResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          compartmentName,
		Description:   compartmentDescription,
	}
	var headers http.Header
	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}

	return c.createIdentityResource(resourceCompartments, createRequest, headers)
}

// GetCompartment returns the compartment identified by compartmentID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetCompartment(compartmentID string) (compartment *IdentityResource, e error) {
	compartment, e = c.getIdentityResource(resourceCompartments, compartmentID)
	return
}

// ListCompartments returns a list of compartments. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listCompartments
func (c *Client) ListCompartments(options ...ListOptions) (response *ListResourceResponse, e error) {
	return c.listIdentityResources(resourceCompartments, options...)
}

// UpdateCompartment updates the description of a compartment.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updateCompartment
func (c *Client) UpdateCompartment(compartmentID, description string, options ...Options) (compartment *IdentityResource, e error) {

	headers := getUpdateHeaders(options...)

	request := UpdateIdentityResourceRequest{
		Description: description,
	}

	var resp []byte
	if resp, e = c.updateIdentityResource(resourceCompartments, compartmentID, request, headers); e != nil {
		return
	}

	compartment = &IdentityResource{}
	e = json.Unmarshal(resp, compartment)
	return

}
