// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_management_agent "github.com/oracle/oci-go-sdk/v56/managementagent"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentManagementAgentCountDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularManagementAgentManagementAgentCount,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_by": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"has_plugins": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"install_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dimensions": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"availability_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"has_plugins": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"install_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"platform_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
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

func readSingularManagementAgentManagementAgentCount(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentManagementAgentCountDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

type ManagementAgentManagementAgentCountDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_management_agent.ManagementAgentClient
	Res    *oci_management_agent.SummarizeManagementAgentCountsResponse
}

func (s *ManagementAgentManagementAgentCountDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagementAgentManagementAgentCountDataSourceCrud) Get() error {
	request := oci_management_agent.SummarizeManagementAgentCountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		interfaces := groupBy.([]interface{})
		tmp := make([]oci_management_agent.ManagementAgentGroupByEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_management_agent.ManagementAgentGroupByEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("group_by") {
			request.GroupBy = tmp
		}
	}

	if hasPlugins, ok := s.D.GetOkExists("has_plugins"); ok {
		tmp := hasPlugins.(bool)
		request.HasPlugins = &tmp
	}

	if installType, ok := s.D.GetOkExists("install_type"); ok {
		request.InstallType = oci_management_agent.SummarizeManagementAgentCountsInstallTypeEnum(installType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "management_agent")

	response, err := s.Client.SummarizeManagementAgentCounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagementAgentManagementAgentCountDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagementAgentManagementAgentCountDataSource-", ManagementAgentManagementAgentCountDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagementAgentAggregationToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func ManagementAgentAggregationToMap(obj oci_management_agent.ManagementAgentAggregation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Count != nil {
		result["count"] = int(*obj.Count)
	}

	if obj.Dimensions != nil {
		result["dimensions"] = []interface{}{ManagementAgentAggregationDimensionsToMap(obj.Dimensions)}
	}

	return result
}

func ManagementAgentAggregationDimensionsToMap(obj *oci_management_agent.ManagementAgentAggregationDimensions) map[string]interface{} {
	result := map[string]interface{}{}

	result["availability_status"] = string(obj.AvailabilityStatus)

	if obj.HasPlugins != nil {
		result["has_plugins"] = bool(*obj.HasPlugins)
	}

	result["install_type"] = string(obj.InstallType)

	result["platform_type"] = string(obj.PlatformType)

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
