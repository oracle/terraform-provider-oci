// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_oce "github.com/oracle/oci-go-sdk/oce"
)

func init() {
	RegisterDatasource("oci_oce_oce_instances", OceOceInstancesDataSource())
}

func OceOceInstancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readOceOceInstances,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			"oce_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(OceOceInstanceResource()),
			},
		},
	}
}

func readOceOceInstances(d *schema.ResourceData, m interface{}) error {
	sync := &OceOceInstancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).oceInstanceClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "oce")

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

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		oceInstance := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AdminEmail != nil {
			oceInstance["admin_email"] = *r.AdminEmail
		}

		if r.DefinedTags != nil {
			oceInstance["defined_tags"] = definedTagsToMap(r.DefinedTags)
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

		oceInstance["instance_usage_type"] = r.InstanceUsageType

		if r.Name != nil {
			oceInstance["name"] = *r.Name
		}

		if r.ObjectStorageNamespace != nil {
			oceInstance["object_storage_namespace"] = *r.ObjectStorageNamespace
		}

		oceInstance["service"] = genericMapToJsonMap(r.Service)

		oceInstance["state"] = r.LifecycleState

		if r.StateMessage != nil {
			oceInstance["state_message"] = *r.StateMessage
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
		resources = ApplyFilters(f.(*schema.Set), resources, OceOceInstancesDataSource().Schema["oce_instances"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("oce_instances", resources); err != nil {
		return err
	}

	return nil
}
