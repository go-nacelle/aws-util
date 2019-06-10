package awsutil

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/acm"
	"github.com/aws/aws-sdk-go/service/acmpca"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/backup"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/chime"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go/service/cloudsearch"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codecommit"
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/dax"
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/glacier"
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/greengrass"
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/iot"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"github.com/aws/aws-sdk-go/service/iotevents"
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/kafka"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/macie"
	"github.com/aws/aws-sdk-go/service/managedblockchain"
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplacemetering"
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediastore"
	"github.com/aws/aws-sdk-go/service/mediastoredata"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mobileanalytics"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworkscm"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/pi"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/simpledb"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/snowball"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/storagegateway"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/textract"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/workdocs"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/go-nacelle/nacelle"
)

func NewACMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("acm", func(sess *session.Session) interface{} {
		return acm.New(sess)
	})
}

func NewACMPCAServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("acmpca", func(sess *session.Session) interface{} {
		return acmpca.New(sess)
	})
}

func NewAlexaForBusinessServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("alexaforbusiness", func(sess *session.Session) interface{} {
		return alexaforbusiness.New(sess)
	})
}

func NewAmplifyServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("amplify", func(sess *session.Session) interface{} {
		return amplify.New(sess)
	})
}

func NewAPIGatewayServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("apigateway", func(sess *session.Session) interface{} {
		return apigateway.New(sess)
	})
}

func NewApiGatewayManagementApiServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("apigatewaymanagementapi", func(sess *session.Session) interface{} {
		return apigatewaymanagementapi.New(sess)
	})
}

func NewApiGatewayV2ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("apigatewayv2", func(sess *session.Session) interface{} {
		return apigatewayv2.New(sess)
	})
}

func NewApplicationAutoScalingServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("applicationautoscaling", func(sess *session.Session) interface{} {
		return applicationautoscaling.New(sess)
	})
}

func NewApplicationDiscoveryServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("applicationdiscoveryservice", func(sess *session.Session) interface{} {
		return applicationdiscoveryservice.New(sess)
	})
}

func NewAppMeshServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("appmesh", func(sess *session.Session) interface{} {
		return appmesh.New(sess)
	})
}

func NewAppStreamServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("appstream", func(sess *session.Session) interface{} {
		return appstream.New(sess)
	})
}

func NewAppSyncServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("appsync", func(sess *session.Session) interface{} {
		return appsync.New(sess)
	})
}

func NewAthenaServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("athena", func(sess *session.Session) interface{} {
		return athena.New(sess)
	})
}

func NewAutoScalingServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("autoscaling", func(sess *session.Session) interface{} {
		return autoscaling.New(sess)
	})
}

func NewAutoScalingPlansServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("autoscalingplans", func(sess *session.Session) interface{} {
		return autoscalingplans.New(sess)
	})
}

func NewBackupServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("backup", func(sess *session.Session) interface{} {
		return backup.New(sess)
	})
}

func NewBatchServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("batch", func(sess *session.Session) interface{} {
		return batch.New(sess)
	})
}

func NewBudgetsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("budgets", func(sess *session.Session) interface{} {
		return budgets.New(sess)
	})
}

func NewChimeServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("chime", func(sess *session.Session) interface{} {
		return chime.New(sess)
	})
}

func NewCloud9ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloud9", func(sess *session.Session) interface{} {
		return cloud9.New(sess)
	})
}

func NewCloudDirectoryServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("clouddirectory", func(sess *session.Session) interface{} {
		return clouddirectory.New(sess)
	})
}

func NewCloudFormationServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudformation", func(sess *session.Session) interface{} {
		return cloudformation.New(sess)
	})
}

func NewCloudFrontServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudfront", func(sess *session.Session) interface{} {
		return cloudfront.New(sess)
	})
}

func NewCloudHSMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudhsm", func(sess *session.Session) interface{} {
		return cloudhsm.New(sess)
	})
}

func NewCloudHSMV2ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudhsmv2", func(sess *session.Session) interface{} {
		return cloudhsmv2.New(sess)
	})
}

func NewCloudSearchServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudsearch", func(sess *session.Session) interface{} {
		return cloudsearch.New(sess)
	})
}

func NewCloudSearchDomainServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudsearchdomain", func(sess *session.Session) interface{} {
		return cloudsearchdomain.New(sess)
	})
}

func NewCloudTrailServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudtrail", func(sess *session.Session) interface{} {
		return cloudtrail.New(sess)
	})
}

func NewCloudWatchServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudwatch", func(sess *session.Session) interface{} {
		return cloudwatch.New(sess)
	})
}

func NewCloudWatchEventsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudwatchevents", func(sess *session.Session) interface{} {
		return cloudwatchevents.New(sess)
	})
}

func NewCloudWatchLogsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cloudwatchlogs", func(sess *session.Session) interface{} {
		return cloudwatchlogs.New(sess)
	})
}

func NewCodeBuildServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("codebuild", func(sess *session.Session) interface{} {
		return codebuild.New(sess)
	})
}

func NewCodeCommitServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("codecommit", func(sess *session.Session) interface{} {
		return codecommit.New(sess)
	})
}

func NewCodeDeployServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("codedeploy", func(sess *session.Session) interface{} {
		return codedeploy.New(sess)
	})
}

func NewCodePipelineServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("codepipeline", func(sess *session.Session) interface{} {
		return codepipeline.New(sess)
	})
}

func NewCodeStarServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("codestar", func(sess *session.Session) interface{} {
		return codestar.New(sess)
	})
}

func NewCognitoIdentityServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cognitoidentity", func(sess *session.Session) interface{} {
		return cognitoidentity.New(sess)
	})
}

func NewCognitoIdentityProviderServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cognitoidentityprovider", func(sess *session.Session) interface{} {
		return cognitoidentityprovider.New(sess)
	})
}

func NewCognitoSyncServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("cognitosync", func(sess *session.Session) interface{} {
		return cognitosync.New(sess)
	})
}

func NewComprehendServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("comprehend", func(sess *session.Session) interface{} {
		return comprehend.New(sess)
	})
}

func NewComprehendMedicalServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("comprehendmedical", func(sess *session.Session) interface{} {
		return comprehendmedical.New(sess)
	})
}

func NewConfigServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("configservice", func(sess *session.Session) interface{} {
		return configservice.New(sess)
	})
}

func NewConnectServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("connect", func(sess *session.Session) interface{} {
		return connect.New(sess)
	})
}

func NewCostandUsageReportServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("costandusagereportservice", func(sess *session.Session) interface{} {
		return costandusagereportservice.New(sess)
	})
}

func NewCostExplorerServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("costexplorer", func(sess *session.Session) interface{} {
		return costexplorer.New(sess)
	})
}

func NewDatabaseMigrationServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("databasemigrationservice", func(sess *session.Session) interface{} {
		return databasemigrationservice.New(sess)
	})
}

func NewDataPipelineServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("datapipeline", func(sess *session.Session) interface{} {
		return datapipeline.New(sess)
	})
}

func NewDataSyncServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("datasync", func(sess *session.Session) interface{} {
		return datasync.New(sess)
	})
}

func NewDAXServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("dax", func(sess *session.Session) interface{} {
		return dax.New(sess)
	})
}

func NewDeviceFarmServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("devicefarm", func(sess *session.Session) interface{} {
		return devicefarm.New(sess)
	})
}

func NewDirectConnectServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("directconnect", func(sess *session.Session) interface{} {
		return directconnect.New(sess)
	})
}

func NewDirectoryServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("directoryservice", func(sess *session.Session) interface{} {
		return directoryservice.New(sess)
	})
}

func NewDLMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("dlm", func(sess *session.Session) interface{} {
		return dlm.New(sess)
	})
}

func NewDocDBServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("docdb", func(sess *session.Session) interface{} {
		return docdb.New(sess)
	})
}

func NewDynamoDBServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("dynamodb", func(sess *session.Session) interface{} {
		return dynamodb.New(sess)
	})
}

func NewDynamoDBStreamsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("dynamodbstreams", func(sess *session.Session) interface{} {
		return dynamodbstreams.New(sess)
	})
}

func NewEC2ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ec2", func(sess *session.Session) interface{} {
		return ec2.New(sess)
	})
}

func NewECRServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ecr", func(sess *session.Session) interface{} {
		return ecr.New(sess)
	})
}

func NewECSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ecs", func(sess *session.Session) interface{} {
		return ecs.New(sess)
	})
}

func NewEFSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("efs", func(sess *session.Session) interface{} {
		return efs.New(sess)
	})
}

func NewEKSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("eks", func(sess *session.Session) interface{} {
		return eks.New(sess)
	})
}

func NewElastiCacheServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("elasticache", func(sess *session.Session) interface{} {
		return elasticache.New(sess)
	})
}

func NewElasticBeanstalkServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("elasticbeanstalk", func(sess *session.Session) interface{} {
		return elasticbeanstalk.New(sess)
	})
}

func NewElasticsearchServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("elasticsearchservice", func(sess *session.Session) interface{} {
		return elasticsearchservice.New(sess)
	})
}

func NewElasticTranscoderServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("elastictranscoder", func(sess *session.Session) interface{} {
		return elastictranscoder.New(sess)
	})
}

func NewELBServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("elb", func(sess *session.Session) interface{} {
		return elb.New(sess)
	})
}

func NewELBV2ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("elbv2", func(sess *session.Session) interface{} {
		return elbv2.New(sess)
	})
}

func NewEMRServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("emr", func(sess *session.Session) interface{} {
		return emr.New(sess)
	})
}

func NewFirehoseServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("firehose", func(sess *session.Session) interface{} {
		return firehose.New(sess)
	})
}

func NewFMSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("fms", func(sess *session.Session) interface{} {
		return fms.New(sess)
	})
}

func NewFSxServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("fsx", func(sess *session.Session) interface{} {
		return fsx.New(sess)
	})
}

func NewGameLiftServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("gamelift", func(sess *session.Session) interface{} {
		return gamelift.New(sess)
	})
}

func NewGlacierServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("glacier", func(sess *session.Session) interface{} {
		return glacier.New(sess)
	})
}

func NewGlobalAcceleratorServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("globalaccelerator", func(sess *session.Session) interface{} {
		return globalaccelerator.New(sess)
	})
}

func NewGlueServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("glue", func(sess *session.Session) interface{} {
		return glue.New(sess)
	})
}

func NewGreengrassServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("greengrass", func(sess *session.Session) interface{} {
		return greengrass.New(sess)
	})
}

func NewGroundStationServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("groundstation", func(sess *session.Session) interface{} {
		return groundstation.New(sess)
	})
}

func NewGuardDutyServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("guardduty", func(sess *session.Session) interface{} {
		return guardduty.New(sess)
	})
}

func NewHealthServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("health", func(sess *session.Session) interface{} {
		return health.New(sess)
	})
}

func NewIAMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iam", func(sess *session.Session) interface{} {
		return iam.New(sess)
	})
}

func NewInspectorServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("inspector", func(sess *session.Session) interface{} {
		return inspector.New(sess)
	})
}

func NewIoTServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iot", func(sess *session.Session) interface{} {
		return iot.New(sess)
	})
}

func NewIoT1ClickDevicesServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iot1clickdevicesservice", func(sess *session.Session) interface{} {
		return iot1clickdevicesservice.New(sess)
	})
}

func NewIoT1ClickProjectsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iot1clickprojects", func(sess *session.Session) interface{} {
		return iot1clickprojects.New(sess)
	})
}

func NewIoTAnalyticsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iotanalytics", func(sess *session.Session) interface{} {
		return iotanalytics.New(sess)
	})
}

func NewIoTDataPlaneServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iotdataplane", func(sess *session.Session) interface{} {
		return iotdataplane.New(sess)
	})
}

func NewIoTEventsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iotevents", func(sess *session.Session) interface{} {
		return iotevents.New(sess)
	})
}

func NewIoTEventsDataServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ioteventsdata", func(sess *session.Session) interface{} {
		return ioteventsdata.New(sess)
	})
}

func NewIoTJobsDataPlaneServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iotjobsdataplane", func(sess *session.Session) interface{} {
		return iotjobsdataplane.New(sess)
	})
}

func NewIoTThingsGraphServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("iotthingsgraph", func(sess *session.Session) interface{} {
		return iotthingsgraph.New(sess)
	})
}

func NewKafkaServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kafka", func(sess *session.Session) interface{} {
		return kafka.New(sess)
	})
}

func NewKinesisServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kinesis", func(sess *session.Session) interface{} {
		return kinesis.New(sess)
	})
}

func NewKinesisAnalyticsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kinesisanalytics", func(sess *session.Session) interface{} {
		return kinesisanalytics.New(sess)
	})
}

func NewKinesisAnalyticsV2ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kinesisanalyticsv2", func(sess *session.Session) interface{} {
		return kinesisanalyticsv2.New(sess)
	})
}

func NewKinesisVideoServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kinesisvideo", func(sess *session.Session) interface{} {
		return kinesisvideo.New(sess)
	})
}

func NewKinesisVideoArchivedMediaServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kinesisvideoarchivedmedia", func(sess *session.Session) interface{} {
		return kinesisvideoarchivedmedia.New(sess)
	})
}

func NewKinesisVideoMediaServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kinesisvideomedia", func(sess *session.Session) interface{} {
		return kinesisvideomedia.New(sess)
	})
}

func NewKMSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("kms", func(sess *session.Session) interface{} {
		return kms.New(sess)
	})
}

func NewLambdaServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("lambda", func(sess *session.Session) interface{} {
		return lambda.New(sess)
	})
}

func NewLexModelBuildingServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("lexmodelbuildingservice", func(sess *session.Session) interface{} {
		return lexmodelbuildingservice.New(sess)
	})
}

func NewLexRuntimeServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("lexruntimeservice", func(sess *session.Session) interface{} {
		return lexruntimeservice.New(sess)
	})
}

func NewLicenseManagerServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("licensemanager", func(sess *session.Session) interface{} {
		return licensemanager.New(sess)
	})
}

func NewLightsailServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("lightsail", func(sess *session.Session) interface{} {
		return lightsail.New(sess)
	})
}

func NewMachineLearningServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("machinelearning", func(sess *session.Session) interface{} {
		return machinelearning.New(sess)
	})
}

func NewMacieServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("macie", func(sess *session.Session) interface{} {
		return macie.New(sess)
	})
}

func NewManagedBlockchainServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("managedblockchain", func(sess *session.Session) interface{} {
		return managedblockchain.New(sess)
	})
}

func NewMarketplaceCommerceAnalyticsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("marketplacecommerceanalytics", func(sess *session.Session) interface{} {
		return marketplacecommerceanalytics.New(sess)
	})
}

func NewMarketplaceEntitlementServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("marketplaceentitlementservice", func(sess *session.Session) interface{} {
		return marketplaceentitlementservice.New(sess)
	})
}

func NewMarketplaceMeteringServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("marketplacemetering", func(sess *session.Session) interface{} {
		return marketplacemetering.New(sess)
	})
}

func NewMediaConnectServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediaconnect", func(sess *session.Session) interface{} {
		return mediaconnect.New(sess)
	})
}

func NewMediaConvertServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediaconvert", func(sess *session.Session) interface{} {
		return mediaconvert.New(sess)
	})
}

func NewMediaLiveServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("medialive", func(sess *session.Session) interface{} {
		return medialive.New(sess)
	})
}

func NewMediaPackageServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediapackage", func(sess *session.Session) interface{} {
		return mediapackage.New(sess)
	})
}

func NewMediaPackageVodServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediapackagevod", func(sess *session.Session) interface{} {
		return mediapackagevod.New(sess)
	})
}

func NewMediaStoreServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediastore", func(sess *session.Session) interface{} {
		return mediastore.New(sess)
	})
}

func NewMediaStoreDataServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediastoredata", func(sess *session.Session) interface{} {
		return mediastoredata.New(sess)
	})
}

func NewMediaTailorServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mediatailor", func(sess *session.Session) interface{} {
		return mediatailor.New(sess)
	})
}

func NewMigrationHubServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("migrationhub", func(sess *session.Session) interface{} {
		return migrationhub.New(sess)
	})
}

func NewMobileServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mobile", func(sess *session.Session) interface{} {
		return mobile.New(sess)
	})
}

func NewMobileAnalyticsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mobileanalytics", func(sess *session.Session) interface{} {
		return mobileanalytics.New(sess)
	})
}

func NewMQServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mq", func(sess *session.Session) interface{} {
		return mq.New(sess)
	})
}

func NewMTurkServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("mturk", func(sess *session.Session) interface{} {
		return mturk.New(sess)
	})
}

func NewNeptuneServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("neptune", func(sess *session.Session) interface{} {
		return neptune.New(sess)
	})
}

func NewOpsWorksServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("opsworks", func(sess *session.Session) interface{} {
		return opsworks.New(sess)
	})
}

func NewOpsWorksCMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("opsworkscm", func(sess *session.Session) interface{} {
		return opsworkscm.New(sess)
	})
}

func NewOrganizationsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("organizations", func(sess *session.Session) interface{} {
		return organizations.New(sess)
	})
}

func NewPIServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("pi", func(sess *session.Session) interface{} {
		return pi.New(sess)
	})
}

func NewPinpointServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("pinpoint", func(sess *session.Session) interface{} {
		return pinpoint.New(sess)
	})
}

func NewPinpointEmailServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("pinpointemail", func(sess *session.Session) interface{} {
		return pinpointemail.New(sess)
	})
}

func NewPinpointSMSVoiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("pinpointsmsvoice", func(sess *session.Session) interface{} {
		return pinpointsmsvoice.New(sess)
	})
}

func NewPollyServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("polly", func(sess *session.Session) interface{} {
		return polly.New(sess)
	})
}

func NewPricingServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("pricing", func(sess *session.Session) interface{} {
		return pricing.New(sess)
	})
}

func NewQuickSightServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("quicksight", func(sess *session.Session) interface{} {
		return quicksight.New(sess)
	})
}

func NewRAMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ram", func(sess *session.Session) interface{} {
		return ram.New(sess)
	})
}

func NewRDSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("rds", func(sess *session.Session) interface{} {
		return rds.New(sess)
	})
}

func NewRDSDataServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("rdsdataservice", func(sess *session.Session) interface{} {
		return rdsdataservice.New(sess)
	})
}

func NewRedshiftServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("redshift", func(sess *session.Session) interface{} {
		return redshift.New(sess)
	})
}

func NewRekognitionServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("rekognition", func(sess *session.Session) interface{} {
		return rekognition.New(sess)
	})
}

func NewResourceGroupsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("resourcegroups", func(sess *session.Session) interface{} {
		return resourcegroups.New(sess)
	})
}

func NewResourceGroupsTaggingAPIServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("resourcegroupstaggingapi", func(sess *session.Session) interface{} {
		return resourcegroupstaggingapi.New(sess)
	})
}

func NewRoboMakerServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("robomaker", func(sess *session.Session) interface{} {
		return robomaker.New(sess)
	})
}

func NewRoute53ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("route53", func(sess *session.Session) interface{} {
		return route53.New(sess)
	})
}

func NewRoute53DomainsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("route53domains", func(sess *session.Session) interface{} {
		return route53domains.New(sess)
	})
}

func NewRoute53ResolverServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("route53resolver", func(sess *session.Session) interface{} {
		return route53resolver.New(sess)
	})
}

func NewS3ServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("s3", func(sess *session.Session) interface{} {
		return s3.New(sess)
	})
}

func NewS3ControlServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("s3control", func(sess *session.Session) interface{} {
		return s3control.New(sess)
	})
}

func NewSageMakerServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sagemaker", func(sess *session.Session) interface{} {
		return sagemaker.New(sess)
	})
}

func NewSageMakerRuntimeServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sagemakerruntime", func(sess *session.Session) interface{} {
		return sagemakerruntime.New(sess)
	})
}

func NewSecretsManagerServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("secretsmanager", func(sess *session.Session) interface{} {
		return secretsmanager.New(sess)
	})
}

func NewSecurityHubServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("securityhub", func(sess *session.Session) interface{} {
		return securityhub.New(sess)
	})
}

func NewServerlessApplicationRepositoryServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("serverlessapplicationrepository", func(sess *session.Session) interface{} {
		return serverlessapplicationrepository.New(sess)
	})
}

func NewServiceCatalogServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("servicecatalog", func(sess *session.Session) interface{} {
		return servicecatalog.New(sess)
	})
}

func NewServiceDiscoveryServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("servicediscovery", func(sess *session.Session) interface{} {
		return servicediscovery.New(sess)
	})
}

func NewSESServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ses", func(sess *session.Session) interface{} {
		return ses.New(sess)
	})
}

func NewSFNServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sfn", func(sess *session.Session) interface{} {
		return sfn.New(sess)
	})
}

func NewShieldServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("shield", func(sess *session.Session) interface{} {
		return shield.New(sess)
	})
}

func NewSignerServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("signer", func(sess *session.Session) interface{} {
		return signer.New(sess)
	})
}

func NewSimpleDBServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("simpledb", func(sess *session.Session) interface{} {
		return simpledb.New(sess)
	})
}

func NewSMSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sms", func(sess *session.Session) interface{} {
		return sms.New(sess)
	})
}

func NewSnowballServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("snowball", func(sess *session.Session) interface{} {
		return snowball.New(sess)
	})
}

func NewSNSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sns", func(sess *session.Session) interface{} {
		return sns.New(sess)
	})
}

func NewSQSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sqs", func(sess *session.Session) interface{} {
		return sqs.New(sess)
	})
}

func NewSSMServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("ssm", func(sess *session.Session) interface{} {
		return ssm.New(sess)
	})
}

func NewStorageGatewayServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("storagegateway", func(sess *session.Session) interface{} {
		return storagegateway.New(sess)
	})
}

func NewSTSServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("sts", func(sess *session.Session) interface{} {
		return sts.New(sess)
	})
}

func NewSupportServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("support", func(sess *session.Session) interface{} {
		return support.New(sess)
	})
}

func NewSWFServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("swf", func(sess *session.Session) interface{} {
		return swf.New(sess)
	})
}

func NewTextractServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("textract", func(sess *session.Session) interface{} {
		return textract.New(sess)
	})
}

func NewTranscribeServiceServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("transcribeservice", func(sess *session.Session) interface{} {
		return transcribeservice.New(sess)
	})
}

func NewTransferServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("transfer", func(sess *session.Session) interface{} {
		return transfer.New(sess)
	})
}

func NewTranslateServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("translate", func(sess *session.Session) interface{} {
		return translate.New(sess)
	})
}

func NewWAFServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("waf", func(sess *session.Session) interface{} {
		return waf.New(sess)
	})
}

func NewWAFRegionalServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("wafregional", func(sess *session.Session) interface{} {
		return wafregional.New(sess)
	})
}

func NewWorkDocsServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("workdocs", func(sess *session.Session) interface{} {
		return workdocs.New(sess)
	})
}

func NewWorkLinkServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("worklink", func(sess *session.Session) interface{} {
		return worklink.New(sess)
	})
}

func NewWorkMailServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("workmail", func(sess *session.Session) interface{} {
		return workmail.New(sess)
	})
}

func NewWorkSpacesServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("workspaces", func(sess *session.Session) interface{} {
		return workspaces.New(sess)
	})
}

func NewXRayServiceInitializer() nacelle.Initializer {
	return NewServiceInitializer("xray", func(sess *session.Session) interface{} {
		return xray.New(sess)
	})
}
