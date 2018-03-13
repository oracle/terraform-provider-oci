// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/oracle/terraform-provider-oci/crud"

	"log"
)

func ObjectHeadDataSource() *schema.Resource {
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
			"content-length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"content-type": {
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
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.HeadObject(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response

	return nil
}

func (s *ObjectHeadDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.OpcMeta != nil {
		if err := s.D.Set("metadata", s.Res.OpcMeta); err != nil {
			log.Printf("Unable to set `metadata`. Error %q", err)
		}
	}

	if s.Res.ContentLength != nil {
		s.D.Set("content-length", *s.Res.ContentLength)
	}

	if s.Res.ContentType != nil {
		s.D.Set("content-type", *s.Res.ContentType)
	}

	return
}
