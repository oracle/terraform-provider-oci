// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceIdentityConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["identity_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BdsBdsInstanceIdentityConfigurationResource(), fieldMap, readSingularBdsBdsInstanceIdentityConfiguration)
}

func readSingularBdsBdsInstanceIdentityConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsBdsInstanceIdentityConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.ReadResource(sync)
}

type BdsBdsInstanceIdentityConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetIdentityConfigurationResponse
}

func (s *BdsBdsInstanceIdentityConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsBdsInstanceIdentityConfigurationDataSourceCrud) Get() error {
	request := oci_bds.GetIdentityConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if identityConfigurationId, ok := s.D.GetOkExists("identity_configuration_id"); ok {
		tmp := identityConfigurationId.(string)
		request.IdentityConfigurationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bds")

	response, err := s.Client.GetIdentityConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsBdsInstanceIdentityConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ConfidentialApplicationId != nil {
		s.D.Set("confidential_application_id", *s.Res.ConfidentialApplicationId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IamUserSyncConfiguration != nil {
		s.D.Set("iam_user_sync_configuration", []interface{}{IamUserSyncConfigurationToMap(s.Res.IamUserSyncConfiguration)})
	} else {
		s.D.Set("iam_user_sync_configuration", nil)
	}

	if s.Res.IdentityDomainId != nil {
		s.D.Set("identity_domain_id", *s.Res.IdentityDomainId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.UpstConfiguration != nil {
		s.D.Set("upst_configuration", []interface{}{UpstConfigurationToMap(s.Res.UpstConfiguration)})
	} else {
		s.D.Set("upst_configuration", nil)
	}

	return nil
}
