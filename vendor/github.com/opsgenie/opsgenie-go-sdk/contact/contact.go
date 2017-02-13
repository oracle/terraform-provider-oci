/*
Copyright 2016. All rights reserved.
Use of this source code is governed by a Apache Software
license that can be found in the LICENSE file.
*/

//Package contact provides requests and response structures to achieve Contact API actions.
package contact

// CreateContactRequest provides necessary parameter structure for creating contact
type CreateContactRequest struct {
	APIKey 		string `json:"apiKey,omitempty"`
	Method 		string `json:"method,omitempty"`
	To 		string `json:"to,omitempty"`
	Username 	string `json:"username,omitempty"`
}

// UpdateContactRequest provides necessary parameter structure for updating a contact
type UpdateContactRequest struct {
	APIKey 		string `json:"apiKey,omitempty"`
	Id     		string `json:"id,omitempty"`
	To 		string `json:"to,omitempty"`
	Username 	string `json:"username,omitempty"`
}

// DeleteContactRequest provides necessary parameter structure for deleting a contact
type DeleteContactRequest struct {
	APIKey 		string `url:"apiKey,omitempty"`
	Id     		string `url:"id,omitempty"`
	Username 	string `url:"username,omitempty"`
}

// GetContactRequest provides necessary parameter structure for requesting contact information
type GetContactRequest struct {
	APIKey 		string `url:"apiKey,omitempty"`
	Id     		string `url:"id,omitempty"`
	Username 	string `url:"username,omitempty"`
}