package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceCoreShape() *schema.Resource {
	return &schema.Resource{
		Read:   readShape,
		Schema: map[string]*schema.Schema{},
	}
}

func readShape(d *schema.ResourceData, m interface{}) (e error) {
	return
}
