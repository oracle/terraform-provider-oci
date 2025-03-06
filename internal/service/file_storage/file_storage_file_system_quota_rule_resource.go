// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageFileSystemQuotaRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFileStorageFileSystemQuotaRule,
		Read:     readFileStorageFileSystemQuotaRule,
		Update:   updateFileStorageFileSystemQuotaRule,
		Delete:   deleteFileStorageFileSystemQuotaRule,
		Schema: map[string]*schema.Schema{
			// Required
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_hard_quota": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"principal_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"quota_limit_in_gigabytes": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Optional
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"principal_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"quota_rule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"are_violators_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			// Computed
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFileStorageFileSystemQuotaRule(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemQuotaRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.CreateResource(d, sync)
}

func readFileStorageFileSystemQuotaRule(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemQuotaRuleResourceCrud{}
	sync.D = d
	attributes := d.State().Attributes
	for key := range attributes {
		if value, ok := d.GetOk(key); ok {
			fmt.Printf("Key: %s, Value: %v\n", key, value)
		} else {
			fmt.Printf("Key: %s not found or is nil.\n", key)
		}
	}
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

func updateFileStorageFileSystemQuotaRule(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemQuotaRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFileStorageFileSystemQuotaRule(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageFileSystemQuotaRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FileStorageFileSystemQuotaRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_file_storage.FileStorageClient
	Res                    *oci_file_storage.QuotaRule
	DisableNotFoundRetries bool
}

func (s *FileStorageFileSystemQuotaRuleResourceCrud) ID() string {
	return GetFileSystemQuotaRuleCompositeId(s.D.Get("file_system_id").(string), *s.Res.Id)
}

func (s *FileStorageFileSystemQuotaRuleResourceCrud) Create() error {
	request := oci_file_storage.CreateQuotaRuleRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if isHardQuota, ok := s.D.GetOkExists("is_hard_quota"); ok {
		tmp := isHardQuota.(bool)
		request.IsHardQuota = &tmp
	}

	if principalId, ok := s.D.GetOkExists("principal_id"); ok {
		tmp := principalId.(int)
		request.PrincipalId = &tmp
	}

	if principalType, ok := s.D.GetOkExists("principal_type"); ok {
		request.PrincipalType = oci_file_storage.CreateQuotaRuleDetailsPrincipalTypeEnum(principalType.(string))
	}

	if quotaLimitInGigabytes, ok := s.D.GetOkExists("quota_limit_in_gigabytes"); ok {
		tmp := quotaLimitInGigabytes.(int)
		request.QuotaLimitInGigabytes = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.CreateQuotaRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.QuotaRule
	return nil
}

func (s *FileStorageFileSystemQuotaRuleResourceCrud) Get() error {
	request := oci_file_storage.GetQuotaRuleRequest{}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if quotaRuleId, ok := s.D.GetOkExists("quota_rule_id"); ok {
		tmp := quotaRuleId.(string)
		request.QuotaRuleId = &tmp
	}

	fileSystemId, quotaRuleId, err := parseFileSystemQuotaRuleCompositeId(s.D.Id())
	if err == nil {
		request.FileSystemId = &fileSystemId
		request.QuotaRuleId = &quotaRuleId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.GetQuotaRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.QuotaRule
	return nil
}

func (s *FileStorageFileSystemQuotaRuleResourceCrud) Update() error {
	request := oci_file_storage.UpdateQuotaRuleRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if quotaLimitInGigabytes, ok := s.D.GetOkExists("quota_limit_in_gigabytes"); ok {
		tmp := quotaLimitInGigabytes.(int)
		request.QuotaLimitInGigabytes = &tmp
	}

	if quotaRuleId, ok := s.D.GetOkExists("quota_rule_id"); ok {
		tmp := quotaRuleId.(string)
		request.QuotaRuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	response, err := s.Client.UpdateQuotaRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.QuotaRule
	return nil
}

func (s *FileStorageFileSystemQuotaRuleResourceCrud) Delete() error {
	request := oci_file_storage.DeleteQuotaRuleRequest{}

	if fileSystemId, ok := s.D.GetOkExists("file_system_id"); ok {
		tmp := fileSystemId.(string)
		request.FileSystemId = &tmp
	}

	if quotaRuleId, ok := s.D.GetOkExists("quota_rule_id"); ok {
		tmp := quotaRuleId.(string)
		request.QuotaRuleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "file_storage")

	_, err := s.Client.DeleteQuotaRule(context.Background(), request)
	return err
}

func (s *FileStorageFileSystemQuotaRuleResourceCrud) SetData() error {

	fileSystemId, _, err := parseFileSystemQuotaRuleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("file_system_id", &fileSystemId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSystemId != nil {
		s.D.Set("file_system_id", *s.Res.FileSystemId)
	}

	if s.Res.Id != nil {
		s.D.Set("quota_rule_id", *s.Res.Id)
	}

	if s.Res.IsHardQuota != nil {
		s.D.Set("is_hard_quota", *s.Res.IsHardQuota)
	}

	if s.Res.PrincipalId != nil {
		s.D.Set("principal_id", *s.Res.PrincipalId)
	}

	if s.Res.PrincipalType != "" {
		s.D.Set("principal_type", s.Res.PrincipalType)
	}

	if s.Res.QuotaLimitInGigabytes != nil {
		s.D.Set("quota_limit_in_gigabytes", *s.Res.QuotaLimitInGigabytes)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	attributes := s.D.State().Attributes
	for key := range attributes {
		if value, ok := s.D.GetOk(key); ok {
			fmt.Printf("Key: %s, Value: %v\n", key, value)
		} else {
			fmt.Printf("Key: %s not found or is nil.\n", key)
		}
	}
	return nil
}

func GetFileSystemQuotaRuleCompositeId(fileSystemId string, quotaRuleId string) string {
	fileSystemId = url.PathEscape(fileSystemId)
	quotaRuleId = url.PathEscape(quotaRuleId)
	compositeId := "fileSystems/" + fileSystemId + "/quotaRules/" + quotaRuleId
	return compositeId
}

func parseFileSystemQuotaRuleCompositeId(compositeId string) (fileSystemId string, quotaRuleId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("fileSystems/.*/quotaRules/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	fileSystemId, _ = url.PathUnescape(parts[1])
	quotaRuleId, _ = url.PathUnescape(parts[3])

	return
}
