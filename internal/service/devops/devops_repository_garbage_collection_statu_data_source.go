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
//func DevopsRepositoryGarbageCollectionStatuDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularDevopsRepositoryGarbageCollectionStatu,
//		Schema: map[string]*schema.Schema{
//			"repository_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			// Computed
//			"status": {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//		},
//	}
//}
//
//func readSingularDevopsRepositoryGarbageCollectionStatu(d *schema.ResourceData, m interface{}) error {
//	sync := &DevopsRepositoryGarbageCollectionStatuDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).DevopsClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//type DevopsRepositoryGarbageCollectionStatuDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_devops.DevopsClient
//	Res    *oci_devops.GetGarbageCollectionStatusResponse
//}
//
//func (s *DevopsRepositoryGarbageCollectionStatuDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DevopsRepositoryGarbageCollectionStatuDataSourceCrud) Get() error {
//	request := oci_devops.GetGarbageCollectionStatusRequest{}
//
//	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
//		tmp := repositoryId.(string)
//		request.RepositoryId = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")
//
//	response, err := s.Client.GetGarbageCollectionStatus(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *DevopsRepositoryGarbageCollectionStatuDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryGarbageCollectionStatuDataSource-", DevopsRepositoryGarbageCollectionStatuDataSource(), s.D))
//
//	s.D.Set("status", s.Res.Status)
//
//	return nil
//}
