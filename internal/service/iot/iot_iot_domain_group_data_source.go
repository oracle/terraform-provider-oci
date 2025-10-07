// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotDomainGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["iot_domain_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(IotIotDomainGroupResource(), fieldMap, readSingularIotIotDomainGroupWithContext)
}

func readSingularIotIotDomainGroupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type IotIotDomainGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_iot.IotClient
	Res    *oci_iot.GetIotDomainGroupResponse
}

func (s *IotIotDomainGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IotIotDomainGroupDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotDomainGroupRequest{}

	if iotDomainGroupId, ok := s.D.GetOkExists("iot_domain_group_id"); ok {
		tmp := iotDomainGroupId.(string)
		request.IotDomainGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "iot")

	response, err := s.Client.GetIotDomainGroup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *IotIotDomainGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataHost != nil {
		s.D.Set("data_host", *s.Res.DataHost)
	}

	s.D.Set("db_allow_listed_vcn_ids", s.Res.DbAllowListedVcnIds)

	if s.Res.DbConnectionString != nil {
		s.D.Set("db_connection_string", *s.Res.DbConnectionString)
	}

	if s.Res.DbTokenScope != nil {
		s.D.Set("db_token_scope", *s.Res.DbTokenScope)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

	return nil
}
