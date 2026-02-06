// Copyright (c) 2017, 2026, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_config

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_apm_config "github.com/oracle/oci-go-sdk/v65/apmconfig"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmConfigDataFileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createApmConfigDataFileWithContext,
		ReadContext:   readApmConfigDataFileWithContext,
		UpdateContext: updateApmConfigDataFileWithContext,
		DeleteContext: deleteApmConfigDataFileWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"apm_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_file_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"source"},
				StateFunc: func(body interface{}) string {
					v := body.(string)
					if v == "" {
						return ""
					}
					h := md5.Sum([]byte(v))
					return hex.EncodeToString(h[:])
				},
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"content"},
				StateFunc:     tfresource.GetSourceFileState,
				ValidateFunc:  tfresource.ValidateSourceValue,
			},
			"content_length": {
				Type: schema.TypeInt,
				// The length will be computed
				Computed: true,
			},
			"time_last_modified": {
				Type: schema.TypeString,
				// The length will be computed
				Computed: true,
			},
			"content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_encoding": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_md5": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				Type:         schema.TypeMap,
				Elem:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateLowerCaseKeysInMetadata,
			},
		},
	}
}

func createApmConfigDataFileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApmConfigDataFileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readApmConfigDataFileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApmConfigDataFileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateApmConfigDataFileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApmConfigDataFileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteApmConfigDataFileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ApmConfigDataFileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ConfigClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

// The struct defined in the SDK doesn't fully represent the dataFile and metadata
type DataFile struct {
	ApmDomainId         string
	ApmType             string
	DataFileName        string
	DataFileGetResponse oci_apm_config.GetDataFileResponse
}

type ApmConfigDataFileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_config.ConfigClient
	Res                    *DataFile
	DisableNotFoundRetries bool
}

func (s *ApmConfigDataFileResourceCrud) ID() string {
	return GetDataFileCompositeId(s.D.Get("apm_domain_id").(string), s.D.Get("apm_type").(string), s.D.Get("data_file_name").(string))
}

func (s *ApmConfigDataFileResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_apm_config.PutDataFileRequest{}

	if contentDisposition, ok := s.D.GetOkExists("content_disposition"); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentEncoding, ok := s.D.GetOkExists("content_encoding"); ok {
		tmp := contentEncoding.(string)
		request.ContentEncoding = &tmp
	}

	if contentLanguage, ok := s.D.GetOkExists("content_language"); ok {
		tmp := contentLanguage.(string)
		request.ContentLanguage = &tmp
	}

	if contentMD5, ok := s.D.GetOkExists("content_md5"); ok {
		tmp := contentMD5.(string)
		request.ContentMD5 = &tmp
	}

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		tmp := contentType.(string)
		request.ContentType = &tmp
	}

	// Use a source file when defined
	if source, ok := s.D.GetOkExists("source"); ok {
		sourceFile, err := os.Open(source.(string))
		if err != nil {
			return fmt.Errorf("the specified source is not available: %q", err)
		}
		defer tfresource.SafeClose(sourceFile, &err)
		request.PutDataFileBody = ioutil.NopCloser(sourceFile)
	} else if content, ok := s.D.GetOkExists("content"); ok {
		tmp := []byte(content.(string))
		request.PutDataFileBody = ioutil.NopCloser(bytes.NewReader(tmp))
	}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if apmType, ok := s.D.GetOkExists("apm_type"); ok {
		tmp := apmType.(string)
		request.ApmType = &tmp
	}

	if dataFileName, ok := s.D.GetOkExists("data_file_name"); ok {
		tmp := dataFileName.(string)
		request.DataFileName = &tmp
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		tmp, err := metaDataObjectMapToMetaDataJsonString(metadata.(map[string]interface{}))
		if err != nil {
			return fmt.Errorf("could not convert map to string %q", err)
		}
		request.Metadata = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	_, err := s.Client.PutDataFile(ctx, request)
	if err != nil {
		return err
	}
	id := GetDataFileCompositeId(*request.ApmDomainId, *request.ApmType, *request.DataFileName)
	s.D.SetId(id)

	return s.GetWithContext(ctx)
}

func (s *ApmConfigDataFileResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_apm_config.GetDataFileRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if apmType, ok := s.D.GetOkExists("apm_type"); ok {
		tmp := apmType.(string)
		request.ApmType = &tmp
	}

	if dataFileName, ok := s.D.GetOkExists("data_file_name"); ok {
		tmp := dataFileName.(string)
		request.DataFileName = &tmp
	}

	apmDomainId, apmType, dataFileName, err := parseDataFileCompositeId(s.D.Id())
	if err == nil {
		request.DataFileName = &dataFileName
		request.ApmType = &apmType
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	response, err := s.Client.GetDataFile(ctx, request)
	if err != nil {
		return err
	}

	// We must store the response along with the identifiers that aren't returned in the GetResponse.
	s.Res = &DataFile{
		DataFileGetResponse: response,
		DataFileName:        *request.DataFileName,
		ApmDomainId:         *request.ApmDomainId,
		ApmType:             *request.ApmType,
	}

	return nil
}

func (s *ApmConfigDataFileResourceCrud) UpdateWithContext(ctx context.Context) error {
	return s.CreateWithContext(ctx)
}

func (s *ApmConfigDataFileResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_apm_config.DeleteDataFileRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if apmType, ok := s.D.GetOkExists("apm_type"); ok {
		tmp := apmType.(string)
		request.ApmType = &tmp
	}

	if dataFileName, ok := s.D.GetOkExists("data_file_name"); ok {
		tmp := dataFileName.(string)
		request.DataFileName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_config")

	_, err := s.Client.DeleteDataFile(ctx, request)
	return err
}

func (s *ApmConfigDataFileResourceCrud) SetData() error {
	s.D.Set("apm_domain_id", s.Res.ApmDomainId)
	s.D.Set("apm_type", s.Res.ApmType)
	s.D.Set("data_file_name", s.Res.DataFileName)

	contentReader := s.Res.DataFileGetResponse.Content
	if contentReader != nil {
		contentArray, err := ioutil.ReadAll(contentReader)
		if err != nil {
			log.Printf("Unable to read 'content' from response. Error: %q", err)
			return err
		}
		h := md5.Sum(contentArray)
		s.D.Set("content", hex.EncodeToString(h[:]))
	}

	if s.Res.DataFileGetResponse.ContentType != nil {
		s.D.Set("content_type", *s.Res.DataFileGetResponse.ContentType)
	}

	if s.Res.DataFileGetResponse.ContentLength != nil {
		s.D.Set("content_length", *s.Res.DataFileGetResponse.ContentLength)
	}

	if s.Res.DataFileGetResponse.ContentMd5 != nil {
		s.D.Set("content_md5", *s.Res.DataFileGetResponse.ContentMd5)
	}

	if s.Res.DataFileGetResponse.ContentEncoding != nil {
		s.D.Set("content_encoding", *s.Res.DataFileGetResponse.ContentEncoding)
	}

	if s.Res.DataFileGetResponse.ContentLanguage != nil {
		s.D.Set("content_language", *s.Res.DataFileGetResponse.ContentLanguage)
	}

	if s.Res.DataFileGetResponse.ContentDisposition != nil {
		s.D.Set("content_disposition", *s.Res.DataFileGetResponse.ContentDisposition)
	}

	if s.Res.DataFileGetResponse.LastModified != nil {
		s.D.Set("time_last_modified", s.Res.DataFileGetResponse.LastModified.String())
	}

	if s.Res.DataFileGetResponse.Metadata != nil {
		metadata, err := jsonStringToMetaDataStringMap(*s.Res.DataFileGetResponse.Metadata)
		if err != nil {
			return fmt.Errorf("could not set metadata map: %q", err)
		}
		s.D.Set("metadata", metadata)
	}

	return nil
}

func GetDataFileCompositeId(apmDomainId string, apmType string, dataFileName string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	apmType = url.PathEscape(apmType)
	dataFileName = url.PathEscape(dataFileName)
	compositeId := "dataFiles/" + dataFileName + "/apmDomainId/" + apmDomainId + "/apmType/" + apmType
	return compositeId
}

func parseDataFileCompositeId(compositeId string) (apmDomainId string, apmType string, dataFileName string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("dataFiles/.*/apmDomainId/.*/apmType/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	dataFileName, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])
	apmType, _ = url.PathUnescape(parts[5])

	return
}

func DataFileSummaryToMap(obj oci_apm_config.DataFileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApmType != nil {
		result["apm_type"] = string(*obj.ApmType)
	}

	if obj.Md5 != nil {
		result["md5"] = string(*obj.Md5)
	}

	result["metadata"] = obj.Metadata

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = strconv.FormatInt(*obj.SizeInBytes, 10)
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}

// The SDK will return all 'metadata' header keys as lowercase, regardless of how they're specified in the config.
//
// To avoid unnecessary diffs and updates, we need to ensure all config keys for 'metadata' are lowercase.
// This avoids a conflict where our config has a non-lowercase key (e.g. MyKey) while the state file has a lowercase
// key (e.g. mykey) from the SDK.
//
// If we ran a 'terraform plan' on this without any config changes, Terraform will detect a diff between state and
// config; even though nothing changed in the state file.
func validateLowerCaseKeysInMetadata(raw interface{}, fieldName string) ([]string, []error) {
	metadataMap, ok := raw.(map[string]interface{})
	if !ok {
		return nil, []error{fmt.Errorf("Could not convert '%s' to map during validation.", fieldName)}
	}

	errors := []error{}
	for key := range metadataMap {
		lowercaseKey := strings.ToLower(key)
		if key != lowercaseKey {
			errors = append(errors, fmt.Errorf("All '%s' keys must be lowercase. Please specify '%s' as '%s'", fieldName, key, lowercaseKey))
		}
	}

	return nil, errors
}

func metaDataObjectMapToMetaDataJsonString(rm map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(rm)
	if err != nil {
		return "", err
	}

	jsonString := string(jsonBytes)
	return jsonString, nil
}

func jsonStringToMetaDataStringMap(jsonStr string) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal([]byte(jsonStr), &result)
	return result, err
}
