// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpSupportedHostShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOcvpSupportedHostShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initial_host_shape_name": {
				Type:          schema.TypeString,
				ConflictsWith: []string{"sddc_type"},
				Optional:      true,
			},
			"is_single_host_sddc_supported": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"sddc_type"},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sddc_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"is_single_host_sddc_supported", "initial_host_shape_name"},
				Deprecated:    tfresource.FieldDeprecatedForAnother("sddc_type", "is_single_host_sddc_supported"),
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//  Required

						// Optional

						// Computed
						"default_ocpu_count": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_single_host_sddc_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_support_monthly_commitment": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_support_monthly_sku": {
							Type:       schema.TypeBool,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("is_support_monthly_sku", "is_support_monthly_commitment"),
						},
						"is_support_shielded_instances": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_family": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_ocpu_count": {
							Type:     schema.TypeList,
							Computed: true,
							MinItems: 0,
							Elem: &schema.Schema{
								Type: schema.TypeFloat,
							},
						},
						"supported_operations": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"supported_sddc_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Deprecated: tfresource.FieldDeprecatedForAnother("supported_sddc_types", "is_single_host_sddc_supported"),
						},
						"supported_vmware_software_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func readOcvpSupportedHostShapes(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSupportedHostShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()

	return tfresource.ReadResource(sync)
}

type OcvpSupportedHostShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.SddcClient
	Res    *oci_ocvp.ListSupportedHostShapesResponse
}

func (s *OcvpSupportedHostShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpSupportedHostShapesDataSourceCrud) Get() error {
	request := oci_ocvp.ListSupportedHostShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if initialHostShapeName, ok := s.D.GetOkExists("initial_host_shape_name"); ok {
		tmp := initialHostShapeName.(string)
		request.InitialHostShapeName = &tmp
	}

	if isSingleHostSddcSupported, ok := s.D.GetOkExists("is_single_host_sddc_supported"); ok {
		tmp := isSingleHostSddcSupported.(bool)
		request.IsSingleHostSddcSupported = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if sddcType, ok := s.D.GetOkExists("sddc_type"); ok {
		if sddcType == "NON_PRODUCTION" {
			tmp := true
			request.IsSingleHostSddcSupported = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.ListSupportedHostShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSupportedHostShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OcvpSupportedHostShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OcvpSupportedHostShapesDataSource-", OcvpSupportedHostShapesDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, SupportedHostShapeSummaryToMap(item))
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, OcvpSupportedHostShapesDataSource().Schema["items"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("items", items); err != nil {
		return err
	}

	return nil
}

func SupportedHostShapeSummaryToMap(obj oci_ocvp.SupportedHostShapeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultOcpuCount != nil {
		result["default_ocpu_count"] = float32(*obj.DefaultOcpuCount)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsSingleHostSddcSupported != nil {
		result["is_single_host_sddc_supported"] = bool(*obj.IsSingleHostSddcSupported)
		supportedSddcTypes := []string{"PRODUCTION"}
		if *obj.IsSingleHostSddcSupported {
			supportedSddcTypes = append(supportedSddcTypes, "NON_PRODUCTION")
		}
		result["supported_sddc_types"] = supportedSddcTypes
	}

	if obj.IsSupportMonthlyCommitment != nil {
		result["is_support_monthly_commitment"] = bool(*obj.IsSupportMonthlyCommitment)
		result["is_support_monthly_sku"] = bool(*obj.IsSupportMonthlyCommitment)
	}

	if obj.IsSupportShieldedInstances != nil {
		result["is_support_shielded_instances"] = bool(*obj.IsSupportShieldedInstances)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ShapeFamily != nil {
		result["shape_family"] = string(*obj.ShapeFamily)
	}

	result["supported_ocpu_count"] = obj.SupportedOcpuCount

	result["supported_operations"] = obj.SupportedOperations

	result["supported_vmware_software_versions"] = obj.SupportedVmwareSoftwareVersions

	return result
}
