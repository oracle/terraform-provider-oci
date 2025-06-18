// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package wlms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_wlms "github.com/oracle/oci-go-sdk/v65/wlms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WlmsManagedInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readWlmsManagedInstances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"plugin_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_instance_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"configuration": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"discovery_interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"domain_search_paths": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_arch": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"os_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"plugin_status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"server_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_updated": {
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

func readWlmsManagedInstances(d *schema.ResourceData, m interface{}) error {
	sync := &WlmsManagedInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WeblogicManagementServiceClient()

	return tfresource.ReadResource(sync)
}

type WlmsManagedInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_wlms.WeblogicManagementServiceClient
	Res    *oci_wlms.ListManagedInstancesResponse
}

func (s *WlmsManagedInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *WlmsManagedInstancesDataSourceCrud) Get() error {
	request := oci_wlms.ListManagedInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if pluginStatus, ok := s.D.GetOkExists("plugin_status"); ok {
		request.PluginStatus = oci_wlms.ListManagedInstancesPluginStatusEnum(pluginStatus.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "wlms")

	response, err := s.Client.ListManagedInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListManagedInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *WlmsManagedInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("WlmsManagedInstancesDataSource-", WlmsManagedInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	managedInstance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ManagedInstanceSummaryToMap(item))
	}
	managedInstance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, WlmsManagedInstancesDataSource().Schema["managed_instance_collection"].Elem.(*schema.Resource).Schema)
		managedInstance["items"] = items
	}

	resources = append(resources, managedInstance)
	if err := s.D.Set("managed_instance_collection", resources); err != nil {
		return err
	}

	return nil
}

func ManagedInstanceConfigurationToMap(obj *oci_wlms.ManagedInstanceConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DiscoveryInterval != nil {
		result["discovery_interval"] = int(*obj.DiscoveryInterval)
	}

	result["domain_search_paths"] = obj.DomainSearchPaths

	return result
}

func ManagedInstanceSummaryToMap(obj oci_wlms.ManagedInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HostName != nil {
		result["host_name"] = string(*obj.HostName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.PluginStatus != nil {
		result["plugin_status"] = string(*obj.PluginStatus)
	}

	if obj.ServerCount != nil {
		result["server_count"] = int(*obj.ServerCount)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
