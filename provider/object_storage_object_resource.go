// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/crud"

	"crypto/md5"
	"encoding/hex"

	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

func ObjectResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createObject,
		Read:     readObject,
		Update:   updateObject,
		Delete:   deleteObject,
		Schema: map[string]*schema.Schema{
			// @CODEGEN 2/2018:
			// Previous provider doesn't provide an Update method and sets all non-Computed fields to ForceNew.
			// This was done because the same PutObject() call can be used to create and modify existing objects.
			//
			// For generated code, we removed the ForceNew attribute from non-Computed fields and added
			// an Update method which calls the Create() method. This should have the same effect as the
			// previous behavior; while minimizing diffs between this and the generated code.

			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"content": {
				Type: schema.TypeString,
				// @CODEGEN 2/2018: content is optional and stored as checksum to avoid bloating the state file
				// Generator was setting it as required.
				Optional: true,
				StateFunc: func(body interface{}) string {
					v := body.(string)
					if v == "" {
						return ""
					}
					h := md5.Sum([]byte(v))
					return hex.EncodeToString(h[:])
				},
			},
			"content_encoding": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"content_length": {
				Type: schema.TypeInt,
				// @CODEGEN 2/2018: this was generated as Required, we will compute the length from the 'content'
				Computed: true,
			},
			"content_md5": {
				Type: schema.TypeString,
				// @CODEGEN 2/2018: this was generated as Optional, we will set it from the service's response
				Computed: true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata": {
				// @CODEGEN 2/2018: This should be a map[string]string. Spec doesn't specify this correctly and
				// generates it as a TypeString
				Type:         schema.TypeMap,
				Elem:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateLowerCaseKeysInMetadata,
			},

			// @CODEGEN 2/2018: Removed 'range' field from generated code, as it's not currently supported

			// Computed

		},
	}
}

func createObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.CreateResource(d, sync)
}

func readObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.ReadResource(sync)
}

func updateObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient

	return crud.UpdateResource(d, sync)
}

func deleteObject(d *schema.ResourceData, m interface{}) error {
	sync := &ObjectResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).objectStorageClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

// @CODEGEN 2/2018: The existing provider stores a GetObject response along with the
// namespace, bucket, and object name. There's no struct to represent this in SDK, so
// we define our own.
type ObjectStorageObject struct {
	NamespaceName string
	BucketName    string
	ObjectName    string
	oci_object_storage.GetObjectResponse
}

type ObjectResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_object_storage.ObjectStorageClient
	Res                    *ObjectStorageObject
	DisableNotFoundRetries bool
}

// @CODEGEN 2/2018: The existing provider returns a custom Id in following format:
// "tfobm-object-<namespace_name>/<bucket_name>/<object_name>"
func (s *ObjectResourceCrud) ID() string {
	return "tfobm-object-" + s.Res.NamespaceName + "/" + s.Res.BucketName + "/" + s.Res.ObjectName
}

func (s *ObjectResourceCrud) Create() error {
	request := oci_object_storage.PutObjectRequest{}

	if contentEncoding, ok := s.D.GetOkExists("content_encoding"); ok {
		tmp := contentEncoding.(string)
		request.ContentEncoding = &tmp
	}

	if contentLanguage, ok := s.D.GetOkExists("content_language"); ok {
		tmp := contentLanguage.(string)
		request.ContentLanguage = &tmp
	}

	// @CODEGEN 2/2018: Generator code allow you to set the content_length and
	// content_md5 fields from the schema. These are omitted from the existing provider
	// resource because they can be computed.

	if contentType, ok := s.D.GetOkExists("content_type"); ok {
		tmp := contentType.(string)
		request.ContentType = &tmp
	}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if content, ok := s.D.GetOkExists("content"); ok {
		// @CODEGEN 2/2018: The generator doesn't yet handle strings that should be converted to byte arrays.
		tmp := []byte(content.(string))
		tmpLength := len(tmp)
		request.ContentLength = &tmpLength
		request.PutObjectBody = ioutil.NopCloser(bytes.NewBuffer(tmp))
	} else {
		tmp := 0
		request.ContentLength = &tmp
		request.PutObjectBody = ioutil.NopCloser(bytes.NewBuffer([]byte{}))
	}

	if metadata, ok := s.D.GetOkExists("metadata"); ok {
		request.OpcMeta = resourceObjectStorageMapToMetadata(metadata.(map[string]interface{}))
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	// @CODEGEN 2/2018: SDK doesn't have response and retry arguments for PutObject. Removed them for now.
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.PutObject(context.Background(), request)
	if err != nil {
		return err
	}

	// @CODEGEN 2/2018: PutObject() call doesn't return an object. Instead, use existing
	// Get() implementation to retrieve the state of the object.
	return s.Get()
}

func (s *ObjectResourceCrud) Get() error {
	request := oci_object_storage.GetObjectRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if request.NamespaceName == nil || request.BucketName == nil || request.ObjectName == nil {
		return fmt.Errorf("'namespace', 'bucket', or 'object' identifiers are missing")
	}

	// TODO: May be better to use HeadObject() to retrieve status of the object. For large content, doesn't make sense
	// to call Get() all the time
	// @CODEGEN 2/2018: SDK is missing retry arguments for GetObject. Removed them for now.
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	response, err := s.Client.GetObject(context.Background(), request)
	if err != nil {
		return err
	}

	// @CODEGEN 2/2018: We must store the response along with the identifiers that aren't
	// returned in the GetResponse.
	s.Res = &ObjectStorageObject{
		GetObjectResponse: response,
		NamespaceName:     *request.NamespaceName,
		BucketName:        *request.BucketName,
		ObjectName:        *request.ObjectName,
	}

	return nil
}

func (s *ObjectResourceCrud) Update() error {
	// @CODEGEN 2/2018: Update is not supported in the old provider
	return s.Create()
}

func (s *ObjectResourceCrud) Delete() error {
	request := oci_object_storage.DeleteObjectRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "object_storage")

	_, err := s.Client.DeleteObject(context.Background(), request)
	return err
}

func (s *ObjectResourceCrud) SetData() {
	s.D.Set("namespace", s.Res.NamespaceName)
	s.D.Set("bucket", s.Res.BucketName)
	s.D.Set("object", s.Res.ObjectName)

	contentReader := s.Res.Content
	contentArray, err := ioutil.ReadAll(contentReader)
	if err != nil {
		log.Printf("Unable to read 'content' from response. Error: %q", err)
		return
	}
	s.D.Set("content", contentArray)

	if s.Res.ContentEncoding != nil {
		s.D.Set("content_encoding", *s.Res.ContentEncoding)
	}

	if s.Res.ContentLanguage != nil {
		s.D.Set("content_language", *s.Res.ContentLanguage)
	}

	if s.Res.ContentLength != nil {
		s.D.Set("content_length", *s.Res.ContentLength)
	}

	if s.Res.ContentMd5 != nil {
		s.D.Set("content_md5", *s.Res.ContentMd5)
	}

	if s.Res.ContentType != nil {
		s.D.Set("content_type", *s.Res.ContentType)
	}

	if s.Res.OpcMeta != nil {
		// Note: regardless of what we sent to the SDK, the keys we get back from OpcMeta will always be
		// converted to lower case
		if err := s.D.Set("metadata", s.Res.OpcMeta); err != nil {
			log.Printf("Unable to set 'metadata'. Error: %q", err)
			return
		}
	}
}

// @CODEGEN 2/2018: Remove generated mapToObjectSummary as it's not being called

func objectSummaryToMap(obj oci_object_storage.ObjectSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Md5 != nil {
		result["md5"] = string(*obj.Md5)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Size != nil {
		result["size"] = strconv.FormatInt(int64(*obj.Size), 10)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
