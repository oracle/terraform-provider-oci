package main

import "github.com/hashicorp/terraform/helper/schema"

func ResourceCoreInstances() *schema.Resource {
	return &schema.Resource{
		Read: readInstances,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"instances": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     resourceCoreInstance(),
			},
		},
	}
}

func resourceCoreInstance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeMap,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": &schema.Schema{
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
		},
	}
}

func readInstances(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(BareMetalClient)
	reader := &InstancesSync{
		D:      d,
		Client: client,
	}

	return readResource(reader)
}
