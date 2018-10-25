package objectstore

const (
	awsS3HostFmt  = "https://s3-%s.amazonaws.com"
	googleGCSHost = "https://storage.googleapis.com"
)

// ProviderType enum for different providers
type ProviderType string

const (
	// ProviderTypeGCS captures enum value "GCS"
	ProviderTypeGCS ProviderType = "GCS"
	// ProviderTypeS3 captures enum value "S3"
	ProviderTypeS3 ProviderType = "S3"
	// ProviderTypeAzure captures enum value "Azure"
	ProviderTypeAzure ProviderType = "Azure"
)

// SecretType enum for different providers
type SecretType string

const (
	// SecretTypeAwsAccessKey captures enum value "AwsAccessKey"
	SecretTypeAwsAccessKey SecretType = "AwsAccessKey"
	// SecretTypeGcpServiceAccountKey captures enum value "GcpServiceAccountKey"
	SecretTypeGcpServiceAccountKey SecretType = "GcpServiceAccountKey"
	// SecretTypeAzStorageAccount captures enum value "AzStorageAccount"
	SecretTypeAzStorageAccount SecretType = "AzStorageAccount"
)
