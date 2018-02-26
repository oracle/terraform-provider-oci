// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BucketsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBuckets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"page": {
				Type:       schema.TypeString,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("page"),
			},
			"bucket_summaries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     BucketResource(),
			},
		},
	}
}

func readBuckets(d *schema.ResourceData, m interface{}) error {
	sync := &BucketsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
}

type BucketsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListBucketsResponse
}

func (s *BucketsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BucketsDataSourceCrud) Get() error {
	request := oci_object_storage.ListBucketsRequest{}

	if compartmentId, ok := s.D.GetOk("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOk("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	if namespace, ok := s.D.GetOk("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if page, ok := s.D.GetOk("page"); ok {
		tmp := page.(string)
		request.Page = &tmp
	}

	response, err := s.Client.ListBuckets(context.Background(), request, getRetryOptions(false, "object_storage")...)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBuckets(context.Background(), request, getRetryOptions(false, "object_storage")...)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BucketsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bucket := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			bucket["created_by"] = *r.CreatedBy
		}

		if r.Etag != nil {
			bucket["etag"] = *r.Etag
		}

		if r.Name != nil {
			bucket["name"] = *r.Name
		}

		if r.Namespace != nil {
			bucket["namespace"] = *r.Namespace
		}

		bucket["time_created"] = r.TimeCreated.String()

		resources = append(resources, bucket)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("bucket_summaries", resources); err != nil {
		panic(err)
	}

	return
}
