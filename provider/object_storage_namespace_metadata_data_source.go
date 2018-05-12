// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func NamespaceMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNamespaceMetadata2,
		Schema: map[string]*schema.Schema{
			// Required
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"default_s3compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_swift_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readNamespaceMetadata2(d *schema.ResourceData, m interface{}) error {
	sync := &NamespaceMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
}

type NamespaceMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetNamespaceMetadataResponse
}

func (s *NamespaceMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NamespaceMetadataDataSourceCrud) Get() error {
	request := oci_object_storage.GetNamespaceMetadataRequest{}

	tmp := s.D.Get("namespace").(string)
	request.NamespaceName = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

	response, err := s.Client.GetNamespaceMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NamespaceMetadataDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.DefaultS3CompartmentId != nil {
		s.D.Set("default_s3compartment_id", *s.Res.DefaultS3CompartmentId)
	}

	if s.Res.DefaultSwiftCompartmentId != nil {
		s.D.Set("default_swift_compartment_id", *s.Res.DefaultSwiftCompartmentId)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	return
}
