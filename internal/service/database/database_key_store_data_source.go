// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseKeyStoreDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["key_store_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseKeyStoreResource(), fieldMap, readSingularDatabaseKeyStore)
}

func readSingularDatabaseKeyStore(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseKeyStoreDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseKeyStoreDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetKeyStoreResponse
}

func (s *DatabaseKeyStoreDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseKeyStoreDataSourceCrud) Get() error {
	request := oci_database.GetKeyStoreRequest{}

	if keyStoreId, ok := s.D.GetOkExists("key_store_id"); ok {
		tmp := keyStoreId.(string)
		request.KeyStoreId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetKeyStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseKeyStoreDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	associatedDatabases := []interface{}{}
	for _, item := range s.Res.AssociatedDatabases {
		associatedDatabases = append(associatedDatabases, KeyStoreAssociatedDatabaseDetailsToMap(item))
	}
	s.D.Set("associated_databases", associatedDatabases)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TypeDetails != nil {
		typeDetailsArray := []interface{}{}
		if typeDetailsMap := KeyStoreTypeDetailsToMap(&s.Res.TypeDetails, true); typeDetailsMap != nil {
			typeDetailsArray = append(typeDetailsArray, typeDetailsMap)
		}
		s.D.Set("type_details", typeDetailsArray)
	} else {
		s.D.Set("type_details", nil)
	}

	return nil
}
