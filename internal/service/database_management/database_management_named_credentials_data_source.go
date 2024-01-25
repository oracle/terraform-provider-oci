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

func DatabaseManagementNamedCredentialsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseManagementNamedCredentials,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"associated_resource": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"named_credential_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DatabaseManagementNamedCredentialResource()),
						},
					},
				},
			},
		},
	}
}

func readDatabaseManagementNamedCredentials(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseManagementNamedCredentialsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbManagementClient()

	return tfresource.ReadResource(sync)
}

type DatabaseManagementNamedCredentialsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database_management.DbManagementClient
	Res    *oci_database_management.ListNamedCredentialsResponse
}

func (s *DatabaseManagementNamedCredentialsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseManagementNamedCredentialsDataSourceCrud) Get() error {
	request := oci_database_management.ListNamedCredentialsRequest{}

	if associatedResource, ok := s.D.GetOkExists("associated_resource"); ok {
		tmp := associatedResource.(string)
		request.AssociatedResource = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if scope, ok := s.D.GetOkExists("scope"); ok {
		request.Scope = oci_database_management.ListNamedCredentialsScopeEnum(scope.(string))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		request.Type = oci_database_management.ListNamedCredentialsTypeEnum(type_.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_management")

	response, err := s.Client.ListNamedCredentials(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNamedCredentials(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseManagementNamedCredentialsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseManagementNamedCredentialsDataSource-", DatabaseManagementNamedCredentialsDataSource(), s.D))
	resources := []map[string]interface{}{}
	namedCredential := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NamedCredentialSummaryToMap(item))
	}
	namedCredential["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseManagementNamedCredentialsDataSource().Schema["named_credential_collection"].Elem.(*schema.Resource).Schema)
		namedCredential["items"] = items
	}

	resources = append(resources, namedCredential)
	if err := s.D.Set("named_credential_collection", resources); err != nil {
		return err
	}

	return nil
}
