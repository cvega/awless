/* Copyright 2017 WALLIX

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

// DO NOT EDIT
// This file was automatically generated with go generate
package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/wallix/awless/logger"
)

type CopyImage struct {
	_            string `action:"copy" entity:"image" awsAPI:"ec2" awsCall:"CopyImage" awsInput:"ec2.CopyImageInput" awsOutput:"ec2.CopyImageOutput" awsDryRun:""`
	logger       *logger.Logger
	api          ec2iface.EC2API
	Name         *string `awsName:"Name" awsType:"awsstr" templateName:"name" required:""`
	SourceId     *string `awsName:"SourceImageId" awsType:"awsstr" templateName:"source-id" required:""`
	SourceRegion *string `awsName:"SourceRegion" awsType:"awsstr" templateName:"source-region" required:""`
	Encrypted    *bool   `awsName:"Encrypted" awsType:"awsbool" templateName:"encrypted"`
	Description  *string `awsName:"Description" awsType:"awsstr" templateName:"description"`
}

func (cmd *CopyImage) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CopyImage) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CopyImageOutput).ImageId)
}

type ImportImage struct {
	_            string `action:"import" entity:"image" awsAPI:"ec2" awsCall:"ImportImage" awsInput:"ec2.ImportImageInput" awsOutput:"ec2.ImportImageOutput" awsDryRun:""`
	logger       *logger.Logger
	api          ec2iface.EC2API
	Architecture *string   `awsName:"Architecture" awsType:"awsstr" templateName:"architecture"`
	Description  *string   `awsName:"Description" awsType:"awsstr" templateName:"description"`
	License      *string   `awsName:"LicenseType" awsType:"awsstr" templateName:"license"`
	Platform     *string   `awsName:"Platform" awsType:"awsstr" templateName:"platform"`
	Role         *string   `awsName:"RoleName" awsType:"awsstr" templateName:"role"`
	Snapshot     *struct{} `awsName:"DiskContainers[0]SnapshotId" awsType:"awsslicestruct" templateName:"snapshot"`
	Url          *struct{} `awsName:"DiskContainers[0]Url" awsType:"awsslicestruct" templateName:"url"`
	Bucket       *struct{} `awsName:"DiskContainers[0]UserBucket.S3Bucket" awsType:"awsslicestruct" templateName:"bucket"`
	S3object     *struct{} `awsName:"DiskContainers[0]UserBucket.S3Key" awsType:"awsslicestruct" templateName:"s3object"`
}

func (cmd *ImportImage) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *ImportImage) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.ImportImageOutput).ImportTaskId)
}

type DeleteImage struct {
	_               string `action:"delete" entity:"image" awsAPI:"ec2"`
	logger          *logger.Logger
	api             ec2iface.EC2API
	Id              *struct{} `templateName:"id" required:""`
	DeleteSnapshots *struct{} `templateName:"delete-snapshots" required:""`
}
