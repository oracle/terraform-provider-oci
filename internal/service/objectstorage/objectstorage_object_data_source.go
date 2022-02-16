// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package objectstorage

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/v58/objectstorage"
)

func ObjectStorageObjectDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularObjectStorageObject,
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"http_response_cache_control": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_content_disposition": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_content_encoding": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_content_language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_content_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_response_expires": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_length_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				//default value is 1MB
				Default: 1048576,
			},
			"base64_encode_content": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"version_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"cache_control": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_disposition": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_encoding": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_language": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_length": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_md5": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularObjectStorageObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectStorageObjectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ObjectStorageClient()

	return tfresource.ReadResource(sync)
}

type ObjectStorageObjectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetObjectResponse
}

func (s *ObjectStorageObjectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectStorageObjectDataSourceCrud) Get() error {

	headObjectRequest := &oci_object_storage.HeadObjectRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		bucketName := bucket.(string)
		headObjectRequest.BucketName = &bucketName
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		namespaceName := namespace.(string)
		headObjectRequest.NamespaceName = &namespaceName
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		objectName := object.(string)
		headObjectRequest.ObjectName = &objectName
	}

	headObjectRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "object_storage")

	if versionId, ok := s.D.GetOkExists("version_id"); ok {
		tmp := versionId.(string)
		headObjectRequest.VersionId = &tmp
	}

	headObjectResponse, err := s.Client.HeadObject(context.Background(), *headObjectRequest)
	if err != nil {
		return err
	}

	if contentLengthLimit, ok := s.D.GetOkExists("content_length_limit"); ok {
		tmpInt64 := int64(contentLengthLimit.(int))

		if tmpInt64 < *headObjectResponse.ContentLength {
			return fmt.Errorf("the requested object's content length is %v the limit is set to %v", *headObjectResponse.ContentLength, tmpInt64)
		}

	}

	request := oci_object_storage.GetObjectRequest{}
	request.NamespaceName = headObjectRequest.NamespaceName
	request.BucketName = headObjectRequest.BucketName
	request.ObjectName = headObjectRequest.ObjectName

	if versionId, ok := s.D.GetOkExists("version_id"); ok {
		tmp := versionId.(string)
		request.VersionId = &tmp
	}

	if httpResponseCacheControl, ok := s.D.GetOkExists("http_response_cache_control"); ok {
		tmp := httpResponseCacheControl.(string)
		request.HttpResponseCacheControl = &tmp
	}

	if httpResponseContentDisposition, ok := s.D.GetOkExists("http_response_content_disposition"); ok {
		tmp := httpResponseContentDisposition.(string)
		request.HttpResponseContentDisposition = &tmp
	}

	if httpResponseContentEncoding, ok := s.D.GetOkExists("http_response_content_encoding"); ok {
		tmp := httpResponseContentEncoding.(string)
		request.HttpResponseContentEncoding = &tmp
	}

	if httpResponseContentLanguage, ok := s.D.GetOkExists("http_response_content_language"); ok {
		tmp := httpResponseContentLanguage.(string)
		request.HttpResponseContentLanguage = &tmp
	}

	if httpResponseContentType, ok := s.D.GetOkExists("http_response_content_type"); ok {
		tmp := httpResponseContentType.(string)
		request.HttpResponseContentType = &tmp
	}

	if httpResponseExpires, ok := s.D.GetOkExists("http_response_expires"); ok {
		tmp := httpResponseExpires.(string)
		request.HttpResponseExpires = &tmp
	}

	response, err := s.Client.GetObject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectStorageObjectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ObjectStorageObjectDataSource-", ObjectStorageObjectDataSource(), s.D))

	base64EncodeContent := false
	if tmp, ok := s.D.GetOkExists("base64_encode_content"); ok {
		base64EncodeContent = tmp.(bool)
	}

	contentReader := s.Res.Content
	contentArray, err := ioutil.ReadAll(contentReader)
	if err != nil {
		log.Printf("unable to read 'content' from response. Error: %v", err)
	} else if base64EncodeContent {
		// This use case is for v0.12, where content should be base64 encoded to avoid
		// being normalized before setting in state.
		s.D.Set("content", base64.StdEncoding.EncodeToString(contentArray))
	} else {
		s.D.Set("content", string(contentArray))
	}

	if s.Res.CacheControl != nil {
		s.D.Set("cache_control", *s.Res.CacheControl)
	}

	if s.Res.ContentDisposition != nil {
		s.D.Set("content_disposition", *s.Res.ContentDisposition)
	}

	if s.Res.ContentEncoding != nil {
		s.D.Set("content_encoding", *s.Res.ContentEncoding)
	}

	if s.Res.ContentLanguage != nil {
		s.D.Set("content_language", *s.Res.ContentLanguage)
	}

	if s.Res.ContentLength != nil {
		s.D.Set("content_length", strconv.FormatInt(*s.Res.ContentLength, 10))
	}

	if s.Res.ContentMd5 != nil {
		s.D.Set("content_md5", *s.Res.ContentMd5)
	}

	if s.Res.ContentType != nil {
		s.D.Set("content_type", *s.Res.ContentType)
	}

	if s.Res.VersionId != nil {
		s.D.Set("version_id", *s.Res.VersionId)
	}

	s.D.Set("storage_tier", string(s.Res.StorageTier))

	if s.Res.OpcMeta != nil {
		if err := s.D.Set("metadata", s.Res.OpcMeta); err != nil {
			log.Printf("unable to set 'metadata'. Error: %v", err)
		}
	}

	return nil
}
