// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v56/objectstorage"
)

func ObjectStorageObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectStorageObjects,
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
			// @CODEGEN 2/2018: The result type is defined and generated from spec is a ListObject.
			// But the actual result we want to return is the array of ObjectSummary objects within
			// the ListObject response. Manually define this for now.
			"objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						// @CODEGEN 2/2018: The spec says 'size' is an int64 but it's being treated as a
						// string in existing provider. Make it string to avoid breaking change.
						// @CODEGEN 8/2018: The codegen now honors int64, but till HCL2 is released, we will continue to
						// treat it as a string to accommodate larger values. HCL2TODO: This can be plainly changed to
						// the new type when HCL2 is released
						"size": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"md5": {
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
						"etag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_tier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"archival_state": {
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

func readObjectStorageObjects(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStorageObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListObjects
}

func (s *ObjectStorageObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

var listObjectsFields = "name,size,md5,timeCreated,timeModified,etag,storageTier,archivalState"

func (s *ObjectStorageObjectsDataSourceCrud) Get() error {
	request := oci_object_storage.ListObjectsRequest{
		// @CODEGEN 2/2018: Need to specify all the fields we want from the ObjectSummaries
		Fields: &listObjectsFields,
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
		tmp := fields.(string)
		request.Fields = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
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

	// @CODEGEN 2/2018: Preserve the custom logic to extract the ObjectSummary results from ListObjects response
	// and to handle pagination.
	for {
		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

		response, err := s.Client.ListObjects(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res = &oci_object_storage.ListObjects{Objects: []oci_object_storage.ObjectSummary{}}
		for _, objectSummary := range response.Objects {
			s.Res.Objects = append(s.Res.Objects, objectSummary)
		}

		if response.NextStartWith == nil || *response.NextStartWith == "" {
			break
		}

		request.Start = response.NextStartWith
	}
	return nil
}

func (s *ObjectStorageObjectsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ObjectStorageObjectsDataSource-", ObjectStorageObjectsDataSource(), s.D))

	// @CODGEN 2/2018: We generate a call to set 'next_start_with' field. It's not
	// necessary to store it, as it's a pagination token that's handled in the Get() call.

	objects := []map[string]interface{}{}
	for _, item := range s.Res.Objects {
		objects = append(objects, ObjectSummaryToMap(item))
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		objects = tfresource.ApplyFilters(f.(*schema.Set), objects, ObjectStorageObjectsDataSource().Schema["objects"].Elem.(*schema.Resource).Schema)
	}

	s.D.Set("objects", objects)

	s.D.Set("prefixes", s.Res.Prefixes)

	return nil
}
