//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package streamanalytics

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2016-03-01/streamanalytics"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BindingType = original.BindingType

const (
	BindingTypeFunctionRetrieveDefaultDefinitionParameters BindingType = original.BindingTypeFunctionRetrieveDefaultDefinitionParameters
	BindingTypeMicrosoftMachineLearningWebService          BindingType = original.BindingTypeMicrosoftMachineLearningWebService
	BindingTypeMicrosoftStreamAnalyticsJavascriptUdf       BindingType = original.BindingTypeMicrosoftStreamAnalyticsJavascriptUdf
)

type CompatibilityLevel = original.CompatibilityLevel

const (
	OneFullStopZero CompatibilityLevel = original.OneFullStopZero
)

type Encoding = original.Encoding

const (
	UTF8 Encoding = original.UTF8
)

type EventsOutOfOrderPolicy = original.EventsOutOfOrderPolicy

const (
	Adjust EventsOutOfOrderPolicy = original.Adjust
	Drop   EventsOutOfOrderPolicy = original.Drop
)

type JSONOutputSerializationFormat = original.JSONOutputSerializationFormat

const (
	Array         JSONOutputSerializationFormat = original.Array
	LineSeparated JSONOutputSerializationFormat = original.LineSeparated
)

type OutputErrorPolicy = original.OutputErrorPolicy

const (
	OutputErrorPolicyDrop OutputErrorPolicy = original.OutputErrorPolicyDrop
	OutputErrorPolicyStop OutputErrorPolicy = original.OutputErrorPolicyStop
)

type OutputStartMode = original.OutputStartMode

const (
	CustomTime          OutputStartMode = original.CustomTime
	JobStartTime        OutputStartMode = original.JobStartTime
	LastOutputEventTime OutputStartMode = original.LastOutputEventTime
)

type SkuName = original.SkuName

const (
	Standard SkuName = original.Standard
)

type Type = original.Type

const (
	TypeAvro          Type = original.TypeAvro
	TypeCsv           Type = original.TypeCsv
	TypeJSON          Type = original.TypeJSON
	TypeSerialization Type = original.TypeSerialization
)

type TypeBasicFunctionBinding = original.TypeBasicFunctionBinding

const (
	TypeFunctionBinding                       TypeBasicFunctionBinding = original.TypeFunctionBinding
	TypeMicrosoftMachineLearningWebService    TypeBasicFunctionBinding = original.TypeMicrosoftMachineLearningWebService
	TypeMicrosoftStreamAnalyticsJavascriptUdf TypeBasicFunctionBinding = original.TypeMicrosoftStreamAnalyticsJavascriptUdf
)

type TypeBasicFunctionProperties = original.TypeBasicFunctionProperties

const (
	TypeFunctionProperties TypeBasicFunctionProperties = original.TypeFunctionProperties
	TypeScalar             TypeBasicFunctionProperties = original.TypeScalar
)

type TypeBasicInputProperties = original.TypeBasicInputProperties

const (
	TypeInputProperties TypeBasicInputProperties = original.TypeInputProperties
	TypeReference       TypeBasicInputProperties = original.TypeReference
	TypeStream          TypeBasicInputProperties = original.TypeStream
)

type TypeBasicOutputDataSource = original.TypeBasicOutputDataSource

const (
	TypeMicrosoftDataLakeAccounts   TypeBasicOutputDataSource = original.TypeMicrosoftDataLakeAccounts
	TypeMicrosoftServiceBusEventHub TypeBasicOutputDataSource = original.TypeMicrosoftServiceBusEventHub
	TypeMicrosoftServiceBusQueue    TypeBasicOutputDataSource = original.TypeMicrosoftServiceBusQueue
	TypeMicrosoftServiceBusTopic    TypeBasicOutputDataSource = original.TypeMicrosoftServiceBusTopic
	TypeMicrosoftSQLServerDatabase  TypeBasicOutputDataSource = original.TypeMicrosoftSQLServerDatabase
	TypeMicrosoftStorageBlob        TypeBasicOutputDataSource = original.TypeMicrosoftStorageBlob
	TypeMicrosoftStorageDocumentDB  TypeBasicOutputDataSource = original.TypeMicrosoftStorageDocumentDB
	TypeMicrosoftStorageTable       TypeBasicOutputDataSource = original.TypeMicrosoftStorageTable
	TypeOutputDataSource            TypeBasicOutputDataSource = original.TypeOutputDataSource
	TypePowerBI                     TypeBasicOutputDataSource = original.TypePowerBI
)

type TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSource

const (
	TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob     TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSourceTypeMicrosoftStorageBlob
	TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource TypeBasicReferenceInputDataSource = original.TypeBasicReferenceInputDataSourceTypeReferenceInputDataSource
)

type TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSource

const (
	TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs     TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftDevicesIotHubs
	TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftServiceBusEventHub
	TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob        TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeMicrosoftStorageBlob
	TypeBasicStreamInputDataSourceTypeStreamInputDataSource       TypeBasicStreamInputDataSource = original.TypeBasicStreamInputDataSourceTypeStreamInputDataSource
)

type UdfType = original.UdfType

const (
	Scalar UdfType = original.Scalar
)

type AvroSerialization = original.AvroSerialization
type AzureDataLakeStoreOutputDataSource = original.AzureDataLakeStoreOutputDataSource
type AzureDataLakeStoreOutputDataSourceProperties = original.AzureDataLakeStoreOutputDataSourceProperties
type AzureMachineLearningWebServiceFunctionBinding = original.AzureMachineLearningWebServiceFunctionBinding
type AzureMachineLearningWebServiceFunctionBindingProperties = original.AzureMachineLearningWebServiceFunctionBindingProperties
type AzureMachineLearningWebServiceFunctionBindingRetrievalProperties = original.AzureMachineLearningWebServiceFunctionBindingRetrievalProperties
type AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters = original.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters
type AzureMachineLearningWebServiceInputColumn = original.AzureMachineLearningWebServiceInputColumn
type AzureMachineLearningWebServiceInputs = original.AzureMachineLearningWebServiceInputs
type AzureMachineLearningWebServiceOutputColumn = original.AzureMachineLearningWebServiceOutputColumn
type AzureSQLDatabaseDataSourceProperties = original.AzureSQLDatabaseDataSourceProperties
type AzureSQLDatabaseOutputDataSource = original.AzureSQLDatabaseOutputDataSource
type AzureSQLDatabaseOutputDataSourceProperties = original.AzureSQLDatabaseOutputDataSourceProperties
type AzureTableOutputDataSource = original.AzureTableOutputDataSource
type AzureTableOutputDataSourceProperties = original.AzureTableOutputDataSourceProperties
type BaseClient = original.BaseClient
type BasicFunctionBinding = original.BasicFunctionBinding
type BasicFunctionProperties = original.BasicFunctionProperties
type BasicFunctionRetrieveDefaultDefinitionParameters = original.BasicFunctionRetrieveDefaultDefinitionParameters
type BasicInputProperties = original.BasicInputProperties
type BasicOutputDataSource = original.BasicOutputDataSource
type BasicReferenceInputDataSource = original.BasicReferenceInputDataSource
type BasicSerialization = original.BasicSerialization
type BasicStreamInputDataSource = original.BasicStreamInputDataSource
type BlobDataSourceProperties = original.BlobDataSourceProperties
type BlobOutputDataSource = original.BlobOutputDataSource
type BlobOutputDataSourceProperties = original.BlobOutputDataSourceProperties
type BlobReferenceInputDataSource = original.BlobReferenceInputDataSource
type BlobReferenceInputDataSourceProperties = original.BlobReferenceInputDataSourceProperties
type BlobStreamInputDataSource = original.BlobStreamInputDataSource
type BlobStreamInputDataSourceProperties = original.BlobStreamInputDataSourceProperties
type CsvSerialization = original.CsvSerialization
type CsvSerializationProperties = original.CsvSerializationProperties
type DiagnosticCondition = original.DiagnosticCondition
type Diagnostics = original.Diagnostics
type DocumentDbOutputDataSource = original.DocumentDbOutputDataSource
type DocumentDbOutputDataSourceProperties = original.DocumentDbOutputDataSourceProperties
type ErrorResponse = original.ErrorResponse
type EventHubDataSourceProperties = original.EventHubDataSourceProperties
type EventHubOutputDataSource = original.EventHubOutputDataSource
type EventHubOutputDataSourceProperties = original.EventHubOutputDataSourceProperties
type EventHubStreamInputDataSource = original.EventHubStreamInputDataSource
type EventHubStreamInputDataSourceProperties = original.EventHubStreamInputDataSourceProperties
type Function = original.Function
type FunctionBinding = original.FunctionBinding
type FunctionInput = original.FunctionInput
type FunctionListResult = original.FunctionListResult
type FunctionListResultIterator = original.FunctionListResultIterator
type FunctionListResultPage = original.FunctionListResultPage
type FunctionOutput = original.FunctionOutput
type FunctionProperties = original.FunctionProperties
type FunctionRetrieveDefaultDefinitionParameters = original.FunctionRetrieveDefaultDefinitionParameters
type FunctionsClient = original.FunctionsClient
type FunctionsTestFuture = original.FunctionsTestFuture
type Input = original.Input
type InputListResult = original.InputListResult
type InputListResultIterator = original.InputListResultIterator
type InputListResultPage = original.InputListResultPage
type InputProperties = original.InputProperties
type InputsClient = original.InputsClient
type InputsTestFuture = original.InputsTestFuture
type IoTHubStreamInputDataSource = original.IoTHubStreamInputDataSource
type IoTHubStreamInputDataSourceProperties = original.IoTHubStreamInputDataSourceProperties
type JSONSerialization = original.JSONSerialization
type JSONSerializationProperties = original.JSONSerializationProperties
type JavaScriptFunctionBinding = original.JavaScriptFunctionBinding
type JavaScriptFunctionBindingProperties = original.JavaScriptFunctionBindingProperties
type JavaScriptFunctionBindingRetrievalProperties = original.JavaScriptFunctionBindingRetrievalProperties
type JavaScriptFunctionRetrieveDefaultDefinitionParameters = original.JavaScriptFunctionRetrieveDefaultDefinitionParameters
type OAuthBasedDataSourceProperties = original.OAuthBasedDataSourceProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type Output = original.Output
type OutputDataSource = original.OutputDataSource
type OutputListResult = original.OutputListResult
type OutputListResultIterator = original.OutputListResultIterator
type OutputListResultPage = original.OutputListResultPage
type OutputProperties = original.OutputProperties
type OutputsClient = original.OutputsClient
type OutputsTestFuture = original.OutputsTestFuture
type PowerBIOutputDataSource = original.PowerBIOutputDataSource
type PowerBIOutputDataSourceProperties = original.PowerBIOutputDataSourceProperties
type ReferenceInputDataSource = original.ReferenceInputDataSource
type ReferenceInputProperties = original.ReferenceInputProperties
type Resource = original.Resource
type ResourceTestStatus = original.ResourceTestStatus
type ScalarFunctionConfiguration = original.ScalarFunctionConfiguration
type ScalarFunctionProperties = original.ScalarFunctionProperties
type Serialization = original.Serialization
type ServiceBusDataSourceProperties = original.ServiceBusDataSourceProperties
type ServiceBusQueueOutputDataSource = original.ServiceBusQueueOutputDataSource
type ServiceBusQueueOutputDataSourceProperties = original.ServiceBusQueueOutputDataSourceProperties
type ServiceBusTopicOutputDataSource = original.ServiceBusTopicOutputDataSource
type ServiceBusTopicOutputDataSourceProperties = original.ServiceBusTopicOutputDataSourceProperties
type Sku = original.Sku
type StartStreamingJobParameters = original.StartStreamingJobParameters
type StorageAccount = original.StorageAccount
type StreamInputDataSource = original.StreamInputDataSource
type StreamInputProperties = original.StreamInputProperties
type StreamingJob = original.StreamingJob
type StreamingJobListResult = original.StreamingJobListResult
type StreamingJobListResultIterator = original.StreamingJobListResultIterator
type StreamingJobListResultPage = original.StreamingJobListResultPage
type StreamingJobProperties = original.StreamingJobProperties
type StreamingJobsClient = original.StreamingJobsClient
type StreamingJobsCreateOrReplaceFuture = original.StreamingJobsCreateOrReplaceFuture
type StreamingJobsDeleteFuture = original.StreamingJobsDeleteFuture
type StreamingJobsStartFuture = original.StreamingJobsStartFuture
type StreamingJobsStopFuture = original.StreamingJobsStopFuture
type SubResource = original.SubResource
type SubscriptionQuota = original.SubscriptionQuota
type SubscriptionQuotaProperties = original.SubscriptionQuotaProperties
type SubscriptionQuotasListResult = original.SubscriptionQuotasListResult
type SubscriptionsClient = original.SubscriptionsClient
type Transformation = original.Transformation
type TransformationProperties = original.TransformationProperties
type TransformationsClient = original.TransformationsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewFunctionListResultIterator(page FunctionListResultPage) FunctionListResultIterator {
	return original.NewFunctionListResultIterator(page)
}
func NewFunctionListResultPage(cur FunctionListResult, getNextPage func(context.Context, FunctionListResult) (FunctionListResult, error)) FunctionListResultPage {
	return original.NewFunctionListResultPage(cur, getNextPage)
}
func NewFunctionsClient(subscriptionID string) FunctionsClient {
	return original.NewFunctionsClient(subscriptionID)
}
func NewFunctionsClientWithBaseURI(baseURI string, subscriptionID string) FunctionsClient {
	return original.NewFunctionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInputListResultIterator(page InputListResultPage) InputListResultIterator {
	return original.NewInputListResultIterator(page)
}
func NewInputListResultPage(cur InputListResult, getNextPage func(context.Context, InputListResult) (InputListResult, error)) InputListResultPage {
	return original.NewInputListResultPage(cur, getNextPage)
}
func NewInputsClient(subscriptionID string) InputsClient {
	return original.NewInputsClient(subscriptionID)
}
func NewInputsClientWithBaseURI(baseURI string, subscriptionID string) InputsClient {
	return original.NewInputsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOutputListResultIterator(page OutputListResultPage) OutputListResultIterator {
	return original.NewOutputListResultIterator(page)
}
func NewOutputListResultPage(cur OutputListResult, getNextPage func(context.Context, OutputListResult) (OutputListResult, error)) OutputListResultPage {
	return original.NewOutputListResultPage(cur, getNextPage)
}
func NewOutputsClient(subscriptionID string) OutputsClient {
	return original.NewOutputsClient(subscriptionID)
}
func NewOutputsClientWithBaseURI(baseURI string, subscriptionID string) OutputsClient {
	return original.NewOutputsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStreamingJobListResultIterator(page StreamingJobListResultPage) StreamingJobListResultIterator {
	return original.NewStreamingJobListResultIterator(page)
}
func NewStreamingJobListResultPage(cur StreamingJobListResult, getNextPage func(context.Context, StreamingJobListResult) (StreamingJobListResult, error)) StreamingJobListResultPage {
	return original.NewStreamingJobListResultPage(cur, getNextPage)
}
func NewStreamingJobsClient(subscriptionID string) StreamingJobsClient {
	return original.NewStreamingJobsClient(subscriptionID)
}
func NewStreamingJobsClientWithBaseURI(baseURI string, subscriptionID string) StreamingJobsClient {
	return original.NewStreamingJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransformationsClient(subscriptionID string) TransformationsClient {
	return original.NewTransformationsClient(subscriptionID)
}
func NewTransformationsClientWithBaseURI(baseURI string, subscriptionID string) TransformationsClient {
	return original.NewTransformationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleBindingTypeValues() []BindingType {
	return original.PossibleBindingTypeValues()
}
func PossibleCompatibilityLevelValues() []CompatibilityLevel {
	return original.PossibleCompatibilityLevelValues()
}
func PossibleEncodingValues() []Encoding {
	return original.PossibleEncodingValues()
}
func PossibleEventsOutOfOrderPolicyValues() []EventsOutOfOrderPolicy {
	return original.PossibleEventsOutOfOrderPolicyValues()
}
func PossibleJSONOutputSerializationFormatValues() []JSONOutputSerializationFormat {
	return original.PossibleJSONOutputSerializationFormatValues()
}
func PossibleOutputErrorPolicyValues() []OutputErrorPolicy {
	return original.PossibleOutputErrorPolicyValues()
}
func PossibleOutputStartModeValues() []OutputStartMode {
	return original.PossibleOutputStartModeValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleTypeBasicFunctionBindingValues() []TypeBasicFunctionBinding {
	return original.PossibleTypeBasicFunctionBindingValues()
}
func PossibleTypeBasicFunctionPropertiesValues() []TypeBasicFunctionProperties {
	return original.PossibleTypeBasicFunctionPropertiesValues()
}
func PossibleTypeBasicInputPropertiesValues() []TypeBasicInputProperties {
	return original.PossibleTypeBasicInputPropertiesValues()
}
func PossibleTypeBasicOutputDataSourceValues() []TypeBasicOutputDataSource {
	return original.PossibleTypeBasicOutputDataSourceValues()
}
func PossibleTypeBasicReferenceInputDataSourceValues() []TypeBasicReferenceInputDataSource {
	return original.PossibleTypeBasicReferenceInputDataSourceValues()
}
func PossibleTypeBasicStreamInputDataSourceValues() []TypeBasicStreamInputDataSource {
	return original.PossibleTypeBasicStreamInputDataSourceValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUdfTypeValues() []UdfType {
	return original.PossibleUdfTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
