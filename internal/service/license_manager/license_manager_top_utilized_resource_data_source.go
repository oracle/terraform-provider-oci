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
//func LicenseManagerTopUtilizedResourceDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularLicenseManagerTopUtilizedResource,
//		Schema: map[string]*schema.Schema{
//			"compartment_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"is_compartment_id_in_subtree": {
//				Type:     schema.TypeBool,
//				Optional: true,
//			},
//			"resource_unit_type": {
//				Type:     schema.TypeString,
//				Optional: true,
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
//						"total_units": {
//							Type:     schema.TypeFloat,
//							Computed: true,
//						},
//						"unit_type": {
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
//func readSingularLicenseManagerTopUtilizedResource(d *schema.ResourceData, m interface{}) error {
//	sync := &LicenseManagerTopUtilizedResourceDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).LicenseManagerClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type LicenseManagerTopUtilizedResourceDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_license_manager.LicenseManagerClient
//	Res    *oci_license_manager.ListTopUtilizedResourcesResponse
//}
//
//func (s *LicenseManagerTopUtilizedResourceDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *LicenseManagerTopUtilizedResourceDataSourceCrud) Get() error {
//	request := oci_license_manager.ListTopUtilizedResourcesRequest{}
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
//	if resourceUnitType, ok := s.D.GetOkExists("resource_unit_type"); ok {
//		request.ResourceUnitType = oci_license_manager.ListTopUtilizedResourcesResourceUnitTypeEnum(resourceUnitType.(string))
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "license_manager")
//
//	response, err := s.Client.ListTopUtilizedResources(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *LicenseManagerTopUtilizedResourceDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("LicenseManagerTopUtilizedResourceDataSource-", LicenseManagerTopUtilizedResourceDataSource(), s.D))
//
//	items := []interface{}{}
//	for _, item := range s.Res.Items {
//		items = append(items, TopUtilizedResourceSummaryToMap(item))
//	}
//	s.D.Set("items", items)
//
//	return nil
//}
//
//func TopUtilizedResourceSummaryToMap(obj oci_license_manager.TopUtilizedResourceSummary) map[string]interface{} {
//	result := map[string]interface{}{}
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
//	if obj.TotalUnits != nil {
//		result["total_units"] = float64(*obj.TotalUnits)
//	}
//
//	result["unit_type"] = string(obj.UnitType)
//
//	return result
//}
