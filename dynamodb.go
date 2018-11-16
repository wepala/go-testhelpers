package testhelpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

//setup dynamo db
type MockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
	functions map[string]interface{}
}

type BatchGetItemFunc func(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
type BatchGetItemWithContextFunc func(context aws.Context, input *dynamodb.BatchGetItemInput, options ...request.Option) (*dynamodb.BatchGetItemOutput, error)
type BatchGetItemRequestFunc func(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput)
type BatchGetItemPagesFunc func(input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool) error
type BatchGetItemPagesWithContextFunc func(context aws.Context, input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool, options ...request.Option) error
type BatchWriteItemFunc func(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
type BatchWriteItemWithContextFunc func(context aws.Context, input *dynamodb.BatchWriteItemInput, options ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
type BatchWriteItemRequestFunc func(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput)
type CreateBackupFunc func(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error)
type CreateBackupWithContextFunc func(context aws.Context, input *dynamodb.CreateBackupInput, options ...request.Option) (*dynamodb.CreateBackupOutput, error)
type CreateBackupRequestFunc func(input *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput)
type CreateGlobalTableFunc func(input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error)
type CreateGlobalTableWithContextFunc func(context aws.Context, input *dynamodb.CreateGlobalTableInput, options ...request.Option) (*dynamodb.CreateGlobalTableOutput, error)
type CreateGlobalTableRequestFunc func(input *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput)
type CreateTableFunc func(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
type CreateTableWithContextFunc func(context aws.Context, input *dynamodb.CreateTableInput, options ...request.Option) (*dynamodb.CreateTableOutput, error)
type CreateTableRequestFunc func(input *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput)
type DeleteBackupFunc func(input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error)
type DeleteBackupWithContextFunc func(context aws.Context, input *dynamodb.DeleteBackupInput, options ...request.Option) (*dynamodb.DeleteBackupOutput, error)
type DeleteBackupRequestFunc func(input *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput)
type DeleteItemFunc func(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
type DeleteItemWithContextFunc func(context aws.Context, input *dynamodb.DeleteItemInput, options ...request.Option) (*dynamodb.DeleteItemOutput, error)
type DeleteItemRequestFunc func(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput)
type DeleteTableFunc func(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)
type DeleteTableWithContextFunc func(context aws.Context, input *dynamodb.DeleteTableInput, options ...request.Option) (*dynamodb.DeleteTableOutput, error)
type DeleteTableRequestFunc func(input *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput)
type DescribeBackupFunc func(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error)
type DescribeBackupWithContextFunc func(context aws.Context, input *dynamodb.DescribeBackupInput, options ...request.Option) (*dynamodb.DescribeBackupOutput, error)
type DescribeBackupRequestFunc func(input *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput)
type DescribeContinuousBackupsFunc func(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error)
type DescribeContinuousBackupsWithContextFunc func(context aws.Context, input *dynamodb.DescribeContinuousBackupsInput, options ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error)
type DescribeContinuousBackupsRequestFunc func(input *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput)
type DescribeGlobalTableFunc func(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error)
type DescribeGlobalTableWithContextFunc func(context aws.Context, input *dynamodb.DescribeGlobalTableInput, options ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error)
type DescribeGlobalTableRequestFunc func(input *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput)
type DescribeGlobalTableSettingsFunc func(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
type DescribeGlobalTableSettingsWithContextFunc func(context aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, options ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
type DescribeGlobalTableSettingsRequestFunc func(input *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput)
type DescribeLimitsFunc func(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error)
type DescribeLimitsWithContextFunc func(context aws.Context, input *dynamodb.DescribeLimitsInput, options ...request.Option) (*dynamodb.DescribeLimitsOutput, error)
type DescribeLimitsRequestFunc func(input *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput)
type DescribeTableFunc func(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
type DescribeTableWithContextFunc func(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.Option) (*dynamodb.DescribeTableOutput, error)
type DescribeTableRequestFunc func(input *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput)
type DescribeTimeToLiveFunc func(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error)
type DescribeTimeToLiveWithContextFunc func(context aws.Context, input *dynamodb.DescribeTimeToLiveInput, options ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error)
type DescribeTimeToLiveRequestFunc func(input *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput)
type GetItemFunc func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
type GetItemWithContextFunc func(context aws.Context, input *dynamodb.GetItemInput, options ...request.Option) (*dynamodb.GetItemOutput, error)
type GetItemRequestFunc func(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput)
type ListBackupsFunc func(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error)
type ListBackupsWithContextFunc func(context aws.Context, input *dynamodb.ListBackupsInput, options ...request.Option) (*dynamodb.ListBackupsOutput, error)
type ListBackupsRequestFunc func(input *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput)
type ListGlobalTablesFunc func(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error)
type ListGlobalTablesWithContextFunc func(context aws.Context, input *dynamodb.ListGlobalTablesInput, options ...request.Option) (*dynamodb.ListGlobalTablesOutput, error)
type ListGlobalTablesRequestFunc func(input *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput)
type ListTablesFunc func(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)
type ListTablesWithContextFunc func(context aws.Context, input *dynamodb.ListTablesInput, options ...request.Option) (*dynamodb.ListTablesOutput, error)
type ListTablesRequestFunc func(input *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput)
type ListTablesPagesFunc func(input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool) error
type ListTablesPagesWithContextFunc func(context aws.Context, input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool, options ...request.Option) error
type ListTagsOfResourceFunc func(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error)
type ListTagsOfResourceWithContextFunc func(context aws.Context, input *dynamodb.ListTagsOfResourceInput, options ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error)
type ListTagsOfResourceRequestFunc func(input *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput)
type PutItemFunc func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
type PutItemWithContextFunc func(context aws.Context, input *dynamodb.PutItemInput, options ...request.Option) (*dynamodb.PutItemOutput, error)
type PutItemRequestFunc func(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput)
type QueryFunc func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
type QueryWithContextFunc func(context aws.Context, input *dynamodb.QueryInput, options ...request.Option) (*dynamodb.QueryOutput, error)
type QueryRequestFunc func(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput)
type QueryPagesFunc func(input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool) error
type QueryPagesWithContextFunc func(context aws.Context, input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool, options ...request.Option) error
type RestoreTableFromBackupFunc func(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error)
type RestoreTableFromBackupWithContextFunc func(context aws.Context, input *dynamodb.RestoreTableFromBackupInput, options ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error)
type RestoreTableFromBackupRequestFunc func(input *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput)
type RestoreTableToPointInTimeFunc func(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error)
type RestoreTableToPointInTimeWithContextFunc func(context aws.Context, input *dynamodb.RestoreTableToPointInTimeInput, options ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error)
type RestoreTableToPointInTimeRequestFunc func(input *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput)
type ScanFunc func(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
type ScanWithContextFunc func(context aws.Context, input *dynamodb.ScanInput, options ...request.Option) (*dynamodb.ScanOutput, error)
type ScanRequestFunc func(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput)
type ScanPagesFunc func(input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool) error
type ScanPagesWithContextFunc func(context aws.Context, input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool, options ...request.Option) error
type TagResourceFunc func(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error)
type TagResourceWithContextFunc func(context aws.Context, input *dynamodb.TagResourceInput, options ...request.Option) (*dynamodb.TagResourceOutput, error)
type TagResourceRequestFunc func(input *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput)
type UntagResourceFunc func(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error)
type UntagResourceWithContextFunc func(context aws.Context, input *dynamodb.UntagResourceInput, options ...request.Option) (*dynamodb.UntagResourceOutput, error)
type UntagResourceRequestFunc func(input *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput)
type UpdateContinuousBackupsFunc func(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error)
type UpdateContinuousBackupsWithContextFunc func(context aws.Context, input *dynamodb.UpdateContinuousBackupsInput, options ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error)
type UpdateContinuousBackupsRequestFunc func(input *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput)
type UpdateGlobalTableFunc func(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error)
type UpdateGlobalTableWithContextFunc func(context aws.Context, input *dynamodb.UpdateGlobalTableInput, options ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error)
type UpdateGlobalTableRequestFunc func(input *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput)
type UpdateGlobalTableSettingsFunc func(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
type UpdateGlobalTableSettingsWithContextFunc func(context aws.Context, input *dynamodb.UpdateGlobalTableSettingsInput, options ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
type UpdateGlobalTableSettingsRequestFunc func(input *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput)
type UpdateItemFunc func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
type UpdateItemWithContextFunc func(context aws.Context, input *dynamodb.UpdateItemInput, options ...request.Option) (*dynamodb.UpdateItemOutput, error)
type UpdateItemRequestFunc func(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput)
type UpdateTableFunc func(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
type UpdateTableWithContextFunc func(context aws.Context, input *dynamodb.UpdateTableInput, options ...request.Option) (*dynamodb.UpdateTableOutput, error)
type UpdateTableRequestFunc func(input *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput)
type UpdateTimeToLiveFunc func(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error)
type UpdateTimeToLiveWithContextFunc func(context aws.Context, input *dynamodb.UpdateTimeToLiveInput, options ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error)
type UpdateTimeToLiveRequestFunc func(input *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput)
type WaitUntilTableExistsFunc func(input *dynamodb.DescribeTableInput) error
type WaitUntilTableExistsWithContextFunc func(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error
type WaitUntilTableNotExistsFunc func(input *dynamodb.DescribeTableInput) error
type WaitUntilTableNotExistsWithContextFunc func(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error

func (mdb *MockDynamoDBClient) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return mdb.functions["BatchGetItem"].(BatchGetItemFunc)(input)
}

func (mdb *MockDynamoDBClient) BatchGetItemWithContext(context aws.Context, input *dynamodb.BatchGetItemInput, options ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return mdb.functions["BatchGetItemWithContext"].(BatchGetItemWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return mdb.functions["BatchGetItemRequest"].(BatchGetItemRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) BatchGetItemPages(input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return mdb.functions["BatchGetItemPages"].(BatchGetItemPagesFunc)(input, callback)
}

func (mdb *MockDynamoDBClient) BatchGetItemPagesWithContext(context aws.Context, input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["BatchGetItemPagesWithContext"].(BatchGetItemPagesWithContextFunc)(context, input, callback)
}

func (mdb *MockDynamoDBClient) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return mdb.functions["BatchWriteItem"].(BatchWriteItemFunc)(input)
}

func (mdb *MockDynamoDBClient) BatchWriteItemWithContext(context aws.Context, input *dynamodb.BatchWriteItemInput, options ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return mdb.functions["BatchWriteItemWithContext"].(BatchWriteItemWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return mdb.functions["BatchWriteItemRequest"].(BatchWriteItemRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) CreateBackup(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return mdb.functions["CreateBackup"].(CreateBackupFunc)(input)
}

//func (mdb *MockDynamoDBClient) CreateBackupWithContext (context aws.Context, input  *dynamodb.CreateBackupInput, options ...request.Option) (*dynamodb.CreateBackupOutput, error) {
//  return mdb.functions["CreateBackupWithContext"].(CreateBackupWithContextFunc)(context aws.Context, input  *dynamodb.CreateBackupInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) CreateBackupRequest (input *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
//  return mdb.functions["CreateBackupRequest"].(CreateBackupRequestFunc)(input *dynamodb.CreateBackupInput)
//}
//
//func (mdb *MockDynamoDBClient) CreateGlobalTable (input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
//  return mdb.functions["CreateGlobalTable"].(CreateGlobalTableFunc)(input *dynamodb.CreateGlobalTableInput)
//}
//
//func (mdb *MockDynamoDBClient) CreateGlobalTableWithContext (context aws.Context, input  *dynamodb.CreateGlobalTableInput, options ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
//  return mdb.functions["CreateGlobalTableWithContext"].(CreateGlobalTableWithContextFunc)(context aws.Context, input  *dynamodb.CreateGlobalTableInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) CreateGlobalTableRequest (input *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
//  return mdb.functions["CreateGlobalTableRequest"].(CreateGlobalTableRequestFunc)(input *dynamodb.CreateGlobalTableInput)
//}
//
//func (mdb *MockDynamoDBClient) CreateTable (input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
//  return mdb.functions["CreateTable"].(CreateTableFunc)(input *dynamodb.CreateTableInput)
//}
//
//func (mdb *MockDynamoDBClient) CreateTableWithContext (context aws.Context, input  *dynamodb.CreateTableInput, options ...request.Option) (*dynamodb.CreateTableOutput, error) {
//  return mdb.functions["CreateTableWithContext"].(CreateTableWithContextFunc)(context aws.Context, input  *dynamodb.CreateTableInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) CreateTableRequest (input *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
//  return mdb.functions["CreateTableRequest"].(CreateTableRequestFunc)(input *dynamodb.CreateTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DeleteBackup (input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
//  return mdb.functions["DeleteBackup"].(DeleteBackupFunc)(input *dynamodb.DeleteBackupInput)
//}
//
//func (mdb *MockDynamoDBClient) DeleteBackupWithContext (context aws.Context, input  *dynamodb.DeleteBackupInput, options ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
//  return mdb.functions["DeleteBackupWithContext"].(DeleteBackupWithContextFunc)(context aws.Context, input  *dynamodb.DeleteBackupInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DeleteBackupRequest (input *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
//  return mdb.functions["DeleteBackupRequest"].(DeleteBackupRequestFunc)(input *dynamodb.DeleteBackupInput)
//}
//
//func (mdb *MockDynamoDBClient) DeleteItem (input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
//  return mdb.functions["DeleteItem"].(DeleteItemFunc)(input *dynamodb.DeleteItemInput)
//}
//
//func (mdb *MockDynamoDBClient) DeleteItemWithContext (context aws.Context, input *dynamodb.DeleteItemInput, options ...request.Option) (*dynamodb.DeleteItemOutput, error) {
//  return mdb.functions["DeleteItemWithContext"].(DeleteItemWithContextFunc)(context aws.Context, input *dynamodb.DeleteItemInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DeleteItemRequest (input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
//  return mdb.functions["DeleteItemRequest"].(DeleteItemRequestFunc)(input *dynamodb.DeleteItemInput)
//}
//
//func (mdb *MockDynamoDBClient) DeleteTable (input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
//  return mdb.functions["DeleteTable"].(DeleteTableFunc)(input *dynamodb.DeleteTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DeleteTableWithContext (context aws.Context, input *dynamodb.DeleteTableInput, options ...request.Option) (*dynamodb.DeleteTableOutput, error) {
//  return mdb.functions["DeleteTableWithContext"].(DeleteTableWithContextFunc)(context aws.Context, input *dynamodb.DeleteTableInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DeleteTableRequest (input *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
//  return mdb.functions["DeleteTableRequest"].(DeleteTableRequestFunc)(input *dynamodb.DeleteTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeBackup (input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
//  return mdb.functions["DescribeBackup"].(DescribeBackupFunc)(input *dynamodb.DescribeBackupInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeBackupWithContext (context aws.Context, input *dynamodb.DescribeBackupInput, options ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
//  return mdb.functions["DescribeBackupWithContext"].(DescribeBackupWithContextFunc)(context aws.Context, input *dynamodb.DescribeBackupInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeBackupRequest (input *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
//  return mdb.functions["DescribeBackupRequest"].(DescribeBackupRequestFunc)(input *dynamodb.DescribeBackupInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeContinuousBackups (input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
//  return mdb.functions["DescribeContinuousBackups"].(DescribeContinuousBackupsFunc)(input *dynamodb.DescribeContinuousBackupsInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeContinuousBackupsWithContext (context aws.Context, input *dynamodb.DescribeContinuousBackupsInput, options ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
//  return mdb.functions["DescribeContinuousBackupsWithContext"].(DescribeContinuousBackupsWithContextFunc)(context aws.Context, input *dynamodb.DescribeContinuousBackupsInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeContinuousBackupsRequest (input *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
//  return mdb.functions["DescribeContinuousBackupsRequest"].(DescribeContinuousBackupsRequestFunc)(input *dynamodb.DescribeContinuousBackupsInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeGlobalTable (input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
//  return mdb.functions["DescribeGlobalTable"].(DescribeGlobalTableFunc)(input *dynamodb.DescribeGlobalTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeGlobalTableWithContext (context aws.Context, input *dynamodb.DescribeGlobalTableInput, options ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
//  return mdb.functions["DescribeGlobalTableWithContext"].(DescribeGlobalTableWithContextFunc)(context aws.Context, input *dynamodb.DescribeGlobalTableInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeGlobalTableRequest (input *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
//  return mdb.functions["DescribeGlobalTableRequest"].(DescribeGlobalTableRequestFunc)(input *dynamodb.DescribeGlobalTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeGlobalTableSettings (input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
//  return mdb.functions["DescribeGlobalTableSettings"].(DescribeGlobalTableSettingsFunc)(input *dynamodb.DescribeGlobalTableSettingsInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeGlobalTableSettingsWithContext (context aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, options ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
//  return mdb.functions["DescribeGlobalTableSettingsWithContext"].(DescribeGlobalTableSettingsWithContextFunc)(context aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeGlobalTableSettingsRequest (input *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
//  return mdb.functions["DescribeGlobalTableSettingsRequest"].(DescribeGlobalTableSettingsRequestFunc)(input *dynamodb.DescribeGlobalTableSettingsInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeLimits (input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
//  return mdb.functions["DescribeLimits"].(DescribeLimitsFunc)(input *dynamodb.DescribeLimitsInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeLimitsWithContext (context aws.Context, input *dynamodb.DescribeLimitsInput, options ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
//  return mdb.functions["DescribeLimitsWithContext"].(DescribeLimitsWithContextFunc)(context aws.Context, input *dynamodb.DescribeLimitsInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeLimitsRequest (input *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
//  return mdb.functions["DescribeLimitsRequest"].(DescribeLimitsRequestFunc)(input *dynamodb.DescribeLimitsInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeTable (input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
//  return mdb.functions["DescribeTable"].(DescribeTableFunc)(input *dynamodb.DescribeTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeTableWithContext (context aws.Context, input *dynamodb.DescribeTableInput, options ...request.Option) (*dynamodb.DescribeTableOutput, error) {
//  return mdb.functions["DescribeTableWithContext"].(DescribeTableWithContextFunc)(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeTableRequest (input *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
//  return mdb.functions["DescribeTableRequest"].(DescribeTableRequestFunc)(input *dynamodb.DescribeTableInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeTimeToLive (input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
//  return mdb.functions["DescribeTimeToLive"].(DescribeTimeToLiveFunc)(input *dynamodb.DescribeTimeToLiveInput)
//}
//
//func (mdb *MockDynamoDBClient) DescribeTimeToLiveWithContext (context aws.Context, input *dynamodb.DescribeTimeToLiveInput, options ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
//  return mdb.functions["DescribeTimeToLiveWithContext"].(DescribeTimeToLiveWithContextFunc)(context aws.Context, input *dynamodb.DescribeTimeToLiveInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) DescribeTimeToLiveRequest (input *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
//  return mdb.functions["DescribeTimeToLiveRequest"].(DescribeTimeToLiveRequestFunc)(input *dynamodb.DescribeTimeToLiveInput)
//}
//
//func (mdb *MockDynamoDBClient) GetItem (input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
//  return mdb.functions["GetItem"].(GetItemFunc)(input *dynamodb.GetItemInput)
//}
//
//func (mdb *MockDynamoDBClient) GetItemWithContext (context aws.Context, input *dynamodb.GetItemInput, options ...request.Option) (*dynamodb.GetItemOutput, error) {
//  return mdb.functions["GetItemWithContext"].(GetItemWithContextFunc)(context aws.Context, input *dynamodb.GetItemInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) GetItemRequest (input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
//  return mdb.functions["GetItemRequest"].(GetItemRequestFunc)(input *dynamodb.GetItemInput)
//}
//
//func (mdb *MockDynamoDBClient) ListBackups (input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
//  return mdb.functions["ListBackups"].(ListBackupsFunc)(input *dynamodb.ListBackupsInput)
//}
//
//func (mdb *MockDynamoDBClient) ListBackupsWithContext (context aws.Context, input *dynamodb.ListBackupsInput, options ...request.Option) (*dynamodb.ListBackupsOutput, error) {
//  return mdb.functions["ListBackupsWithContext"].(ListBackupsWithContextFunc)(context aws.Context, input *dynamodb.ListBackupsInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) ListBackupsRequest (input *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
//  return mdb.functions["ListBackupsRequest"].(ListBackupsRequestFunc)(input *dynamodb.ListBackupsInput)
//}
//
//func (mdb *MockDynamoDBClient) ListGlobalTables (input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
//  return mdb.functions["ListGlobalTables"].(ListGlobalTablesFunc)(input *dynamodb.ListGlobalTablesInput)
//}
//
//func (mdb *MockDynamoDBClient) ListGlobalTablesWithContext (context aws.Context, input *dynamodb.ListGlobalTablesInput, options ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
//  return mdb.functions["ListGlobalTablesWithContext"].(ListGlobalTablesWithContextFunc)(context aws.Context, input *dynamodb.ListGlobalTablesInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) ListGlobalTablesRequest (input *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
//  return mdb.functions["ListGlobalTablesRequest"].(ListGlobalTablesRequestFunc)(input *dynamodb.ListGlobalTablesInput)
//}
//
//func (mdb *MockDynamoDBClient) ListTables (input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
//  return mdb.functions["ListTables"].(ListTablesFunc)(input *dynamodb.ListTablesInput)
//}
//
//func (mdb *MockDynamoDBClient) ListTablesWithContext (context aws.Context, input  *dynamodb.ListTablesInput, options ...request.Option) (*dynamodb.ListTablesOutput, error) {
//  return mdb.functions["ListTablesWithContext"].(ListTablesWithContextFunc)(context aws.Context, input  *dynamodb.ListTablesInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) ListTablesRequest (input *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
//  return mdb.functions["ListTablesRequest"].(ListTablesRequestFunc)(input *dynamodb.ListTablesInput)
//}
//
//func (mdb *MockDynamoDBClient) ListTablesPages (input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool) error {
//  return mdb.functions["ListTablesPages"].(ListTablesPagesFunc)(input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool)
//}
//
//func (mdb *MockDynamoDBClient) ListTablesPagesWithContext (context aws.Context, input  *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool, options ...request.Option) error {
//  return mdb.functions["ListTablesPagesWithContext"].(ListTablesPagesWithContextFunc)(context aws.Context, input  *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool)
//}
//
//func (mdb *MockDynamoDBClient) ListTagsOfResource (input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
//  return mdb.functions["ListTagsOfResource"].(ListTagsOfResourceFunc)(input *dynamodb.ListTagsOfResourceInput)
//}
//
//func (mdb *MockDynamoDBClient) ListTagsOfResourceWithContext (context aws.Context, input *dynamodb.ListTagsOfResourceInput, options ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
//  return mdb.functions["ListTagsOfResourceWithContext"].(ListTagsOfResourceWithContextFunc)(context aws.Context, input *dynamodb.ListTagsOfResourceInput, options ...request.Option)
//}
//
//func (mdb *MockDynamoDBClient) ListTagsOfResourceRequest (input *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
//  return mdb.functions["ListTagsOfResourceRequest"].(ListTagsOfResourceRequestFunc)(input *dynamodb.ListTagsOfResourceInput)
//}

func (mdb *MockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return mdb.functions["PutItem"].(PutItemFunc)(input)
}

func (mdb *MockDynamoDBClient) PutItemWithContext(context aws.Context, input *dynamodb.PutItemInput, options ...request.Option) (*dynamodb.PutItemOutput, error) {
	return mdb.functions["PutItemWithContext"].(PutItemWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return mdb.functions["PutItemRequest"].(PutItemRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return mdb.functions["Query"].(QueryFunc)(input)
}

func (mdb *MockDynamoDBClient) QueryWithContext(context aws.Context, input *dynamodb.QueryInput, options ...request.Option) (*dynamodb.QueryOutput, error) {
	return mdb.functions["QueryWithContext"].(QueryWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return mdb.functions["QueryRequest"].(QueryRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) QueryPages(input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool) error {
	return mdb.functions["QueryPages"].(QueryPagesFunc)(input, callback)
}

func (mdb *MockDynamoDBClient) QueryPagesWithContext(context aws.Context, input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["QueryPagesWithContext"].(QueryPagesWithContextFunc)(context, input, callback)
}

func (mdb *MockDynamoDBClient) RestoreTableFromBackup(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return mdb.functions["RestoreTableFromBackup"].(RestoreTableFromBackupFunc)(input)
}

func (mdb *MockDynamoDBClient) RestoreTableFromBackupWithContext(context aws.Context, input *dynamodb.RestoreTableFromBackupInput, options ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return mdb.functions["RestoreTableFromBackupWithContext"].(RestoreTableFromBackupWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) RestoreTableFromBackupRequest(input *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	return mdb.functions["RestoreTableFromBackupRequest"].(RestoreTableFromBackupRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) RestoreTableToPointInTime(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return mdb.functions["RestoreTableToPointInTime"].(RestoreTableToPointInTimeFunc)(input)
}

func (mdb *MockDynamoDBClient) RestoreTableToPointInTimeWithContext(context aws.Context, input *dynamodb.RestoreTableToPointInTimeInput, options ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return mdb.functions["RestoreTableToPointInTimeWithContext"].(RestoreTableToPointInTimeWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) RestoreTableToPointInTimeRequest(input *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	return mdb.functions["RestoreTableToPointInTimeRequest"].(RestoreTableToPointInTimeRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return mdb.functions["Scan"].(ScanFunc)(input)
}

func (mdb *MockDynamoDBClient) ScanWithContext(context aws.Context, input *dynamodb.ScanInput, options ...request.Option) (*dynamodb.ScanOutput, error) {
	return mdb.functions["ScanWithContext"].(ScanWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return mdb.functions["ScanRequest"].(ScanRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) ScanPages(input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool) error {
	return mdb.functions["ScanPages"].(ScanPagesFunc)(input, callback)
}

func (mdb *MockDynamoDBClient) ScanPagesWithContext(context aws.Context, input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["ScanPagesWithContext"].(ScanPagesWithContextFunc)(context, input, callback)
}

func (mdb *MockDynamoDBClient) TagResource(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return mdb.functions["TagResource"].(TagResourceFunc)(input)
}

func (mdb *MockDynamoDBClient) TagResourceWithContext(context aws.Context, input *dynamodb.TagResourceInput, options ...request.Option) (*dynamodb.TagResourceOutput, error) {
	return mdb.functions["TagResourceWithContext"].(TagResourceWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) TagResourceRequest(input *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	return mdb.functions["TagResourceRequest"].(TagResourceRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UntagResource(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return mdb.functions["UntagResource"].(UntagResourceFunc)(input)
}

func (mdb *MockDynamoDBClient) UntagResourceWithContext(context aws.Context, input *dynamodb.UntagResourceInput, options ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	return mdb.functions["UntagResourceWithContext"].(UntagResourceWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UntagResourceRequest(input *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	return mdb.functions["UntagResourceRequest"].(UntagResourceRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateContinuousBackups(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return mdb.functions["UpdateContinuousBackups"].(UpdateContinuousBackupsFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateContinuousBackupsWithContext(context aws.Context, input *dynamodb.UpdateContinuousBackupsInput, options ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return mdb.functions["UpdateContinuousBackupsWithContext"].(UpdateContinuousBackupsWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateContinuousBackupsRequest(input *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	return mdb.functions["UpdateContinuousBackupsRequest"].(UpdateContinuousBackupsRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTable(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return mdb.functions["UpdateGlobalTable"].(UpdateGlobalTableFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableWithContext(context aws.Context, input *dynamodb.UpdateGlobalTableInput, options ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	return mdb.functions["UpdateGlobalTableWithContext"].(UpdateGlobalTableWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableRequest(input *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	return mdb.functions["UpdateGlobalTableRequest"].(UpdateGlobalTableRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableSettings(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return mdb.functions["UpdateGlobalTableSettings"].(UpdateGlobalTableSettingsFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableSettingsWithContext(context aws.Context, input *dynamodb.UpdateGlobalTableSettingsInput, options ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return mdb.functions["UpdateGlobalTableSettingsWithContext"].(UpdateGlobalTableSettingsWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableSettingsRequest(input *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	return mdb.functions["UpdateGlobalTableSettingsRequest"].(UpdateGlobalTableSettingsRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return mdb.functions["UpdateItem"].(UpdateItemFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateItemWithContext(context aws.Context, input *dynamodb.UpdateItemInput, options ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return mdb.functions["UpdateItemWithContext"].(UpdateItemWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return mdb.functions["UpdateItemRequest"].(UpdateItemRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateTable(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return mdb.functions["UpdateTable"].(UpdateTableFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateTableWithContext(context aws.Context, input *dynamodb.UpdateTableInput, options ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	return mdb.functions["UpdateTableWithContext"].(UpdateTableWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateTableRequest(input *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return mdb.functions["UpdateTableRequest"].(UpdateTableRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateTimeToLive(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return mdb.functions["UpdateTimeToLive"].(UpdateTimeToLiveFunc)(input)
}

func (mdb *MockDynamoDBClient) UpdateTimeToLiveWithContext(context aws.Context, input *dynamodb.UpdateTimeToLiveInput, options ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return mdb.functions["UpdateTimeToLiveWithContext"].(UpdateTimeToLiveWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateTimeToLiveRequest(input *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	return mdb.functions["UpdateTimeToLiveRequest"].(UpdateTimeToLiveRequestFunc)(input)
}

func (mdb *MockDynamoDBClient) WaitUntilTableExists(input *dynamodb.DescribeTableInput) error {
	return mdb.functions["WaitUntilTableExists"].(WaitUntilTableExistsFunc)(input)
}

func (mdb *MockDynamoDBClient) WaitUntilTableExistsWithContext(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error {
	return mdb.functions["WaitUntilTableExistsWithContext"].(WaitUntilTableExistsWithContextFunc)(context, input, options...)
}

func (mdb *MockDynamoDBClient) WaitUntilTableNotExists(input *dynamodb.DescribeTableInput) error {
	return mdb.functions["WaitUntilTableNotExists"].(WaitUntilTableNotExistsFunc)(input)
}

func (mdb *MockDynamoDBClient) WaitUntilTableNotExistsWithContext(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error {
	return mdb.functions["WaitUntilTableNotExistsWithContext"].(WaitUntilTableNotExistsWithContextFunc)(context, input, options...)
}

func NewMockDynamoDBClient(functions map[string]interface{}) dynamodbiface.DynamoDBAPI {
	return &MockDynamoDBClient{
		functions: functions,
	}
}
