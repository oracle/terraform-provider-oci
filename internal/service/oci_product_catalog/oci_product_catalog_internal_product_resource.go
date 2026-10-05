// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci_product_catalog

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_oci_product_catalog "github.com/oracle/oci-go-sdk/v65/ociproductcatalog"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OciProductCatalogInternalProductResource() *schema.Resource {
	// For internal product, backend update/delete are no-ops. Update refreshes
	// state via GET and delete only clears Terraform state.

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createOciProductCatalogInternalProductWithContext,
		ReadContext:   readOciProductCatalogInternalProductWithContext,
		UpdateContext: updateOciProductCatalogInternalProductWithContext,
		DeleteContext: deleteOciProductCatalogInternalProductWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Expose the raw product OCID to enable data-source chaining in tests and user configs
			"product_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"backfill_eligibility": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_excluded": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"limits": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"public_limit_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"public_service_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"meters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// Optional
			"skus": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						// Optional

						"bpart_number": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
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

func createOciProductCatalogInternalProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalProductResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductInternalClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readOciProductCatalogInternalProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalProductResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductInternalClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateOciProductCatalogInternalProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalProductResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductInternalClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteOciProductCatalogInternalProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// No-op delete: this resource is managed internally and should not trigger
	// any backend delete operation. Clear the state so Terraform treats it as
	// removed and return success.
	d.SetId("")
	return nil
}

type OciProductCatalogInternalProductResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_oci_product_catalog.ProductInternalClient
	Res                    *oci_oci_product_catalog.Product
	DisableNotFoundRetries bool
}

func (s *OciProductCatalogInternalProductResourceCrud) ID() string {
	return GetInternalProductCompositeId(*s.Res.Id)
}

func (s *OciProductCatalogInternalProductResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *OciProductCatalogInternalProductResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oci_product_catalog.ProductLifecycleStateInactive),
		string(oci_oci_product_catalog.ProductLifecycleStateActive),
		string(oci_oci_product_catalog.ProductLifecycleStateNeedsAttention),
	}
}

func (s *OciProductCatalogInternalProductResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *OciProductCatalogInternalProductResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *OciProductCatalogInternalProductResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.CreateInternalProductRequest{}

	if backfillEligibility, ok := s.D.GetOkExists("backfill_eligibility"); ok {
		request.BackfillEligibility = oci_oci_product_catalog.CreateProductDetailsBackfillEligibilityEnum(backfillEligibility.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if isExcluded, ok := s.D.GetOkExists("is_excluded"); ok {
		tmp := isExcluded.(bool)
		request.IsExcluded = &tmp
	}

	if limits, ok := s.D.GetOkExists("limits"); ok {
		interfaces := limits.([]interface{})
		// Build only non-empty limits to avoid sending empty objects
		tmp := make([]oci_oci_product_catalog.Limit, 0, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "limits", stateDataIndex)
			converted, err := s.mapToLimit(fieldKeyFormat)
			if err != nil {
				return err
			}
			if converted.PublicLimitName != nil || converted.PublicServiceName != nil {
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Limits = tmp
		}
	}

	if meters, ok := s.D.GetOkExists("meters"); ok {
		interfaces := meters.([]interface{})
		// Build only non-empty meters to avoid sending empty objects
		tmp := make([]oci_oci_product_catalog.Meter, 0, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "meters", stateDataIndex)
			converted, err := s.mapToMeter(fieldKeyFormat)
			if err != nil {
				return err
			}
			if converted.Name != nil {
				tmp = append(tmp, converted)
			}
		}
		if len(tmp) > 0 {
			request.Meters = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		if strings.TrimSpace(tmp) != "" {
			request.Name = &tmp
		}
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	response, err := s.Client.CreateInternalProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Product
	return nil
}

func (s *OciProductCatalogInternalProductResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.GetInternalProductRequest{}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	productId, err := parseInternalProductCompositeId(s.D.Id())
	if err == nil {
		request.ProductId = &productId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	response, err := s.Client.GetInternalProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Product
	return nil
}

func (s *OciProductCatalogInternalProductResourceCrud) UpdateWithContext(ctx context.Context) error {
	// No-op update: backend does not support updates. Refresh current state via GET.
	request := oci_oci_product_catalog.GetInternalProductRequest{}

	// Derive product ID from state composite ID if available.
	if s.D.Id() != "" {
		if pid, err := parseInternalProductCompositeId(s.D.Id()); err == nil {
			request.ProductId = &pid
		}
	}

	// Fallback to explicit attribute if present.
	if request.ProductId == nil {
		if v, ok := s.D.GetOkExists("product_id"); ok {
			tmp := v.(string)
			request.ProductId = &tmp
		}
	}

	if request.ProductId == nil {
		return fmt.Errorf("missing product ID: unable to refresh internal product during update")
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	response, err := s.Client.GetInternalProduct(ctx, request)
	if err != nil {
		return err
	}
	s.Res = &response.Product
	return nil
}

func (s *OciProductCatalogInternalProductResourceCrud) DeleteWithContext(ctx context.Context) error {
	//This is a no-op functionality for internal calls

	//request := oci_oci_product_catalog.DeleteInternalProductRequest{}
	//
	//if productId, ok := s.D.GetOkExists("product_id"); ok {
	//	tmp := productId.(string)
	//	request.ProductId = &tmp
	//}
	//
	//request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")
	//
	//_, err := s.Client.DeleteInternalProduct(ctx, request)
	s.D.SetId("")
	return nil
}

func (s *OciProductCatalogInternalProductResourceCrud) SetData() error {

	_, err := parseInternalProductCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	// Also set the raw product OCID for downstream references
	if s.Res.Id != nil {
		s.D.Set("product_id", *s.Res.Id)
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
		skus = append(skus, InternalSkuSummaryToMap(item))
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

func GetInternalProductCompositeId(productId string) string {
	productId = url.PathEscape(productId)
	compositeId := "internal/product/" + productId + ""
	return compositeId
}

func parseInternalProductCompositeId(compositeId string) (productId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("internal/product/.*", compositeId)
	if !match || len(parts) != 3 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	// compositeId format: "internal/product/<productId>"
	// parts => ["internal", "product", "<productId>"]
	productId, _ = url.PathUnescape(parts[2])

	return
}

func (s *OciProductCatalogInternalProductResourceCrud) mapToLimit(fieldKeyFormat string) (oci_oci_product_catalog.Limit, error) {
	result := oci_oci_product_catalog.Limit{}

	if publicLimitName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_limit_name")); ok {
		tmp := publicLimitName.(string)
		result.PublicLimitName = &tmp
	}

	if publicServiceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_service_name")); ok {
		tmp := publicServiceName.(string)
		result.PublicServiceName = &tmp
	}

	return result, nil
}

func InternalLimitToMap(obj oci_oci_product_catalog.Limit) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PublicLimitName != nil {
		result["public_limit_name"] = string(*obj.PublicLimitName)
	}

	if obj.PublicServiceName != nil {
		result["public_service_name"] = string(*obj.PublicServiceName)
	}

	return result
}

func (s *OciProductCatalogInternalProductResourceCrud) mapToMeter(fieldKeyFormat string) (oci_oci_product_catalog.Meter, error) {
	result := oci_oci_product_catalog.Meter{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func InternalMeterToMap(obj oci_oci_product_catalog.Meter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func InternalSkuSummaryToMap(obj oci_oci_product_catalog.SkuSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BpartNumber != nil {
		result["bpart_number"] = string(*obj.BpartNumber)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	return result
}

func (s *OciProductCatalogInternalProductResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_oci_product_catalog.ChangeInternalProductCompartmentRequest{}

	if newCompartmentId, ok := s.D.GetOkExists("new_compartment_id"); ok {
		tmp := newCompartmentId.(string)
		changeCompartmentRequest.NewCompartmentId = &tmp
	}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		changeCompartmentRequest.ProductId = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	_, err := s.Client.ChangeInternalProductCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
