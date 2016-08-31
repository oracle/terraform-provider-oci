package baremetal

// CreateCompartment create a new compartment.

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/CreateCompartment
func (c *Client) CreateCompartment(compartmentName, compartmentDescription string, options ...Options) (compartment *IdentityResource, e error) {
	body := CreateIdentityResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          compartmentName,
		Description:   compartmentDescription,
	}

	return c.createIdentityResource(resourceCompartments, body, options)
}

// GetCompartment returns the compartment identified by compartmentID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/GetCompartment
func (c *Client) GetCompartment(compartmentID string) (compartment *IdentityResource, e error) {
	compartment, e = c.getIdentityResource(resourceCompartments, compartmentID)
	return
}

// ListCompartments returns a list of compartments. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/ListCompartments
func (c *Client) ListCompartments(options ...Options) (response *ListResourceResponses, e error) {
	return c.listIdentityResources(resourceCompartments, options...)
}

// UpdateCompartment updates the description of a compartment.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Compartment/UpdateCompartment
func (c *Client) UpdateCompartment(compartmentID, description string, options ...Options) (compartment *IdentityResource, e error) {
	body := UpdateIdentityResourceRequest{
		Description: description,
	}

	return c.updateIdentityResource(resourceCompartments, compartmentID, body, options)
}
