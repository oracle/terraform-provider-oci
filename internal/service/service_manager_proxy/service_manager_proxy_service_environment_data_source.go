// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package service_manager_proxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_manager_proxy "github.com/oracle/oci-go-sdk/v58/servicemanagerproxy"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ServiceManagerProxyServiceEnvironmentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularServiceManagerProxyServiceEnvironment,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_environment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"console_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_definition": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"short_display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"service_environment_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"environment_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularServiceManagerProxyServiceEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceManagerProxyServiceEnvironmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceManagerProxyClient()

	return tfresource.ReadResource(sync)
}

type ServiceManagerProxyServiceEnvironmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_manager_proxy.ServiceManagerProxyClient
	Res    *oci_service_manager_proxy.GetServiceEnvironmentResponse
}

func (s *ServiceManagerProxyServiceEnvironmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceManagerProxyServiceEnvironmentDataSourceCrud) Get() error {
	request := oci_service_manager_proxy.GetServiceEnvironmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if serviceEnvironmentId, ok := s.D.GetOkExists("service_environment_id"); ok {
		tmp := serviceEnvironmentId.(string)
		request.ServiceEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_manager_proxy")

	response, err := s.Client.GetServiceEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ServiceManagerProxyServiceEnvironmentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("compartment_id", *s.Res.CompartmentId)

	if s.Res.ConsoleUrl != nil {
		s.D.Set("console_url", *s.Res.ConsoleUrl)
	}

	if s.Res.ServiceDefinition != nil {
		s.D.Set("service_definition", []interface{}{ServiceDefinitionToMap(s.Res.ServiceDefinition)})
	} else {
		s.D.Set("service_definition", nil)
	}

	serviceEnvironmentEndpoints := []interface{}{}
	for _, item := range s.Res.ServiceEnvironmentEndpoints {
		serviceEnvironmentEndpoints = append(serviceEnvironmentEndpoints, ServiceEnvironmentEndPointOverviewToMap(item))
	}
	s.D.Set("service_environment_endpoints", serviceEnvironmentEndpoints)

	s.D.Set("status", s.Res.Status)

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	return nil
}

func ServiceDefinitionToMap(obj *oci_service_manager_proxy.ServiceDefinition) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ShortDisplayName != nil {
		result["short_display_name"] = string(*obj.ShortDisplayName)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func ServiceEnvironmentEndPointOverviewToMap(obj oci_service_manager_proxy.ServiceEnvironmentEndPointOverview) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["environment_type"] = string(obj.EnvironmentType)

	if obj.Url != nil {
		result["url"] = string(*obj.Url)
	}

	return result
}

func ServiceEnvironmentSummaryToMap(obj oci_service_manager_proxy.ServiceEnvironmentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ConsoleUrl != nil {
		result["console_url"] = string(*obj.ConsoleUrl)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ServiceDefinition != nil {
		result["service_definition"] = []interface{}{ServiceDefinitionToMap(obj.ServiceDefinition)}
	}

	serviceEnvironmentEndpoints := []interface{}{}
	for _, item := range obj.ServiceEnvironmentEndpoints {
		serviceEnvironmentEndpoints = append(serviceEnvironmentEndpoints, ServiceEnvironmentEndPointOverviewToMap(item))
	}
	result["service_environment_endpoints"] = serviceEnvironmentEndpoints

	result["status"] = string(obj.Status)

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	return result
}
