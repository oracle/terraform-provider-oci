// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseKeyStoresDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseKeyStores,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_stores": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseKeyStoreResource()),
			},
		},
	}
}

func readDatabaseKeyStores(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseKeyStoresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseKeyStoresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListKeyStoresResponse
}

func (s *DatabaseKeyStoresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseKeyStoresDataSourceCrud) Get() error {
	request := oci_database.ListKeyStoresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListKeyStores(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKeyStores(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseKeyStoresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseKeyStoresDataSource-", DatabaseKeyStoresDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		keyStore := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		associatedDatabases := []interface{}{}
		for _, item := range r.AssociatedDatabases {
			associatedDatabases = append(associatedDatabases, KeyStoreAssociatedDatabaseDetailsToMap(item))
		}
		keyStore["associated_databases"] = associatedDatabases

		if r.DefinedTags != nil {
			keyStore["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			keyStore["display_name"] = *r.DisplayName
		}

		keyStore["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			keyStore["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			keyStore["lifecycle_details"] = *r.LifecycleDetails
		}

		keyStore["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			keyStore["time_created"] = r.TimeCreated.String()
		}

		if r.TypeDetails != nil {
			typeDetailsArray := []interface{}{}
			if typeDetailsMap := KeyStoreTypeDetailsToMap(&r.TypeDetails, true); typeDetailsMap != nil {
				typeDetailsArray = append(typeDetailsArray, typeDetailsMap)
			}
			keyStore["type_details"] = typeDetailsArray
		} else {
			keyStore["type_details"] = nil
		}

		resources = append(resources, keyStore)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseKeyStoresDataSource().Schema["key_stores"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("key_stores", resources); err != nil {
		return err
	}

	return nil
}
