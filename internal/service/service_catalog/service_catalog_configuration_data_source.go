// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v65/servicecatalog"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ServiceCatalogConfigurationDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularServiceCatalogConfigurationWithContext,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"is_service_catalog_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularServiceCatalogConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ServiceCatalogConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceCatalogClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ServiceCatalogConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.GetConfigurationResponse
}

func (s *ServiceCatalogConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogConfigurationDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_service_catalog.GetConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_catalog")

	response, err := s.Client.GetConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceCatalogConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceCatalogConfigurationDataSource-", ServiceCatalogConfigurationDataSource(), s.D))

	s.D.Set("is_service_catalog_mode", s.Res.IsServiceCatalogMode)

	return nil
}
