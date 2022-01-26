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

func ObjectStorageNamespaceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularObjectStorageNamespace,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			// @CODEGEN 2/2018: No computed values are generated because the generator doesn't know what to do with
			// responses that aren't reference types.
			// In this case, the response from service is a string so manually add a field for it.
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularObjectStorageNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageNamespaceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStorageNamespaceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetNamespaceResponse
}

func (s *ObjectStorageNamespaceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageNamespaceDataSourceCrud) Get() error {
	request := oci_object_storage.GetNamespaceRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

	response, err := s.Client.GetNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectStorageNamespaceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ObjectStorageNamespaceDataSource-", ObjectStorageNamespaceDataSource(), s.D))

	if s.Res.Value != nil {
		s.D.Set("namespace", *s.Res.Value)
	}

	return nil
}
