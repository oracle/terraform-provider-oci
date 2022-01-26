//// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
//// Licensed under the Mozilla Public License v2.0
//
package devops

//
//import (
//	"context"
//
//	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
//	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"
//)
//
//func DevopsExternalGitRepositoryDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularDevopsExternalGitRepository,
//		Schema: map[string]*schema.Schema{
//			"connection_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"repository_name_contains": {
//				Type:     schema.TypeString,
//				Optional: true,
//			},
//			// Computed
//			"repositories": {
//				Type:     schema.TypeList,
//				Computed: true,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						// Required
//
//						// Optional
//
//						// Computed
//						"http_url": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"repository_id": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"repository_name": {
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
//func readSingularDevopsExternalGitRepository(d *schema.ResourceData, m interface{}) error {
//	sync := &DevopsExternalGitRepositoryDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).DevopsClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type DevopsExternalGitRepositoryDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_devops.DevopsClient
//	Res    *oci_devops.ListExternalGitRepositoriesResponse
//}
//
//func (s *DevopsExternalGitRepositoryDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DevopsExternalGitRepositoryDataSourceCrud) Get() error {
//	request := oci_devops.ListExternalGitRepositoriesRequest{}
//
//	if connectionId, ok := s.D.GetOkExists("connection_id"); ok {
//		tmp := connectionId.(string)
//		request.ConnectionId = &tmp
//	}
//
//	if repositoryNameContains, ok := s.D.GetOkExists("repository_name_contains"); ok {
//		tmp := repositoryNameContains.(string)
//		request.RepositoryNameContains = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")
//
//	response, err := s.Client.ListExternalGitRepositories(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *DevopsExternalGitRepositoryDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsExternalGitRepositoryDataSource-", DevopsExternalGitRepositoryDataSource(), s.D))
//
//	repositories := []interface{}{}
//	for _, item := range s.Res.Repositories {
//		repositories = append(repositories, ExternalGitRepositorySummaryToMap(item))
//	}
//	s.D.Set("repositories", repositories)
//
//	return nil
//}
//
//func ExternalGitRepositorySummaryToMap(obj oci_devops.ExternalGitRepositorySummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.HttpUrl != nil {
//		result["http_url"] = string(*obj.HttpUrl)
//	}
//
//	if obj.RepositoryId != nil {
//		result["repository_id"] = string(*obj.RepositoryId)
//	}
//
//	if obj.RepositoryName != nil {
//		result["repository_name"] = string(*obj.RepositoryName)
//	}
//
//	return result
//}
