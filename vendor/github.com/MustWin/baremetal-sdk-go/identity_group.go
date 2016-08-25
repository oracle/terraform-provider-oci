package baremetal

// CreateGroup create a new group. groupName MUST be supplied and MUST be
// unique. groupDescription is optional. You MAY supply one option with an
// idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createGroup
func (c *Client) CreateGroup(groupName, groupDescription string, options ...Options) (response *IdentityResource, e error) {
	body := CreateIdentityResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          groupName,
		Description:   groupDescription,
	}

	return c.createIdentityResource(resourceGroups, body, options)
}

// DeleteGroup removes a group identified by groupID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteGroup
func (c *Client) DeleteGroup(groupID string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceGroups,
		options: opts,
		ids:     urlParts{groupID},
	}
	return c.identityApi.deleteRequest(reqOpts)
}

// GetGroup returns a group identified by groupID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getGroup
func (c *Client) GetGroup(groupID string) (group *IdentityResource, e error) {
	group, e = c.getIdentityResource(resourceGroups, groupID)
	return
}

// ListGroups returns a list of Groups in a tenancy. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listGroups
func (c *Client) ListGroups(options ...Options) (response *ListResourceResponses, e error) {
	return c.listIdentityResources(resourceGroups, options...)
}

// UpdateGroup updates the description of a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updateGroup
func (c *Client) UpdateGroup(groupID, description string, options ...Options) (group *IdentityResource, e error) {
	body := UpdateIdentityResourceRequest{
		Description: description,
	}

	return c.updateIdentityResource(resourceGroups, groupID, body, options)
}
