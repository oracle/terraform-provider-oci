// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v29/objectstorage"
)

func init() {
	RegisterDatasource("oci_objectstorage_replication_policies", ObjectStorageReplicationPoliciesDataSource())
}

func ObjectStorageReplicationPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStorageReplicationPolicies,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(ObjectStorageReplicationPolicyResource()),
			},
		},
	}
}

func readObjectStorageReplicationPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageReplicationPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient()

	return ReadResource(sync)
}

type ObjectStorageReplicationPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListReplicationPoliciesResponse
}

func (s *ObjectStorageReplicationPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageReplicationPoliciesDataSourceCrud) Get() error {
	request := oci_object_storage.ListReplicationPoliciesRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.ListReplicationPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReplicationPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ObjectStorageReplicationPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ObjectStorageReplicationPoliciesDataSource-", ObjectStorageReplicationPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		replicationPolicy := map[string]interface{}{}

		if r.DestinationBucketName != nil {
			replicationPolicy["destination_bucket_name"] = *r.DestinationBucketName
		}

		if r.DestinationRegionName != nil {
			replicationPolicy["destination_region_name"] = *r.DestinationRegionName
		}

		if r.Id != nil {
			replicationPolicy["id"] = *r.Id
		}

		if r.Name != nil {
			replicationPolicy["name"] = *r.Name
		}

		replicationPolicy["status"] = r.Status

		if r.StatusMessage != nil {
			replicationPolicy["status_message"] = *r.StatusMessage
		}

		if r.TimeCreated != nil {
			replicationPolicy["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastSync != nil {
			replicationPolicy["time_last_sync"] = r.TimeLastSync.String()
		}

		resources = append(resources, replicationPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ObjectStorageReplicationPoliciesDataSource().Schema["replication_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("replication_policies", resources); err != nil {
		return err
	}

	return nil
}
