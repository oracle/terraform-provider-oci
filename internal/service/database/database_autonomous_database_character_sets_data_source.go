// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousDatabaseCharacterSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabaseCharacterSets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"character_set_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_dedicated": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"is_shared": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"autonomous_database_character_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabaseCharacterSets(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseCharacterSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseCharacterSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseCharacterSetsResponse
}

func (s *DatabaseAutonomousDatabaseCharacterSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseCharacterSetsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseCharacterSetsRequest{}

	if characterSetType, ok := s.D.GetOkExists("character_set_type"); ok {
		request.CharacterSetType = oci_database.ListAutonomousDatabaseCharacterSetsCharacterSetTypeEnum(characterSetType.(string))
	}

	if isDedicated, ok := s.D.GetOkExists("is_dedicated"); ok {
		tmp := isDedicated.(bool)
		request.IsDedicated = &tmp
	}

	if isShared, ok := s.D.GetOkExists("is_shared"); ok {
		tmp := isShared.(bool)
		request.IsShared = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseCharacterSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousDatabaseCharacterSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseCharacterSetsDataSource-", DatabaseAutonomousDatabaseCharacterSetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabaseCharacterSet := map[string]interface{}{}

		if r.Name != nil {
			autonomousDatabaseCharacterSet["name"] = *r.Name
		}

		resources = append(resources, autonomousDatabaseCharacterSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabaseCharacterSetsDataSource().Schema["autonomous_database_character_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_database_character_sets", resources); err != nil {
		return err
	}

	return nil
}
