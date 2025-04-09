package tfresource

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"gopkg.in/yaml.v3"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	MlApplicationInstanceTimeout = &schema.ResourceTimeout{
		Create: &ThirtyMinutes,
		Update: &ThirtyMinutes,
		Delete: &ThirtyMinutes,
	}
)

// ValidateMLApplicationPackage Validate the structure of the ml_application_package
func ValidateMLApplicationPackage(val interface{}, key string) ([]string, []error) {
	packageDetails, ok := val.(map[string]interface{})
	if !ok {
		return nil, []error{fmt.Errorf("%s must be a map, got %T", key, val)}
	}
	var errors []error

	if len(packageDetails) == 0 {
		return nil, errors
	}

	// Validate source_type
	sourceType, sourceTypeExists := packageDetails["source_type"].(string)
	if !sourceTypeExists {
		errors = append(errors, fmt.Errorf("%s must contain 'source_type'", key))
	} else {
		validSourceTypes := []string{"local", "object_storage_download", "object_storage"}
		if !stringInSlice(sourceType, validSourceTypes) {
			errors = append(errors, fmt.Errorf("source_type must be one of: %v", validSourceTypes))
		}
	}

	// Validate path or uri based on source_type
	if sourceType == "local" {
		path, pathExists := packageDetails["path"].(string)
		if !pathExists {
			errors = append(errors, fmt.Errorf("path must be specified when source_type is 'local'"))
		} else if !strings.HasPrefix(path, "file://") {
			errors = append(errors, fmt.Errorf("path must start with 'file://', got '%s'", path))
		}
	} else if sourceType == "object_storage_download" || sourceType == "object_storage" {
		uri, uriExists := packageDetails["uri"].(string)
		if !uriExists {
			errors = append(errors, fmt.Errorf("uri must be specified when source_type is '%s'", sourceType))
		} else if !strings.HasPrefix(uri, "https://") {
			errors = append(errors, fmt.Errorf("uri must start with 'https://', got '%s'", uri))
		}
	}

	return nil, errors
}

func SupressPackageUpload(key string, old string, new string, d *schema.ResourceData) bool {
	if key == "ml_application_package.path" {
		raw, ok := d.GetOk("ml_application_package")
		if ok {
			packageDetails, ok := raw.(map[string]interface{})
			if ok {
				oldSplit := strings.SplitN(old, " ", 2)
				packageDetails["path"] = oldSplit[0]
				d.Set("ml_application_package", packageDetails)
				return oldSplit[0] == new
			}
		}
	}
	// Always do not suppress
	return false
}

// Helper function to check if a string exists in a slice
func stringInSlice(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}

// SetPackagePath Customize Diff to update the path with the package version
func SetPackagePath(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	log.Printf("[DEBUG] CustomizeDiff: Processing ml_application_package")
	raw, ok := d.GetOk("ml_application_package")
	if !ok {
		// Set empty map. If we don't set empty map, it throws "After applying this test step, the plan was not empty." error.
		if err := d.SetNew("ml_application_package", make(map[string]interface{})); err != nil {
			return fmt.Errorf("failed to initialise ml_application_package with empty map: %w", err)
		}
		return nil
	}
	packageDetails, ok := raw.(map[string]interface{})
	if !ok {
		return fmt.Errorf("ml_application_package must be a map")
	}
	sourceType, sourceTypeExists := packageDetails["source_type"].(string)
	if !sourceTypeExists {
		return fmt.Errorf("source_type is a required field in ml_application_package")
	}

	if sourceType == "local" {
		// Process the path field
		path, pathExists := packageDetails["path"].(string)
		if !pathExists {
			return fmt.Errorf("path is a required field in ml_application_package")
		}

		localPath := strings.TrimPrefix(path, "file://")
		packageVersion, err := processLocalZip(localPath)
		if err != nil {
			return fmt.Errorf("failed to process package version from path: %w", err)
		}

		// Update path with package version
		updatedPath := fmt.Sprintf("%s::%s", path, packageVersion)
		packageDetails["path"] = updatedPath

		// Update resource diff
		if err := d.SetNew("ml_application_package", packageDetails); err != nil {
			return fmt.Errorf("failed to update ml_application_package: %w", err)
		}
	}

	return nil
}

// Extract the package version from a ZIP file
func processLocalZip(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open ZIP file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}

	zipReader, err := zip.NewReader(file, fileInfo.Size())
	if err != nil {
		return "", fmt.Errorf("failed to create zip reader: %w", err)
	}

	packageVersion, err := getPacakgeVersion(zipReader)
	if err != nil {
		return "", fmt.Errorf("failed to extract package version: %w", err)
	}

	return packageVersion, nil
}

// Extract the package version from descriptor.yaml inside the ZIP
func getPacakgeVersion(zipReader *zip.Reader) (string, error) {
	for _, file := range zipReader.File {
		if file.Name == "descriptor.yaml" {
			rc, err := file.Open()
			if err != nil {
				return "", fmt.Errorf("failed to open descriptor.yaml: %w", err)
			}
			defer rc.Close()

			var data map[string]interface{}
			decoder := yaml.NewDecoder(rc)
			if err := decoder.Decode(&data); err != nil {
				return "", fmt.Errorf("failed to parse YAML: %w", err)
			}

			packageVersion, ok := data["packageVersion"]
			if !ok {
				return "", fmt.Errorf("packageVersion not found in descriptor.yaml")
			}

			switch v := packageVersion.(type) {
			case string:
				return v, nil
			case float64:
				return fmt.Sprintf("%.1f", v), nil
			default:
				return "", fmt.Errorf("invalid packageVersion type: %T", v)
			}
		}
	}
	return "", fmt.Errorf("descriptor.yaml not found in zip")
}

// Retrieve a package from a signed URL
func GetPackage(provider *oci_common.ConfigurationProvider, url string) (io.ReadCloser, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	request.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	signer := oci_common.DefaultRequestSigner(*provider)
	if err = signer.Sign(request); err != nil {
		return nil, 0, fmt.Errorf("failed to sign HTTP request: %w", err)
	}

	httpClient := &http.Client{Timeout: 90 * time.Second}
	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("HTTP request failed: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		return nil, 0, fmt.Errorf("HTTP request returned non-2xx status: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	packageResponseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read response body: %w", err)
	}
	return io.NopCloser(bytes.NewReader(packageResponseBody)), int64(len(packageResponseBody)), nil
}
