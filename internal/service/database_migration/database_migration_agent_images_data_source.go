// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package database_migration

//
//import (
//	"context"
//
//	"github.com/oracle/terraform-provider-oci/internal/client"
//	"github.com/oracle/terraform-provider-oci/internal/tfresource"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
//)
//
//func DatabaseMigrationAgentImagesDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readDatabaseMigrationAgentImages,
//		Schema: map[string]*schema.Schema{
//			"filter": tfresource.DataSourceFiltersSchema(),
//			"agent_image_collection": {
//				Type:     schema.TypeList,
//				Computed: true,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						// Required
//
//						// Optional
//
//						// Computed
//						"items": {
//							Type:     schema.TypeList,
//							Computed: true,
//							Elem: &schema.Resource{
//								Schema: map[string]*schema.Schema{
//									// Required
//
//									// Optional
//
//									// Computed
//									"download_url": {
//										Type:     schema.TypeString,
//										Computed: true,
//									},
//									"version": {
//										Type:     schema.TypeString,
//										Computed: true,
//									},
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//}
//
//func readDatabaseMigrationAgentImages(d *schema.ResourceData, m interface{}) error {
//	sync := &DatabaseMigrationAgentImagesDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).DatabaseMigrationClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type DatabaseMigrationAgentImagesDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_database_migration.DatabaseMigrationClient
//	Res    *oci_database_migration.ListAgentImagesResponse
//}
//
//func (s *DatabaseMigrationAgentImagesDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DatabaseMigrationAgentImagesDataSourceCrud) Get() error {
//	request := oci_database_migration.ListAgentImagesRequest{}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database_migration")
//
//	response, err := s.Client.ListAgentImages(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	request.Page = s.Res.OpcNextPage
//
//	for request.Page != nil {
//		listResponse, err := s.Client.ListAgentImages(context.Background(), request)
//		if err != nil {
//			return err
//		}
//
//		s.Res.Items = append(s.Res.Items, listResponse.Items...)
//		request.Page = listResponse.OpcNextPage
//	}
//
//	return nil
//}
//
//func (s *DatabaseMigrationAgentImagesDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseMigrationAgentImagesDataSource-", DatabaseMigrationAgentImagesDataSource(), s.D))
//	resources := []map[string]interface{}{}
//	agentImage := map[string]interface{}{}
//
//	items := []interface{}{}
//	for _, item := range s.Res.Items {
//		items = append(items, AgentImageSummaryToMap(item))
//	}
//	agentImage["items"] = items
//
//	if f, fOk := s.D.GetOkExists("filter"); fOk {
//		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DatabaseMigrationAgentImagesDataSource().Schema["agent_image_collection"].Elem.(*schema.Resource).Schema)
//		agentImage["items"] = items
//	}
//
//	resources = append(resources, agentImage)
//	if err := s.D.Set("agent_image_collection", resources); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func AgentImageSummaryToMap(obj oci_database_migration.AgentImageSummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.DownloadUrl != nil {
//		result["download_url"] = string(*obj.DownloadUrl)
//	}
//
//	if obj.Version != nil {
//		result["version"] = string(*obj.Version)
//	}
//
//	return result
//}
