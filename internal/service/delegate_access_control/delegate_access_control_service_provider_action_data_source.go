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

func DelegateAccessControlServiceProviderActionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDelegateAccessControlServiceProviderAction,
		Schema: map[string]*schema.Schema{
			"service_provider_action_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularDelegateAccessControlServiceProviderAction(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlServiceProviderActionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlServiceProviderActionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.GetServiceProviderActionResponse
}

func (s *DelegateAccessControlServiceProviderActionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlServiceProviderActionDataSourceCrud) Get() error {
	request := oci_delegate_access_control.GetServiceProviderActionRequest{}

	if serviceProviderActionId, ok := s.D.GetOkExists("service_provider_action_id"); ok {
		tmp := serviceProviderActionId.(string)
		request.ServiceProviderActionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.GetServiceProviderAction(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DelegateAccessControlServiceProviderActionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Component != nil {
		s.D.Set("component", *s.Res.Component)
	}

	if s.Res.CustomerDisplayName != nil {
		s.D.Set("customer_display_name", *s.Res.CustomerDisplayName)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	properties := []interface{}{}
	for _, item := range s.Res.Properties {
		properties = append(properties, ServiceProviderActionPropertiesToMap(item))
	}
	s.D.Set("properties", properties)

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("service_provider_service_types", s.Res.ServiceProviderServiceTypes)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
