// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package license_manager

//import (
//	"context"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	oci_license_manager "github.com/oracle/oci-go-sdk/v65/licensemanager"
//
//	"github.com/oracle/terraform-provider-oci/internal/client"
//	"github.com/oracle/terraform-provider-oci/internal/tfresource"
//)
//
//func LicenseManagerProductLicenseConsumerDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularLicenseManagerProductLicenseConsumer,
//		Schema: map[string]*schema.Schema{
//			"compartment_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"is_compartment_id_in_subtree": {
//				Type:     schema.TypeBool,
//				Optional: true,
//			},
//			"product_license_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			// Computed
//			"items": {
//				Type:     schema.TypeList,
//				Computed: true,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						// Required
//
//						// Optional
//
//						// Computed
//						"are_all_options_available": {
//							Type:     schema.TypeBool,
//							Computed: true,
//						},
//						"is_base_license_available": {
//							Type:     schema.TypeBool,
//							Computed: true,
//						},
//						"license_unit_type": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"license_units_consumed": {
//							Type:     schema.TypeFloat,
//							Computed: true,
//						},
//						"missing_products": {
//							Type:     schema.TypeList,
//							Computed: true,
//							Elem: &schema.Resource{
//								Schema: map[string]*schema.Schema{
//									// Required
//
//									// Optional
//
//									// Computed
//									"category": {
//										Type:     schema.TypeString,
//										Computed: true,
//									},
//									"count": {
//										Type:     schema.TypeFloat,
//										Computed: true,
//									},
//									"name": {
//										Type:     schema.TypeString,
//										Computed: true,
//									},
//								},
//							},
//						},
//						"product_name": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"resource_compartment_id": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"resource_compartment_name": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"resource_id": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"resource_name": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"resource_unit_count": {
//							Type:     schema.TypeFloat,
//							Computed: true,
//						},
//						"resource_unit_type": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//					},
//				},
//			},
//		},
//	}
//}
//
//func readSingularLicenseManagerProductLicenseConsumer(d *schema.ResourceData, m interface{}) error {
//	sync := &LicenseManagerProductLicenseConsumerDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).LicenseManagerClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type LicenseManagerProductLicenseConsumerDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_license_manager.LicenseManagerClient
//	Res    *oci_license_manager.ListProductLicenseConsumersResponse
//}
//
//func (s *LicenseManagerProductLicenseConsumerDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *LicenseManagerProductLicenseConsumerDataSourceCrud) Get() error {
//	request := oci_license_manager.ListProductLicenseConsumersRequest{}
//
//	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
//		tmp := compartmentId.(string)
//		request.CompartmentId = &tmp
//	}
//
//	if isCompartmentIdInSubtree, ok := s.D.GetOkExists("is_compartment_id_in_subtree"); ok {
//		tmp := isCompartmentIdInSubtree.(bool)
//		request.IsCompartmentIdInSubtree = &tmp
//	}
//
//	if productLicenseId, ok := s.D.GetOkExists("product_license_id"); ok {
//		tmp := productLicenseId.(string)
//		request.ProductLicenseId = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")
//
//	response, err := s.Client.ListProductLicenseConsumers(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *LicenseManagerProductLicenseConsumerDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerProductLicenseConsumerDataSource-", LicenseManagerProductLicenseConsumerDataSource(), s.D))
//
//	items := []interface{}{}
//	for _, item := range s.Res.Items {
//		items = append(items, ProductLicenseConsumerSummaryToMap(item))
//	}
//	s.D.Set("items", items)
//
//	return nil
//}
//
//func ProductToMap(obj oci_license_manager.Product) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	result["category"] = string(obj.Category)
//
//	if obj.Count != nil {
//		result["count"] = float64(*obj.Count)
//	}
//
//	if obj.Name != nil {
//		result["name"] = string(*obj.Name)
//	}
//
//	return result
//}
//
//func ProductLicenseConsumerSummaryToMap(obj oci_license_manager.ProductLicenseConsumerSummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.AreAllOptionsAvailable != nil {
//		result["are_all_options_available"] = bool(*obj.AreAllOptionsAvailable)
//	}
//
//	if obj.IsBaseLicenseAvailable != nil {
//		result["is_base_license_available"] = bool(*obj.IsBaseLicenseAvailable)
//	}
//
//	result["license_unit_type"] = string(obj.LicenseUnitType)
//
//	if obj.LicenseUnitsConsumed != nil {
//		result["license_units_consumed"] = float64(*obj.LicenseUnitsConsumed)
//	}
//
//	missingProducts := []interface{}{}
//	for _, item := range obj.MissingProducts {
//		missingProducts = append(missingProducts, ProductToMap(item))
//	}
//	result["missing_products"] = missingProducts
//
//	if obj.ProductName != nil {
//		result["product_name"] = string(*obj.ProductName)
//	}
//
//	if obj.ResourceCompartmentId != nil {
//		result["resource_compartment_id"] = string(*obj.ResourceCompartmentId)
//	}
//
//	if obj.ResourceCompartmentName != nil {
//		result["resource_compartment_name"] = string(*obj.ResourceCompartmentName)
//	}
//
//	if obj.ResourceId != nil {
//		result["resource_id"] = string(*obj.ResourceId)
//	}
//
//	if obj.ResourceName != nil {
//		result["resource_name"] = string(*obj.ResourceName)
//	}
//
//	if obj.ResourceUnitCount != nil {
//		result["resource_unit_count"] = float64(*obj.ResourceUnitCount)
//	}
//
//	result["resource_unit_type"] = string(obj.ResourceUnitType)
//
//	return result
//}
