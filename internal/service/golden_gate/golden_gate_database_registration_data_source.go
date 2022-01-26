// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v56/goldengate"
)

func GoldenGateDatabaseRegistrationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["database_registration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(GoldenGateDatabaseRegistrationResource(), fieldMap, readSingularGoldenGateDatabaseRegistration)
}

func readSingularGoldenGateDatabaseRegistration(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDatabaseRegistrationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.ReadResource(sync)
}

type GoldenGateDatabaseRegistrationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.GetDatabaseRegistrationResponse
}

func (s *GoldenGateDatabaseRegistrationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDatabaseRegistrationDataSourceCrud) Get() error {
	request := oci_golden_gate.GetDatabaseRegistrationRequest{}

	if databaseRegistrationId, ok := s.D.GetOkExists("database_registration_id"); ok {
		tmp := databaseRegistrationId.(string)
		request.DatabaseRegistrationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "golden_gate")

	response, err := s.Client.GetDatabaseRegistration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GoldenGateDatabaseRegistrationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AliasName != nil {
		s.D.Set("alias_name", *s.Res.AliasName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
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

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RcePrivateIp != nil {
		s.D.Set("rce_private_ip", *s.Res.RcePrivateIp)
	}

	if s.Res.SecretCompartmentId != nil {
		s.D.Set("secret_compartment_id", *s.Res.SecretCompartmentId)
	}

	if s.Res.SecretId != nil {
		s.D.Set("secret_id", *s.Res.SecretId)
	}

	s.D.Set("session_mode", s.Res.SessionMode)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Username != nil {
		s.D.Set("username", *s.Res.Username)
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}
