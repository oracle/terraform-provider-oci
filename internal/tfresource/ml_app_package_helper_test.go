package tfresource

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestUnitValidateMLApplicationPackage(t *testing.T) {
	cases := []struct {
		name     string
		args     interface{}
		wantErrs []string
	}{
		{
			name: "Valid Package Details with source_type as object_storage_download",
			args: map[string]interface{}{
				"source_type": "object_storage_download",
				"uri":         "https://objectstorage.us-ashburn-1.oraclecloud.com/n/ociodscdev/b/Artifact/o/ml-app-package-1.8.zip",
			},
			wantErrs: nil, // No errors expected
		},
		{
			name: "Valid Package Details with source_type as local",
			args: map[string]interface{}{
				"source_type": "local",
				"path":        "file://./ml-app-package-1.4.zip::1.4",
			},
			wantErrs: nil, // No errors expected
		},
		{
			name: "Invalid Package Details - Incorrect Path",
			args: map[string]interface{}{
				"source_type": "local",
				"path":        "https://test.com",
			},
			wantErrs: []string{"path must start with 'file://', got 'https://test.com'"},
		},
		{
			name: "Invalid Package Details - Path absent",
			args: map[string]interface{}{
				"source_type": "local",
			},
			wantErrs: []string{"path must be specified when source_type is 'local'"},
		},
		{
			name: "Invalid Package Details - Incorrect URI",
			args: map[string]interface{}{
				"source_type": "object_storage_download",
				"uri":         "file://test.com",
			},
			wantErrs: []string{"uri must start with 'https://', got 'file://test.com'"},
		},
		{
			name: "Invalid Package Details - Uri absent",
			args: map[string]interface{}{
				"source_type": "object_storage_download",
			},
			wantErrs: []string{"uri must be specified when source_type is 'object_storage_download'"},
		},
		{
			name:     "Invalid Package Details - Not a map",
			args:     []int{1, 2, 3},
			wantErrs: []string{"ml_application_package must be a map, got []int"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, errs := ValidateMLApplicationPackage(tc.args, "ml_application_package")

			// Check error messages
			if len(errs) != len(tc.wantErrs) {
				t.Fatalf("expected %d errors, got %d", len(tc.wantErrs), len(errs))
			}
			for i, err := range errs {
				if err.Error() != tc.wantErrs[i] {
					t.Errorf("expected error: %q, got: %q", tc.wantErrs[i], err.Error())
				}
			}
		})
	}
}

func TestUnitSupressPackageUpload(t *testing.T) {
	cases := []struct {
		name         string
		keyArgument  string
		old          string
		new          string
		returnStatus bool
	}{
		{
			name:         "SupressPackageUpload Valid use case when path contains some text in the end",
			keyArgument:  "ml_application_package.path",
			old:          "file://./ml-app-package-1.4.zip::1.4 test",
			new:          "file://./ml-app-package-1.4.zip::1.4",
			returnStatus: true, // No errors expected
		},
		{
			name:         "SupressPackageUpload Valid use case without additional text in the end",
			keyArgument:  "ml_application_package.path",
			old:          "file://./ml-app-package-1.4.zip::1.4",
			new:          "file://./ml-app-package-1.4.zip::1.4",
			returnStatus: true, // No errors expected
		},
		{
			name:         "SupressPackageUpload Valid use case for keys except ml_application_package.path",
			keyArgument:  "ml_application_package.source_type",
			old:          "local",
			new:          "local",
			returnStatus: false, // No errors expected
		},
	}

	resourceSchema := map[string]*schema.Schema{
		"ml_application_package": {
			Type:     schema.TypeMap,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}

	resourceData := schema.TestResourceDataRaw(&testing.T{}, resourceSchema, map[string]interface{}{
		"ml_application_package": map[string]interface{}{
			"source_type": "local",
			"path":        "file://./ml-app-package-1.4.zip::1.4",
		},
	},
	)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := SupressPackageUpload(tc.keyArgument, tc.old, tc.new, resourceData)

			// Check error messages
			if res != tc.returnStatus {
				t.Fatalf("expected %t errors, got %t", tc.returnStatus, res)
			}
		})
	}
}

func TestUnitStringInSlice(t *testing.T) {
	tests := []struct {
		name   string
		str    string
		list   []string
		expect bool
	}{
		{
			name:   "String present in list",
			str:    "apple",
			list:   []string{"apple", "banana", "cherry"},
			expect: true,
		},
		{
			name:   "String not present in list",
			str:    "grape",
			list:   []string{"apple", "banana", "cherry"},
			expect: false,
		},
		{
			name:   "Empty list",
			str:    "apple",
			list:   []string{},
			expect: false,
		},
		{
			name:   "Empty string in list",
			str:    "",
			list:   []string{"apple", "", "cherry"},
			expect: true,
		},
		{
			name:   "String is empty and not in list",
			str:    "",
			list:   []string{"apple", "banana", "cherry"},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stringInSlice(tt.str, tt.list)
			if result != tt.expect {
				t.Errorf("stringInSlice(%q, %v) = %v; want %v", tt.str, tt.list, result, tt.expect)
			}
		})
	}
}

func TestUnitProcessLocalZip(t *testing.T) {
	cases := []struct {
		name                  string
		path                  string
		packageVersion        string
		expectedError         string
		descriptorFileContent string
	}{
		{
			name:                  "Valid Package File",
			path:                  "package1.zip",
			packageVersion:        "1.9",
			descriptorFileContent: "descriptorSchemaVersion: 1.0\nmlApplicationVersion: 1.0\npackageVersion: 1.9\n",
			expectedError:         "", // No errors expected
		},
		{
			name:                  "InValid Package File",
			path:                  "package2.zip",
			packageVersion:        "1.9",
			descriptorFileContent: "descriptorSchemaVersion: 1.0\nmlApplicationVersion: 1.0\npackageVersion: 1.9\n",
			expectedError:         "failed to open ZIP file: open package2.zip: no such file or directory",
		},
		{
			name:                  "Valid Package File without package version",
			path:                  "package2.zip",
			packageVersion:        "1.9",
			descriptorFileContent: "descriptorSchemaVersion: 1.0\nmlApplicationVersion: 1.0\n",
			expectedError:         "failed to extract package version: packageVersion not found in descriptor.yaml",
		},
	}

	for _, tc := range cases {
		createZipWithTextFile(tc.path, tc.descriptorFileContent)

		t.Run(tc.name, func(t *testing.T) {
			packageVersion, error := processLocalZip(tc.path)

			// Check error messages
			if error != nil && error.Error() != tc.expectedError {
				t.Fatalf("expected %s errors, got %s", tc.expectedError, error.Error())
			} else if packageVersion != "" && packageVersion != tc.packageVersion {
				t.Fatalf("expected %s package version, got %s", tc.packageVersion, packageVersion)
			}
		})
		os.Remove(tc.path)

	}
}

func TestUnitGetPacakgeVersion(t *testing.T) {
	cases := []struct {
		name                  string
		path                  string
		packageVersion        string
		expectedError         string
		descriptorFileContent string
	}{
		{
			name:                  "Valid Package File",
			path:                  "package1.zip",
			packageVersion:        "1.9",
			descriptorFileContent: "descriptorSchemaVersion: 1.0\nmlApplicationVersion: 1.0\npackageVersion: 1.9\n",
			expectedError:         "", // No errors expected
		},
		{
			name:                  "Valid Package File without package version",
			path:                  "package2.zip",
			packageVersion:        "1.9",
			descriptorFileContent: "descriptorSchemaVersion: 1.0\nmlApplicationVersion: 1.0\n",
			expectedError:         "packageVersion not found in descriptor.yaml",
		},
	}

	for _, tc := range cases {
		createZipWithTextFile(tc.path, tc.descriptorFileContent)
		file, err := os.Open(tc.path)
		if err != nil {
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)

		fileInfo, err := file.Stat()
		if err != nil {
			return
		}

		zipReader, err := zip.NewReader(file, fileInfo.Size())
		if err != nil {
			return
		}

		t.Run(tc.name, func(t *testing.T) {
			packageVersion, error := getPacakgeVersion(zipReader)

			// Check error messages
			if error != nil && error.Error() != tc.expectedError {
				t.Fatalf("expected %s errors, got %s", tc.expectedError, error.Error())
			} else if packageVersion != "" && packageVersion != tc.packageVersion {
				t.Fatalf("expected %s package version, got %s", tc.packageVersion, packageVersion)
			}
		})
		os.Remove(tc.path)
	}
}

func createZipWithTextFile(filename string, descriptorFileContent string) error {
	// Step 1: Create a text file with some content
	textContent := []byte(descriptorFileContent)
	textFileName := "descriptor.yaml"

	// Step 2: Create a buffer to hold the ZIP file data
	var zipBuffer bytes.Buffer
	zipWriter := zip.NewWriter(&zipBuffer)

	// Step 3: Create a new file inside the ZIP archive
	textFileInZip, err := zipWriter.Create(textFileName)
	if err != nil {
		return fmt.Errorf("failed to create file inside zip: %v", err)
	}

	// Step 4: Write the text content to the file inside the zip
	_, err = textFileInZip.Write(textContent)
	if err != nil {
		return fmt.Errorf("failed to write content to zip file: %v", err)
	}

	// Step 5: Close the ZIP writer to flush the data into the buffer
	err = zipWriter.Close()
	if err != nil {
		return fmt.Errorf("failed to close zip writer: %v", err)
	}

	// Step 6: Create the zip file on disk
	outputFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer outputFile.Close()

	// Step 7: Write the zipBuffer content into the final zip file
	_, err = io.Copy(outputFile, &zipBuffer)
	if err != nil {
		return fmt.Errorf("failed to write zip content to file: %v", err)
	}

	fmt.Println("ZIP file created successfully!")
	return nil
}
