// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"
)

func ObjectStorageReplicationSourcesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStorageReplicationSources,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_sources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"policy_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_bucket_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_region_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readObjectStorageReplicationSources(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageReplicationSourcesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStorageReplicationSourcesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListReplicationSourcesResponse
}

func (s *ObjectStorageReplicationSourcesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageReplicationSourcesDataSourceCrud) Get() error {
	request := oci_object_storage.ListReplicationSourcesRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

	response, err := s.Client.ListReplicationSources(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReplicationSources(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ObjectStorageReplicationSourcesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ObjectStorageReplicationSourcesDataSource-", ObjectStorageReplicationSourcesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		replicationSource := map[string]interface{}{}

		if r.PolicyName != nil {
			replicationSource["policy_name"] = *r.PolicyName
		}

		if r.SourceBucketName != nil {
			replicationSource["source_bucket_name"] = *r.SourceBucketName
		}

		if r.SourceRegionName != nil {
			replicationSource["source_region_name"] = *r.SourceRegionName
		}

		resources = append(resources, replicationSource)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, ObjectStorageReplicationSourcesDataSource().Schema["replication_sources"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("replication_sources", resources); err != nil {
		return err
	}

	return nil
}
