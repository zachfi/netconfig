// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlgorithmSpec string

// Enum values for AlgorithmSpec
const (
	AlgorithmSpecRsaesPkcs1V15   AlgorithmSpec = "RSAES_PKCS1_V1_5"
	AlgorithmSpecRsaesOaepSha1   AlgorithmSpec = "RSAES_OAEP_SHA_1"
	AlgorithmSpecRsaesOaepSha256 AlgorithmSpec = "RSAES_OAEP_SHA_256"
)

// Values returns all known values for AlgorithmSpec. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlgorithmSpec) Values() []AlgorithmSpec {
	return []AlgorithmSpec{
		"RSAES_PKCS1_V1_5",
		"RSAES_OAEP_SHA_1",
		"RSAES_OAEP_SHA_256",
	}
}

type ConnectionErrorCodeType string

// Enum values for ConnectionErrorCodeType
const (
	ConnectionErrorCodeTypeInvalidCredentials       ConnectionErrorCodeType = "INVALID_CREDENTIALS"
	ConnectionErrorCodeTypeClusterNotFound          ConnectionErrorCodeType = "CLUSTER_NOT_FOUND"
	ConnectionErrorCodeTypeNetworkErrors            ConnectionErrorCodeType = "NETWORK_ERRORS"
	ConnectionErrorCodeTypeInternalError            ConnectionErrorCodeType = "INTERNAL_ERROR"
	ConnectionErrorCodeTypeInsufficientCloudhsmHsms ConnectionErrorCodeType = "INSUFFICIENT_CLOUDHSM_HSMS"
	ConnectionErrorCodeTypeUserLockedOut            ConnectionErrorCodeType = "USER_LOCKED_OUT"
	ConnectionErrorCodeTypeUserNotFound             ConnectionErrorCodeType = "USER_NOT_FOUND"
	ConnectionErrorCodeTypeUserLoggedIn             ConnectionErrorCodeType = "USER_LOGGED_IN"
	ConnectionErrorCodeTypeSubnetNotFound           ConnectionErrorCodeType = "SUBNET_NOT_FOUND"
)

// Values returns all known values for ConnectionErrorCodeType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionErrorCodeType) Values() []ConnectionErrorCodeType {
	return []ConnectionErrorCodeType{
		"INVALID_CREDENTIALS",
		"CLUSTER_NOT_FOUND",
		"NETWORK_ERRORS",
		"INTERNAL_ERROR",
		"INSUFFICIENT_CLOUDHSM_HSMS",
		"USER_LOCKED_OUT",
		"USER_NOT_FOUND",
		"USER_LOGGED_IN",
		"SUBNET_NOT_FOUND",
	}
}

type ConnectionStateType string

// Enum values for ConnectionStateType
const (
	ConnectionStateTypeConnected     ConnectionStateType = "CONNECTED"
	ConnectionStateTypeConnecting    ConnectionStateType = "CONNECTING"
	ConnectionStateTypeFailed        ConnectionStateType = "FAILED"
	ConnectionStateTypeDisconnected  ConnectionStateType = "DISCONNECTED"
	ConnectionStateTypeDisconnecting ConnectionStateType = "DISCONNECTING"
)

// Values returns all known values for ConnectionStateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionStateType) Values() []ConnectionStateType {
	return []ConnectionStateType{
		"CONNECTED",
		"CONNECTING",
		"FAILED",
		"DISCONNECTED",
		"DISCONNECTING",
	}
}

type CustomerMasterKeySpec string

// Enum values for CustomerMasterKeySpec
const (
	CustomerMasterKeySpecRsa2048          CustomerMasterKeySpec = "RSA_2048"
	CustomerMasterKeySpecRsa3072          CustomerMasterKeySpec = "RSA_3072"
	CustomerMasterKeySpecRsa4096          CustomerMasterKeySpec = "RSA_4096"
	CustomerMasterKeySpecEccNistP256      CustomerMasterKeySpec = "ECC_NIST_P256"
	CustomerMasterKeySpecEccNistP384      CustomerMasterKeySpec = "ECC_NIST_P384"
	CustomerMasterKeySpecEccNistP521      CustomerMasterKeySpec = "ECC_NIST_P521"
	CustomerMasterKeySpecEccSecgP256k1    CustomerMasterKeySpec = "ECC_SECG_P256K1"
	CustomerMasterKeySpecSymmetricDefault CustomerMasterKeySpec = "SYMMETRIC_DEFAULT"
)

// Values returns all known values for CustomerMasterKeySpec. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CustomerMasterKeySpec) Values() []CustomerMasterKeySpec {
	return []CustomerMasterKeySpec{
		"RSA_2048",
		"RSA_3072",
		"RSA_4096",
		"ECC_NIST_P256",
		"ECC_NIST_P384",
		"ECC_NIST_P521",
		"ECC_SECG_P256K1",
		"SYMMETRIC_DEFAULT",
	}
}

type DataKeyPairSpec string

// Enum values for DataKeyPairSpec
const (
	DataKeyPairSpecRsa2048       DataKeyPairSpec = "RSA_2048"
	DataKeyPairSpecRsa3072       DataKeyPairSpec = "RSA_3072"
	DataKeyPairSpecRsa4096       DataKeyPairSpec = "RSA_4096"
	DataKeyPairSpecEccNistP256   DataKeyPairSpec = "ECC_NIST_P256"
	DataKeyPairSpecEccNistP384   DataKeyPairSpec = "ECC_NIST_P384"
	DataKeyPairSpecEccNistP521   DataKeyPairSpec = "ECC_NIST_P521"
	DataKeyPairSpecEccSecgP256k1 DataKeyPairSpec = "ECC_SECG_P256K1"
)

// Values returns all known values for DataKeyPairSpec. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataKeyPairSpec) Values() []DataKeyPairSpec {
	return []DataKeyPairSpec{
		"RSA_2048",
		"RSA_3072",
		"RSA_4096",
		"ECC_NIST_P256",
		"ECC_NIST_P384",
		"ECC_NIST_P521",
		"ECC_SECG_P256K1",
	}
}

type DataKeySpec string

// Enum values for DataKeySpec
const (
	DataKeySpecAes256 DataKeySpec = "AES_256"
	DataKeySpecAes128 DataKeySpec = "AES_128"
)

// Values returns all known values for DataKeySpec. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DataKeySpec) Values() []DataKeySpec {
	return []DataKeySpec{
		"AES_256",
		"AES_128",
	}
}

type EncryptionAlgorithmSpec string

// Enum values for EncryptionAlgorithmSpec
const (
	EncryptionAlgorithmSpecSymmetricDefault EncryptionAlgorithmSpec = "SYMMETRIC_DEFAULT"
	EncryptionAlgorithmSpecRsaesOaepSha1    EncryptionAlgorithmSpec = "RSAES_OAEP_SHA_1"
	EncryptionAlgorithmSpecRsaesOaepSha256  EncryptionAlgorithmSpec = "RSAES_OAEP_SHA_256"
)

// Values returns all known values for EncryptionAlgorithmSpec. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionAlgorithmSpec) Values() []EncryptionAlgorithmSpec {
	return []EncryptionAlgorithmSpec{
		"SYMMETRIC_DEFAULT",
		"RSAES_OAEP_SHA_1",
		"RSAES_OAEP_SHA_256",
	}
}

type ExpirationModelType string

// Enum values for ExpirationModelType
const (
	ExpirationModelTypeKeyMaterialExpires       ExpirationModelType = "KEY_MATERIAL_EXPIRES"
	ExpirationModelTypeKeyMaterialDoesNotExpire ExpirationModelType = "KEY_MATERIAL_DOES_NOT_EXPIRE"
)

// Values returns all known values for ExpirationModelType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExpirationModelType) Values() []ExpirationModelType {
	return []ExpirationModelType{
		"KEY_MATERIAL_EXPIRES",
		"KEY_MATERIAL_DOES_NOT_EXPIRE",
	}
}

type GrantOperation string

// Enum values for GrantOperation
const (
	GrantOperationDecrypt                             GrantOperation = "Decrypt"
	GrantOperationEncrypt                             GrantOperation = "Encrypt"
	GrantOperationGenerateDataKey                     GrantOperation = "GenerateDataKey"
	GrantOperationGenerateDataKeyWithoutPlaintext     GrantOperation = "GenerateDataKeyWithoutPlaintext"
	GrantOperationReEncryptFrom                       GrantOperation = "ReEncryptFrom"
	GrantOperationReEncryptTo                         GrantOperation = "ReEncryptTo"
	GrantOperationSign                                GrantOperation = "Sign"
	GrantOperationVerify                              GrantOperation = "Verify"
	GrantOperationGetPublicKey                        GrantOperation = "GetPublicKey"
	GrantOperationCreateGrant                         GrantOperation = "CreateGrant"
	GrantOperationRetireGrant                         GrantOperation = "RetireGrant"
	GrantOperationDescribeKey                         GrantOperation = "DescribeKey"
	GrantOperationGenerateDataKeyPair                 GrantOperation = "GenerateDataKeyPair"
	GrantOperationGenerateDataKeyPairWithoutPlaintext GrantOperation = "GenerateDataKeyPairWithoutPlaintext"
)

// Values returns all known values for GrantOperation. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GrantOperation) Values() []GrantOperation {
	return []GrantOperation{
		"Decrypt",
		"Encrypt",
		"GenerateDataKey",
		"GenerateDataKeyWithoutPlaintext",
		"ReEncryptFrom",
		"ReEncryptTo",
		"Sign",
		"Verify",
		"GetPublicKey",
		"CreateGrant",
		"RetireGrant",
		"DescribeKey",
		"GenerateDataKeyPair",
		"GenerateDataKeyPairWithoutPlaintext",
	}
}

type KeyManagerType string

// Enum values for KeyManagerType
const (
	KeyManagerTypeAws      KeyManagerType = "AWS"
	KeyManagerTypeCustomer KeyManagerType = "CUSTOMER"
)

// Values returns all known values for KeyManagerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KeyManagerType) Values() []KeyManagerType {
	return []KeyManagerType{
		"AWS",
		"CUSTOMER",
	}
}

type KeySpec string

// Enum values for KeySpec
const (
	KeySpecRsa2048          KeySpec = "RSA_2048"
	KeySpecRsa3072          KeySpec = "RSA_3072"
	KeySpecRsa4096          KeySpec = "RSA_4096"
	KeySpecEccNistP256      KeySpec = "ECC_NIST_P256"
	KeySpecEccNistP384      KeySpec = "ECC_NIST_P384"
	KeySpecEccNistP521      KeySpec = "ECC_NIST_P521"
	KeySpecEccSecgP256k1    KeySpec = "ECC_SECG_P256K1"
	KeySpecSymmetricDefault KeySpec = "SYMMETRIC_DEFAULT"
)

// Values returns all known values for KeySpec. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (KeySpec) Values() []KeySpec {
	return []KeySpec{
		"RSA_2048",
		"RSA_3072",
		"RSA_4096",
		"ECC_NIST_P256",
		"ECC_NIST_P384",
		"ECC_NIST_P521",
		"ECC_SECG_P256K1",
		"SYMMETRIC_DEFAULT",
	}
}

type KeyState string

// Enum values for KeyState
const (
	KeyStateCreating               KeyState = "Creating"
	KeyStateEnabled                KeyState = "Enabled"
	KeyStateDisabled               KeyState = "Disabled"
	KeyStatePendingDeletion        KeyState = "PendingDeletion"
	KeyStatePendingImport          KeyState = "PendingImport"
	KeyStatePendingReplicaDeletion KeyState = "PendingReplicaDeletion"
	KeyStateUnavailable            KeyState = "Unavailable"
	KeyStateUpdating               KeyState = "Updating"
)

// Values returns all known values for KeyState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (KeyState) Values() []KeyState {
	return []KeyState{
		"Creating",
		"Enabled",
		"Disabled",
		"PendingDeletion",
		"PendingImport",
		"PendingReplicaDeletion",
		"Unavailable",
		"Updating",
	}
}

type KeyUsageType string

// Enum values for KeyUsageType
const (
	KeyUsageTypeSignVerify     KeyUsageType = "SIGN_VERIFY"
	KeyUsageTypeEncryptDecrypt KeyUsageType = "ENCRYPT_DECRYPT"
)

// Values returns all known values for KeyUsageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (KeyUsageType) Values() []KeyUsageType {
	return []KeyUsageType{
		"SIGN_VERIFY",
		"ENCRYPT_DECRYPT",
	}
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeRaw    MessageType = "RAW"
	MessageTypeDigest MessageType = "DIGEST"
)

// Values returns all known values for MessageType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MessageType) Values() []MessageType {
	return []MessageType{
		"RAW",
		"DIGEST",
	}
}

type MultiRegionKeyType string

// Enum values for MultiRegionKeyType
const (
	MultiRegionKeyTypePrimary MultiRegionKeyType = "PRIMARY"
	MultiRegionKeyTypeReplica MultiRegionKeyType = "REPLICA"
)

// Values returns all known values for MultiRegionKeyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MultiRegionKeyType) Values() []MultiRegionKeyType {
	return []MultiRegionKeyType{
		"PRIMARY",
		"REPLICA",
	}
}

type OriginType string

// Enum values for OriginType
const (
	OriginTypeAwsKms      OriginType = "AWS_KMS"
	OriginTypeExternal    OriginType = "EXTERNAL"
	OriginTypeAwsCloudhsm OriginType = "AWS_CLOUDHSM"
)

// Values returns all known values for OriginType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OriginType) Values() []OriginType {
	return []OriginType{
		"AWS_KMS",
		"EXTERNAL",
		"AWS_CLOUDHSM",
	}
}

type SigningAlgorithmSpec string

// Enum values for SigningAlgorithmSpec
const (
	SigningAlgorithmSpecRsassaPssSha256      SigningAlgorithmSpec = "RSASSA_PSS_SHA_256"
	SigningAlgorithmSpecRsassaPssSha384      SigningAlgorithmSpec = "RSASSA_PSS_SHA_384"
	SigningAlgorithmSpecRsassaPssSha512      SigningAlgorithmSpec = "RSASSA_PSS_SHA_512"
	SigningAlgorithmSpecRsassaPkcs1V15Sha256 SigningAlgorithmSpec = "RSASSA_PKCS1_V1_5_SHA_256"
	SigningAlgorithmSpecRsassaPkcs1V15Sha384 SigningAlgorithmSpec = "RSASSA_PKCS1_V1_5_SHA_384"
	SigningAlgorithmSpecRsassaPkcs1V15Sha512 SigningAlgorithmSpec = "RSASSA_PKCS1_V1_5_SHA_512"
	SigningAlgorithmSpecEcdsaSha256          SigningAlgorithmSpec = "ECDSA_SHA_256"
	SigningAlgorithmSpecEcdsaSha384          SigningAlgorithmSpec = "ECDSA_SHA_384"
	SigningAlgorithmSpecEcdsaSha512          SigningAlgorithmSpec = "ECDSA_SHA_512"
)

// Values returns all known values for SigningAlgorithmSpec. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SigningAlgorithmSpec) Values() []SigningAlgorithmSpec {
	return []SigningAlgorithmSpec{
		"RSASSA_PSS_SHA_256",
		"RSASSA_PSS_SHA_384",
		"RSASSA_PSS_SHA_512",
		"RSASSA_PKCS1_V1_5_SHA_256",
		"RSASSA_PKCS1_V1_5_SHA_384",
		"RSASSA_PKCS1_V1_5_SHA_512",
		"ECDSA_SHA_256",
		"ECDSA_SHA_384",
		"ECDSA_SHA_512",
	}
}

type WrappingKeySpec string

// Enum values for WrappingKeySpec
const (
	WrappingKeySpecRsa2048 WrappingKeySpec = "RSA_2048"
)

// Values returns all known values for WrappingKeySpec. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WrappingKeySpec) Values() []WrappingKeySpec {
	return []WrappingKeySpec{
		"RSA_2048",
	}
}
