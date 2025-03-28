// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneSubscriptionMappingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createTenantmanagercontrolplaneSubscriptionMapping,
		Read:     readTenantmanagercontrolplaneSubscriptionMapping,
		Delete:   deleteTenantmanagercontrolplaneSubscriptionMapping,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"is_explicitly_assigned": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_terminated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createTenantmanagercontrolplaneSubscriptionMapping(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSubscriptionMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.CreateResource(d, sync)
}

func readTenantmanagercontrolplaneSubscriptionMapping(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSubscriptionMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()

	return tfresource.ReadResource(sync)
}

func deleteTenantmanagercontrolplaneSubscriptionMapping(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneSubscriptionMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OrganizationsSubscriptionClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type TenantmanagercontrolplaneSubscriptionMappingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_tenantmanagercontrolplane.SubscriptionClient
	Res                    *oci_tenantmanagercontrolplane.SubscriptionMapping
	DisableNotFoundRetries bool
}

func (s *TenantmanagercontrolplaneSubscriptionMappingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *TenantmanagercontrolplaneSubscriptionMappingResourceCrud) Create() error {
	request := oci_tenantmanagercontrolplane.CreateSubscriptionMappingRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane")

	response, err := s.Client.CreateSubscriptionMapping(context.Background(), request)
	if err != nil {
		return err
	}

	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	s.Res = &response.SubscriptionMapping

	return nil
}

func (s *TenantmanagercontrolplaneSubscriptionMappingResourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.GetSubscriptionMappingRequest{}

	tmp := s.D.Id()
	request.SubscriptionMappingId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane")

	response, err := s.Client.GetSubscriptionMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SubscriptionMapping
	return nil
}

func (s *TenantmanagercontrolplaneSubscriptionMappingResourceCrud) Delete() error {
	request := oci_tenantmanagercontrolplane.DeleteSubscriptionMappingRequest{}

	tmp := s.D.Id()
	request.SubscriptionMappingId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "tenantmanagercontrolplane")

	_, err := s.Client.DeleteSubscriptionMapping(context.Background(), request)
	return err
}

func (s *TenantmanagercontrolplaneSubscriptionMappingResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsExplicitlyAssigned != nil {
		s.D.Set("is_explicitly_assigned", *s.Res.IsExplicitlyAssigned)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubscriptionId != nil {
		s.D.Set("subscription_id", *s.Res.SubscriptionId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTerminated != nil {
		s.D.Set("time_terminated", s.Res.TimeTerminated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func SubscriptionMappingSummaryToMap(obj oci_tenantmanagercontrolplane.SubscriptionMappingSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsExplicitlyAssigned != nil {
		result["is_explicitly_assigned"] = bool(*obj.IsExplicitlyAssigned)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeTerminated != nil {
		result["time_terminated"] = obj.TimeTerminated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
