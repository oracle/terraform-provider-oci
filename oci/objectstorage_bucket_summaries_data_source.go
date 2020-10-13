// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v27/objectstorage"
)

func init() {
	RegisterDatasource("oci_objectstorage_bucket_summaries", ObjectStorageBucketsDataSource())
}

func ObjectStorageBucketsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStorageBuckets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
				Elem:     GetDataSourceItemSchema(ObjectStorageBucketResource()),
			},
		},
	}
}

func readObjectStorageBuckets(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageBucketsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient()

	return ReadResource(sync)
}

type ObjectStorageBucketsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListBucketsResponse
}

func (s *ObjectStorageBucketsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageBucketsDataSourceCrud) Get() error {
	request := oci_object_storage.ListBucketsRequest{
		// @CODEGEN 6/2018: Need to specify all the fields we want from the BucketSummaries
		Fields: oci_object_storage.GetListBucketsFieldsEnumValues(),
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.ListBuckets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBuckets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ObjectStorageBucketsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		bucket := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
			"namespace":      *r.Namespace,
		}

		if r.CreatedBy != nil {
			bucket["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			bucket["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.Etag != nil {
			bucket["etag"] = *r.Etag
		}

		bucket["freeform_tags"] = r.FreeformTags

		if r.Name != nil {
			bucket["name"] = *r.Name
		}

		if r.TimeCreated != nil {
			bucket["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, bucket)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ObjectStorageBucketsDataSource().Schema["bucket_summaries"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("bucket_summaries", resources); err != nil {
		return err
	}

	return nil
}
