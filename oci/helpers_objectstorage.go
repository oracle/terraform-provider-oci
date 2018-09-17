// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"strings"

	"github.com/oracle/oci-go-sdk/common"
	oci_object_storage "github.com/oracle/oci-go-sdk/objectstorage"
)

const defaultFilePartSize int64 = 128 * 1024 * 1024 // 128MB
const defaultNumberOfGoroutines = 10
const maxPartSize int64 = 50 * 1024 * 1024 * 1024
const maxCount int64 = 10000

type MultipartUploadData struct {
	NamespaceName       *string                                 `mandatory:"true"`
	BucketName          *string                                 `mandatory:"true"`
	ObjectName          *string                                 `mandatory:"true"`
	ObjectStorageClient *oci_object_storage.ObjectStorageClient `mandatory:"true"`
	SourcePath          *string                                 `mandatory:"true"`
	SourceInfo          *os.FileInfo                            `mandatory:"true"`
	ContentMD5          *string
	ContentType         *string
	ContentLanguage     *string
	ContentEncoding     *string
	Metadata            map[string]interface{}
	OpcClientRequestID  *string
	RequestMetadata     common.RequestMetadata
}

type objectStorageUploadPartResponse struct {
	response   oci_object_storage.UploadPartResponse
	partNumber *int
	error      error
}

type objectStorageMultiPartUploadContext struct {
	client                  oci_object_storage.ObjectStorageClient
	sourceBlocks            chan objectStorageSourceBlock
	osUploadPartResponses   chan objectStorageUploadPartResponse
	wg                      *sync.WaitGroup
	multipartUploadResponse oci_object_storage.CreateMultipartUploadResponse
	multipartUploadRequest  oci_object_storage.CreateMultipartUploadRequest
}

type objectStorageSourceBlock struct {
	section     *io.SectionReader
	blockNumber *int
}

func resourceObjectStorageMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func resourceObjectStorageMapToOPCMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result["opc-meta-"+k] = v.(string)
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

func validateSourceValue(i interface{}, k string) (s []string, es []error) {
	v, ok := i.(string)
	if !ok {
		es = append(es, fmt.Errorf("expected type of %s to be string", k))
		return
	}
	info, err := os.Stat(v)
	if err != nil {
		es = append(es, fmt.Errorf("cannot get file information for the specified source: %s", k))
		return
	}
	if info.Size() > maxCount*maxPartSize {
		es = append(es, fmt.Errorf("the specified source: %s file is too large", k))
	}
	return
}

// Borrowed from https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/
func safeClose(c io.Closer, err *error) {
	if cerr := c.Close(); cerr != nil && *err == nil {
		*err = cerr
	}
}

func singlePartUpload(multipartUploadData MultipartUploadData) (string, error) {

	sourcePath := *multipartUploadData.SourcePath
	sourceInfo := *multipartUploadData.SourceInfo

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return "", err
	}

	defer safeClose(sourceFile, &err)

	tmpSize := sourceInfo.Size()

	putObjectRequest := &oci_object_storage.PutObjectRequest{
		ContentEncoding: multipartUploadData.ContentEncoding,
		ContentLanguage: multipartUploadData.ContentLanguage,
		ContentType:     multipartUploadData.ContentType,
		BucketName:      multipartUploadData.BucketName,
		ContentLength:   &tmpSize,
		PutObjectBody:   ioutil.NopCloser(sourceFile),
		OpcMeta:         resourceObjectStorageMapToMetadata(multipartUploadData.Metadata),
		NamespaceName:   multipartUploadData.NamespaceName,
		ObjectName:      multipartUploadData.ObjectName,
	}
	putObjectRequest.RequestMetadata.RetryPolicy = multipartUploadData.RequestMetadata.RetryPolicy

	_, err = multipartUploadData.ObjectStorageClient.PutObject(context.Background(), *putObjectRequest)
	if err != nil {
		return "", err
	}

	id := getId(*putObjectRequest.NamespaceName, *putObjectRequest.BucketName, *putObjectRequest.ObjectName)

	return id, nil
}

func MultiPartUpload(multipartUploadData MultipartUploadData) (string, error) {

	sourceInfo := *multipartUploadData.SourceInfo

	if sourceInfo.Size() > defaultFilePartSize {
		return multiPartUploadImpl(multipartUploadData)
	}

	return singlePartUpload(multipartUploadData)
}

func multiPartUploadImpl(multipartUploadData MultipartUploadData) (string, error) {

	multipartUploadRequest := &oci_object_storage.CreateMultipartUploadRequest{
		NamespaceName:   multipartUploadData.NamespaceName,
		BucketName:      multipartUploadData.BucketName,
		RequestMetadata: multipartUploadData.RequestMetadata,
		CreateMultipartUploadDetails: oci_object_storage.CreateMultipartUploadDetails{
			ContentEncoding: multipartUploadData.ContentEncoding,
			ContentLanguage: multipartUploadData.ContentLanguage,
			ContentType:     multipartUploadData.ContentType,
			Object:          multipartUploadData.ObjectName,
			Metadata:        resourceObjectStorageMapToOPCMetadata(multipartUploadData.Metadata),
		},
	}
	source := multipartUploadData.SourcePath
	client := multipartUploadData.ObjectStorageClient

	file, err := os.Open(*source)
	if err != nil {
		return "", fmt.Errorf("error opening source file for upload %q: %s", source, err)
	}
	defer safeClose(file, &err)

	sourceBlocks, err := objectMultiPartSplit(file)
	if err != nil {
		return "", fmt.Errorf("error splitting source file for upload %q: %s", source, err)
	}

	multipartUploadResponse, err := client.CreateMultipartUpload(context.Background(), *multipartUploadRequest)
	if err != nil {
		return "", fmt.Errorf("error creating object in the Oracle cloud %q: %s", source, err)
	}

	workerCount := defaultNumberOfGoroutines

	osUploadPartResponses := make(chan objectStorageUploadPartResponse, len(sourceBlocks))
	sourceBlocksChan := make(chan objectStorageSourceBlock, len(sourceBlocks))

	wg := &sync.WaitGroup{}
	wg.Add(len(sourceBlocks))

	for _, sourceBlock := range sourceBlocks {
		sourceBlocksChan <- sourceBlock
	}
	close(sourceBlocksChan)

	for i := 0; i < workerCount; i++ {
		go uploadPartsWorker(objectStorageMultiPartUploadContext{
			client: *client,
			wg:     wg,
			multipartUploadResponse: multipartUploadResponse,
			multipartUploadRequest:  *multipartUploadRequest,
			sourceBlocks:            sourceBlocksChan,
			osUploadPartResponses:   osUploadPartResponses,
		})
	}

	wg.Wait()

	close(osUploadPartResponses)

	commitMultipartUploadPartDetails := make([]oci_object_storage.CommitMultipartUploadPartDetails, len(sourceBlocks))

	osResponseIndex := 0
	var uploadPartRespErr error
	for osUploadPartResponse := range osUploadPartResponses {
		if osUploadPartResponse.error != nil {
			uploadPartRespErr = osUploadPartResponse.error
			break
		}

		commitMultipartUploadPartDetails[osResponseIndex] = oci_object_storage.CommitMultipartUploadPartDetails{
			PartNum: osUploadPartResponse.partNumber,
			Etag:    osUploadPartResponse.response.ETag,
		}
		osResponseIndex++
	}

	if uploadPartRespErr != nil {
		// just aborting the multi upload for now; but the service itself will handle the same request again
		abortMultipartUploadRequest := oci_object_storage.AbortMultipartUploadRequest{
			NamespaceName:      multipartUploadResponse.Namespace,
			BucketName:         multipartUploadResponse.Bucket,
			ObjectName:         multipartUploadResponse.Object,
			UploadId:           multipartUploadResponse.MultipartUpload.UploadId,
			OpcClientRequestId: multipartUploadResponse.OpcClientRequestId,
			RequestMetadata:    multipartUploadRequest.RequestMetadata,
		}

		_, err := client.AbortMultipartUpload(context.Background(), abortMultipartUploadRequest)

		if err != nil {
			log.Println("[WARN] Aborting the multi part upload failed")
		}

		return "", fmt.Errorf("failed to upload object parts of %q to the Oracle cloud: %s", source, uploadPartRespErr)
	}

	commitMultipartUploadRequest := oci_object_storage.CommitMultipartUploadRequest{
		UploadId:           multipartUploadResponse.MultipartUpload.UploadId,
		NamespaceName:      multipartUploadResponse.Namespace,
		BucketName:         multipartUploadResponse.Bucket,
		ObjectName:         multipartUploadResponse.Object,
		OpcClientRequestId: multipartUploadResponse.OpcClientRequestId,
		RequestMetadata:    multipartUploadRequest.RequestMetadata,
	}
	commitMultipartUploadRequest.PartsToCommit = commitMultipartUploadPartDetails

	_, err = client.CommitMultipartUpload(context.Background(), commitMultipartUploadRequest)
	if err != nil {
		return "", fmt.Errorf("failed to commit multi part upload of %q to the service: %s", source, err)
	}

	id := getId(*commitMultipartUploadRequest.NamespaceName, *commitMultipartUploadRequest.BucketName, *commitMultipartUploadRequest.ObjectName)

	return id, nil
}

func objectMultiPartSplit(file *os.File) ([]objectStorageSourceBlock, error) {

	info, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get FileInfo for the source %q: %s", file.Name(), err)
	}

	offsets, limits, err := splitSizeToOffsetsAndLimits(info.Size())
	sourceBlocks := make([]objectStorageSourceBlock, len(offsets))
	for index := 0; index < len(offsets); index++ {
		tmpIndex := index + 1
		sourceBlocks[index] = objectStorageSourceBlock{
			section:     io.NewSectionReader(file, offsets[index], limits[index]),
			blockNumber: &tmpIndex,
		}
	}

	return sourceBlocks, nil
}

func splitSizeToOffsetsAndLimits(infoSize int64) ([]int64, []int64, error) {
	partSize := defaultFilePartSize
	remainingPart := int64(0)

	totalNumber := infoSize / partSize
	if infoSize%partSize > 0 {
		remainingPart = 1
	}

	if totalNumber+remainingPart > maxCount {
		remainingPart = 0
		if infoSize%maxCount > 0 {
			remainingPart = 1
		}

		partSize = infoSize/maxCount + remainingPart

		if partSize > maxPartSize {
			return nil, nil, fmt.Errorf("the %v size of the source object is more than the service limit", infoSize)
		}
	}

	totalNumber = infoSize / partSize
	if infoSize%partSize > 0 {
		totalNumber++
	}

	offsets := make([]int64, totalNumber)
	limits := make([]int64, totalNumber)
	index := 0
	for offset := int64(0); offset < infoSize; offset += partSize {
		tmpLimit := infoSize - offset
		if partSize < tmpLimit {
			tmpLimit = partSize
		}
		offsets[index] = offset
		limits[index] = tmpLimit
		index++
	}

	return offsets, limits, nil
}

func uploadPartsWorker(ctx objectStorageMultiPartUploadContext) {
	for sourceBlock := range ctx.sourceBlocks {

		block := make([]byte, sourceBlock.section.Size())
		_, err := sourceBlock.section.Read(block)
		if err != nil && err != io.EOF {
			if sourceBlock.blockNumber != nil {
				log.Printf("[ERROR] failed to read source file section %v: %s\n", *sourceBlock.blockNumber, err)
			}

			osUploadPartResponse := &objectStorageUploadPartResponse{
				error: err,
			}
			ctx.osUploadPartResponses <- *osUploadPartResponse
			ctx.wg.Done()
			continue
		}
		tmpLength := int64(len(block))
		uploadPartRequest := &oci_object_storage.UploadPartRequest{
			UploadId:        ctx.multipartUploadResponse.UploadId,
			ObjectName:      ctx.multipartUploadResponse.Object,
			NamespaceName:   ctx.multipartUploadResponse.Namespace,
			BucketName:      ctx.multipartUploadResponse.Bucket,
			RequestMetadata: ctx.multipartUploadRequest.RequestMetadata,
			ContentLength:   &tmpLength,
			UploadPartBody:  ioutil.NopCloser(bytes.NewBuffer(block)),
			UploadPartNum:   sourceBlock.blockNumber,
		}

		uploadPartResponse, err := ctx.client.UploadPart(context.Background(), *uploadPartRequest)

		osUploadPartResponse := &objectStorageUploadPartResponse{
			response:   uploadPartResponse,
			error:      err,
			partNumber: uploadPartRequest.UploadPartNum,
		}

		ctx.osUploadPartResponses <- *osUploadPartResponse
		ctx.wg.Done()
	}
}
