// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ObjectsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjects,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			// @CODEGEN 2/2018: 'fields' field omitted from existing provider.
			"limit": {
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: crud.FieldDeprecated("limit"),
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
					},
				},
			},
		},
	}
}

func readObjects(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
}

type ObjectsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.ListObjects
}

func (s *ObjectsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

var listObjectsFields = "name,size,md5,timeCreated"

func (s *ObjectsDataSourceCrud) Get() error {
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

	// @CODEGEN 2/2018: Preserve the custom logic to extract the ObjectSummary results from ListObjects response
	// and to handle pagination.
	for {
		request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

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

func (s *ObjectsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	// @CODGEN 2/2018: We generate a call to set 'next_start_with' field. It's not
	// necessary to store it, as it's a pagination token that's handled in the Get() call.

	objects := []map[string]interface{}{}
	for _, item := range s.Res.Objects {
		objects = append(objects, ObjectSummaryToMap(item))
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		objects = ApplyFilters(f.(*schema.Set), objects, ObjectsDataSource().Schema["objects"].Elem.(*schema.Resource).Schema)
	}

	s.D.Set("objects", objects)

	s.D.Set("prefixes", s.Res.Prefixes)

	return
}
