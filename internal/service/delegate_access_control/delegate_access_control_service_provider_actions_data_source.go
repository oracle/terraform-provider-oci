// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlServiceProviderActionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDelegateAccessControlServiceProviderActions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_provider_service_type": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_provider_action_summary_collection": {
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
									"component": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"customer_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"properties": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"resource_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_provider_service_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"state": {
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

func readDelegateAccessControlServiceProviderActions(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlServiceProviderActionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlServiceProviderActionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.ListServiceProviderActionsResponse
}

func (s *DelegateAccessControlServiceProviderActionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlServiceProviderActionsDataSourceCrud) Get() error {
	request := oci_delegate_access_control.ListServiceProviderActionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
		request.ResourceType = oci_delegate_access_control.ListServiceProviderActionsResourceTypeEnum(resourceType.(string))
	}

	if serviceProviderServiceType, ok := s.D.GetOkExists("service_provider_service_type"); ok {
		interfaces := serviceProviderServiceType.([]interface{})
		tmp := make([]oci_delegate_access_control.ServiceProviderServiceTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_delegate_access_control.ServiceProviderServiceTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("service_provider_service_type") {
			request.ServiceProviderServiceType = tmp
		}
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_delegate_access_control.ServiceProviderActionLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.ListServiceProviderActions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceProviderActions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DelegateAccessControlServiceProviderActionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DelegateAccessControlServiceProviderActionsDataSource-", DelegateAccessControlServiceProviderActionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	serviceProviderAction := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceProviderActionSummaryToMap(item))
	}
	serviceProviderAction["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DelegateAccessControlServiceProviderActionsDataSource().Schema["service_provider_action_summary_collection"].Elem.(*schema.Resource).Schema)
		serviceProviderAction["items"] = items
	}

	resources = append(resources, serviceProviderAction)
	if err := s.D.Set("service_provider_action_summary_collection", resources); err != nil {
		return err
	}

	return nil
}

func ServiceProviderActionPropertiesToMap(obj oci_delegate_access_control.ServiceProviderActionProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func ServiceProviderActionSummaryToMap(obj oci_delegate_access_control.ServiceProviderActionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Component != nil {
		result["component"] = string(*obj.Component)
	}

	if obj.CustomerDisplayName != nil {
		result["customer_display_name"] = string(*obj.CustomerDisplayName)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["resource_type"] = string(obj.ResourceType)

	result["service_provider_service_types"] = obj.ServiceProviderServiceTypes

	result["state"] = string(obj.LifecycleState)

	return result
}
