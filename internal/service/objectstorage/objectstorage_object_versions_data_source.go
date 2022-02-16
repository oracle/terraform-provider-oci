// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"strconv"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"
)

func ObjectStorageObjectVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStorageObjectVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"delimiter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"end": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"archival_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_delete_marker": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"md5": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_tier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readObjectStorageObjectVersions(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStorageObjectVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListObjectVersionsResponse
}

func (s *ObjectStorageObjectVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageObjectVersionsDataSourceCrud) Get() error {
	request := oci_object_storage.ListObjectVersionsRequest{
		Fields: oci_object_storage.ListObjectVersionsFieldsEnum(listObjectsFields),
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if delimiter, ok := s.D.GetOkExists("delimiter"); ok {
		tmp := delimiter.(string)
		request.Delimiter = &tmp
	}

	if end, ok := s.D.GetOkExists("end"); ok {
		tmp := end.(string)
		request.End = &tmp
	}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		request.Fields = oci_object_storage.ListObjectVersionsFieldsEnum(fields.(string))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if prefix, ok := s.D.GetOkExists("prefix"); ok {
		tmp := prefix.(string)
		request.Prefix = &tmp
	}

	if start, ok := s.D.GetOkExists("start"); ok {
		tmp := start.(string)
		request.Start = &tmp
	}

	if startAfter, ok := s.D.GetOkExists("start_after"); ok {
		tmp := startAfter.(string)
		request.StartAfter = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

	response, err := s.Client.ListObjectVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectStorageObjectVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ObjectStorageObjectVersionsDataSource-", ObjectStorageObjectVersionsDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ObjectVersionSummaryToMap(item))
	}
	s.D.Set("items", items)

	s.D.Set("prefixes", s.Res.Prefixes)

	return nil
}

func ObjectVersionSummaryToMap(obj oci_object_storage.ObjectVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["archival_state"] = string(obj.ArchivalState)

	if obj.Etag != nil {
		result["etag"] = string(*obj.Etag)
	}

	if obj.IsDeleteMarker != nil {
		result["is_delete_marker"] = bool(*obj.IsDeleteMarker)
	}

	if obj.Md5 != nil {
		result["md5"] = string(*obj.Md5)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Size != nil {
		result["size"] = strconv.FormatInt(*obj.Size, 10)
	}

	result["storage_tier"] = string(obj.StorageTier)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeModified != nil {
		result["time_modified"] = obj.TimeModified.Format(time.RFC3339Nano)
	}

	if obj.VersionId != nil {
		result["version_id"] = string(*obj.VersionId)
	}

	return result
}
