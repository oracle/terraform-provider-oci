// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v65/objectstorage"

	"log"
)

func ObjectStorageObjectHeadDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readObjectHead,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"content_type": {
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
			"metadata": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func readObjectHead(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectHeadDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	return tfresource.ReadResource(sync)
}

type ObjectHeadDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.HeadObjectResponse
}

func (s *ObjectHeadDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectHeadDataSourceCrud) Get() error {
	request := oci_object_storage.HeadObjectRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

	response, err := s.Client.HeadObject(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response

	return nil
}

func (s *ObjectHeadDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceID())

	if s.Res.OpcMeta != nil {
		if err := s.D.Set("metadata", s.Res.OpcMeta); err != nil {
			log.Printf("Unable to set `metadata`. Error %q", err)
		}
	}

	if s.Res.ContentLength != nil {
		s.D.Set("content_length", *s.Res.ContentLength)
	}

	if s.Res.ContentType != nil {
		s.D.Set("content_type", *s.Res.ContentType)
	}

	if s.Res.ETag != nil {
		s.D.Set("etag", *s.Res.ETag)
	}

	s.D.Set("storage_tier", string(s.Res.StorageTier))

	s.D.Set("archival_state", string(s.Res.ArchivalState))

	return nil
}
