package baremetal

// CreateUser is used to create a user. userName MUST be unique. description

// contains a comment about the user. The caller can supply 0 or 1 options. Options
// MAY contain an idempotency token.
// The caller specifies this token so that subsequent calls to create user will
// be idempotent. The token expires after 30 minutes.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createUser
func (c *Client) CreateUser(userName, userDescription string, options ...Options) (user *IdentityResource, e error) {
	body := CreateIdentityResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          userName,
		Description:   userDescription,
	}

	return c.createIdentityResource(resourceUsers, body, options)
}

// DeleteUser deletes a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteUser
func (c *Client) DeleteUser(userID string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceUsers,
		options: opts,
		ids:     urlParts{userID},
	}
	return c.identityApi.deleteRequest(reqOpts)
}

// GetUser returns a user identified by userID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetUser(userID string) (user *IdentityResource, e error) {
	user, e = c.getIdentityResource(resourceUsers, userID)
	return
}

// ListUsers returns an array of users for the current tenancy.  The requestor
// MAY supply paging options.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
func (c *Client) ListUsers(options ...Options) (response *ListResourceResponse, e error) {
	return c.listIdentityResources(resourceUsers, options...)
}

func (c *Client) UpdateUser(userID, userDescription string, options ...Options) (user *IdentityResource, e error) {
	body := UpdateIdentityResourceRequest{
		Description: userDescription,
	}

	return c.updateIdentityResource(resourceUsers, userID, body, options)
}
