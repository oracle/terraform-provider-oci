package main

import (
	"fmt"
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
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIdentityUserCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name")
	if name == nil {
		return fmt.Errorf("Name cannot be empty")
	}

	description := d.Get("description")
	if description == nil {
		return fmt.Errorf("Description cannot be empty")
	}
	return nil
}

func resourceIdentityUserRead(d *schema.ResourceData, m interface{}) error {
	userID := d.Get("user_id")
	if userID == nil {
		return fmt.Errorf("user_id cannot be empty")
	}
	return nil
}

func resourceIdentityUserUpdate(d *schema.ResourceData, m interface{}) error {
	userID := d.Get("user_id")
	if userID == nil {
		return fmt.Errorf("user_id cannot be empty")
	}
	_ = d.Get("name")
	_ = d.Get("description")

	return nil
}

func resourceIdentityUserDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
