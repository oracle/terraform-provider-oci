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

func FileStorageFileSystemQuotaRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["file_system_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["quota_rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	fieldMap["are_violators_only"] = &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(FileStorageFileSystemQuotaRuleResource(), fieldMap, readSingularFileStorageFileSystemQuotaRule)
}

func readSingularFileStorageFileSystemQuotaRule(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemQuotaRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageFileSystemQuotaRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.GetQuotaRuleResponse
}

func (s *FileStorageFileSystemQuotaRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageFileSystemQuotaRuleDataSourceCrud) Get() error {
	request := oci_file_storage.GetQuotaRuleRequest{}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if quotaRuleId, ok := s.D.GetOkExists("quota_rule_id"); ok {
		tmp := quotaRuleId.(string)
		request.QuotaRuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.GetQuotaRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *FileStorageFileSystemQuotaRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.IsHardQuota != nil {
		s.D.Set("is_hard_quota", *s.Res.IsHardQuota)
	}

	if s.Res.PrincipalId != nil {
		s.D.Set("principal_id", *s.Res.PrincipalId)
	}

	s.D.Set("principal_type", s.Res.PrincipalType)

	if s.Res.QuotaLimitInGigabytes != nil {
		s.D.Set("quota_limit_in_gigabytes", *s.Res.QuotaLimitInGigabytes)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
