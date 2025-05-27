// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementFleetCredentialDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["fleet_credential_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["fleet_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FleetAppsManagementFleetCredentialResource(), fieldMap, readSingularFleetAppsManagementFleetCredential)
}

func readSingularFleetAppsManagementFleetCredential(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementFleetCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

type FleetAppsManagementFleetCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_fleet_apps_management.FleetAppsManagementClient
	Res    *oci_fleet_apps_management.GetFleetCredentialResponse
}

func (s *FleetAppsManagementFleetCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FleetAppsManagementFleetCredentialDataSourceCrud) Get() error {
	request := oci_fleet_apps_management.GetFleetCredentialRequest{}

	if fleetCredentialId, ok := s.D.GetOkExists("fleet_credential_id"); ok {
		tmp := fleetCredentialId.(string)
		request.FleetCredentialId = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "fleet_apps_management")

	response, err := s.Client.GetFleetCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FleetAppsManagementFleetCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EntitySpecifics != nil {
		entitySpecificsArray := []interface{}{}
		if entitySpecificsMap := CredentialEntitySpecificDetailsToMap(&s.Res.EntitySpecifics); entitySpecificsMap != nil {
			entitySpecificsArray = append(entitySpecificsArray, entitySpecificsMap)
		}
		s.D.Set("entity_specifics", entitySpecificsArray)
	} else {
		s.D.Set("entity_specifics", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Password != nil {
		passwordArray := []interface{}{}
		if passwordMap := CredentialDetailsToMap(&s.Res.Password); passwordMap != nil {
			passwordArray = append(passwordArray, passwordMap)
		}
		s.D.Set("password", passwordArray)
	} else {
		s.D.Set("password", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.User != nil {
		userArray := []interface{}{}
		if userMap := CredentialDetailsToMap(&s.Res.User); userMap != nil {
			userArray = append(userArray, userMap)
		}
		s.D.Set("user", userArray)
	} else {
		s.D.Set("user", nil)
	}

	return nil
}
