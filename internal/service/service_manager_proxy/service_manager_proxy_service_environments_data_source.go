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

func ServiceManagerProxyServiceEnvironmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceManagerProxyServiceEnvironments,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_environment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_environment_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_environment_collection": {
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
									"console_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_definition": {
										Type:     schema.TypeList,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
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
									"defined_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
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

func readServiceManagerProxyServiceEnvironments(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceManagerProxyServiceEnvironmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ServiceManagerProxyClient()

	return tfresource.ReadResource(sync)
}

type ServiceManagerProxyServiceEnvironmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_manager_proxy.ServiceManagerProxyClient
	Res    *oci_service_manager_proxy.ListServiceEnvironmentsResponse
}

func (s *ServiceManagerProxyServiceEnvironmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceManagerProxyServiceEnvironmentsDataSourceCrud) Get() error {
	request := oci_service_manager_proxy.ListServiceEnvironmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if serviceEnvironmentId, ok := s.D.GetOkExists("id"); ok {
		tmp := serviceEnvironmentId.(string)
		request.ServiceEnvironmentId = &tmp
	}

	if serviceEnvironmentType, ok := s.D.GetOkExists("service_environment_type"); ok {
		tmp := serviceEnvironmentType.(string)
		request.ServiceEnvironmentType = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "service_manager_proxy")

	response, err := s.Client.ListServiceEnvironments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceEnvironments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceManagerProxyServiceEnvironmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ServiceManagerProxyServiceEnvironmentsDataSource-", ServiceManagerProxyServiceEnvironmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	serviceEnvironment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceEnvironmentSummaryToMap(item))
	}
	serviceEnvironment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ServiceManagerProxyServiceEnvironmentsDataSource().Schema["service_environment_collection"].Elem.(*schema.Resource).Schema)
		serviceEnvironment["items"] = items
	}

	resources = append(resources, serviceEnvironment)
	if err := s.D.Set("service_environment_collection", resources); err != nil {
		return err
	}

	return nil
}
