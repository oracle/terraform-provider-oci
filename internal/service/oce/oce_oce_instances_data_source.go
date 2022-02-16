// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oce

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_oce "github.com/oracle/oci-go-sdk/v58/oce"
)

func OceOceInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOceOceInstances,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oce_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(OceOceInstanceResource()),
			},
		},
	}
}

func readOceOceInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OceInstanceClient()

	return tfresource.ReadResource(sync)
}

type OceOceInstancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_oce.OceInstanceClient
	Res    *oci_oce.ListOceInstancesResponse
}

func (s *OceOceInstancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OceOceInstancesDataSourceCrud) Get() error {
	request := oci_oce.ListOceInstancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_oce.ListOceInstancesLifecycleStateEnum(state.(string))
	}

	if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
		tmp := tenancyId.(string)
		request.TenancyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "oce")

	response, err := s.Client.ListOceInstances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOceInstances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *OceOceInstancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("OceOceInstancesDataSource-", OceOceInstancesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		oceInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AdminEmail != nil {
			oceInstance["admin_email"] = *r.AdminEmail
		}

		if r.DefinedTags != nil {
			oceInstance["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			oceInstance["description"] = *r.Description
		}

		oceInstance["freeform_tags"] = r.FreeformTags

		if r.Guid != nil {
			oceInstance["guid"] = *r.Guid
		}

		if r.Id != nil {
			oceInstance["id"] = *r.Id
		}

		if r.IdcsTenancy != nil {
			oceInstance["idcs_tenancy"] = *r.IdcsTenancy
		}

		oceInstance["instance_access_type"] = r.InstanceAccessType

		oceInstance["instance_license_type"] = r.InstanceLicenseType

		oceInstance["instance_usage_type"] = r.InstanceUsageType

		if r.Name != nil {
			oceInstance["name"] = *r.Name
		}

		if r.ObjectStorageNamespace != nil {
			oceInstance["object_storage_namespace"] = *r.ObjectStorageNamespace
		}

		oceInstance["service"] = tfresource.GenericMapToJsonMap(r.Service)

		oceInstance["state"] = r.LifecycleState

		if r.StateMessage != nil {
			oceInstance["state_message"] = *r.StateMessage
		}

		if r.SystemTags != nil {
			oceInstance["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TenancyId != nil {
			oceInstance["tenancy_id"] = *r.TenancyId
		}

		if r.TenancyName != nil {
			oceInstance["tenancy_name"] = *r.TenancyName
		}

		if r.TimeCreated != nil {
			oceInstance["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			oceInstance["time_updated"] = r.TimeUpdated.String()
		}

		oceInstance["upgrade_schedule"] = r.UpgradeSchedule

		if r.WafPrimaryDomain != nil {
			oceInstance["waf_primary_domain"] = *r.WafPrimaryDomain
		}

		resources = append(resources, oceInstance)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, OceOceInstancesDataSource().Schema["oce_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("oce_instances", resources); err != nil {
		return err
	}

	return nil
}
