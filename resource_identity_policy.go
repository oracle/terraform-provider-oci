package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const fiveMinutes = 5 * time.Minute

var policySchema = map[string]*schema.Schema{
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
	"statements": &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	},
}

func ResourceIdentityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: createPolicy,
		Read:   readPolicy,
		Update: updatePolicy,
		Delete: deletePolicy,
		Schema: policySchema,
	}
}

func createPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	statements := d.Get("statements").([]string)

	var policy *baremtlsdk.Policy
	if policy, e = client.CreatePolicy(name, description, statements); e != nil {
		return
	}

	d.SetId(policy.ID)
	setPolicyData(d, policy)

	if policy.State != baremtlsdk.ResourceCreated {
		monitor := &resource.StateChangeConf{
			Pending: []string{baremtlsdk.ResourceCreating},
			Target:  []string{baremtlsdk.ResourceCreated},
			Refresh: func() (p interface{}, state string, e error) {
				if p, e = client.GetPolicy(policy.ID); e != nil {
					return
				}
				state = p.(*baremtlsdk.Policy).State
				return
			},
			Timeout: fiveMinutes,
		}

		var res interface{}
		if res, e = monitor.WaitForState(); e != nil {
			return
		}

		policy = res.(*baremtlsdk.Policy)
		setPolicyData(d, policy)
	}

	return
}

func setPolicyData(d *schema.ResourceData, policy *baremtlsdk.Policy) {
	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("compartment_id", policy.CompartmentID)
	d.Set("state", policy.State)
	d.Set("time_modified", policy.TimeModified.String())
	d.Set("time_created", policy.TimeCreated.String())
	d.Set("statements", policy.Statements)
}

func readPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)

	var policy *baremtlsdk.Policy
	if policy, e = client.GetPolicy(d.Id()); e != nil {
		return
	}

	setPolicyData(d, policy)

	return
}

func updatePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	description := d.Get("description").(string)
	statements := d.Get("statements").([]string)
	d.Partial(true)

	var policy *baremtlsdk.Policy
	if policy, e = client.UpdatePolicy(d.Id(), description, statements); e != nil {
		return
	}

	d.Partial(false)
	setPolicyData(d, policy)
	return
}

func deletePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	destroyResource(d, client.DeletePolicy)
	return
}
