// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"

	"github.com/oracle/terraform-provider-oci/crud"
)

func NamespacesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNamespaces,
		Schema: map[string]*schema.Schema{
			// @CODEGEN: Removed the filter field. Generator recognizes this as a multiple datasource, but it's
			// actually a singular datasource and should never require a client-side filter.

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

func readNamespaces(d *schema.ResourceData, m interface{}) error {
	sync := &NamespacesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
}

type NamespacesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetNamespaceResponse
}

func (s *NamespacesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NamespacesDataSourceCrud) Get() error {
	request := oci_object_storage.GetNamespaceRequest{}

	response, err := s.Client.GetNamespace(context.Background(), request, getRetryOptions(false, "object_storage")...)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NamespacesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	if s.Res.Value != nil {
		s.D.Set("namespace", *s.Res.Value)
	}

	return
}
