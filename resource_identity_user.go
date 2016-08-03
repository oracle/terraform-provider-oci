package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceIdentityUser exposes a IdentityUser Resource
func ResourceIdentityUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceIdentityUserCreate,
		Read:   resourceIdentityUserRead,
		Update: resourceIdentityUserUpdate,
		Delete: resourceIdentityUserDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIdentityUserCreate(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	var user *baremtlsdk.Resource
	if user, e = client.CreateUser(name, description); e != nil {
		return
	}

	// Set the id and set any fields that were returned by the API.
	d.SetId(user.ID)
	setResourceData(d, user)

	if user.State != baremtlsdk.ResourceCreated {
		user, e = waitForUserStateRefresh(d, client)
	}

	return
}

func resourceIdentityUserRead(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	var user *baremtlsdk.Resource
	if user, e = client.GetUser(d.Id()); e != nil {
		return
	}

	setResourceData(d, user)

	return
}

func resourceIdentityUserUpdate(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	desc := d.Get("description").(string)

	d.Partial(true)
	var user *baremtlsdk.Resource
	if user, e = client.UpdateUser(d.Id(), desc); e != nil {
		return
	}
	d.Partial(false)

	// Capture any upstream changes, like time_modified.
	setResourceData(d, user)

	return
}

func resourceIdentityUserDelete(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	if e = client.DeleteUser(d.Id()); e != nil {
		return
	}

	return
}

func userStateRefreshFunc(client BareMetalClient, d *schema.ResourceData) resource.StateRefreshFunc {
	return func() (user interface{}, s string, e error) {
		if user, e = client.GetUser(d.Id()); e != nil {
			return nil, "", e
		}
		s = user.(*baremtlsdk.Resource).State
		return
	}
}

func setResourceData(d *schema.ResourceData, user *baremtlsdk.Resource) {
	d.Set("name", user.Name)
	d.Set("description", user.Description)
	d.Set("compartment_id", user.CompartmentID)
	d.Set("state", user.State)
	d.Set("time_modified", user.TimeModified.String())
	d.Set("time_created", user.TimeCreated.String())
}

func waitForUserStateRefresh(d *schema.ResourceData, c BareMetalClient) (user *baremtlsdk.Resource, e error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{baremtlsdk.ResourceCreating},
		Target:  []string{baremtlsdk.ResourceCreated},
		Refresh: userStateRefreshFunc(c, d),
		Timeout: 5 * time.Minute,
	}

	rawu, err := stateConf.WaitForState()
	user = rawu.(*baremtlsdk.Resource)
	if e = err; e != nil {
		return
	}

	// Fields may have changed during polling, set them again.
	setResourceData(d, user)

	return
}
