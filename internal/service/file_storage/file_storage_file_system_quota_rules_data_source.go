// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageFileSystemQuotaRulesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageFileSystemQuotaRules,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"are_violators_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"principal_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"principal_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quota_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(FileStorageFileSystemQuotaRuleResource()),
			},
		},
	}
}

func readFileStorageFileSystemQuotaRules(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemQuotaRulesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageFileSystemQuotaRulesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListQuotaRulesResponse
}

func (s *FileStorageFileSystemQuotaRulesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageFileSystemQuotaRulesDataSourceCrud) Get() error {
	request := oci_file_storage.ListQuotaRulesRequest{}

	if areViolatorsOnly, ok := s.D.GetOkExists("are_violators_only"); ok {
		tmp := areViolatorsOnly.(bool)
		request.AreViolatorsOnly = &tmp
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if principalId, ok := s.D.GetOk("principal_id"); ok {
		tmp := principalId.(int)
		request.PrincipalId = &tmp
	}

	if principalType, ok := s.D.GetOkExists("principal_type"); ok {
		request.PrincipalType = oci_file_storage.ListQuotaRulesPrincipalTypeEnum(principalType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListQuotaRules(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListQuotaRules(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageFileSystemQuotaRulesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageFileSystemQuotaRulesDataSource-", FileStorageFileSystemQuotaRulesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		fileSystemQuotaRule := map[string]interface{}{
			"file_system_id": *r.FileSystemId,
			"quota_rule_id":  *r.Id,
		}

		if r.AreViolatorsOnly != nil {
			fileSystemQuotaRule["are_violators_only"] = *r.AreViolatorsOnly
		}

		if r.DisplayName != nil {
			fileSystemQuotaRule["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			fileSystemQuotaRule["id"] = *r.Id
		}

		if r.IsHardQuota != nil {
			fileSystemQuotaRule["is_hard_quota"] = *r.IsHardQuota
		}

		if r.PrincipalId != nil {
			fileSystemQuotaRule["principal_id"] = *r.PrincipalId
		}

		if r.QuotaLimitInGigabytes != nil {
			fileSystemQuotaRule["quota_limit_in_gigabytes"] = *r.QuotaLimitInGigabytes
		}

		if r.TimeCreated != nil {
			fileSystemQuotaRule["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			fileSystemQuotaRule["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, fileSystemQuotaRule)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageFileSystemQuotaRulesDataSource().Schema["quota_rules"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("quota_rules", resources); err != nil {
		return err
	}

	return nil
}
