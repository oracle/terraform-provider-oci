// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database_management

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseManagementNamedCredentialDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["named_credential_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseManagementNamedCredentialResource(), fieldMap, readSingularDatabaseManagementNamedCredential)
}

func readSingularDatabaseManagementNamedCredential(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementNamedCredentialDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementNamedCredentialDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.GetNamedCredentialResponse
}

func (s *DatabaseManagementNamedCredentialDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementNamedCredentialDataSourceCrud) Get() error {
	request := oci_database_management.GetNamedCredentialRequest{}

	if namedCredentialId, ok := s.D.GetOkExists("named_credential_id"); ok {
		tmp := namedCredentialId.(string)
		request.NamedCredentialId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.GetNamedCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseManagementNamedCredentialDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssociatedResource != nil {
		s.D.Set("associated_resource", *s.Res.AssociatedResource)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Content != nil {
		contentArray := []interface{}{}
		if contentMap := NamedCredentialContentToMap(&s.Res.Content); contentMap != nil {
			contentArray = append(contentArray, contentMap)
		}
		s.D.Set("content", contentArray)
	} else {
		s.D.Set("content", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("scope", s.Res.Scope)

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

	s.D.Set("type", s.Res.Type)

	return nil
}
