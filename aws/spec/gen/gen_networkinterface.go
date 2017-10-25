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

type CreateNetworkinterface struct {
	_              string `action:"create" entity:"networkinterface" awsAPI:"ec2" awsCall:"CreateNetworkInterface" awsInput:"ec2.CreateNetworkInterfaceInput" awsOutput:"ec2.CreateNetworkInterfaceOutput" awsDryRun:""`
	logger         *logger.Logger
	api            ec2iface.EC2API
	Subnet         *string   `awsName:"SubnetId" awsType:"awsstr" templateName:"subnet" required:""`
	Description    *string   `awsName:"Description" awsType:"awsstr" templateName:"description"`
	Securitygroups *[]string `awsName:"Groups" awsType:"awsstringslice" templateName:"securitygroups"`
	Privateip      *string   `awsName:"PrivateIpAddress" awsType:"awsstr" templateName:"privateip"`
}

func (cmd *CreateNetworkinterface) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateNetworkinterface) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.CreateNetworkInterfaceOutput).NetworkInterface.NetworkInterfaceId)
}

type DeleteNetworkinterface struct {
	_      string `action:"delete" entity:"networkinterface" awsAPI:"ec2" awsCall:"DeleteNetworkInterface" awsInput:"ec2.DeleteNetworkInterfaceInput" awsOutput:"ec2.DeleteNetworkInterfaceOutput" awsDryRun:""`
	logger *logger.Logger
	api    ec2iface.EC2API
	Id     *string `awsName:"NetworkInterfaceId" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteNetworkinterface) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

type AttachNetworkinterface struct {
	_           string `action:"attach" entity:"networkinterface" awsAPI:"ec2" awsCall:"AttachNetworkInterface" awsInput:"ec2.AttachNetworkInterfaceInput" awsOutput:"ec2.AttachNetworkInterfaceOutput" awsDryRun:""`
	logger      *logger.Logger
	api         ec2iface.EC2API
	Id          *string `awsName:"NetworkInterfaceId" awsType:"awsstr" templateName:"id" required:""`
	Instance    *string `awsName:"InstanceId" awsType:"awsstr" templateName:"instance" required:""`
	DeviceIndex *int64  `awsName:"DeviceIndex" awsType:"awsint64" templateName:"device-index" required:""`
}

func (cmd *AttachNetworkinterface) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *AttachNetworkinterface) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*ec2.AttachNetworkInterfaceOutput).AttachmentId)
}

type DetachNetworkinterface struct {
	_          string `action:"detach" entity:"networkinterface" awsAPI:"ec2"`
	logger     *logger.Logger
	api        ec2iface.EC2API
	Attachment *string `awsName:"AttachmentId" awsType:"awsstr" templateName:"attachment"`
	Instance   *string `awsName:"InstanceId" awsType:"awsstr" templateName:"instance"`
	Id         *string `awsName:"NetworkInterfaceId" awsType:"awsstr" templateName:"id"`
	Force      *bool   `awsName:"Force" awsType:"awsbool" templateName:"force"`
}

type CheckNetworkinterface struct {
	_       string `action:"check" entity:"networkinterface" awsAPI:"ec2"`
	logger  *logger.Logger
	api     ec2iface.EC2API
	Id      *struct{} `templateName:"id" required:""`
	State   *struct{} `templateName:"state" required:""`
	Timeout *struct{} `templateName:"timeout" required:""`
}
