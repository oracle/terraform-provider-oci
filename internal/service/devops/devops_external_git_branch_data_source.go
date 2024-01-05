// // Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// // Licensed under the Mozilla Public License v2.0
package devops

//
//import (
//	"context"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
//)
//
//func DevopsExternalGitBranchDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularDevopsExternalGitBranch,
//		Schema: map[string]*schema.Schema{
//			"branch_name_contains": {
//				Type:     schema.TypeString,
//				Optional: true,
//			},
//			"connection_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"repository_name_contains": {
//				Type:     schema.TypeString,
//				Optional: true,
//			},
//			// Computed
//			"branches": {
//				Type:     schema.TypeList,
//				Computed: true,
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						// Required
//
//						// Optional
//
//						// Computed
//						"branch_name": {
//							Type:     schema.TypeString,
//							Computed: true,
//						},
//						"commit_id": {
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
//func readSingularDevopsExternalGitBranch(d *schema.ResourceData, m interface{}) error {
//	sync := &DevopsExternalGitBranchDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).DevopsClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type DevopsExternalGitBranchDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_devops.DevopsClient
//	Res    *oci_devops.ListExternalGitBranchesResponse
//}
//
//func (s *DevopsExternalGitBranchDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DevopsExternalGitBranchDataSourceCrud) Get() error {
//	request := oci_devops.ListExternalGitBranchesRequest{}
//
//	if branchNameContains, ok := s.D.GetOkExists("branch_name_contains"); ok {
//		tmp := branchNameContains.(string)
//		request.BranchNameContains = &tmp
//	}
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
//	response, err := s.Client.ListExternalGitBranches(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *DevopsExternalGitBranchDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsExternalGitBranchDataSource-", DevopsExternalGitBranchDataSource(), s.D))
//
//	branches := []interface{}{}
//	for _, item := range s.Res.Branches {
//		branches = append(branches, ExternalGitBranchSummaryToMap(item))
//	}
//	s.D.Set("branches", branches)
//
//	return nil
//}
//
//func ExternalGitBranchSummaryToMap(obj oci_devops.ExternalGitBranchSummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.BranchName != nil {
//		result["branch_name"] = string(*obj.BranchName)
//	}
//
//	if obj.CommitId != nil {
//		result["commit_id"] = string(*obj.CommitId)
//	}
//
//	return result
//}
