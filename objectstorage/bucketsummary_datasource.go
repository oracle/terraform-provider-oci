// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package objectstorage

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func BucketSummaryDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readBucketSummaries,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket_summaries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"limit": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func readBucketSummaries(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &BucketSummaryDatasourceCrud{}
	reader.D = d
	reader.Client = client
	return crud.ReadResource(reader)
}
