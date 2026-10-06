// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci_product_catalog

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_oci_product_catalog "github.com/oracle/oci-go-sdk/v65/ociproductcatalog"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OciProductCatalogInternalAdminProductDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["product_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(OciProductCatalogInternalAdminProductResource(), fieldMap, readSingularOciProductCatalogInternalAdminProductWithContext)
}

func readSingularOciProductCatalogInternalAdminProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalAdminProductDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductAdminClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OciProductCatalogInternalAdminProductDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oci_product_catalog.ProductAdminClient
	Res    *oci_oci_product_catalog.GetAdminProductResponse
}

func (s *OciProductCatalogInternalAdminProductDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OciProductCatalogInternalAdminProductDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.GetAdminProductRequest{}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oci_product_catalog")

	response, err := s.Client.GetAdminProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OciProductCatalogInternalAdminProductDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsExcluded != nil {
		s.D.Set("is_excluded", *s.Res.IsExcluded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	limits := []interface{}{}
	for _, item := range s.Res.Limits {
		limits = append(limits, AdminLimitToMap(item))
	}
	s.D.Set("limits", limits)

	skus := []interface{}{}
	for _, item := range s.Res.MeterwithSKUs {
		skus = append(skus, InternalAdminMeterWithSkuToMap(item))
	}
	s.D.Set("skus", skus)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ServiceName != nil {
		s.D.Set("service_name", *s.Res.ServiceName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExcluded != nil {
		s.D.Set("time_excluded", s.Res.TimeExcluded.String())
	}

	if s.Res.TimeLaunched != nil {
		s.D.Set("time_launched", s.Res.TimeLaunched.String())
	}

	if s.Res.TimeReady != nil {
		s.D.Set("time_ready", s.Res.TimeReady.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
