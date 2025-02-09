/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DBInstanceParameters defines the desired state of DBInstance
type DBInstanceParameters struct {
	// Region is which region the DBInstance will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The amount of storage (in gibibytes) to allocate for the DB instance.
	//
	// Type: Integer
	//
	// Amazon Aurora
	//
	// Not applicable. Aurora cluster volumes automatically grow as the amount of
	// data in your database increases, though you are only charged for the space
	// that you use in an Aurora cluster volume.
	//
	// MySQL
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	//    * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	//
	//    * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//
	//    * Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// MariaDB
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	//    * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	//
	//    * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//
	//    * Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// PostgreSQL
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	//    * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	//
	//    * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//
	//    * Magnetic storage (standard): Must be an integer from 5 to 3072.
	//
	// Oracle
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	//    * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.
	//
	//    * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.
	//
	//    * Magnetic storage (standard): Must be an integer from 10 to 3072.
	//
	// SQL Server
	//
	// Constraints to the amount of storage for each storage type are the following:
	//
	//    * General Purpose (SSD) storage (gp2): Enterprise and Standard editions:
	//    Must be an integer from 200 to 16384. Web and Express editions: Must be
	//    an integer from 20 to 16384.
	//
	//    * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must
	//    be an integer from 200 to 16384. Web and Express editions: Must be an
	//    integer from 100 to 16384.
	//
	//    * Magnetic storage (standard): Enterprise and Standard editions: Must
	//    be an integer from 200 to 1024. Web and Express editions: Must be an integer
	//    from 20 to 1024.
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty"`
	// A value that indicates whether minor engine upgrades are applied automatically
	// to the DB instance during the maintenance window. By default, minor engine
	// upgrades are applied automatically.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty"`
	// The Availability Zone (AZ) where the database will be created. For information
	// on AWS Regions and Availability Zones, see Regions and Availability Zones
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS
	// Region.
	//
	// Example: us-east-1d
	//
	// Constraint: The AvailabilityZone parameter can't be specified if the DB instance
	// is a Multi-AZ deployment. The specified Availability Zone must be in the
	// same AWS Region as the current endpoint.
	//
	// If you're creating a DB instance in an RDS on VMware environment, specify
	// the identifier of the custom Availability Zone to create the DB instance
	// in.
	//
	// For more information about RDS on VMware, see the RDS on VMware User Guide.
	// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
	AvailabilityZone *string `json:"availabilityZone,omitempty"`
	// The number of days for which automated backups are retained. Setting this
	// parameter to a positive number enables backups. Setting this parameter to
	// 0 disables automated backups.
	//
	// Amazon Aurora
	//
	// Not applicable. The retention period for automated backups is managed by
	// the DB cluster.
	//
	// Default: 1
	//
	// Constraints:
	//
	//    * Must be a value from 0 to 35
	//
	//    * Can't be set to 0 if the DB instance is a source to read replicas
	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty"`
	// For supported engines, indicates that the DB instance should be associated
	// with the specified CharacterSet.
	//
	// Amazon Aurora
	//
	// Not applicable. The character set is managed by the DB cluster. For more
	// information, see CreateDBCluster.
	CharacterSetName *string `json:"characterSetName,omitempty"`
	// A value that indicates whether to copy tags from the DB instance to snapshots
	// of the DB instance. By default, tags are not copied.
	//
	// Amazon Aurora
	//
	// Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting
	// this value for an Aurora DB instance has no effect on the DB cluster setting.
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty"`
	// The identifier of the DB cluster that the instance will belong to.
	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`
	// The compute and memory capacity of the DB instance, for example, db.m4.large.
	// Not all DB instance classes are available in all AWS Regions, or for all
	// database engines. For the full list of DB instance classes, and availability
	// for your engine, see DB Instance Class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)
	// in the Amazon RDS User Guide.
	// +kubebuilder:validation:Required
	DBInstanceClass *string `json:"dbInstanceClass"`
	// The meaning of this parameter differs according to the database engine you
	// use.
	//
	// MySQL
	//
	// The name of the database to create when the DB instance is created. If this
	// parameter isn't specified, no database is created in the DB instance.
	//
	// Constraints:
	//
	//    * Must contain 1 to 64 letters or numbers.
	//
	//    * Must begin with a letter. Subsequent characters can be letters, underscores,
	//    or digits (0-9).
	//
	//    * Can't be a word reserved by the specified database engine
	//
	// MariaDB
	//
	// The name of the database to create when the DB instance is created. If this
	// parameter isn't specified, no database is created in the DB instance.
	//
	// Constraints:
	//
	//    * Must contain 1 to 64 letters or numbers.
	//
	//    * Must begin with a letter. Subsequent characters can be letters, underscores,
	//    or digits (0-9).
	//
	//    * Can't be a word reserved by the specified database engine
	//
	// PostgreSQL
	//
	// The name of the database to create when the DB instance is created. If this
	// parameter isn't specified, a database named postgres is created in the DB
	// instance.
	//
	// Constraints:
	//
	//    * Must contain 1 to 63 letters, numbers, or underscores.
	//
	//    * Must begin with a letter. Subsequent characters can be letters, underscores,
	//    or digits (0-9).
	//
	//    * Can't be a word reserved by the specified database engine
	//
	// Oracle
	//
	// The Oracle System ID (SID) of the created DB instance. If you specify null,
	// the default value ORCL is used. You can't specify the string NULL, or any
	// other reserved word, for DBName.
	//
	// Default: ORCL
	//
	// Constraints:
	//
	//    * Can't be longer than 8 characters
	//
	// SQL Server
	//
	// Not applicable. Must be null.
	//
	// Amazon Aurora MySQL
	//
	// The name of the database to create when the primary DB instance of the Aurora
	// MySQL DB cluster is created. If this parameter isn't specified for an Aurora
	// MySQL DB cluster, no database is created in the DB cluster.
	//
	// Constraints:
	//
	//    * It must contain 1 to 64 alphanumeric characters.
	//
	//    * It can't be a word reserved by the database engine.
	//
	// Amazon Aurora PostgreSQL
	//
	// The name of the database to create when the primary DB instance of the Aurora
	// PostgreSQL DB cluster is created. If this parameter isn't specified for an
	// Aurora PostgreSQL DB cluster, a database named postgres is created in the
	// DB cluster.
	//
	// Constraints:
	//
	//    * It must contain 1 to 63 alphanumeric characters.
	//
	//    * It must begin with a letter or an underscore. Subsequent characters
	//    can be letters, underscores, or digits (0 to 9).
	//
	//    * It can't be a word reserved by the database engine.
	DBName *string `json:"dbName,omitempty"`
	// The name of the DB parameter group to associate with this DB instance. If
	// you do not specify a value, then the default DB parameter group for the specified
	// DB engine and version is used.
	//
	// Constraints:
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	DBParameterGroupName *string `json:"dbParameterGroupName,omitempty"`
	// A DB subnet group to associate with this DB instance.
	//
	// If there is no DB subnet group, then it is a non-VPC DB instance.
	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty"`
	// A value that indicates whether the DB instance has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled. For more information, see Deleting a DB
	// Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).
	//
	// Amazon Aurora
	//
	// Not applicable. You can enable or disable deletion protection for the DB
	// cluster. For more information, see CreateDBCluster. DB instances in a DB
	// cluster can be deleted even when deletion protection is enabled for the DB
	// cluster.
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
	// The Active Directory directory ID to create the DB instance in. Currently,
	// only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can
	// be created in an Active Directory Domain.
	//
	// For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)
	// in the Amazon RDS User Guide.
	Domain *string `json:"domain,omitempty"`
	// Specify the name of the IAM role to be used when making API calls to the
	// Directory Service.
	DomainIAMRoleName *string `json:"domainIAMRoleName,omitempty"`
	// The list of log types that need to be enabled for exporting to CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Relational Database Service User Guide.
	//
	// Amazon Aurora
	//
	// Not applicable. CloudWatch Logs exports are managed by the DB cluster.
	//
	// MariaDB
	//
	// Possible values are audit, error, general, and slowquery.
	//
	// Microsoft SQL Server
	//
	// Possible values are agent and error.
	//
	// MySQL
	//
	// Possible values are audit, error, general, and slowquery.
	//
	// Oracle
	//
	// Possible values are alert, audit, listener, trace, and oemagent.
	//
	// PostgreSQL
	//
	// Possible values are postgresql and upgrade.
	EnableCloudwatchLogsExports []*string `json:"enableCloudwatchLogsExports,omitempty"`
	// A value that indicates whether to enable a customer-owned IP address (CoIP)
	// for an RDS on Outposts DB instance.
	//
	// A CoIP provides local or external connectivity to resources in your Outpost
	// subnets through your on-premises network. For some use cases, a CoIP can
	// provide lower latency for connections to the DB instance from outside of
	// its virtual private cloud (VPC) on your local network.
	//
	// For more information about RDS on Outposts, see Working with Amazon RDS on
	// AWS Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)
	// in the Amazon RDS User Guide.
	//
	// For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
	// in the AWS Outposts User Guide.
	EnableCustomerOwnedIP *bool `json:"enableCustomerOwnedIP,omitempty"`
	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	//
	// This setting doesn't apply to Amazon Aurora. Mapping AWS IAM accounts to
	// database accounts is managed by the DB cluster.
	//
	// For more information, see IAM Database Authentication for MySQL and PostgreSQL
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon RDS User Guide.
	EnableIAMDatabaseAuthentication *bool `json:"enableIAMDatabaseAuthentication,omitempty"`
	// A value that indicates whether to enable Performance Insights for the DB
	// instance.
	//
	// For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)
	// in the Amazon Relational Database Service User Guide.
	EnablePerformanceInsights *bool `json:"enablePerformanceInsights,omitempty"`
	// The name of the database engine to be used for this instance.
	//
	// Not every database engine is available for every AWS Region.
	//
	// Valid Values:
	//
	//    * aurora (for MySQL 5.6-compatible Aurora)
	//
	//    * aurora-mysql (for MySQL 5.7-compatible Aurora)
	//
	//    * aurora-postgresql
	//
	//    * mariadb
	//
	//    * mysql
	//
	//    * oracle-ee
	//
	//    * oracle-se2
	//
	//    * oracle-se1
	//
	//    * oracle-se
	//
	//    * postgres
	//
	//    * sqlserver-ee
	//
	//    * sqlserver-se
	//
	//    * sqlserver-ex
	//
	//    * sqlserver-web
	// +kubebuilder:validation:Required
	Engine *string `json:"engine"`
	// The version number of the database engine to use.
	//
	// For a list of valid engine versions, use the DescribeDBEngineVersions action.
	//
	// The following are the database engines and links to information about the
	// major and minor versions that are available with Amazon RDS. Not every database
	// engine is available for every AWS Region.
	//
	// Amazon Aurora
	//
	// Not applicable. The version number of the database engine to be used by the
	// DB instance is managed by the DB cluster.
	//
	// MariaDB
	//
	// See MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt)
	// in the Amazon RDS User Guide.
	//
	// Microsoft SQL Server
	//
	// See Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport)
	// in the Amazon RDS User Guide.
	//
	// MySQL
	//
	// See MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt)
	// in the Amazon RDS User Guide.
	//
	// Oracle
	//
	// See Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html)
	// in the Amazon RDS User Guide.
	//
	// PostgreSQL
	//
	// See Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts)
	// in the Amazon RDS User Guide.
	EngineVersion *string `json:"engineVersion,omitempty"`
	// The amount of Provisioned IOPS (input/output operations per second) to be
	// initially allocated for the DB instance. For information about valid Iops
	// values, see Amazon RDS Provisioned IOPS Storage to Improve Performance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS)
	// in the Amazon RDS User Guide.
	//
	// Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, must
	// be a multiple between .5 and 50 of the storage amount for the DB instance.
	// For SQL Server DB instances, must be a multiple between 1 and 50 of the storage
	// amount for the DB instance.
	IOPS *int64 `json:"iops,omitempty"`
	// The AWS KMS key identifier for an encrypted DB instance.
	//
	// The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name
	// for the AWS KMS customer master key (CMK). To use a CMK in a different AWS
	// account, specify the key ARN or alias ARN.
	//
	// Amazon Aurora
	//
	// Not applicable. The AWS KMS key identifier is managed by the DB cluster.
	// For more information, see CreateDBCluster.
	//
	// If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId
	// parameter, then Amazon RDS uses your default CMK. There is a default CMK
	// for your AWS account. Your AWS account has a different default CMK for each
	// AWS Region.
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// License model information for this DB instance.
	//
	// Valid values: license-included | bring-your-own-license | general-public-license
	LicenseModel *string `json:"licenseModel,omitempty"`
	// The name for the master user.
	//
	// Amazon Aurora
	//
	// Not applicable. The name for the master user is managed by the DB cluster.
	//
	// MariaDB
	//
	// Constraints:
	//
	//    * Required for MariaDB.
	//
	//    * Must be 1 to 16 letters or numbers.
	//
	//    * Can't be a reserved word for the chosen database engine.
	//
	// Microsoft SQL Server
	//
	// Constraints:
	//
	//    * Required for SQL Server.
	//
	//    * Must be 1 to 128 letters or numbers.
	//
	//    * The first character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	//
	// MySQL
	//
	// Constraints:
	//
	//    * Required for MySQL.
	//
	//    * Must be 1 to 16 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	//
	// Oracle
	//
	// Constraints:
	//
	//    * Required for Oracle.
	//
	//    * Must be 1 to 30 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	//
	// PostgreSQL
	//
	// Constraints:
	//
	//    * Required for PostgreSQL.
	//
	//    * Must be 1 to 63 letters or numbers.
	//
	//    * First character must be a letter.
	//
	//    * Can't be a reserved word for the chosen database engine.
	MasterUsername *string `json:"masterUsername,omitempty"`
	// The upper limit to which Amazon RDS can automatically scale the storage of
	// the DB instance.
	//
	// For more information about this setting, including limitations that apply
	// to it, see Managing capacity automatically with Amazon RDS storage autoscaling
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)
	// in the Amazon RDS User Guide.
	MaxAllocatedStorage *int64 `json:"maxAllocatedStorage,omitempty"`
	// The interval, in seconds, between points when Enhanced Monitoring metrics
	// are collected for the DB instance. To disable collecting Enhanced Monitoring
	// metrics, specify 0. The default is 0.
	//
	// If MonitoringRoleArn is specified, then you must also set MonitoringInterval
	// to a value other than 0.
	//
	// Valid Values: 0, 1, 5, 10, 15, 30, 60
	MonitoringInterval *int64 `json:"monitoringInterval,omitempty"`
	// The ARN for the IAM role that permits RDS to send enhanced monitoring metrics
	// to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.
	// For information on creating a monitoring role, go to Setting Up and Enabling
	// Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)
	// in the Amazon RDS User Guide.
	//
	// If MonitoringInterval is set to a value other than 0, then you must supply
	// a MonitoringRoleArn value.
	MonitoringRoleARN *string `json:"monitoringRoleARN,omitempty"`
	// A value that indicates whether the DB instance is a Multi-AZ deployment.
	// You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ
	// deployment.
	MultiAZ *bool `json:"multiAZ,omitempty"`
	// The name of the NCHAR character set for the Oracle DB instance.
	NcharCharacterSetName *string `json:"ncharCharacterSetName,omitempty"`
	// A value that indicates that the DB instance should be associated with the
	// specified option group.
	//
	// Permanent options, such as the TDE option for Oracle Advanced Security TDE,
	// can't be removed from an option group. Also, that option group can't be removed
	// from a DB instance once it is associated with a DB instance
	OptionGroupName *string `json:"optionGroupName,omitempty"`
	// The AWS KMS key identifier for encryption of Performance Insights data.
	//
	// The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name
	// for the AWS KMS customer master key (CMK).
	//
	// If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon
	// RDS uses your default CMK. There is a default CMK for your AWS account. Your
	// AWS account has a different default CMK for each AWS Region.
	PerformanceInsightsKMSKeyID *string `json:"performanceInsightsKMSKeyID,omitempty"`
	// The amount of time, in days, to retain Performance Insights data. Valid values
	// are 7 or 731 (2 years).
	PerformanceInsightsRetentionPeriod *int64 `json:"performanceInsightsRetentionPeriod,omitempty"`
	// The port number on which the database accepts connections.
	//
	// MySQL
	//
	// Default: 3306
	//
	// Valid values: 1150-65535
	//
	// Type: Integer
	//
	// MariaDB
	//
	// Default: 3306
	//
	// Valid values: 1150-65535
	//
	// Type: Integer
	//
	// PostgreSQL
	//
	// Default: 5432
	//
	// Valid values: 1150-65535
	//
	// Type: Integer
	//
	// Oracle
	//
	// Default: 1521
	//
	// Valid values: 1150-65535
	//
	// SQL Server
	//
	// Default: 1433
	//
	// Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and
	// 49152-49156.
	//
	// Amazon Aurora
	//
	// Default: 3306
	//
	// Valid values: 1150-65535
	//
	// Type: Integer
	Port *int64 `json:"port,omitempty"`
	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. For more
	// information, see The Backup Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)
	// in the Amazon RDS User Guide.
	//
	// Amazon Aurora
	//
	// Not applicable. The daily time range for creating automated backups is managed
	// by the DB cluster.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region. To see the time blocks available, see Adjusting
	// the Preferred DB Instance Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow)
	// in the Amazon RDS User Guide.
	//
	// Constraints:
	//
	//    * Must be in the format hh24:mi-hh24:mi.
	//
	//    * Must be in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty"`
	// The time range each week during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). For more information, see Amazon RDS Maintenance
	// Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).
	//
	// Format: ddd:hh24:mi-ddd:hh24:mi
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week.
	//
	// Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`
	// The number of CPU cores and the number of threads per core for the DB instance
	// class of the DB instance.
	ProcessorFeatures []*ProcessorFeature `json:"processorFeatures,omitempty"`
	// A value that specifies the order in which an Aurora Replica is promoted to
	// the primary instance after a failure of the existing primary instance. For
	// more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance)
	// in the Amazon Aurora User Guide.
	//
	// Default: 1
	//
	// Valid Values: 0 - 15
	PromotionTier *int64 `json:"promotionTier,omitempty"`
	// A value that indicates whether the DB instance is publicly accessible.
	//
	// When the DB instance is publicly accessible, its DNS endpoint resolves to
	// the private IP address from within the DB instance's VPC, and to the public
	// IP address from outside of the DB instance's VPC. Access to the DB instance
	// is ultimately controlled by the security group it uses, and that public access
	// is not permitted if the security group assigned to the DB instance doesn't
	// permit it.
	//
	// When the DB instance isn't publicly accessible, it is an internal DB instance
	// with a DNS name that resolves to a private IP address.
	//
	// Default: The default behavior varies depending on whether DBSubnetGroupName
	// is specified.
	//
	// If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified,
	// the following applies:
	//
	//    * If the default VPC in the target region doesn’t have an Internet gateway
	//    attached to it, the DB instance is private.
	//
	//    * If the default VPC in the target region has an Internet gateway attached
	//    to it, the DB instance is public.
	//
	// If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified,
	// the following applies:
	//
	//    * If the subnets are part of a VPC that doesn’t have an Internet gateway
	//    attached to it, the DB instance is private.
	//
	//    * If the subnets are part of a VPC that has an Internet gateway attached
	//    to it, the DB instance is public.
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty"`
	// A value that indicates whether the DB instance is encrypted. By default,
	// it isn't encrypted.
	//
	// Amazon Aurora
	//
	// Not applicable. The encryption for DB instances is managed by the DB cluster.
	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`
	// Specifies the storage type to be associated with the DB instance.
	//
	// Valid values: standard | gp2 | io1
	//
	// If you specify io1, you must also include a value for the Iops parameter.
	//
	// Default: io1 if the Iops parameter is specified, otherwise gp2
	StorageType *string `json:"storageType,omitempty"`
	// Tags to assign to the DB instance.
	Tags []*Tag `json:"tags,omitempty"`
	// The ARN from the key store with which to associate the instance for TDE encryption.
	TDECredentialARN *string `json:"tdeCredentialARN,omitempty"`
	// The password for the given ARN from the key store in order to access the
	// device.
	TDECredentialPassword *string `json:"tdeCredentialPassword,omitempty"`
	// The time zone of the DB instance. The time zone parameter is currently supported
	// only by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).
	Timezone                   *string `json:"timezone,omitempty"`
	CustomDBInstanceParameters `json:",inline"`
}

// DBInstanceSpec defines the desired state of DBInstance
type DBInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DBInstanceParameters `json:"forProvider"`
}

// DBInstanceObservation defines the observed state of DBInstance
type DBInstanceObservation struct {
	// The AWS Identity and Access Management (IAM) roles associated with the DB
	// instance.
	AssociatedRoles []*DBInstanceRole `json:"associatedRoles,omitempty"`
	// The identifier of the CA certificate for this DB instance.
	CACertificateIdentifier *string `json:"caCertificateIdentifier,omitempty"`
	// Specifies whether a customer-owned IP address (CoIP) is enabled for an RDS
	// on Outposts DB instance.
	//
	// A CoIP provides local or external connectivity to resources in your Outpost
	// subnets through your on-premises network. For some use cases, a CoIP can
	// provide lower latency for connections to the DB instance from outside of
	// its virtual private cloud (VPC) on your local network.
	//
	// For more information about RDS on Outposts, see Working with Amazon RDS on
	// AWS Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)
	// in the Amazon RDS User Guide.
	//
	// For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
	// in the AWS Outposts User Guide.
	CustomerOwnedIPEnabled *bool `json:"customerOwnedIPEnabled,omitempty"`
	// The Amazon Resource Name (ARN) for the DB instance.
	DBInstanceARN *string `json:"dbInstanceARN,omitempty"`
	// The list of replicated automated backups associated with the DB instance.
	DBInstanceAutomatedBackupsReplications []*DBInstanceAutomatedBackupsReplication `json:"dbInstanceAutomatedBackupsReplications,omitempty"`
	// Contains a user-supplied database identifier. This identifier is the unique
	// key that identifies a DB instance.
	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty"`
	// Specifies the current state of this database.
	//
	// For information about DB instance statuses, see DB Instance Status (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.DBInstance.Status.html)
	// in the Amazon RDS User Guide.
	DBInstanceStatus *string `json:"dbInstanceStatus,omitempty"`
	// Provides the list of DB parameter groups applied to this DB instance.
	DBParameterGroups []*DBParameterGroupStatus_SDK `json:"dbParameterGroups,omitempty"`
	// A list of DB security group elements containing DBSecurityGroup.Name and
	// DBSecurityGroup.Status subelements.
	DBSecurityGroups []*DBSecurityGroupMembership `json:"dbSecurityGroups,omitempty"`
	// Specifies information on the subnet group associated with the DB instance,
	// including the name, description, and subnets in the subnet group.
	DBSubnetGroup *DBSubnetGroup `json:"dbSubnetGroup,omitempty"`
	// Specifies the port that the DB instance listens on. If the DB instance is
	// part of a DB cluster, this can be a different port than the DB cluster port.
	DBInstancePort *int64 `json:"dbInstancePort,omitempty"`
	// The AWS Region-unique, immutable identifier for the DB instance. This identifier
	// is found in AWS CloudTrail log entries whenever the AWS KMS customer master
	// key (CMK) for the DB instance is accessed.
	DBIResourceID *string `json:"dbiResourceID,omitempty"`
	// The Active Directory Domain membership records associated with the DB instance.
	DomainMemberships []*DomainMembership `json:"domainMemberships,omitempty"`
	// A list of log types that this DB instance is configured to export to CloudWatch
	// Logs.
	//
	// Log types vary by DB engine. For information about the log types for each
	// DB engine, see Amazon RDS Database Log Files (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html)
	// in the Amazon RDS User Guide.
	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty"`
	// Specifies the connection endpoint.
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch Logs log stream that
	// receives the Enhanced Monitoring metrics data for the DB instance.
	EnhancedMonitoringResourceARN *string `json:"enhancedMonitoringResourceARN,omitempty"`
	// True if mapping of AWS Identity and Access Management (IAM) accounts to database
	// accounts is enabled, and otherwise false.
	//
	// IAM database authentication can be enabled for the following database engines
	//
	//    * For MySQL 5.6, minor version 5.6.34 or higher
	//
	//    * For MySQL 5.7, minor version 5.7.16 or higher
	//
	//    * Aurora 5.6 or higher. To enable IAM database authentication for Aurora,
	//    see DBCluster Type.
	IAMDatabaseAuthenticationEnabled *bool `json:"iamDatabaseAuthenticationEnabled,omitempty"`
	// Provides the date and time the DB instance was created.
	InstanceCreateTime *metav1.Time `json:"instanceCreateTime,omitempty"`
	// Specifies the latest time to which a database can be restored with point-in-time
	// restore.
	LatestRestorableTime *metav1.Time `json:"latestRestorableTime,omitempty"`
	// Specifies the listener connection endpoint for SQL Server Always On.
	ListenerEndpoint *Endpoint `json:"listenerEndpoint,omitempty"`
	// Provides the list of option group memberships for this DB instance.
	OptionGroupMemberships []*OptionGroupMembership `json:"optionGroupMemberships,omitempty"`
	// A value that specifies that changes to the DB instance are pending. This
	// element is only included when changes are pending. Specific changes are identified
	// by subelements.
	PendingModifiedValues *PendingModifiedValues `json:"pendingModifiedValues,omitempty"`
	// True if Performance Insights is enabled for the DB instance, and otherwise
	// false.
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty"`
	// Contains one or more identifiers of Aurora DB clusters to which the RDS DB
	// instance is replicated as a read replica. For example, when you create an
	// Aurora read replica of an RDS MySQL DB instance, the Aurora MySQL DB cluster
	// for the Aurora read replica is shown. This output does not contain information
	// about cross region Aurora read replicas.
	//
	// Currently, each RDS DB instance can have only one Aurora read replica.
	ReadReplicaDBClusterIdentifiers []*string `json:"readReplicaDBClusterIdentifiers,omitempty"`
	// Contains one or more identifiers of the read replicas associated with this
	// DB instance.
	ReadReplicaDBInstanceIdentifiers []*string `json:"readReplicaDBInstanceIdentifiers,omitempty"`
	// Contains the identifier of the source DB instance if this DB instance is
	// a read replica.
	ReadReplicaSourceDBInstanceIdentifier *string `json:"readReplicaSourceDBInstanceIdentifier,omitempty"`
	// The open mode of an Oracle read replica. The default is open-read-only. For
	// more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html)
	// in the Amazon RDS User Guide.
	//
	// This attribute is only supported in RDS for Oracle.
	ReplicaMode *string `json:"replicaMode,omitempty"`
	// If present, specifies the name of the secondary Availability Zone for a DB
	// instance with multi-AZ support.
	SecondaryAvailabilityZone *string `json:"secondaryAvailabilityZone,omitempty"`
	// The status of a read replica. If the instance isn't a read replica, this
	// is blank.
	StatusInfos []*DBInstanceStatusInfo `json:"statusInfos,omitempty"`

	TagList []*Tag `json:"tagList,omitempty"`
	// Provides a list of VPC security group elements that the DB instance belongs
	// to.
	VPCSecurityGroups []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

// DBInstanceStatus defines the observed state of DBInstance.
type DBInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DBInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DBInstance is the Schema for the DBInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DBInstanceSpec   `json:"spec"`
	Status            DBInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBInstanceList contains a list of DBInstances
type DBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBInstance `json:"items"`
}

// Repository type metadata.
var (
	DBInstanceKind             = "DBInstance"
	DBInstanceGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DBInstanceKind}.String()
	DBInstanceKindAPIVersion   = DBInstanceKind + "." + GroupVersion.String()
	DBInstanceGroupVersionKind = GroupVersion.WithKind(DBInstanceKind)
)

func init() {
	SchemeBuilder.Register(&DBInstance{}, &DBInstanceList{})
}
