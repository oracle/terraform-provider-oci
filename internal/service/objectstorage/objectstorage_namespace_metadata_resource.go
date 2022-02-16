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

func ObjectStorageNamespaceMetadataResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: tfresource.DefaultTimeout,
		Create:   readNamespaceMetadata,
		Read:     readNamespaceMetadata,
		Update:   updateNamespaceMetadata,
		Delete:   deleteNamespaceMetadata,
		Schema: map[string]*schema.Schema{
			// Required
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"default_s3compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_swift_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			// Computed
		},
	}
}

func readNamespaceMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &NamespaceMetadataResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	return tfresource.ReadResource(sync)
}

func updateNamespaceMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &NamespaceMetadataResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteNamespaceMetadata(d *schema.ResourceData, m interface{}) error {
	return nil
}

type NamespaceMetadataResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *oci_object_storage.NamespaceMetadata
	DisableNotFoundRetries bool
}

func (s *NamespaceMetadataResourceCrud) ID() string {
	return *s.Res.Namespace
}

func (s *NamespaceMetadataResourceCrud) Get() error {
	request := oci_object_storage.GetNamespaceMetadataRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetNamespaceMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.D.SetId(*response.NamespaceMetadata.Namespace)
	s.Res = &response.NamespaceMetadata
	return nil
}

func (s *NamespaceMetadataResourceCrud) Update() error {
	request := oci_object_storage.UpdateNamespaceMetadataRequest{}

	if defaultS3CompartmentId, ok := s.D.GetOkExists("default_s3compartment_id"); ok {
		tmp := defaultS3CompartmentId.(string)
		request.DefaultS3CompartmentId = &tmp
	}

	if defaultSwiftCompartmentId, ok := s.D.GetOkExists("default_swift_compartment_id"); ok {
		tmp := defaultSwiftCompartmentId.(string)
		request.DefaultSwiftCompartmentId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.UpdateNamespaceMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NamespaceMetadata
	return nil
}

func (s *NamespaceMetadataResourceCrud) SetData() error {
	if s.Res.DefaultS3CompartmentId != nil {
		s.D.Set("default_s3compartment_id", *s.Res.DefaultS3CompartmentId)
	}

	if s.Res.DefaultSwiftCompartmentId != nil {
		s.D.Set("default_swift_compartment_id", *s.Res.DefaultSwiftCompartmentId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	return nil
}
