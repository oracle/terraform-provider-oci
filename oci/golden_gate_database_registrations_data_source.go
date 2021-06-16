// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v42/goldengate"
)

func init() {
	RegisterDatasource("oci_golden_gate_database_registrations", GoldenGateDatabaseRegistrationsDataSource())
}

func GoldenGateDatabaseRegistrationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readGoldenGateDatabaseRegistrations,
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
			"database_registration_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(GoldenGateDatabaseRegistrationResource()),
						},
					},
				},
			},
		},
	}
}

func readGoldenGateDatabaseRegistrations(d *schema.ResourceData, m interface{}) error {
	sync := &GoldenGateDatabaseRegistrationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).goldenGateClient()

	return ReadResource(sync)
}

type GoldenGateDatabaseRegistrationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_golden_gate.GoldenGateClient
	Res    *oci_golden_gate.ListDatabaseRegistrationsResponse
}

func (s *GoldenGateDatabaseRegistrationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GoldenGateDatabaseRegistrationsDataSourceCrud) Get() error {
	request := oci_golden_gate.ListDatabaseRegistrationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_golden_gate.ListDatabaseRegistrationsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "golden_gate")

	response, err := s.Client.ListDatabaseRegistrations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDatabaseRegistrations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *GoldenGateDatabaseRegistrationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("GoldenGateDatabaseRegistrationsDataSource-", GoldenGateDatabaseRegistrationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	databaseRegistration := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DatabaseRegistrationSummaryToMap(item))
	}
	databaseRegistration["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, GoldenGateDatabaseRegistrationsDataSource().Schema["database_registration_collection"].Elem.(*schema.Resource).Schema)
		databaseRegistration["items"] = items
	}

	resources = append(resources, databaseRegistration)
	if err := s.D.Set("database_registration_collection", resources); err != nil {
		return err
	}

	return nil
}
