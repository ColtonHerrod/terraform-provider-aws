// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	quicksight_sdkv1 "github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceFolderMembership,
			Name:    "Folder Membership",
		},
		{
			Factory: newResourceIAMPolicyAssignment,
			Name:    "IAM Policy Assignment",
		},
		{
			Factory: newResourceIngestion,
			Name:    "Ingestion",
		},
		{
			Factory: newResourceNamespace,
			Name:    "Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory: newResourceRefreshSchedule,
			Name:    "Refresh Schedule",
		},
		{
			Factory: newResourceTemplateAlias,
			Name:    "Template Alias",
		},
		{
			Factory: newResourceVPCConnection,
			Name:    "VPC Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAnalysis,
			TypeName: "aws_quicksight_analysis",
			Name:     "Analysis",
		},
		{
			Factory:  DataSourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
		},
		{
			Factory:  DataSourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
		},
		{
			Factory:  DataSourceTheme,
			TypeName: "aws_quicksight_theme",
			Name:     "Theme",
		},
		{
			Factory:  DataSourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccountSubscription,
			TypeName: "aws_quicksight_account_subscription",
			Name:     "Account Subscription",
		},
		{
			Factory:  ResourceAnalysis,
			TypeName: "aws_quicksight_analysis",
			Name:     "Analysis",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDashboard,
			TypeName: "aws_quicksight_dashboard",
			Name:     "Dashboard",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDataSource,
			TypeName: "aws_quicksight_data_source",
			Name:     "Data Source",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceFolder,
			TypeName: "aws_quicksight_folder",
			Name:     "Folder",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
		},
		{
			Factory:  ResourceGroupMembership,
			TypeName: "aws_quicksight_group_membership",
			Name:     "Group Membership",
		},
		{
			Factory:  ResourceTemplate,
			TypeName: "aws_quicksight_template",
			Name:     "Template",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTheme,
			TypeName: "aws_quicksight_theme",
			Name:     "Theme",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.QuickSight
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*quicksight_sdkv1.QuickSight, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return quicksight_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
