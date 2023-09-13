package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/oracle/oci-go-sdk/v65/common"
)

type servicePrincipalDelegationTokenConfigurationProvider struct {
	servicePrincipalConfigurationProvider servicePrincipalConfigurationProvider
	delegationToken                       string
	region                                *common.Region
}
type servicePrincipalDelegationTokenError struct {
	err error
}

func (ipe servicePrincipalDelegationTokenError) Error() string {
	return fmt.Sprintf("%s\nService principals delegation token authentication can only be used on specific OCI services. Please confirm this code is running on the correct environment", ipe.err.Error())
}

// ServicePrincipalDelegationTokenConfigurationProvider returns a configuration for obo token service principals
func ServicePrincipalDelegationTokenConfigurationProvider(delegationToken *string, tenancyID string,
	cert, key []byte, intermediates [][]byte, passphrase []byte) (common.ConfigurationProvider, error) {
	parametersArePresent, err := parametersArePresent(delegationToken, tenancyID, cert, key, intermediates, passphrase)
	if parametersArePresent {
		return newServicePrincipalDelegationTokenConfigurationProvider(delegationToken, nil, tenancyID, "", cert, key, intermediates, passphrase)
	}
	return nil, servicePrincipalDelegationTokenError{err: err}
}

// ServicePrincipalDelegationTokenConfigurationProviderForRegion returns a configuration for obo token service principals with a given region
func ServicePrincipalDelegationTokenConfigurationProviderForRegion(delegationToken *string, tenancyID string,
	cert, key []byte, intermediates [][]byte, passphrase []byte, region common.Region) (common.ConfigurationProvider, error) {
	parametersArePresent, err := parametersArePresent(delegationToken, tenancyID, cert, key, intermediates, passphrase)
	if parametersArePresent {
		return newServicePrincipalDelegationTokenConfigurationProvider(delegationToken, nil, tenancyID, region, cert, key, intermediates, passphrase)
	}
	return nil, servicePrincipalDelegationTokenError{err: err}
}

// ServicePrincipalDelegationTokenConfigurationProviderWithCustomClient returns a configuration for obo token service principals using a modifier function to modify the HTTPRequestDispatcher
func ServicePrincipalDelegationTokenConfigurationProviderWithCustomClient(delegationToken *string, tenancyID string, region common.Region,
	cert, key []byte, intermediates [][]byte, passphrase []byte, modifier func(common.HTTPRequestDispatcher) (common.HTTPRequestDispatcher, error)) (common.ConfigurationProvider, error) {
	parametersArePresent, err := parametersArePresent(delegationToken, tenancyID, cert, key, intermediates, passphrase)
	if parametersArePresent {
		return newServicePrincipalDelegationTokenConfigurationProvider(delegationToken, modifier, tenancyID, region, cert, key, intermediates, passphrase)
	}
	return nil, servicePrincipalDelegationTokenError{err: err}
}

// ServicePrincipalDelegationTokenwithInstancePrincipalConfigurationProvider returns a S2S configuration provider for obo token by acquiring credentials via instance principals
func ServicePrincipalDelegationTokenwithInstancePrincipalConfigurationProvider(delegationToken *string, region common.Region, modifier func(common.HTTPRequestDispatcher) (common.HTTPRequestDispatcher, error)) (common.ConfigurationProvider, error) {
	return newInstancePrincipalDelegationTokenConfigurationProviderWithPurpose(delegationToken, region, nil, servicePrincipalTokenPurpose)
}

// ServicePrincipalDelegationTokenConfigurationProviderWithCerts returns a configuration for obo token service prinicpals with a given region and hardcoded certificates in lieu of metadata service certs
func ServicePrincipalDelegationTokenConfigurationProviderWithCerts(delegationToken *string, region common.Region, leafCertificate, leafPassphrase, leafPrivateKey []byte, intermediateCertificates [][]byte) (common.ConfigurationProvider, error) {
	leafCertificateRetriever := staticCertificateRetriever{Passphrase: leafPassphrase, CertificatePem: leafCertificate, PrivateKeyPem: leafPrivateKey}

	//The .Refresh() call actually reads the certificates from the inputs
	err := leafCertificateRetriever.Refresh()
	if err != nil {
		return nil, servicePrincipalDelegationTokenError{err: err}
	}
	certificate := leafCertificateRetriever.Certificate()
	tenancyID := extractTenancyIDFromCertificate(certificate)
	fedClient, err := newX509FederationClientWithCerts(region, tenancyID, leafCertificate, leafPassphrase, leafPrivateKey, intermediateCertificates, *newDispatcherModifier(nil), "")
	if err != nil {
		return nil, servicePrincipalDelegationTokenError{err: err}
	}
	keyProvider := servicePrincipalKeyProvider{federationClient: fedClient}

	configurationProvider := servicePrincipalConfigurationProvider{keyProvider: &keyProvider, region: string(region), tenancyID: tenancyID}
	return servicePrincipalDelegationTokenConfigurationProvider{configurationProvider, *delegationToken, &region}, err
}

// ServicePrincipalDelegationTokenConfigurationProviderFromHostCerts returns a configuration for obo token service prinicpals,
// given the region and a pathname to the host's service principal certificate directory.
// The pathname generally follows the pattern "/var/certs/hostclass/${hostclass}/${servicePrincipalName}-identity"
func ServicePrincipalDelegationTokenConfigurationProviderFromHostCerts(delegationToken *string, region common.Region, certDir string) (common.ConfigurationProvider, error) {
	if certDir == "" {
		return nil, servicePrincipalDelegationTokenError{err: fmt.Errorf("empty input string")}
	}
	// Read certs from substrate host.
	leafKey, err := ioutil.ReadFile(path.Join(certDir, "key.pem"))
	if err != nil {
		return nil, servicePrincipalDelegationTokenError{err: fmt.Errorf("error reading leafPrivateKey")}
	}
	leafCert, err := ioutil.ReadFile(path.Join(certDir, "cert.pem"))
	if err != nil {
		return nil, servicePrincipalDelegationTokenError{err: fmt.Errorf("error reading leafCertificate")}
	}
	interCert, err := ioutil.ReadFile(path.Join(certDir, "intermediates.pem"))
	if err != nil {
		return nil, servicePrincipalDelegationTokenError{err: fmt.Errorf("error reading intermediateCertificate")}
	}
	var interCerts [][]byte
	interCerts = append(interCerts, interCert)
	var leafPass = []byte("")

	return ServicePrincipalDelegationTokenConfigurationProviderWithCerts(delegationToken, region, leafCert, leafPass, leafKey, interCerts)

}

func newServicePrincipalDelegationTokenConfigurationProvider(delegationToken *string, modifier func(common.HTTPRequestDispatcher) (common.HTTPRequestDispatcher,
	error), tenancyID string, region common.Region, cert, key []byte, intermediates [][]byte, passphrase []byte) (common.ConfigurationProvider, error) {

	keyProvider, err := newServicePrincipalKeyProvider(tenancyID, string(region), cert, key, intermediates, passphrase, modifier)
	if err != nil {
		return nil, servicePrincipalDelegationTokenError{err: fmt.Errorf("failed to create a new key provider for service principal: %s", err.Error())}
	}

	if len(region) > 0 {
		configurationProvider := servicePrincipalConfigurationProvider{keyProvider: keyProvider, region: string(region), tenancyID: tenancyID}
		return servicePrincipalDelegationTokenConfigurationProvider{configurationProvider, *delegationToken, &region}, err
	}

	configurationProvider := servicePrincipalConfigurationProvider{keyProvider: keyProvider, region: "", tenancyID: tenancyID}
	return servicePrincipalDelegationTokenConfigurationProvider{configurationProvider, *delegationToken, nil}, err
}

func (p servicePrincipalDelegationTokenConfigurationProvider) getServicePrincipalDelegationTokenConfigurationProvider() (servicePrincipalDelegationTokenConfigurationProvider, error) {
	return p, nil
}

func (p servicePrincipalDelegationTokenConfigurationProvider) PrivateRSAKey() (*rsa.PrivateKey, error) {
	return p.servicePrincipalConfigurationProvider.keyProvider.PrivateRSAKey()
}

func (p servicePrincipalDelegationTokenConfigurationProvider) KeyID() (string, error) {
	return p.servicePrincipalConfigurationProvider.keyProvider.KeyID()
}

func (p servicePrincipalDelegationTokenConfigurationProvider) TenancyOCID() (string, error) {

	return p.servicePrincipalConfigurationProvider.tenancyID, nil
}

func (p servicePrincipalDelegationTokenConfigurationProvider) UserOCID() (string, error) {
	return "", nil
}

func (p servicePrincipalDelegationTokenConfigurationProvider) KeyFingerprint() (string, error) {
	return "", nil
}

func (p servicePrincipalDelegationTokenConfigurationProvider) Region() (string, error) {
	return string(*p.region), nil
}

func (p servicePrincipalDelegationTokenConfigurationProvider) AuthType() (common.AuthConfig, error) {
	token := p.delegationToken
	return common.AuthConfig{AuthType: common.ServicePrincipalDelegationToken, IsFromConfigFile: false, OboToken: &token}, nil
}

func (p servicePrincipalDelegationTokenConfigurationProvider) Refreshable() bool {
	return true
}

// Check that all necessary parameters are present, report a specified error if not
func parametersArePresent(delegationToken *string, tenancyID string,
	cert, key []byte, intermediates [][]byte, passphrase []byte) (bool, error) {

	if delegationToken == nil || len(*delegationToken) == 0 {
		return false, fmt.Errorf("failed to create a delagationTokenConfigurationProvider: token is a mandatory input parameter")
	} else if len(tenancyID) == 0 {
		return false, fmt.Errorf("failed to create a delagationTokenConfigurationProvider: tenancy ID is a mandatory input parameter")
	} else if cert == nil {
		return false, fmt.Errorf("failed to create a delagationTokenConfigurationProvider: cert is a mandatory input parameter")
	} else if key == nil {
		return false, fmt.Errorf("failed to create a delagationTokenConfigurationProvider: key is a mandatory input parameter")
	} else if intermediates == nil {
		return false, fmt.Errorf("failed to create a delagationTokenConfigurationProvider: intermediates is a mandatory input parameter")
	} else if passphrase == nil {
		return false, fmt.Errorf("failed to create a delagationTokenConfigurationProvider: passphrase is a mandatory input parameter")
	}

	return true, nil
}
