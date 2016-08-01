package main

import "github.com/hashicorp/terraform/helper/schema"

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

func resourceIdentityUserCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(BareMetalClient)

	name := d.Get("name").(string)
	description := d.Get("description").(string)

	user, err := client.CreateUser(name, description)

	d.SetId(user.ID)

	resourceIdentityUserRead(d, m)

	return err
}

func resourceIdentityUserRead(d *schema.ResourceData, m interface{}) error {
	client := m.(BareMetalClient)

	user, _ := client.GetUser(d.Id())

	d.Set("name", user.Name)
	d.Set("compartment_id", user.CompartmentID)
	d.Set("state", user.State)
	d.Set("time_modified", user.TimeModified.String())
	d.Set("time_created", user.TimeCreated.String())

	return nil
}

func resourceIdentityUserUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(BareMetalClient)

	client.UpdateUser(d.Get("description").(string))

	return nil
}

func resourceIdentityUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
