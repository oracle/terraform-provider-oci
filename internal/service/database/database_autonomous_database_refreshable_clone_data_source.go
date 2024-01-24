// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

//import (
//	"context"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	oci_database "github.com/oracle/oci-go-sdk/v65/database"
//
//	"github.com/oracle/terraform-provider-oci/internal/client"
//	"github.com/oracle/terraform-provider-oci/internal/tfresource"
//)
//
//func DatabaseAutonomousDatabaseRefreshableCloneDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularDatabaseAutonomousDatabaseRefreshableClone,
//		Schema: map[string]*schema.Schema{
//			"autonomous_database_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			// Computed
//			"items": {
//				Type:     schema.TypeList,
//				Computed: true,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						// Required
//
//						// Optional
//
//						// Computed
//						"id": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"region": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//					},
//				},
//			},
//		},
//	}
//}
//
//func readSingularDatabaseAutonomousDatabaseRefreshableClone(d *schema.ResourceData, m interface{}) error {
//	sync := &DatabaseAutonomousDatabaseRefreshableCloneDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).DatabaseClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type DatabaseAutonomousDatabaseRefreshableCloneDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_database.DatabaseClient
//	Res    *oci_database.ListAutonomousDatabaseRefreshableClonesResponse
//}
//
//func (s *DatabaseAutonomousDatabaseRefreshableCloneDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DatabaseAutonomousDatabaseRefreshableCloneDataSourceCrud) Get() error {
//	request := oci_database.ListAutonomousDatabaseRefreshableClonesRequest{}
//
//	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
//		tmp := autonomousDatabaseId.(string)
//		request.AutonomousDatabaseId = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")
//
//	response, err := s.Client.ListAutonomousDatabaseRefreshableClones(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *DatabaseAutonomousDatabaseRefreshableCloneDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseRefreshableCloneDataSource-", DatabaseAutonomousDatabaseRefreshableCloneDataSource(), s.D))
//
//	items := []interface{}{}
//	for _, item := range s.Res.Items {
//		items = append(items, RefreshableCloneSummaryToMap(item))
//	}
//	s.D.Set("items", items)
//
//	return nil
//}
