// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateDeploymentEnvironmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDeploymentEnvironments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_environment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"category": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_cpu_core_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"environment_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_auto_scaling_enabled_by_default": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"max_cpu_core_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"memory_per_ocpu_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"min_cpu_core_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"network_bandwidth_per_ocpu_in_gbps": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"storage_usage_limit_per_ocpu_in_gbs": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readGoldenGateDeploymentEnvironments(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDeploymentEnvironmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDeploymentEnvironmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListDeploymentEnvironmentsResponse
}

func (s *GoldenGateDeploymentEnvironmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDeploymentEnvironmentsDataSourceCrud) Get() error {
	request := oci_golden_gate.ListDeploymentEnvironmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListDeploymentEnvironments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeploymentEnvironments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDeploymentEnvironmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GoldenGateDeploymentEnvironmentsDataSource-", GoldenGateDeploymentEnvironmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	deploymentEnvironment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeploymentEnvironmentSummaryToMap(item))
	}
	deploymentEnvironment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDeploymentEnvironmentsDataSource().Schema["deployment_environment_collection"].Elem.(*schema.Resource).Schema)
		deploymentEnvironment["items"] = items
	}

	resources = append(resources, deploymentEnvironment)
	if err := s.D.Set("deployment_environment_collection", resources); err != nil {
		return err
	}

	return nil
}

func DeploymentEnvironmentSummaryToMap(obj oci_golden_gate.DeploymentEnvironmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["category"] = string(obj.Category)

	if obj.DefaultCpuCoreCount != nil {
		result["default_cpu_core_count"] = int(*obj.DefaultCpuCoreCount)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["environment_type"] = string(obj.EnvironmentType)

	if obj.IsAutoScalingEnabledByDefault != nil {
		result["is_auto_scaling_enabled_by_default"] = bool(*obj.IsAutoScalingEnabledByDefault)
	}

	if obj.MaxCpuCoreCount != nil {
		result["max_cpu_core_count"] = int(*obj.MaxCpuCoreCount)
	}

	if obj.MemoryPerOcpuInGBs != nil {
		result["memory_per_ocpu_in_gbs"] = int(*obj.MemoryPerOcpuInGBs)
	}

	if obj.MinCpuCoreCount != nil {
		result["min_cpu_core_count"] = int(*obj.MinCpuCoreCount)
	}

	if obj.NetworkBandwidthPerOcpuInGbps != nil {
		result["network_bandwidth_per_ocpu_in_gbps"] = int(*obj.NetworkBandwidthPerOcpuInGbps)
	}

	if obj.StorageUsageLimitPerOcpuInGBs != nil {
		result["storage_usage_limit_per_ocpu_in_gbs"] = int(*obj.StorageUsageLimitPerOcpuInGBs)
	}

	return result
}
