package acctest

import (
	"errors"
	"os"
	"testing"
)

func TestUnitGetCompartmentIDForLegacyTests(t *testing.T) {
	type testFormat struct {
		name     string
		output   string
		mockFunc func()
	}
	tests := []testFormat{
		{
			name:   "Test get compartment_ocid",
			output: "test123",
			mockFunc: func() {
				GetEnvSettingWithDefaultVar = func(s string, dv string) string {
					return "test123"
				}
			},
		},
		{
			name:   "Test get compartment_id_for_create",
			output: "test1234",
			mockFunc: func() {
				GetEnvSettingWithDefaultVar = func(s string, dv string) string {
					return "compartment_ocid"
				}
				GetEnvSettingWithBlankDefaultVar = func(s string) string {
					return "test1234"
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if res := GetCompartmentIDForLegacyTests(); res != test.output {
			t.Errorf("expected string %s not equal to output string %s", test.output, res)
		}
	}
}

func TestUnitWriteConfigFile(t *testing.T) {
	type testFormat struct {
		name     string
		gotError bool
		mockFunc func()
	}
	tests := []testFormat{
		{
			name:     "Test writeConfig pass",
			gotError: false,
			mockFunc: func() {

			},
		},
		{
			name:     "Test os.MkdirAll() returns error",
			gotError: true,
			mockFunc: func() {
				osMkdirAllVar = func(path string, perm os.FileMode) error {
					return errors.New("")
				}
			},
		},
		{
			name:     "Test utils.WriteTempFile() returns error",
			gotError: true,
			mockFunc: func() {
				osMkdirAllVar = os.MkdirAll
				utilsWriteTempFileVar = func(data string, originFileName string) (err error) {
					return errors.New("")
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if _, _, err := WriteConfigFile(); (err != nil) != test.gotError {
			t.Errorf("Output error - %q which is not equal to expected error - %t", err, test.gotError)
		}
	}
}

func TestUnitLegacyTestProviderConfig(t *testing.T) {
	type testFormat struct {
		name     string
		output   string
		mockFunc func()
	}
	tests := []testFormat{
		{
			name: "Test providerConfig match with custom config",
			output: ProviderTestConfig() + `variable "compartment_id" {
		default = "test123"
	}
	`,
			mockFunc: func() {
				GetCompartmentIDForLegacyTests = func() string {
					return "test123"
				}
			},
		},
	}
	for _, test := range tests {
		t.Logf("Running %s", test.name)
		test.mockFunc()
		if res := LegacyTestProviderConfig(); res != test.output {
			t.Errorf("expected string %s not equal to output string %s", test.output, res)
		}
	}
}
