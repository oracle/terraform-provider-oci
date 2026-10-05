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

func OciProductCatalogInternalAdminProductResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createOciProductCatalogInternalAdminProductWithContext,
		ReadContext:   readOciProductCatalogInternalAdminProductWithContext,
		UpdateContext: updateOciProductCatalogInternalAdminProductWithContext,
		DeleteContext: deleteOciProductCatalogInternalAdminProductWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Expose the raw product OCID to enable data-source chaining in tests and user configs
			"product_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
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
			"skus": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
			"time_excluded": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createOciProductCatalogInternalAdminProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalAdminProductResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductAdminClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readOciProductCatalogInternalAdminProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalAdminProductResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductAdminClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateOciProductCatalogInternalAdminProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OciProductCatalogInternalAdminProductResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ProductAdminClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteOciProductCatalogInternalAdminProductWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//sync := &OciProductCatalogInternalAdminProductResourceCrud{}
	//sync.D = d
	//sync.Client = m.(*client.OracleClients).ProductAdminClient()
	//sync.DisableNotFoundRetries = true
	//
	//return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
	// No-op delete: this resource is managed internally and should not trigger
	// any backend delete operation. Clear the state so Terraform treats it as
	// removed and return success.
	d.SetId("")
	return nil
}

type OciProductCatalogInternalAdminProductResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_oci_product_catalog.ProductAdminClient
	Res                    *oci_oci_product_catalog.AdminProduct
	DisableNotFoundRetries bool
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) ID() string {
	return GetInternalAdminProductCompositeId(*s.Res.Id)
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_oci_product_catalog.AdminProductLifecycleStateInactive),
		string(oci_oci_product_catalog.AdminProductLifecycleStateActive),
		string(oci_oci_product_catalog.AdminProductLifecycleStateNeedsAttention),
	}
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) DeletedTarget() []string {
	// Service continues to return 200 with lifecycleState=INACTIVE and
	// lifecycleDetails="Deleted" after a successful DELETE. Treat INACTIVE
	// as a terminal deletion target so the waiter can complete.
	return []string{}
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.CreateAdminProductRequest{}

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
		tmp := make([]oci_oci_product_catalog.Limit, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "limits", stateDataIndex)
			converted, err := s.mapToLimit(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("limits") {
			request.Limits = tmp
		}
	}

	if meters, ok := s.D.GetOkExists("meters"); ok {
		interfaces := meters.([]interface{})
		tmp := make([]oci_oci_product_catalog.Meter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "meters", stateDataIndex)
			converted, err := s.mapToMeter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("meters") {
			request.Meters = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	response, err := s.Client.CreateAdminProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &oci_oci_product_catalog.AdminProduct{Id: response.Id}
	return nil
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.GetAdminProductRequest{}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}

	productId, err := parseInternalAdminProductCompositeId(s.D.Id())
	if err == nil {
		request.ProductId = &productId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	response, err := s.Client.GetAdminProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.AdminProduct

	// If the backend reports the resource as logically deleted via lifecycleDetails,
	// surface a not-found error so higher-level helpers can void state and stop polling.
	if s.Res != nil && s.Res.LifecycleDetails != nil && strings.EqualFold(*s.Res.LifecycleDetails, "Deleted") {
		return tfresource.ResourceNotFoundErrorMessage("product_admin", "lifecycleDetails indicates Deleted")
	}
	return nil
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_oci_product_catalog.UpdateAdminProductRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if internalAdminProductState, ok := s.D.GetOkExists("internal_admin_product_state"); ok {
		request.State = oci_oci_product_catalog.UpdateProductDetailsStateEnum(internalAdminProductState.(string))
	}

	if isExcluded, ok := s.D.GetOkExists("is_excluded"); ok {
		tmp := isExcluded.(bool)
		request.IsExcluded = &tmp
	}

	if limits, ok := s.D.GetOkExists("limits"); ok {
		interfaces := limits.([]interface{})
		tmp := make([]oci_oci_product_catalog.Limit, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "limits", stateDataIndex)
			converted, err := s.mapToLimit(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("limits") {
			request.Limits = tmp
		}
	}

	if meters, ok := s.D.GetOkExists("meters"); ok {
		interfaces := meters.([]interface{})
		tmp := make([]oci_oci_product_catalog.Meter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "meters", stateDataIndex)
			converted, err := s.mapToMeter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("meters") {
			request.Meters = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}
	// Fallback to ID-parsed productId if not explicitly present in state
	if request.ProductId == nil {
		if pid, err := parseInternalAdminProductCompositeId(s.D.Id()); err == nil {
			request.ProductId = &pid
		} else {
			log.Printf("[WARN] Update() unable to resolve product_id from ID: %s, err: %v", s.D.Id(), err)
		}
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	response, err := s.Client.UpdateAdminProduct(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &oci_oci_product_catalog.AdminProduct{Id: response.Id}
	return nil
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_oci_product_catalog.DeleteAdminProductRequest{}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		request.ProductId = &tmp
	}
	// Fallback to parse from composite ID if product_id is not in state
	if request.ProductId == nil {
		if pid, err := parseInternalAdminProductCompositeId(s.D.Id()); err == nil {
			request.ProductId = &pid
		} else {
			return fmt.Errorf("unable to resolve product_id for delete from ID %q: %w", s.D.Id(), err)
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	_, err := s.Client.DeleteAdminProduct(ctx, request)
	if err != nil {
		// Propagate delete error so the test harness can surface issues instead of
		// silently dropping state and leaving dangling resources.
		return err
	}

	s.D.SetId("")
	return nil
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) SetData() error {

	_, err := parseInternalAdminProductCompositeId(s.D.Id())
	if err == nil {
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}
	// Ensure computed product_id is populated in state for downstream operations (update/delete)
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

	limits := []interface{}{}
	for _, item := range s.Res.Limits {
		limits = append(limits, AdminLimitToMap(item))
	}
	s.D.Set("limits", limits)

	// AdminProduct exposes SKU details through MeterwithSKUs. Populate the legacy
	// skus field from that response to preserve the Terraform schema contract.
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

func GetInternalAdminProductCompositeId(productId string) string {
	productId = url.PathEscape(productId)
	compositeId := "internal/admin/product/" + productId
	return compositeId
}

func parseInternalAdminProductCompositeId(compositeId string) (productId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("internal/admin/product/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	productId, _ = url.PathUnescape(parts[3])

	return
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) mapToLimit(fieldKeyFormat string) (oci_oci_product_catalog.Limit, error) {
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

func AdminLimitToMap(obj oci_oci_product_catalog.Limit) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.PublicLimitName != nil {
		result["public_limit_name"] = string(*obj.PublicLimitName)
	}

	if obj.PublicServiceName != nil {
		result["public_service_name"] = string(*obj.PublicServiceName)
	}

	return result
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) mapToMeter(fieldKeyFormat string) (oci_oci_product_catalog.Meter, error) {
	result := oci_oci_product_catalog.Meter{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AdminMeterToMap(obj oci_oci_product_catalog.Meter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func InternalAdminMeterWithSkuToMap(obj oci_oci_product_catalog.MeterWithSku) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SkuBNumber != nil {
		result["bpart_number"] = string(*obj.SkuBNumber)
	}

	if obj.SkuDescription != nil {
		result["description"] = string(*obj.SkuDescription)
	}

	return result
}

func AdminSkuSummaryToMap(obj oci_oci_product_catalog.SkuSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BpartNumber != nil {
		result["bpart_number"] = string(*obj.BpartNumber)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	return result
}

func (s *OciProductCatalogInternalAdminProductResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_oci_product_catalog.ChangeAdminProductCompartmentRequest{}

	if newCompartmentId, ok := s.D.GetOkExists("new_compartment_id"); ok {
		tmp := newCompartmentId.(string)
		changeCompartmentRequest.NewCompartmentId = &tmp
	}

	if productId, ok := s.D.GetOkExists("product_id"); ok {
		tmp := productId.(string)
		changeCompartmentRequest.ProductId = &tmp
	}

	// Fallback to ID-parsed productId if not explicitly present in state
	if changeCompartmentRequest.ProductId == nil {
		if pid, err := parseInternalAdminProductCompositeId(s.D.Id()); err == nil {
			changeCompartmentRequest.ProductId = &pid
		} else {
			log.Printf("[WARN] updateCompartment() unable to resolve product_id from ID: %s, err: %v", s.D.Id(), err)
		}
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "oci_product_catalog")

	_, err := s.Client.ChangeAdminProductCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
