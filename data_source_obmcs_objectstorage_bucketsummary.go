// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
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
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func readBucketSummaries(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	reader := &BucketSummaryDatasourceCrud{}
	reader.D = d
	reader.Client = client.client
	return crud.ReadResource(reader)
}

type BucketSummaryDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListBuckets
}

func (s *BucketSummaryDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	namespace := s.D.Get("namespace").(string)

	opts := &baremetal.ListBucketsOptions{}
	if page, ok := s.D.GetOk("page"); ok {
		opts.Page = page.(string)
	}

	if limit, ok := s.D.GetOk("limit"); ok {
		opts.ListOptions.Limit = uint64(limit.(int))
	}

	s.Res = &baremetal.ListBuckets{BucketSummaries: []baremetal.BucketSummary{}}

	for {
		var list *baremetal.ListBuckets
		if list, e = s.Client.ListBuckets(compartmentID, baremetal.Namespace(namespace), opts); e != nil {
			break
		}

		s.Res.BucketSummaries = append(s.Res.BucketSummaries, list.BucketSummaries...)
		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}
	return
}

func (s *BucketSummaryDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, r := range s.Res.BucketSummaries {
			res := map[string]interface{}{
				"namespace":      r.Namespace,
				"name":           r.Name,
				"compartment_id": r.CompartmentID,
				"created_by":     r.CreatedBy,
				"time_created":   r.TimeCreated.String(),
				"etag":           r.ETag,
			}
			resources = append(resources, res)
		}
		s.D.Set("bucket_summaries", resources)
	}
	return
}
