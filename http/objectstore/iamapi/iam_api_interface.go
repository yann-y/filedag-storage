package iamapi

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"
)

// IAMAPI provides an interface to enable mocking the
// iam.IAM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Identity and Access Management.
//    func myFunc(svc iamiface.IAMAPI) bool {
//        // Make svc.AddClientIDToOpenIDConnectProvider request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iam.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIAMClient struct {
//        iamiface.IAMAPI
//    }
//    func (m *mockIAMClient) AddClientIDToOpenIDConnectProvider(input *iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIAMClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IAMAPI interface {
	AddClientIDToOpenIDConnectProvider(*iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
	AddClientIDToOpenIDConnectProviderWithContext(aws.Context, *iam.AddClientIDToOpenIDConnectProviderInput, ...request.Option) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)
	AddClientIDToOpenIDConnectProviderRequest(*iam.AddClientIDToOpenIDConnectProviderInput) (*request.Request, *iam.AddClientIDToOpenIDConnectProviderOutput)

	AddRoleToInstanceProfile(*iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)
	AddRoleToInstanceProfileWithContext(aws.Context, *iam.AddRoleToInstanceProfileInput, ...request.Option) (*iam.AddRoleToInstanceProfileOutput, error)
	AddRoleToInstanceProfileRequest(*iam.AddRoleToInstanceProfileInput) (*request.Request, *iam.AddRoleToInstanceProfileOutput)

	AddUserToGroup(*iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)
	AddUserToGroupWithContext(aws.Context, *iam.AddUserToGroupInput, ...request.Option) (*iam.AddUserToGroupOutput, error)
	AddUserToGroupRequest(*iam.AddUserToGroupInput) (*request.Request, *iam.AddUserToGroupOutput)

	AttachGroupPolicy(*iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)
	AttachGroupPolicyWithContext(aws.Context, *iam.AttachGroupPolicyInput, ...request.Option) (*iam.AttachGroupPolicyOutput, error)
	AttachGroupPolicyRequest(*iam.AttachGroupPolicyInput) (*request.Request, *iam.AttachGroupPolicyOutput)

	AttachRolePolicy(*iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)
	AttachRolePolicyWithContext(aws.Context, *iam.AttachRolePolicyInput, ...request.Option) (*iam.AttachRolePolicyOutput, error)
	AttachRolePolicyRequest(*iam.AttachRolePolicyInput) (*request.Request, *iam.AttachRolePolicyOutput)

	AttachUserPolicy(*iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)
	AttachUserPolicyWithContext(aws.Context, *iam.AttachUserPolicyInput, ...request.Option) (*iam.AttachUserPolicyOutput, error)
	AttachUserPolicyRequest(*iam.AttachUserPolicyInput) (*request.Request, *iam.AttachUserPolicyOutput)

	ChangePassword(*iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)
	ChangePasswordWithContext(aws.Context, *iam.ChangePasswordInput, ...request.Option) (*iam.ChangePasswordOutput, error)
	ChangePasswordRequest(*iam.ChangePasswordInput) (*request.Request, *iam.ChangePasswordOutput)

	CreateAccessKey(*iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)
	CreateAccessKeyWithContext(aws.Context, *iam.CreateAccessKeyInput, ...request.Option) (*iam.CreateAccessKeyOutput, error)
	CreateAccessKeyRequest(*iam.CreateAccessKeyInput) (*request.Request, *iam.CreateAccessKeyOutput)

	CreateAccountAlias(*iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)
	CreateAccountAliasWithContext(aws.Context, *iam.CreateAccountAliasInput, ...request.Option) (*iam.CreateAccountAliasOutput, error)
	CreateAccountAliasRequest(*iam.CreateAccountAliasInput) (*request.Request, *iam.CreateAccountAliasOutput)

	CreateGroup(*iam.CreateGroupInput) (*iam.CreateGroupOutput, error)
	CreateGroupWithContext(aws.Context, *iam.CreateGroupInput, ...request.Option) (*iam.CreateGroupOutput, error)
	CreateGroupRequest(*iam.CreateGroupInput) (*request.Request, *iam.CreateGroupOutput)

	CreateInstanceProfile(*iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)
	CreateInstanceProfileWithContext(aws.Context, *iam.CreateInstanceProfileInput, ...request.Option) (*iam.CreateInstanceProfileOutput, error)
	CreateInstanceProfileRequest(*iam.CreateInstanceProfileInput) (*request.Request, *iam.CreateInstanceProfileOutput)

	CreateLoginProfile(*iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)
	CreateLoginProfileWithContext(aws.Context, *iam.CreateLoginProfileInput, ...request.Option) (*iam.CreateLoginProfileOutput, error)
	CreateLoginProfileRequest(*iam.CreateLoginProfileInput) (*request.Request, *iam.CreateLoginProfileOutput)

	CreateOpenIDConnectProvider(*iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)
	CreateOpenIDConnectProviderWithContext(aws.Context, *iam.CreateOpenIDConnectProviderInput, ...request.Option) (*iam.CreateOpenIDConnectProviderOutput, error)
	CreateOpenIDConnectProviderRequest(*iam.CreateOpenIDConnectProviderInput) (*request.Request, *iam.CreateOpenIDConnectProviderOutput)

	CreatePolicy(*iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)
	CreatePolicyWithContext(aws.Context, *iam.CreatePolicyInput, ...request.Option) (*iam.CreatePolicyOutput, error)
	CreatePolicyRequest(*iam.CreatePolicyInput) (*request.Request, *iam.CreatePolicyOutput)

	CreatePolicyVersion(*iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)
	CreatePolicyVersionWithContext(aws.Context, *iam.CreatePolicyVersionInput, ...request.Option) (*iam.CreatePolicyVersionOutput, error)
	CreatePolicyVersionRequest(*iam.CreatePolicyVersionInput) (*request.Request, *iam.CreatePolicyVersionOutput)

	CreateRole(*iam.CreateRoleInput) (*iam.CreateRoleOutput, error)
	CreateRoleWithContext(aws.Context, *iam.CreateRoleInput, ...request.Option) (*iam.CreateRoleOutput, error)
	CreateRoleRequest(*iam.CreateRoleInput) (*request.Request, *iam.CreateRoleOutput)

	CreateSAMLProvider(*iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)
	CreateSAMLProviderWithContext(aws.Context, *iam.CreateSAMLProviderInput, ...request.Option) (*iam.CreateSAMLProviderOutput, error)
	CreateSAMLProviderRequest(*iam.CreateSAMLProviderInput) (*request.Request, *iam.CreateSAMLProviderOutput)

	CreateServiceLinkedRole(*iam.CreateServiceLinkedRoleInput) (*iam.CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleWithContext(aws.Context, *iam.CreateServiceLinkedRoleInput, ...request.Option) (*iam.CreateServiceLinkedRoleOutput, error)
	CreateServiceLinkedRoleRequest(*iam.CreateServiceLinkedRoleInput) (*request.Request, *iam.CreateServiceLinkedRoleOutput)

	CreateServiceSpecificCredential(*iam.CreateServiceSpecificCredentialInput) (*iam.CreateServiceSpecificCredentialOutput, error)
	CreateServiceSpecificCredentialWithContext(aws.Context, *iam.CreateServiceSpecificCredentialInput, ...request.Option) (*iam.CreateServiceSpecificCredentialOutput, error)
	CreateServiceSpecificCredentialRequest(*iam.CreateServiceSpecificCredentialInput) (*request.Request, *iam.CreateServiceSpecificCredentialOutput)

	CreateUser(*iam.CreateUserInput) (*iam.CreateUserOutput, error)
	CreateUserWithContext(aws.Context, *iam.CreateUserInput, ...request.Option) (*iam.CreateUserOutput, error)
	CreateUserRequest(*iam.CreateUserInput) (*request.Request, *iam.CreateUserOutput)

	CreateVirtualMFADevice(*iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)
	CreateVirtualMFADeviceWithContext(aws.Context, *iam.CreateVirtualMFADeviceInput, ...request.Option) (*iam.CreateVirtualMFADeviceOutput, error)
	CreateVirtualMFADeviceRequest(*iam.CreateVirtualMFADeviceInput) (*request.Request, *iam.CreateVirtualMFADeviceOutput)

	DeactivateMFADevice(*iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)
	DeactivateMFADeviceWithContext(aws.Context, *iam.DeactivateMFADeviceInput, ...request.Option) (*iam.DeactivateMFADeviceOutput, error)
	DeactivateMFADeviceRequest(*iam.DeactivateMFADeviceInput) (*request.Request, *iam.DeactivateMFADeviceOutput)

	DeleteAccessKey(*iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)
	DeleteAccessKeyWithContext(aws.Context, *iam.DeleteAccessKeyInput, ...request.Option) (*iam.DeleteAccessKeyOutput, error)
	DeleteAccessKeyRequest(*iam.DeleteAccessKeyInput) (*request.Request, *iam.DeleteAccessKeyOutput)

	DeleteAccountAlias(*iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)
	DeleteAccountAliasWithContext(aws.Context, *iam.DeleteAccountAliasInput, ...request.Option) (*iam.DeleteAccountAliasOutput, error)
	DeleteAccountAliasRequest(*iam.DeleteAccountAliasInput) (*request.Request, *iam.DeleteAccountAliasOutput)

	DeleteAccountPasswordPolicy(*iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)
	DeleteAccountPasswordPolicyWithContext(aws.Context, *iam.DeleteAccountPasswordPolicyInput, ...request.Option) (*iam.DeleteAccountPasswordPolicyOutput, error)
	DeleteAccountPasswordPolicyRequest(*iam.DeleteAccountPasswordPolicyInput) (*request.Request, *iam.DeleteAccountPasswordPolicyOutput)

	DeleteGroup(*iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)
	DeleteGroupWithContext(aws.Context, *iam.DeleteGroupInput, ...request.Option) (*iam.DeleteGroupOutput, error)
	DeleteGroupRequest(*iam.DeleteGroupInput) (*request.Request, *iam.DeleteGroupOutput)

	DeleteGroupPolicy(*iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)
	DeleteGroupPolicyWithContext(aws.Context, *iam.DeleteGroupPolicyInput, ...request.Option) (*iam.DeleteGroupPolicyOutput, error)
	DeleteGroupPolicyRequest(*iam.DeleteGroupPolicyInput) (*request.Request, *iam.DeleteGroupPolicyOutput)

	DeleteInstanceProfile(*iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileWithContext(aws.Context, *iam.DeleteInstanceProfileInput, ...request.Option) (*iam.DeleteInstanceProfileOutput, error)
	DeleteInstanceProfileRequest(*iam.DeleteInstanceProfileInput) (*request.Request, *iam.DeleteInstanceProfileOutput)

	DeleteLoginProfile(*iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)
	DeleteLoginProfileWithContext(aws.Context, *iam.DeleteLoginProfileInput, ...request.Option) (*iam.DeleteLoginProfileOutput, error)
	DeleteLoginProfileRequest(*iam.DeleteLoginProfileInput) (*request.Request, *iam.DeleteLoginProfileOutput)

	DeleteOpenIDConnectProvider(*iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)
	DeleteOpenIDConnectProviderWithContext(aws.Context, *iam.DeleteOpenIDConnectProviderInput, ...request.Option) (*iam.DeleteOpenIDConnectProviderOutput, error)
	DeleteOpenIDConnectProviderRequest(*iam.DeleteOpenIDConnectProviderInput) (*request.Request, *iam.DeleteOpenIDConnectProviderOutput)

	DeletePolicy(*iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)
	DeletePolicyWithContext(aws.Context, *iam.DeletePolicyInput, ...request.Option) (*iam.DeletePolicyOutput, error)
	DeletePolicyRequest(*iam.DeletePolicyInput) (*request.Request, *iam.DeletePolicyOutput)

	DeletePolicyVersion(*iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)
	DeletePolicyVersionWithContext(aws.Context, *iam.DeletePolicyVersionInput, ...request.Option) (*iam.DeletePolicyVersionOutput, error)
	DeletePolicyVersionRequest(*iam.DeletePolicyVersionInput) (*request.Request, *iam.DeletePolicyVersionOutput)

	DeleteRole(*iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)
	DeleteRoleWithContext(aws.Context, *iam.DeleteRoleInput, ...request.Option) (*iam.DeleteRoleOutput, error)
	DeleteRoleRequest(*iam.DeleteRoleInput) (*request.Request, *iam.DeleteRoleOutput)

	DeleteRolePermissionsBoundary(*iam.DeleteRolePermissionsBoundaryInput) (*iam.DeleteRolePermissionsBoundaryOutput, error)
	DeleteRolePermissionsBoundaryWithContext(aws.Context, *iam.DeleteRolePermissionsBoundaryInput, ...request.Option) (*iam.DeleteRolePermissionsBoundaryOutput, error)
	DeleteRolePermissionsBoundaryRequest(*iam.DeleteRolePermissionsBoundaryInput) (*request.Request, *iam.DeleteRolePermissionsBoundaryOutput)

	DeleteRolePolicy(*iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)
	DeleteRolePolicyWithContext(aws.Context, *iam.DeleteRolePolicyInput, ...request.Option) (*iam.DeleteRolePolicyOutput, error)
	DeleteRolePolicyRequest(*iam.DeleteRolePolicyInput) (*request.Request, *iam.DeleteRolePolicyOutput)

	DeleteSAMLProvider(*iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderWithContext(aws.Context, *iam.DeleteSAMLProviderInput, ...request.Option) (*iam.DeleteSAMLProviderOutput, error)
	DeleteSAMLProviderRequest(*iam.DeleteSAMLProviderInput) (*request.Request, *iam.DeleteSAMLProviderOutput)

	DeleteSSHPublicKey(*iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error)
	DeleteSSHPublicKeyWithContext(aws.Context, *iam.DeleteSSHPublicKeyInput, ...request.Option) (*iam.DeleteSSHPublicKeyOutput, error)
	DeleteSSHPublicKeyRequest(*iam.DeleteSSHPublicKeyInput) (*request.Request, *iam.DeleteSSHPublicKeyOutput)

	DeleteServerCertificate(*iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)
	DeleteServerCertificateWithContext(aws.Context, *iam.DeleteServerCertificateInput, ...request.Option) (*iam.DeleteServerCertificateOutput, error)
	DeleteServerCertificateRequest(*iam.DeleteServerCertificateInput) (*request.Request, *iam.DeleteServerCertificateOutput)

	DeleteServiceLinkedRole(*iam.DeleteServiceLinkedRoleInput) (*iam.DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleWithContext(aws.Context, *iam.DeleteServiceLinkedRoleInput, ...request.Option) (*iam.DeleteServiceLinkedRoleOutput, error)
	DeleteServiceLinkedRoleRequest(*iam.DeleteServiceLinkedRoleInput) (*request.Request, *iam.DeleteServiceLinkedRoleOutput)

	DeleteServiceSpecificCredential(*iam.DeleteServiceSpecificCredentialInput) (*iam.DeleteServiceSpecificCredentialOutput, error)
	DeleteServiceSpecificCredentialWithContext(aws.Context, *iam.DeleteServiceSpecificCredentialInput, ...request.Option) (*iam.DeleteServiceSpecificCredentialOutput, error)
	DeleteServiceSpecificCredentialRequest(*iam.DeleteServiceSpecificCredentialInput) (*request.Request, *iam.DeleteServiceSpecificCredentialOutput)

	DeleteSigningCertificate(*iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)
	DeleteSigningCertificateWithContext(aws.Context, *iam.DeleteSigningCertificateInput, ...request.Option) (*iam.DeleteSigningCertificateOutput, error)
	DeleteSigningCertificateRequest(*iam.DeleteSigningCertificateInput) (*request.Request, *iam.DeleteSigningCertificateOutput)

	DeleteUser(*iam.DeleteUserInput) (*iam.DeleteUserOutput, error)
	DeleteUserWithContext(aws.Context, *iam.DeleteUserInput, ...request.Option) (*iam.DeleteUserOutput, error)
	DeleteUserRequest(*iam.DeleteUserInput) (*request.Request, *iam.DeleteUserOutput)

	DeleteUserPermissionsBoundary(*iam.DeleteUserPermissionsBoundaryInput) (*iam.DeleteUserPermissionsBoundaryOutput, error)
	DeleteUserPermissionsBoundaryWithContext(aws.Context, *iam.DeleteUserPermissionsBoundaryInput, ...request.Option) (*iam.DeleteUserPermissionsBoundaryOutput, error)
	DeleteUserPermissionsBoundaryRequest(*iam.DeleteUserPermissionsBoundaryInput) (*request.Request, *iam.DeleteUserPermissionsBoundaryOutput)

	DeleteUserPolicy(*iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)
	DeleteUserPolicyWithContext(aws.Context, *iam.DeleteUserPolicyInput, ...request.Option) (*iam.DeleteUserPolicyOutput, error)
	DeleteUserPolicyRequest(*iam.DeleteUserPolicyInput) (*request.Request, *iam.DeleteUserPolicyOutput)

	DeleteVirtualMFADevice(*iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)
	DeleteVirtualMFADeviceWithContext(aws.Context, *iam.DeleteVirtualMFADeviceInput, ...request.Option) (*iam.DeleteVirtualMFADeviceOutput, error)
	DeleteVirtualMFADeviceRequest(*iam.DeleteVirtualMFADeviceInput) (*request.Request, *iam.DeleteVirtualMFADeviceOutput)

	DetachGroupPolicy(*iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)
	DetachGroupPolicyWithContext(aws.Context, *iam.DetachGroupPolicyInput, ...request.Option) (*iam.DetachGroupPolicyOutput, error)
	DetachGroupPolicyRequest(*iam.DetachGroupPolicyInput) (*request.Request, *iam.DetachGroupPolicyOutput)

	DetachRolePolicy(*iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)
	DetachRolePolicyWithContext(aws.Context, *iam.DetachRolePolicyInput, ...request.Option) (*iam.DetachRolePolicyOutput, error)
	DetachRolePolicyRequest(*iam.DetachRolePolicyInput) (*request.Request, *iam.DetachRolePolicyOutput)

	DetachUserPolicy(*iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)
	DetachUserPolicyWithContext(aws.Context, *iam.DetachUserPolicyInput, ...request.Option) (*iam.DetachUserPolicyOutput, error)
	DetachUserPolicyRequest(*iam.DetachUserPolicyInput) (*request.Request, *iam.DetachUserPolicyOutput)

	EnableMFADevice(*iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)
	EnableMFADeviceWithContext(aws.Context, *iam.EnableMFADeviceInput, ...request.Option) (*iam.EnableMFADeviceOutput, error)
	EnableMFADeviceRequest(*iam.EnableMFADeviceInput) (*request.Request, *iam.EnableMFADeviceOutput)

	GenerateCredentialReport(*iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)
	GenerateCredentialReportWithContext(aws.Context, *iam.GenerateCredentialReportInput, ...request.Option) (*iam.GenerateCredentialReportOutput, error)
	GenerateCredentialReportRequest(*iam.GenerateCredentialReportInput) (*request.Request, *iam.GenerateCredentialReportOutput)

	GenerateOrganizationsAccessReport(*iam.GenerateOrganizationsAccessReportInput) (*iam.GenerateOrganizationsAccessReportOutput, error)
	GenerateOrganizationsAccessReportWithContext(aws.Context, *iam.GenerateOrganizationsAccessReportInput, ...request.Option) (*iam.GenerateOrganizationsAccessReportOutput, error)
	GenerateOrganizationsAccessReportRequest(*iam.GenerateOrganizationsAccessReportInput) (*request.Request, *iam.GenerateOrganizationsAccessReportOutput)

	GenerateServiceLastAccessedDetails(*iam.GenerateServiceLastAccessedDetailsInput) (*iam.GenerateServiceLastAccessedDetailsOutput, error)
	GenerateServiceLastAccessedDetailsWithContext(aws.Context, *iam.GenerateServiceLastAccessedDetailsInput, ...request.Option) (*iam.GenerateServiceLastAccessedDetailsOutput, error)
	GenerateServiceLastAccessedDetailsRequest(*iam.GenerateServiceLastAccessedDetailsInput) (*request.Request, *iam.GenerateServiceLastAccessedDetailsOutput)

	GetAccessKeyLastUsed(*iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedWithContext(aws.Context, *iam.GetAccessKeyLastUsedInput, ...request.Option) (*iam.GetAccessKeyLastUsedOutput, error)
	GetAccessKeyLastUsedRequest(*iam.GetAccessKeyLastUsedInput) (*request.Request, *iam.GetAccessKeyLastUsedOutput)

	GetAccountAuthorizationDetails(*iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountAuthorizationDetailsWithContext(aws.Context, *iam.GetAccountAuthorizationDetailsInput, ...request.Option) (*iam.GetAccountAuthorizationDetailsOutput, error)
	GetAccountAuthorizationDetailsRequest(*iam.GetAccountAuthorizationDetailsInput) (*request.Request, *iam.GetAccountAuthorizationDetailsOutput)

	GetAccountAuthorizationDetailsPages(*iam.GetAccountAuthorizationDetailsInput, func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool) error
	GetAccountAuthorizationDetailsPagesWithContext(aws.Context, *iam.GetAccountAuthorizationDetailsInput, func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool, ...request.Option) error

	GetAccountPasswordPolicy(*iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountPasswordPolicyWithContext(aws.Context, *iam.GetAccountPasswordPolicyInput, ...request.Option) (*iam.GetAccountPasswordPolicyOutput, error)
	GetAccountPasswordPolicyRequest(*iam.GetAccountPasswordPolicyInput) (*request.Request, *iam.GetAccountPasswordPolicyOutput)

	GetAccountSummary(*iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)
	GetAccountSummaryWithContext(aws.Context, *iam.GetAccountSummaryInput, ...request.Option) (*iam.GetAccountSummaryOutput, error)
	GetAccountSummaryRequest(*iam.GetAccountSummaryInput) (*request.Request, *iam.GetAccountSummaryOutput)

	GetContextKeysForCustomPolicy(*iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForCustomPolicyWithContext(aws.Context, *iam.GetContextKeysForCustomPolicyInput, ...request.Option) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForCustomPolicyRequest(*iam.GetContextKeysForCustomPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse)

	GetContextKeysForPrincipalPolicy(*iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForPrincipalPolicyWithContext(aws.Context, *iam.GetContextKeysForPrincipalPolicyInput, ...request.Option) (*iam.GetContextKeysForPolicyResponse, error)
	GetContextKeysForPrincipalPolicyRequest(*iam.GetContextKeysForPrincipalPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse)

	GetCredentialReport(*iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)
	GetCredentialReportWithContext(aws.Context, *iam.GetCredentialReportInput, ...request.Option) (*iam.GetCredentialReportOutput, error)
	GetCredentialReportRequest(*iam.GetCredentialReportInput) (*request.Request, *iam.GetCredentialReportOutput)

	GetGroup(*iam.GetGroupInput) (*iam.GetGroupOutput, error)
	GetGroupWithContext(aws.Context, *iam.GetGroupInput, ...request.Option) (*iam.GetGroupOutput, error)
	GetGroupRequest(*iam.GetGroupInput) (*request.Request, *iam.GetGroupOutput)

	GetGroupPages(*iam.GetGroupInput, func(*iam.GetGroupOutput, bool) bool) error
	GetGroupPagesWithContext(aws.Context, *iam.GetGroupInput, func(*iam.GetGroupOutput, bool) bool, ...request.Option) error

	GetGroupPolicy(*iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)
	GetGroupPolicyWithContext(aws.Context, *iam.GetGroupPolicyInput, ...request.Option) (*iam.GetGroupPolicyOutput, error)
	GetGroupPolicyRequest(*iam.GetGroupPolicyInput) (*request.Request, *iam.GetGroupPolicyOutput)

	GetInstanceProfile(*iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)
	GetInstanceProfileWithContext(aws.Context, *iam.GetInstanceProfileInput, ...request.Option) (*iam.GetInstanceProfileOutput, error)
	GetInstanceProfileRequest(*iam.GetInstanceProfileInput) (*request.Request, *iam.GetInstanceProfileOutput)

	GetLoginProfile(*iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)
	GetLoginProfileWithContext(aws.Context, *iam.GetLoginProfileInput, ...request.Option) (*iam.GetLoginProfileOutput, error)
	GetLoginProfileRequest(*iam.GetLoginProfileInput) (*request.Request, *iam.GetLoginProfileOutput)

	GetOpenIDConnectProvider(*iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)
	GetOpenIDConnectProviderWithContext(aws.Context, *iam.GetOpenIDConnectProviderInput, ...request.Option) (*iam.GetOpenIDConnectProviderOutput, error)
	GetOpenIDConnectProviderRequest(*iam.GetOpenIDConnectProviderInput) (*request.Request, *iam.GetOpenIDConnectProviderOutput)

	GetOrganizationsAccessReport(*iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error)
	GetOrganizationsAccessReportWithContext(aws.Context, *iam.GetOrganizationsAccessReportInput, ...request.Option) (*iam.GetOrganizationsAccessReportOutput, error)
	GetOrganizationsAccessReportRequest(*iam.GetOrganizationsAccessReportInput) (*request.Request, *iam.GetOrganizationsAccessReportOutput)

	GetPolicy(*iam.GetPolicyInput) (*iam.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *iam.GetPolicyInput, ...request.Option) (*iam.GetPolicyOutput, error)
	GetPolicyRequest(*iam.GetPolicyInput) (*request.Request, *iam.GetPolicyOutput)

	GetPolicyVersion(*iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)
	GetPolicyVersionWithContext(aws.Context, *iam.GetPolicyVersionInput, ...request.Option) (*iam.GetPolicyVersionOutput, error)
	GetPolicyVersionRequest(*iam.GetPolicyVersionInput) (*request.Request, *iam.GetPolicyVersionOutput)

	GetRole(*iam.GetRoleInput) (*iam.GetRoleOutput, error)
	GetRoleWithContext(aws.Context, *iam.GetRoleInput, ...request.Option) (*iam.GetRoleOutput, error)
	GetRoleRequest(*iam.GetRoleInput) (*request.Request, *iam.GetRoleOutput)

	GetRolePolicy(*iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)
	GetRolePolicyWithContext(aws.Context, *iam.GetRolePolicyInput, ...request.Option) (*iam.GetRolePolicyOutput, error)
	GetRolePolicyRequest(*iam.GetRolePolicyInput) (*request.Request, *iam.GetRolePolicyOutput)

	GetSAMLProvider(*iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)
	GetSAMLProviderWithContext(aws.Context, *iam.GetSAMLProviderInput, ...request.Option) (*iam.GetSAMLProviderOutput, error)
	GetSAMLProviderRequest(*iam.GetSAMLProviderInput) (*request.Request, *iam.GetSAMLProviderOutput)

	GetSSHPublicKey(*iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error)
	GetSSHPublicKeyWithContext(aws.Context, *iam.GetSSHPublicKeyInput, ...request.Option) (*iam.GetSSHPublicKeyOutput, error)
	GetSSHPublicKeyRequest(*iam.GetSSHPublicKeyInput) (*request.Request, *iam.GetSSHPublicKeyOutput)

	GetServerCertificate(*iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)
	GetServerCertificateWithContext(aws.Context, *iam.GetServerCertificateInput, ...request.Option) (*iam.GetServerCertificateOutput, error)
	GetServerCertificateRequest(*iam.GetServerCertificateInput) (*request.Request, *iam.GetServerCertificateOutput)

	GetServiceLastAccessedDetails(*iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error)
	GetServiceLastAccessedDetailsWithContext(aws.Context, *iam.GetServiceLastAccessedDetailsInput, ...request.Option) (*iam.GetServiceLastAccessedDetailsOutput, error)
	GetServiceLastAccessedDetailsRequest(*iam.GetServiceLastAccessedDetailsInput) (*request.Request, *iam.GetServiceLastAccessedDetailsOutput)

	GetServiceLastAccessedDetailsWithEntities(*iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error)
	GetServiceLastAccessedDetailsWithEntitiesWithContext(aws.Context, *iam.GetServiceLastAccessedDetailsWithEntitiesInput, ...request.Option) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error)
	GetServiceLastAccessedDetailsWithEntitiesRequest(*iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*request.Request, *iam.GetServiceLastAccessedDetailsWithEntitiesOutput)

	GetServiceLinkedRoleDeletionStatus(*iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetServiceLinkedRoleDeletionStatusWithContext(aws.Context, *iam.GetServiceLinkedRoleDeletionStatusInput, ...request.Option) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error)
	GetServiceLinkedRoleDeletionStatusRequest(*iam.GetServiceLinkedRoleDeletionStatusInput) (*request.Request, *iam.GetServiceLinkedRoleDeletionStatusOutput)

	GetUser(*iam.GetUserInput) (*iam.GetUserOutput, error)
	GetUserWithContext(aws.Context, *iam.GetUserInput, ...request.Option) (*iam.GetUserOutput, error)
	GetUserRequest(*iam.GetUserInput) (*request.Request, *iam.GetUserOutput)

	GetUserPolicy(*iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)
	GetUserPolicyWithContext(aws.Context, *iam.GetUserPolicyInput, ...request.Option) (*iam.GetUserPolicyOutput, error)
	GetUserPolicyRequest(*iam.GetUserPolicyInput) (*request.Request, *iam.GetUserPolicyOutput)

	ListAccessKeys(*iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)
	ListAccessKeysWithContext(aws.Context, *iam.ListAccessKeysInput, ...request.Option) (*iam.ListAccessKeysOutput, error)
	ListAccessKeysRequest(*iam.ListAccessKeysInput) (*request.Request, *iam.ListAccessKeysOutput)

	ListAccessKeysPages(*iam.ListAccessKeysInput, func(*iam.ListAccessKeysOutput, bool) bool) error
	ListAccessKeysPagesWithContext(aws.Context, *iam.ListAccessKeysInput, func(*iam.ListAccessKeysOutput, bool) bool, ...request.Option) error

	ListAccountAliases(*iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)
	ListAccountAliasesWithContext(aws.Context, *iam.ListAccountAliasesInput, ...request.Option) (*iam.ListAccountAliasesOutput, error)
	ListAccountAliasesRequest(*iam.ListAccountAliasesInput) (*request.Request, *iam.ListAccountAliasesOutput)

	ListAccountAliasesPages(*iam.ListAccountAliasesInput, func(*iam.ListAccountAliasesOutput, bool) bool) error
	ListAccountAliasesPagesWithContext(aws.Context, *iam.ListAccountAliasesInput, func(*iam.ListAccountAliasesOutput, bool) bool, ...request.Option) error

	ListAttachedGroupPolicies(*iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedGroupPoliciesWithContext(aws.Context, *iam.ListAttachedGroupPoliciesInput, ...request.Option) (*iam.ListAttachedGroupPoliciesOutput, error)
	ListAttachedGroupPoliciesRequest(*iam.ListAttachedGroupPoliciesInput) (*request.Request, *iam.ListAttachedGroupPoliciesOutput)

	ListAttachedGroupPoliciesPages(*iam.ListAttachedGroupPoliciesInput, func(*iam.ListAttachedGroupPoliciesOutput, bool) bool) error
	ListAttachedGroupPoliciesPagesWithContext(aws.Context, *iam.ListAttachedGroupPoliciesInput, func(*iam.ListAttachedGroupPoliciesOutput, bool) bool, ...request.Option) error

	ListAttachedRolePolicies(*iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesWithContext(aws.Context, *iam.ListAttachedRolePoliciesInput, ...request.Option) (*iam.ListAttachedRolePoliciesOutput, error)
	ListAttachedRolePoliciesRequest(*iam.ListAttachedRolePoliciesInput) (*request.Request, *iam.ListAttachedRolePoliciesOutput)

	ListAttachedRolePoliciesPages(*iam.ListAttachedRolePoliciesInput, func(*iam.ListAttachedRolePoliciesOutput, bool) bool) error
	ListAttachedRolePoliciesPagesWithContext(aws.Context, *iam.ListAttachedRolePoliciesInput, func(*iam.ListAttachedRolePoliciesOutput, bool) bool, ...request.Option) error

	ListAttachedUserPolicies(*iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesWithContext(aws.Context, *iam.ListAttachedUserPoliciesInput, ...request.Option) (*iam.ListAttachedUserPoliciesOutput, error)
	ListAttachedUserPoliciesRequest(*iam.ListAttachedUserPoliciesInput) (*request.Request, *iam.ListAttachedUserPoliciesOutput)

	ListAttachedUserPoliciesPages(*iam.ListAttachedUserPoliciesInput, func(*iam.ListAttachedUserPoliciesOutput, bool) bool) error
	ListAttachedUserPoliciesPagesWithContext(aws.Context, *iam.ListAttachedUserPoliciesInput, func(*iam.ListAttachedUserPoliciesOutput, bool) bool, ...request.Option) error

	ListEntitiesForPolicy(*iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyWithContext(aws.Context, *iam.ListEntitiesForPolicyInput, ...request.Option) (*iam.ListEntitiesForPolicyOutput, error)
	ListEntitiesForPolicyRequest(*iam.ListEntitiesForPolicyInput) (*request.Request, *iam.ListEntitiesForPolicyOutput)

	ListEntitiesForPolicyPages(*iam.ListEntitiesForPolicyInput, func(*iam.ListEntitiesForPolicyOutput, bool) bool) error
	ListEntitiesForPolicyPagesWithContext(aws.Context, *iam.ListEntitiesForPolicyInput, func(*iam.ListEntitiesForPolicyOutput, bool) bool, ...request.Option) error

	ListGroupPolicies(*iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)
	ListGroupPoliciesWithContext(aws.Context, *iam.ListGroupPoliciesInput, ...request.Option) (*iam.ListGroupPoliciesOutput, error)
	ListGroupPoliciesRequest(*iam.ListGroupPoliciesInput) (*request.Request, *iam.ListGroupPoliciesOutput)

	ListGroupPoliciesPages(*iam.ListGroupPoliciesInput, func(*iam.ListGroupPoliciesOutput, bool) bool) error
	ListGroupPoliciesPagesWithContext(aws.Context, *iam.ListGroupPoliciesInput, func(*iam.ListGroupPoliciesOutput, bool) bool, ...request.Option) error

	ListGroups(*iam.ListGroupsInput) (*iam.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *iam.ListGroupsInput, ...request.Option) (*iam.ListGroupsOutput, error)
	ListGroupsRequest(*iam.ListGroupsInput) (*request.Request, *iam.ListGroupsOutput)

	ListGroupsPages(*iam.ListGroupsInput, func(*iam.ListGroupsOutput, bool) bool) error
	ListGroupsPagesWithContext(aws.Context, *iam.ListGroupsInput, func(*iam.ListGroupsOutput, bool) bool, ...request.Option) error

	ListGroupsForUser(*iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)
	ListGroupsForUserWithContext(aws.Context, *iam.ListGroupsForUserInput, ...request.Option) (*iam.ListGroupsForUserOutput, error)
	ListGroupsForUserRequest(*iam.ListGroupsForUserInput) (*request.Request, *iam.ListGroupsForUserOutput)

	ListGroupsForUserPages(*iam.ListGroupsForUserInput, func(*iam.ListGroupsForUserOutput, bool) bool) error
	ListGroupsForUserPagesWithContext(aws.Context, *iam.ListGroupsForUserInput, func(*iam.ListGroupsForUserOutput, bool) bool, ...request.Option) error

	ListInstanceProfileTags(*iam.ListInstanceProfileTagsInput) (*iam.ListInstanceProfileTagsOutput, error)
	ListInstanceProfileTagsWithContext(aws.Context, *iam.ListInstanceProfileTagsInput, ...request.Option) (*iam.ListInstanceProfileTagsOutput, error)
	ListInstanceProfileTagsRequest(*iam.ListInstanceProfileTagsInput) (*request.Request, *iam.ListInstanceProfileTagsOutput)

	ListInstanceProfiles(*iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)
	ListInstanceProfilesWithContext(aws.Context, *iam.ListInstanceProfilesInput, ...request.Option) (*iam.ListInstanceProfilesOutput, error)
	ListInstanceProfilesRequest(*iam.ListInstanceProfilesInput) (*request.Request, *iam.ListInstanceProfilesOutput)

	ListInstanceProfilesPages(*iam.ListInstanceProfilesInput, func(*iam.ListInstanceProfilesOutput, bool) bool) error
	ListInstanceProfilesPagesWithContext(aws.Context, *iam.ListInstanceProfilesInput, func(*iam.ListInstanceProfilesOutput, bool) bool, ...request.Option) error

	ListInstanceProfilesForRole(*iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListInstanceProfilesForRoleWithContext(aws.Context, *iam.ListInstanceProfilesForRoleInput, ...request.Option) (*iam.ListInstanceProfilesForRoleOutput, error)
	ListInstanceProfilesForRoleRequest(*iam.ListInstanceProfilesForRoleInput) (*request.Request, *iam.ListInstanceProfilesForRoleOutput)

	ListInstanceProfilesForRolePages(*iam.ListInstanceProfilesForRoleInput, func(*iam.ListInstanceProfilesForRoleOutput, bool) bool) error
	ListInstanceProfilesForRolePagesWithContext(aws.Context, *iam.ListInstanceProfilesForRoleInput, func(*iam.ListInstanceProfilesForRoleOutput, bool) bool, ...request.Option) error

	ListMFADeviceTags(*iam.ListMFADeviceTagsInput) (*iam.ListMFADeviceTagsOutput, error)
	ListMFADeviceTagsWithContext(aws.Context, *iam.ListMFADeviceTagsInput, ...request.Option) (*iam.ListMFADeviceTagsOutput, error)
	ListMFADeviceTagsRequest(*iam.ListMFADeviceTagsInput) (*request.Request, *iam.ListMFADeviceTagsOutput)

	ListMFADevices(*iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)
	ListMFADevicesWithContext(aws.Context, *iam.ListMFADevicesInput, ...request.Option) (*iam.ListMFADevicesOutput, error)
	ListMFADevicesRequest(*iam.ListMFADevicesInput) (*request.Request, *iam.ListMFADevicesOutput)

	ListMFADevicesPages(*iam.ListMFADevicesInput, func(*iam.ListMFADevicesOutput, bool) bool) error
	ListMFADevicesPagesWithContext(aws.Context, *iam.ListMFADevicesInput, func(*iam.ListMFADevicesOutput, bool) bool, ...request.Option) error

	ListOpenIDConnectProviderTags(*iam.ListOpenIDConnectProviderTagsInput) (*iam.ListOpenIDConnectProviderTagsOutput, error)
	ListOpenIDConnectProviderTagsWithContext(aws.Context, *iam.ListOpenIDConnectProviderTagsInput, ...request.Option) (*iam.ListOpenIDConnectProviderTagsOutput, error)
	ListOpenIDConnectProviderTagsRequest(*iam.ListOpenIDConnectProviderTagsInput) (*request.Request, *iam.ListOpenIDConnectProviderTagsOutput)

	ListOpenIDConnectProviders(*iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListOpenIDConnectProvidersWithContext(aws.Context, *iam.ListOpenIDConnectProvidersInput, ...request.Option) (*iam.ListOpenIDConnectProvidersOutput, error)
	ListOpenIDConnectProvidersRequest(*iam.ListOpenIDConnectProvidersInput) (*request.Request, *iam.ListOpenIDConnectProvidersOutput)

	ListPolicies(*iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)
	ListPoliciesWithContext(aws.Context, *iam.ListPoliciesInput, ...request.Option) (*iam.ListPoliciesOutput, error)
	ListPoliciesRequest(*iam.ListPoliciesInput) (*request.Request, *iam.ListPoliciesOutput)

	ListPoliciesPages(*iam.ListPoliciesInput, func(*iam.ListPoliciesOutput, bool) bool) error
	ListPoliciesPagesWithContext(aws.Context, *iam.ListPoliciesInput, func(*iam.ListPoliciesOutput, bool) bool, ...request.Option) error

	ListPoliciesGrantingServiceAccess(*iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error)
	ListPoliciesGrantingServiceAccessWithContext(aws.Context, *iam.ListPoliciesGrantingServiceAccessInput, ...request.Option) (*iam.ListPoliciesGrantingServiceAccessOutput, error)
	ListPoliciesGrantingServiceAccessRequest(*iam.ListPoliciesGrantingServiceAccessInput) (*request.Request, *iam.ListPoliciesGrantingServiceAccessOutput)

	ListPolicyTags(*iam.ListPolicyTagsInput) (*iam.ListPolicyTagsOutput, error)
	ListPolicyTagsWithContext(aws.Context, *iam.ListPolicyTagsInput, ...request.Option) (*iam.ListPolicyTagsOutput, error)
	ListPolicyTagsRequest(*iam.ListPolicyTagsInput) (*request.Request, *iam.ListPolicyTagsOutput)

	ListPolicyVersions(*iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)
	ListPolicyVersionsWithContext(aws.Context, *iam.ListPolicyVersionsInput, ...request.Option) (*iam.ListPolicyVersionsOutput, error)
	ListPolicyVersionsRequest(*iam.ListPolicyVersionsInput) (*request.Request, *iam.ListPolicyVersionsOutput)

	ListPolicyVersionsPages(*iam.ListPolicyVersionsInput, func(*iam.ListPolicyVersionsOutput, bool) bool) error
	ListPolicyVersionsPagesWithContext(aws.Context, *iam.ListPolicyVersionsInput, func(*iam.ListPolicyVersionsOutput, bool) bool, ...request.Option) error

	ListRolePolicies(*iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)
	ListRolePoliciesWithContext(aws.Context, *iam.ListRolePoliciesInput, ...request.Option) (*iam.ListRolePoliciesOutput, error)
	ListRolePoliciesRequest(*iam.ListRolePoliciesInput) (*request.Request, *iam.ListRolePoliciesOutput)

	ListRolePoliciesPages(*iam.ListRolePoliciesInput, func(*iam.ListRolePoliciesOutput, bool) bool) error
	ListRolePoliciesPagesWithContext(aws.Context, *iam.ListRolePoliciesInput, func(*iam.ListRolePoliciesOutput, bool) bool, ...request.Option) error

	ListRoleTags(*iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error)
	ListRoleTagsWithContext(aws.Context, *iam.ListRoleTagsInput, ...request.Option) (*iam.ListRoleTagsOutput, error)
	ListRoleTagsRequest(*iam.ListRoleTagsInput) (*request.Request, *iam.ListRoleTagsOutput)

	ListRoles(*iam.ListRolesInput) (*iam.ListRolesOutput, error)
	ListRolesWithContext(aws.Context, *iam.ListRolesInput, ...request.Option) (*iam.ListRolesOutput, error)
	ListRolesRequest(*iam.ListRolesInput) (*request.Request, *iam.ListRolesOutput)

	ListRolesPages(*iam.ListRolesInput, func(*iam.ListRolesOutput, bool) bool) error
	ListRolesPagesWithContext(aws.Context, *iam.ListRolesInput, func(*iam.ListRolesOutput, bool) bool, ...request.Option) error

	ListSAMLProviderTags(*iam.ListSAMLProviderTagsInput) (*iam.ListSAMLProviderTagsOutput, error)
	ListSAMLProviderTagsWithContext(aws.Context, *iam.ListSAMLProviderTagsInput, ...request.Option) (*iam.ListSAMLProviderTagsOutput, error)
	ListSAMLProviderTagsRequest(*iam.ListSAMLProviderTagsInput) (*request.Request, *iam.ListSAMLProviderTagsOutput)

	ListSAMLProviders(*iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)
	ListSAMLProvidersWithContext(aws.Context, *iam.ListSAMLProvidersInput, ...request.Option) (*iam.ListSAMLProvidersOutput, error)
	ListSAMLProvidersRequest(*iam.ListSAMLProvidersInput) (*request.Request, *iam.ListSAMLProvidersOutput)

	ListSSHPublicKeys(*iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error)
	ListSSHPublicKeysWithContext(aws.Context, *iam.ListSSHPublicKeysInput, ...request.Option) (*iam.ListSSHPublicKeysOutput, error)
	ListSSHPublicKeysRequest(*iam.ListSSHPublicKeysInput) (*request.Request, *iam.ListSSHPublicKeysOutput)

	ListSSHPublicKeysPages(*iam.ListSSHPublicKeysInput, func(*iam.ListSSHPublicKeysOutput, bool) bool) error
	ListSSHPublicKeysPagesWithContext(aws.Context, *iam.ListSSHPublicKeysInput, func(*iam.ListSSHPublicKeysOutput, bool) bool, ...request.Option) error

	ListServerCertificateTags(*iam.ListServerCertificateTagsInput) (*iam.ListServerCertificateTagsOutput, error)
	ListServerCertificateTagsWithContext(aws.Context, *iam.ListServerCertificateTagsInput, ...request.Option) (*iam.ListServerCertificateTagsOutput, error)
	ListServerCertificateTagsRequest(*iam.ListServerCertificateTagsInput) (*request.Request, *iam.ListServerCertificateTagsOutput)

	ListServerCertificates(*iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)
	ListServerCertificatesWithContext(aws.Context, *iam.ListServerCertificatesInput, ...request.Option) (*iam.ListServerCertificatesOutput, error)
	ListServerCertificatesRequest(*iam.ListServerCertificatesInput) (*request.Request, *iam.ListServerCertificatesOutput)

	ListServerCertificatesPages(*iam.ListServerCertificatesInput, func(*iam.ListServerCertificatesOutput, bool) bool) error
	ListServerCertificatesPagesWithContext(aws.Context, *iam.ListServerCertificatesInput, func(*iam.ListServerCertificatesOutput, bool) bool, ...request.Option) error

	ListServiceSpecificCredentials(*iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListServiceSpecificCredentialsWithContext(aws.Context, *iam.ListServiceSpecificCredentialsInput, ...request.Option) (*iam.ListServiceSpecificCredentialsOutput, error)
	ListServiceSpecificCredentialsRequest(*iam.ListServiceSpecificCredentialsInput) (*request.Request, *iam.ListServiceSpecificCredentialsOutput)

	ListSigningCertificates(*iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)
	ListSigningCertificatesWithContext(aws.Context, *iam.ListSigningCertificatesInput, ...request.Option) (*iam.ListSigningCertificatesOutput, error)
	ListSigningCertificatesRequest(*iam.ListSigningCertificatesInput) (*request.Request, *iam.ListSigningCertificatesOutput)

	ListSigningCertificatesPages(*iam.ListSigningCertificatesInput, func(*iam.ListSigningCertificatesOutput, bool) bool) error
	ListSigningCertificatesPagesWithContext(aws.Context, *iam.ListSigningCertificatesInput, func(*iam.ListSigningCertificatesOutput, bool) bool, ...request.Option) error

	ListUserPolicies(*iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)
	ListUserPoliciesWithContext(aws.Context, *iam.ListUserPoliciesInput, ...request.Option) (*iam.ListUserPoliciesOutput, error)
	ListUserPoliciesRequest(*iam.ListUserPoliciesInput) (*request.Request, *iam.ListUserPoliciesOutput)

	ListUserPoliciesPages(*iam.ListUserPoliciesInput, func(*iam.ListUserPoliciesOutput, bool) bool) error
	ListUserPoliciesPagesWithContext(aws.Context, *iam.ListUserPoliciesInput, func(*iam.ListUserPoliciesOutput, bool) bool, ...request.Option) error

	ListUserTags(*iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error)
	ListUserTagsWithContext(aws.Context, *iam.ListUserTagsInput, ...request.Option) (*iam.ListUserTagsOutput, error)
	ListUserTagsRequest(*iam.ListUserTagsInput) (*request.Request, *iam.ListUserTagsOutput)

	ListUserTagsPages(*iam.ListUserTagsInput, func(*iam.ListUserTagsOutput, bool) bool) error
	ListUserTagsPagesWithContext(aws.Context, *iam.ListUserTagsInput, func(*iam.ListUserTagsOutput, bool) bool, ...request.Option) error

	ListUsers(*iam.ListUsersInput) (*iam.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *iam.ListUsersInput, ...request.Option) (*iam.ListUsersOutput, error)
	ListUsersRequest(*iam.ListUsersInput) (*request.Request, *iam.ListUsersOutput)

	ListUsersPages(*iam.ListUsersInput, func(*iam.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *iam.ListUsersInput, func(*iam.ListUsersOutput, bool) bool, ...request.Option) error

	ListVirtualMFADevices(*iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)
	ListVirtualMFADevicesWithContext(aws.Context, *iam.ListVirtualMFADevicesInput, ...request.Option) (*iam.ListVirtualMFADevicesOutput, error)
	ListVirtualMFADevicesRequest(*iam.ListVirtualMFADevicesInput) (*request.Request, *iam.ListVirtualMFADevicesOutput)

	ListVirtualMFADevicesPages(*iam.ListVirtualMFADevicesInput, func(*iam.ListVirtualMFADevicesOutput, bool) bool) error
	ListVirtualMFADevicesPagesWithContext(aws.Context, *iam.ListVirtualMFADevicesInput, func(*iam.ListVirtualMFADevicesOutput, bool) bool, ...request.Option) error

	PutGroupPolicy(*iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)
	PutGroupPolicyWithContext(aws.Context, *iam.PutGroupPolicyInput, ...request.Option) (*iam.PutGroupPolicyOutput, error)
	PutGroupPolicyRequest(*iam.PutGroupPolicyInput) (*request.Request, *iam.PutGroupPolicyOutput)

	PutRolePermissionsBoundary(*iam.PutRolePermissionsBoundaryInput) (*iam.PutRolePermissionsBoundaryOutput, error)
	PutRolePermissionsBoundaryWithContext(aws.Context, *iam.PutRolePermissionsBoundaryInput, ...request.Option) (*iam.PutRolePermissionsBoundaryOutput, error)
	PutRolePermissionsBoundaryRequest(*iam.PutRolePermissionsBoundaryInput) (*request.Request, *iam.PutRolePermissionsBoundaryOutput)

	PutRolePolicy(*iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)
	PutRolePolicyWithContext(aws.Context, *iam.PutRolePolicyInput, ...request.Option) (*iam.PutRolePolicyOutput, error)
	PutRolePolicyRequest(*iam.PutRolePolicyInput) (*request.Request, *iam.PutRolePolicyOutput)

	PutUserPermissionsBoundary(*iam.PutUserPermissionsBoundaryInput) (*iam.PutUserPermissionsBoundaryOutput, error)
	PutUserPermissionsBoundaryWithContext(aws.Context, *iam.PutUserPermissionsBoundaryInput, ...request.Option) (*iam.PutUserPermissionsBoundaryOutput, error)
	PutUserPermissionsBoundaryRequest(*iam.PutUserPermissionsBoundaryInput) (*request.Request, *iam.PutUserPermissionsBoundaryOutput)

	PutUserPolicy(*iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)
	PutUserPolicyWithContext(aws.Context, *iam.PutUserPolicyInput, ...request.Option) (*iam.PutUserPolicyOutput, error)
	PutUserPolicyRequest(*iam.PutUserPolicyInput) (*request.Request, *iam.PutUserPolicyOutput)

	RemoveClientIDFromOpenIDConnectProvider(*iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
	RemoveClientIDFromOpenIDConnectProviderWithContext(aws.Context, *iam.RemoveClientIDFromOpenIDConnectProviderInput, ...request.Option) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)
	RemoveClientIDFromOpenIDConnectProviderRequest(*iam.RemoveClientIDFromOpenIDConnectProviderInput) (*request.Request, *iam.RemoveClientIDFromOpenIDConnectProviderOutput)

	RemoveRoleFromInstanceProfile(*iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	RemoveRoleFromInstanceProfileWithContext(aws.Context, *iam.RemoveRoleFromInstanceProfileInput, ...request.Option) (*iam.RemoveRoleFromInstanceProfileOutput, error)
	RemoveRoleFromInstanceProfileRequest(*iam.RemoveRoleFromInstanceProfileInput) (*request.Request, *iam.RemoveRoleFromInstanceProfileOutput)

	RemoveUserFromGroup(*iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupWithContext(aws.Context, *iam.RemoveUserFromGroupInput, ...request.Option) (*iam.RemoveUserFromGroupOutput, error)
	RemoveUserFromGroupRequest(*iam.RemoveUserFromGroupInput) (*request.Request, *iam.RemoveUserFromGroupOutput)

	ResetServiceSpecificCredential(*iam.ResetServiceSpecificCredentialInput) (*iam.ResetServiceSpecificCredentialOutput, error)
	ResetServiceSpecificCredentialWithContext(aws.Context, *iam.ResetServiceSpecificCredentialInput, ...request.Option) (*iam.ResetServiceSpecificCredentialOutput, error)
	ResetServiceSpecificCredentialRequest(*iam.ResetServiceSpecificCredentialInput) (*request.Request, *iam.ResetServiceSpecificCredentialOutput)

	ResyncMFADevice(*iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)
	ResyncMFADeviceWithContext(aws.Context, *iam.ResyncMFADeviceInput, ...request.Option) (*iam.ResyncMFADeviceOutput, error)
	ResyncMFADeviceRequest(*iam.ResyncMFADeviceInput) (*request.Request, *iam.ResyncMFADeviceOutput)

	SetDefaultPolicyVersion(*iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionWithContext(aws.Context, *iam.SetDefaultPolicyVersionInput, ...request.Option) (*iam.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionRequest(*iam.SetDefaultPolicyVersionInput) (*request.Request, *iam.SetDefaultPolicyVersionOutput)

	SetSecurityTokenServicePreferences(*iam.SetSecurityTokenServicePreferencesInput) (*iam.SetSecurityTokenServicePreferencesOutput, error)
	SetSecurityTokenServicePreferencesWithContext(aws.Context, *iam.SetSecurityTokenServicePreferencesInput, ...request.Option) (*iam.SetSecurityTokenServicePreferencesOutput, error)
	SetSecurityTokenServicePreferencesRequest(*iam.SetSecurityTokenServicePreferencesInput) (*request.Request, *iam.SetSecurityTokenServicePreferencesOutput)

	SimulateCustomPolicy(*iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulateCustomPolicyWithContext(aws.Context, *iam.SimulateCustomPolicyInput, ...request.Option) (*iam.SimulatePolicyResponse, error)
	SimulateCustomPolicyRequest(*iam.SimulateCustomPolicyInput) (*request.Request, *iam.SimulatePolicyResponse)

	SimulateCustomPolicyPages(*iam.SimulateCustomPolicyInput, func(*iam.SimulatePolicyResponse, bool) bool) error
	SimulateCustomPolicyPagesWithContext(aws.Context, *iam.SimulateCustomPolicyInput, func(*iam.SimulatePolicyResponse, bool) bool, ...request.Option) error

	SimulatePrincipalPolicy(*iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error)
	SimulatePrincipalPolicyWithContext(aws.Context, *iam.SimulatePrincipalPolicyInput, ...request.Option) (*iam.SimulatePolicyResponse, error)
	SimulatePrincipalPolicyRequest(*iam.SimulatePrincipalPolicyInput) (*request.Request, *iam.SimulatePolicyResponse)

	SimulatePrincipalPolicyPages(*iam.SimulatePrincipalPolicyInput, func(*iam.SimulatePolicyResponse, bool) bool) error
	SimulatePrincipalPolicyPagesWithContext(aws.Context, *iam.SimulatePrincipalPolicyInput, func(*iam.SimulatePolicyResponse, bool) bool, ...request.Option) error

	TagInstanceProfile(*iam.TagInstanceProfileInput) (*iam.TagInstanceProfileOutput, error)
	TagInstanceProfileWithContext(aws.Context, *iam.TagInstanceProfileInput, ...request.Option) (*iam.TagInstanceProfileOutput, error)
	TagInstanceProfileRequest(*iam.TagInstanceProfileInput) (*request.Request, *iam.TagInstanceProfileOutput)

	TagMFADevice(*iam.TagMFADeviceInput) (*iam.TagMFADeviceOutput, error)
	TagMFADeviceWithContext(aws.Context, *iam.TagMFADeviceInput, ...request.Option) (*iam.TagMFADeviceOutput, error)
	TagMFADeviceRequest(*iam.TagMFADeviceInput) (*request.Request, *iam.TagMFADeviceOutput)

	TagOpenIDConnectProvider(*iam.TagOpenIDConnectProviderInput) (*iam.TagOpenIDConnectProviderOutput, error)
	TagOpenIDConnectProviderWithContext(aws.Context, *iam.TagOpenIDConnectProviderInput, ...request.Option) (*iam.TagOpenIDConnectProviderOutput, error)
	TagOpenIDConnectProviderRequest(*iam.TagOpenIDConnectProviderInput) (*request.Request, *iam.TagOpenIDConnectProviderOutput)

	TagPolicy(*iam.TagPolicyInput) (*iam.TagPolicyOutput, error)
	TagPolicyWithContext(aws.Context, *iam.TagPolicyInput, ...request.Option) (*iam.TagPolicyOutput, error)
	TagPolicyRequest(*iam.TagPolicyInput) (*request.Request, *iam.TagPolicyOutput)

	TagRole(*iam.TagRoleInput) (*iam.TagRoleOutput, error)
	TagRoleWithContext(aws.Context, *iam.TagRoleInput, ...request.Option) (*iam.TagRoleOutput, error)
	TagRoleRequest(*iam.TagRoleInput) (*request.Request, *iam.TagRoleOutput)

	TagSAMLProvider(*iam.TagSAMLProviderInput) (*iam.TagSAMLProviderOutput, error)
	TagSAMLProviderWithContext(aws.Context, *iam.TagSAMLProviderInput, ...request.Option) (*iam.TagSAMLProviderOutput, error)
	TagSAMLProviderRequest(*iam.TagSAMLProviderInput) (*request.Request, *iam.TagSAMLProviderOutput)

	TagServerCertificate(*iam.TagServerCertificateInput) (*iam.TagServerCertificateOutput, error)
	TagServerCertificateWithContext(aws.Context, *iam.TagServerCertificateInput, ...request.Option) (*iam.TagServerCertificateOutput, error)
	TagServerCertificateRequest(*iam.TagServerCertificateInput) (*request.Request, *iam.TagServerCertificateOutput)

	TagUser(*iam.TagUserInput) (*iam.TagUserOutput, error)
	TagUserWithContext(aws.Context, *iam.TagUserInput, ...request.Option) (*iam.TagUserOutput, error)
	TagUserRequest(*iam.TagUserInput) (*request.Request, *iam.TagUserOutput)

	UntagInstanceProfile(*iam.UntagInstanceProfileInput) (*iam.UntagInstanceProfileOutput, error)
	UntagInstanceProfileWithContext(aws.Context, *iam.UntagInstanceProfileInput, ...request.Option) (*iam.UntagInstanceProfileOutput, error)
	UntagInstanceProfileRequest(*iam.UntagInstanceProfileInput) (*request.Request, *iam.UntagInstanceProfileOutput)

	UntagMFADevice(*iam.UntagMFADeviceInput) (*iam.UntagMFADeviceOutput, error)
	UntagMFADeviceWithContext(aws.Context, *iam.UntagMFADeviceInput, ...request.Option) (*iam.UntagMFADeviceOutput, error)
	UntagMFADeviceRequest(*iam.UntagMFADeviceInput) (*request.Request, *iam.UntagMFADeviceOutput)

	UntagOpenIDConnectProvider(*iam.UntagOpenIDConnectProviderInput) (*iam.UntagOpenIDConnectProviderOutput, error)
	UntagOpenIDConnectProviderWithContext(aws.Context, *iam.UntagOpenIDConnectProviderInput, ...request.Option) (*iam.UntagOpenIDConnectProviderOutput, error)
	UntagOpenIDConnectProviderRequest(*iam.UntagOpenIDConnectProviderInput) (*request.Request, *iam.UntagOpenIDConnectProviderOutput)

	UntagPolicy(*iam.UntagPolicyInput) (*iam.UntagPolicyOutput, error)
	UntagPolicyWithContext(aws.Context, *iam.UntagPolicyInput, ...request.Option) (*iam.UntagPolicyOutput, error)
	UntagPolicyRequest(*iam.UntagPolicyInput) (*request.Request, *iam.UntagPolicyOutput)

	UntagRole(*iam.UntagRoleInput) (*iam.UntagRoleOutput, error)
	UntagRoleWithContext(aws.Context, *iam.UntagRoleInput, ...request.Option) (*iam.UntagRoleOutput, error)
	UntagRoleRequest(*iam.UntagRoleInput) (*request.Request, *iam.UntagRoleOutput)

	UntagSAMLProvider(*iam.UntagSAMLProviderInput) (*iam.UntagSAMLProviderOutput, error)
	UntagSAMLProviderWithContext(aws.Context, *iam.UntagSAMLProviderInput, ...request.Option) (*iam.UntagSAMLProviderOutput, error)
	UntagSAMLProviderRequest(*iam.UntagSAMLProviderInput) (*request.Request, *iam.UntagSAMLProviderOutput)

	UntagServerCertificate(*iam.UntagServerCertificateInput) (*iam.UntagServerCertificateOutput, error)
	UntagServerCertificateWithContext(aws.Context, *iam.UntagServerCertificateInput, ...request.Option) (*iam.UntagServerCertificateOutput, error)
	UntagServerCertificateRequest(*iam.UntagServerCertificateInput) (*request.Request, *iam.UntagServerCertificateOutput)

	UntagUser(*iam.UntagUserInput) (*iam.UntagUserOutput, error)
	UntagUserWithContext(aws.Context, *iam.UntagUserInput, ...request.Option) (*iam.UntagUserOutput, error)
	UntagUserRequest(*iam.UntagUserInput) (*request.Request, *iam.UntagUserOutput)

	UpdateAccessKey(*iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)
	UpdateAccessKeyWithContext(aws.Context, *iam.UpdateAccessKeyInput, ...request.Option) (*iam.UpdateAccessKeyOutput, error)
	UpdateAccessKeyRequest(*iam.UpdateAccessKeyInput) (*request.Request, *iam.UpdateAccessKeyOutput)

	UpdateAccountPasswordPolicy(*iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)
	UpdateAccountPasswordPolicyWithContext(aws.Context, *iam.UpdateAccountPasswordPolicyInput, ...request.Option) (*iam.UpdateAccountPasswordPolicyOutput, error)
	UpdateAccountPasswordPolicyRequest(*iam.UpdateAccountPasswordPolicyInput) (*request.Request, *iam.UpdateAccountPasswordPolicyOutput)

	UpdateAssumeRolePolicy(*iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)
	UpdateAssumeRolePolicyWithContext(aws.Context, *iam.UpdateAssumeRolePolicyInput, ...request.Option) (*iam.UpdateAssumeRolePolicyOutput, error)
	UpdateAssumeRolePolicyRequest(*iam.UpdateAssumeRolePolicyInput) (*request.Request, *iam.UpdateAssumeRolePolicyOutput)

	UpdateGroup(*iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)
	UpdateGroupWithContext(aws.Context, *iam.UpdateGroupInput, ...request.Option) (*iam.UpdateGroupOutput, error)
	UpdateGroupRequest(*iam.UpdateGroupInput) (*request.Request, *iam.UpdateGroupOutput)

	UpdateLoginProfile(*iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)
	UpdateLoginProfileWithContext(aws.Context, *iam.UpdateLoginProfileInput, ...request.Option) (*iam.UpdateLoginProfileOutput, error)
	UpdateLoginProfileRequest(*iam.UpdateLoginProfileInput) (*request.Request, *iam.UpdateLoginProfileOutput)

	UpdateOpenIDConnectProviderThumbprint(*iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
	UpdateOpenIDConnectProviderThumbprintWithContext(aws.Context, *iam.UpdateOpenIDConnectProviderThumbprintInput, ...request.Option) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)
	UpdateOpenIDConnectProviderThumbprintRequest(*iam.UpdateOpenIDConnectProviderThumbprintInput) (*request.Request, *iam.UpdateOpenIDConnectProviderThumbprintOutput)

	UpdateRole(*iam.UpdateRoleInput) (*iam.UpdateRoleOutput, error)
	UpdateRoleWithContext(aws.Context, *iam.UpdateRoleInput, ...request.Option) (*iam.UpdateRoleOutput, error)
	UpdateRoleRequest(*iam.UpdateRoleInput) (*request.Request, *iam.UpdateRoleOutput)

	UpdateRoleDescription(*iam.UpdateRoleDescriptionInput) (*iam.UpdateRoleDescriptionOutput, error)
	UpdateRoleDescriptionWithContext(aws.Context, *iam.UpdateRoleDescriptionInput, ...request.Option) (*iam.UpdateRoleDescriptionOutput, error)
	UpdateRoleDescriptionRequest(*iam.UpdateRoleDescriptionInput) (*request.Request, *iam.UpdateRoleDescriptionOutput)

	UpdateSAMLProvider(*iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderWithContext(aws.Context, *iam.UpdateSAMLProviderInput, ...request.Option) (*iam.UpdateSAMLProviderOutput, error)
	UpdateSAMLProviderRequest(*iam.UpdateSAMLProviderInput) (*request.Request, *iam.UpdateSAMLProviderOutput)

	UpdateSSHPublicKey(*iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error)
	UpdateSSHPublicKeyWithContext(aws.Context, *iam.UpdateSSHPublicKeyInput, ...request.Option) (*iam.UpdateSSHPublicKeyOutput, error)
	UpdateSSHPublicKeyRequest(*iam.UpdateSSHPublicKeyInput) (*request.Request, *iam.UpdateSSHPublicKeyOutput)

	UpdateServerCertificate(*iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)
	UpdateServerCertificateWithContext(aws.Context, *iam.UpdateServerCertificateInput, ...request.Option) (*iam.UpdateServerCertificateOutput, error)
	UpdateServerCertificateRequest(*iam.UpdateServerCertificateInput) (*request.Request, *iam.UpdateServerCertificateOutput)

	UpdateServiceSpecificCredential(*iam.UpdateServiceSpecificCredentialInput) (*iam.UpdateServiceSpecificCredentialOutput, error)
	UpdateServiceSpecificCredentialWithContext(aws.Context, *iam.UpdateServiceSpecificCredentialInput, ...request.Option) (*iam.UpdateServiceSpecificCredentialOutput, error)
	UpdateServiceSpecificCredentialRequest(*iam.UpdateServiceSpecificCredentialInput) (*request.Request, *iam.UpdateServiceSpecificCredentialOutput)

	UpdateSigningCertificate(*iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)
	UpdateSigningCertificateWithContext(aws.Context, *iam.UpdateSigningCertificateInput, ...request.Option) (*iam.UpdateSigningCertificateOutput, error)
	UpdateSigningCertificateRequest(*iam.UpdateSigningCertificateInput) (*request.Request, *iam.UpdateSigningCertificateOutput)

	UpdateUser(*iam.UpdateUserInput) (*iam.UpdateUserOutput, error)
	UpdateUserWithContext(aws.Context, *iam.UpdateUserInput, ...request.Option) (*iam.UpdateUserOutput, error)
	UpdateUserRequest(*iam.UpdateUserInput) (*request.Request, *iam.UpdateUserOutput)

	UploadSSHPublicKey(*iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error)
	UploadSSHPublicKeyWithContext(aws.Context, *iam.UploadSSHPublicKeyInput, ...request.Option) (*iam.UploadSSHPublicKeyOutput, error)
	UploadSSHPublicKeyRequest(*iam.UploadSSHPublicKeyInput) (*request.Request, *iam.UploadSSHPublicKeyOutput)

	UploadServerCertificate(*iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)
	UploadServerCertificateWithContext(aws.Context, *iam.UploadServerCertificateInput, ...request.Option) (*iam.UploadServerCertificateOutput, error)
	UploadServerCertificateRequest(*iam.UploadServerCertificateInput) (*request.Request, *iam.UploadServerCertificateOutput)

	UploadSigningCertificate(*iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
	UploadSigningCertificateWithContext(aws.Context, *iam.UploadSigningCertificateInput, ...request.Option) (*iam.UploadSigningCertificateOutput, error)
	UploadSigningCertificateRequest(*iam.UploadSigningCertificateInput) (*request.Request, *iam.UploadSigningCertificateOutput)

	WaitUntilInstanceProfileExists(*iam.GetInstanceProfileInput) error
	WaitUntilInstanceProfileExistsWithContext(aws.Context, *iam.GetInstanceProfileInput, ...request.WaiterOption) error

	WaitUntilPolicyExists(*iam.GetPolicyInput) error
	WaitUntilPolicyExistsWithContext(aws.Context, *iam.GetPolicyInput, ...request.WaiterOption) error

	WaitUntilRoleExists(*iam.GetRoleInput) error
	WaitUntilRoleExistsWithContext(aws.Context, *iam.GetRoleInput, ...request.WaiterOption) error

	WaitUntilUserExists(*iam.GetUserInput) error
	WaitUntilUserExistsWithContext(aws.Context, *iam.GetUserInput, ...request.WaiterOption) error
}
