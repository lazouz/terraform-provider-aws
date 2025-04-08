// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory:  newDataSourceDataShares,
			TypeName: "aws_redshift_data_shares",
			Name:     "Data Shares",
		},
		{
			Factory:  newDataSourceProducerDataShares,
			TypeName: "aws_redshift_producer_data_shares",
			Name:     "Producer Data Shares",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newResourceDataShareAuthorization,
			TypeName: "aws_redshift_data_share_authorization",
			Name:     "Data Share Authorization",
		},
		{
			Factory:  newResourceDataShareConsumerAssociation,
			TypeName: "aws_redshift_data_share_consumer_association",
			Name:     "Data Share Consumer Association",
		},
		{
			Factory:  newIntegrationResource,
			TypeName: "aws_redshift_integration",
			Name:     "Integration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  newResourceLogging,
			TypeName: "aws_redshift_logging",
			Name:     "Logging",
		},
		{
			Factory:  newResourceSnapshotCopy,
			TypeName: "aws_redshift_snapshot_copy",
			Name:     "Snapshot Copy",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_redshift_cluster",
			Name:     "Cluster",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceClusterCredentials,
			TypeName: "aws_redshift_cluster_credentials",
			Name:     "Cluster Credentials",
		},
		{
			Factory:  dataSourceOrderableCluster,
			TypeName: "aws_redshift_orderable_cluster",
			Name:     "Orderable Cluster",
		},
		{
			Factory:  dataSourceServiceAccount,
			TypeName: "aws_redshift_service_account",
			Name:     "Service Account",
		},
		{
			Factory:  dataSourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
			Name:     "Subnet Group",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAuthenticationProfile,
			TypeName: "aws_redshift_authentication_profile",
			Name:     "Authentication Profile",
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_redshift_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterIAMRoles,
			TypeName: "aws_redshift_cluster_iam_roles",
			Name:     "Cluster IAM Roles",
		},
		{
			Factory:  resourceClusterSnapshot,
			TypeName: "aws_redshift_cluster_snapshot",
			Name:     "Cluster Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceEndpointAccess,
			TypeName: "aws_redshift_endpoint_access",
			Name:     "Endpoint Access",
		},
		{
			Factory:  resourceEndpointAuthorization,
			TypeName: "aws_redshift_endpoint_authorization",
			Name:     "Endpoint Authorization",
		},
		{
			Factory:  resourceEventSubscription,
			TypeName: "aws_redshift_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceHSMClientCertificate,
			TypeName: "aws_redshift_hsm_client_certificate",
			Name:     "HSM Client Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceHSMConfiguration,
			TypeName: "aws_redshift_hsm_configuration",
			Name:     "HSM Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_redshift_parameter_group",
			Name:     "Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourcePartner,
			TypeName: "aws_redshift_partner",
			Name:     "Partner",
		},
		{
			Factory:  resourceResourcePolicy,
			TypeName: "aws_redshift_resource_policy",
			Name:     "Resource Policy",
		},
		{
			Factory:  resourceScheduledAction,
			TypeName: "aws_redshift_scheduled_action",
			Name:     "Scheduled Action",
		},
		{
			Factory:  resourceSnapshotCopyGrant,
			TypeName: "aws_redshift_snapshot_copy_grant",
			Name:     "Snapshot Copy Grant",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSnapshotSchedule,
			TypeName: "aws_redshift_snapshot_schedule",
			Name:     "Snapshot Schedule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSnapshotScheduleAssociation,
			TypeName: "aws_redshift_snapshot_schedule_association",
			Name:     "Snapshot Schedule Association",
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceUsageLimit,
			TypeName: "aws_redshift_usage_limit",
			Name:     "Usage Limit",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Redshift
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*redshift.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*redshift.Options){
		redshift.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return redshift.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*redshift.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*redshift.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *redshift.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*redshift.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
