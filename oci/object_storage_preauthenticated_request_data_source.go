// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func ObjectStoragePreauthenticatedRequestDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularObjectStoragePreauthenticatedRequest,
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"par_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"access_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_expires": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularObjectStoragePreauthenticatedRequest(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStoragePreauthenticatedRequestDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

type ObjectStoragePreauthenticatedRequestDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetPreauthenticatedRequestResponse
}

func (s *ObjectStoragePreauthenticatedRequestDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStoragePreauthenticatedRequestDataSourceCrud) Get() error {
	request := oci_object_storage.GetPreauthenticatedRequestRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if parId, ok := s.D.GetOkExists("par_id"); ok {
		tmp := parId.(string)
		request.ParId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.GetPreauthenticatedRequest(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectStoragePreauthenticatedRequestDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("access_type", s.Res.AccessType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpires != nil {
		s.D.Set("time_expires", s.Res.TimeExpires.String())
	}

	return nil
}
