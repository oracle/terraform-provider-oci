// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func ObjectDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularObject,
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
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

			// Computed
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
			"metadata": {
				Type:     schema.TypeMap,
				Elem:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return ReadResource(sync)
}

type ObjectDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_object_storage.ObjectStorageClient
	Res    *oci_object_storage.GetObjectResponse
}

func (s *ObjectDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ObjectDataSourceCrud) Get() error {

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

	headObjectRequest.RequestMetadata.RetryPolicy = getRetryPolicy(false, "object_storage")

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

	response, err := s.Client.GetObject(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ObjectDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())

	contentReader := s.Res.Content
	contentArray, err := ioutil.ReadAll(contentReader)
	if err != nil {
		log.Printf("unable to read 'content' from response. Error: %v", err)
	} else {
		s.D.Set("content", string(contentArray))
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

	if s.Res.OpcMeta != nil {
		if err := s.D.Set("metadata", s.Res.OpcMeta); err != nil {
			log.Printf("unable to set 'metadata'. Error: %v", err)
		}
	}

	return nil
}
