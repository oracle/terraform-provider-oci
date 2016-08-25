package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func PolicyResource() *schema.Resource {
	policySchema := make(map[string]*schema.Schema)

	for key, value := range identitySchema {
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
		Update: updatePolicy,
		Delete: deletePolicy,
		Schema: policySchema,
	}
}

func createPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &PolicyResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &PolicyResourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}

func updatePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &PolicyResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, sync)
}

func deletePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &PolicyResourceCrud{D: d, Client: client}
	return sync.Delete()
}
