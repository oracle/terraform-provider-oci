// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v35/objectstorage"
)

func init() {
	RegisterDatasource("oci_objectstorage_preauthrequests", ObjectStoragePreauthenticatedRequestsDataSource())
}

func ObjectStoragePreauthenticatedRequestsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStoragePreauthenticatedRequests,
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
			"object_name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"preauthenticated_requests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(ObjectStoragePreauthenticatedRequestResource()),
			},
		},
	}
}

func readObjectStoragePreauthenticatedRequests(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePreauthenticatedRequestsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient()

	return ReadResource(sync)
}

type ObjectStoragePreauthenticatedRequestsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListPreauthenticatedRequestsResponse
}

func (s *ObjectStoragePreauthenticatedRequestsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStoragePreauthenticatedRequestsDataSourceCrud) Get() error {
	request := oci_object_storage.ListPreauthenticatedRequestsRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if objectNamePrefix, ok := s.D.GetOkExists("object_name_prefix"); ok {
		tmp := objectNamePrefix.(string)
		request.ObjectNamePrefix = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.ListPreauthenticatedRequests(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPreauthenticatedRequests(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ObjectStoragePreauthenticatedRequestsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ObjectStoragePreauthenticatedRequestsDataSource-", ObjectStoragePreauthenticatedRequestsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		preauthenticatedRequest := map[string]interface{}{}

		preauthenticatedRequest["access_type"] = r.AccessType

		if r.Id != nil {
			preauthenticatedRequest["id"] = *r.Id
		}

		if r.Name != nil {
			preauthenticatedRequest["name"] = *r.Name
		}

		if r.ObjectName != nil {
			preauthenticatedRequest["object"] = *r.ObjectName
		}

		if r.TimeCreated != nil {
			preauthenticatedRequest["time_created"] = r.TimeCreated.String()
		}

		if r.TimeExpires != nil {
			preauthenticatedRequest["time_expires"] = r.TimeExpires.String()
		}

		resources = append(resources, preauthenticatedRequest)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, ObjectStoragePreauthenticatedRequestsDataSource().Schema["preauthenticated_requests"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("preauthenticated_requests", resources); err != nil {
		return err
	}

	return nil
}
