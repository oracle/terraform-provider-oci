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

func OciProductCatalogProductDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readSingularOciProductCatalogProductWithContext,
		Schema: map[string]*schema.Schema{
			"product_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_excluded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skus": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						// Optional

						"bpart_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_launched": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ready": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularOciProductCatalogProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogProductDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type OciProductCatalogProductDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oci_product_catalog.ProductClient
	Res    *oci_oci_product_catalog.GetProductResponse
}

func (s *OciProductCatalogProductDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OciProductCatalogProductDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.GetProductRequest{}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oci_product_catalog")

	response, err := s.Client.GetProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OciProductCatalogProductDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ServiceName != nil {
		s.D.Set("service_name", *s.Res.ServiceName)
	}

	skus := []interface{}{}
	for _, item := range s.Res.Skus {
		skus = append(skus, ProductSkuSummaryToMap(item))
	}
	s.D.Set("skus", skus)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLaunched != nil {
		s.D.Set("time_launched", s.Res.TimeLaunched.String())
	}

	if s.Res.TimeReady != nil {
		s.D.Set("time_ready", s.Res.TimeReady.String())
	}

	return nil
}

func ProductSkuSummaryToMap(obj oci_oci_product_catalog.SkuSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BpartNumber != nil {
		result["bpart_number"] = string(*obj.BpartNumber)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	return result
}
