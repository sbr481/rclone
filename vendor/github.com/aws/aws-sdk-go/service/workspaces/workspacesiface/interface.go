// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package workspacesiface provides an interface to enable mocking the Amazon WorkSpaces service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package workspacesiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

// WorkSpacesAPI provides an interface to enable mocking the
// workspaces.WorkSpaces service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon WorkSpaces.
//    func myFunc(svc workspacesiface.WorkSpacesAPI) bool {
//        // Make svc.CreateTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := workspaces.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockWorkSpacesClient struct {
//        workspacesiface.WorkSpacesAPI
//    }
//    func (m *mockWorkSpacesClient) CreateTags(input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockWorkSpacesClient{}
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
type WorkSpacesAPI interface {
	CreateTags(*workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error)
	CreateTagsWithContext(aws.Context, *workspaces.CreateTagsInput, ...request.Option) (*workspaces.CreateTagsOutput, error)
	CreateTagsRequest(*workspaces.CreateTagsInput) (*request.Request, *workspaces.CreateTagsOutput)

	CreateWorkspaces(*workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)
	CreateWorkspacesWithContext(aws.Context, *workspaces.CreateWorkspacesInput, ...request.Option) (*workspaces.CreateWorkspacesOutput, error)
	CreateWorkspacesRequest(*workspaces.CreateWorkspacesInput) (*request.Request, *workspaces.CreateWorkspacesOutput)

	DeleteTags(*workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error)
	DeleteTagsWithContext(aws.Context, *workspaces.DeleteTagsInput, ...request.Option) (*workspaces.DeleteTagsOutput, error)
	DeleteTagsRequest(*workspaces.DeleteTagsInput) (*request.Request, *workspaces.DeleteTagsOutput)

	DescribeTags(*workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error)
	DescribeTagsWithContext(aws.Context, *workspaces.DescribeTagsInput, ...request.Option) (*workspaces.DescribeTagsOutput, error)
	DescribeTagsRequest(*workspaces.DescribeTagsInput) (*request.Request, *workspaces.DescribeTagsOutput)

	DescribeWorkspaceBundles(*workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)
	DescribeWorkspaceBundlesWithContext(aws.Context, *workspaces.DescribeWorkspaceBundlesInput, ...request.Option) (*workspaces.DescribeWorkspaceBundlesOutput, error)
	DescribeWorkspaceBundlesRequest(*workspaces.DescribeWorkspaceBundlesInput) (*request.Request, *workspaces.DescribeWorkspaceBundlesOutput)

	DescribeWorkspaceBundlesPages(*workspaces.DescribeWorkspaceBundlesInput, func(*workspaces.DescribeWorkspaceBundlesOutput, bool) bool) error
	DescribeWorkspaceBundlesPagesWithContext(aws.Context, *workspaces.DescribeWorkspaceBundlesInput, func(*workspaces.DescribeWorkspaceBundlesOutput, bool) bool, ...request.Option) error

	DescribeWorkspaceDirectories(*workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
	DescribeWorkspaceDirectoriesWithContext(aws.Context, *workspaces.DescribeWorkspaceDirectoriesInput, ...request.Option) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
	DescribeWorkspaceDirectoriesRequest(*workspaces.DescribeWorkspaceDirectoriesInput) (*request.Request, *workspaces.DescribeWorkspaceDirectoriesOutput)

	DescribeWorkspaceDirectoriesPages(*workspaces.DescribeWorkspaceDirectoriesInput, func(*workspaces.DescribeWorkspaceDirectoriesOutput, bool) bool) error
	DescribeWorkspaceDirectoriesPagesWithContext(aws.Context, *workspaces.DescribeWorkspaceDirectoriesInput, func(*workspaces.DescribeWorkspaceDirectoriesOutput, bool) bool, ...request.Option) error

	DescribeWorkspaces(*workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspacesWithContext(aws.Context, *workspaces.DescribeWorkspacesInput, ...request.Option) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspacesRequest(*workspaces.DescribeWorkspacesInput) (*request.Request, *workspaces.DescribeWorkspacesOutput)

	DescribeWorkspacesPages(*workspaces.DescribeWorkspacesInput, func(*workspaces.DescribeWorkspacesOutput, bool) bool) error
	DescribeWorkspacesPagesWithContext(aws.Context, *workspaces.DescribeWorkspacesInput, func(*workspaces.DescribeWorkspacesOutput, bool) bool, ...request.Option) error

	DescribeWorkspacesConnectionStatus(*workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
	DescribeWorkspacesConnectionStatusWithContext(aws.Context, *workspaces.DescribeWorkspacesConnectionStatusInput, ...request.Option) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
	DescribeWorkspacesConnectionStatusRequest(*workspaces.DescribeWorkspacesConnectionStatusInput) (*request.Request, *workspaces.DescribeWorkspacesConnectionStatusOutput)

	ModifyWorkspaceProperties(*workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error)
	ModifyWorkspacePropertiesWithContext(aws.Context, *workspaces.ModifyWorkspacePropertiesInput, ...request.Option) (*workspaces.ModifyWorkspacePropertiesOutput, error)
	ModifyWorkspacePropertiesRequest(*workspaces.ModifyWorkspacePropertiesInput) (*request.Request, *workspaces.ModifyWorkspacePropertiesOutput)

	RebootWorkspaces(*workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)
	RebootWorkspacesWithContext(aws.Context, *workspaces.RebootWorkspacesInput, ...request.Option) (*workspaces.RebootWorkspacesOutput, error)
	RebootWorkspacesRequest(*workspaces.RebootWorkspacesInput) (*request.Request, *workspaces.RebootWorkspacesOutput)

	RebuildWorkspaces(*workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)
	RebuildWorkspacesWithContext(aws.Context, *workspaces.RebuildWorkspacesInput, ...request.Option) (*workspaces.RebuildWorkspacesOutput, error)
	RebuildWorkspacesRequest(*workspaces.RebuildWorkspacesInput) (*request.Request, *workspaces.RebuildWorkspacesOutput)

	StartWorkspaces(*workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error)
	StartWorkspacesWithContext(aws.Context, *workspaces.StartWorkspacesInput, ...request.Option) (*workspaces.StartWorkspacesOutput, error)
	StartWorkspacesRequest(*workspaces.StartWorkspacesInput) (*request.Request, *workspaces.StartWorkspacesOutput)

	StopWorkspaces(*workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error)
	StopWorkspacesWithContext(aws.Context, *workspaces.StopWorkspacesInput, ...request.Option) (*workspaces.StopWorkspacesOutput, error)
	StopWorkspacesRequest(*workspaces.StopWorkspacesInput) (*request.Request, *workspaces.StopWorkspacesOutput)

	TerminateWorkspaces(*workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
	TerminateWorkspacesWithContext(aws.Context, *workspaces.TerminateWorkspacesInput, ...request.Option) (*workspaces.TerminateWorkspacesOutput, error)
	TerminateWorkspacesRequest(*workspaces.TerminateWorkspacesInput) (*request.Request, *workspaces.TerminateWorkspacesOutput)
}

var _ WorkSpacesAPI = (*workspaces.WorkSpaces)(nil)