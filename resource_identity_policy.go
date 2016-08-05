package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceIdentityPolicy() *schema.Resource {
	policySchema := make(map[string]*schema.Schema)

	for key, value := range resourceSchema {
		policySchema[key] = value
	}

	policySchema["statements"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}

	return &schema.Resource{
		Create: createPolicy,
		Read:   readPolicy,
		Update: updateResource,
		Delete: deletePolicy,
		Schema: policySchema,
	}
}

func createPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &PolicySync{d, client}
	return createResource(d, sync)
}

func readPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &PolicySync{d, client}
	return readResource(sync)
}

func updatePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &PolicySync{d, client}
	return updateResource(d, sync)
}

func deletePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	sync := &PolicySync{d, client}
	return sync.Delete()
}
