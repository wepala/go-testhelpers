package testhelpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// MockDynamoDBClient provides a mock implementation of DynamoDBAPI
//
// To use this client you specify the function you would like to override by appending 'Func' to the original method name (e.g. DynamoDBUpdateItem)
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon DynamoDB.
//    func myFunc(svc dynamodbiface.DynamoDBAPI) bool {
//        // Make svc.BatchGetItem request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := dynamodb.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    func DynamoDBTestMy(t *testing.T) {
//		//setup map of functions you would like to implement. Note the key is the aws sdk method name and the function type is the Me
//		mockedDbFunctions := make(map[string]interface{})
//		mockedDbFunctions["UpdateItem"] = testhelpers.DynamoDBUpdateItem(func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
//
//			key := input.Key["key"]
//			if *key.S != "value" {
//				t.Errorf("Expected token to be 'value', got '%s'", *key.S)
//			}
//			return &dynamodb.UpdateItemOutput{}, nil
//		})
//		mockSvc := testhelpers.NewMockDynamoDBClient(mockedDbFunctions)'
//		myfunc(mockSvc)
//
type MockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
	functions map[string]interface{}
}

type DynamoDBBatchGetItem func(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error)
type DynamoDBBatchGetItemWithContext func(context aws.Context, input *dynamodb.BatchGetItemInput, options ...request.Option) (*dynamodb.BatchGetItemOutput, error)
type DynamoDBBatchGetItemRequest func(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput)
type DynamoDBBatchGetItemPages func(input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool) error
type DynamoDBBatchGetItemPagesWithContext func(context aws.Context, input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool, options ...request.Option) error
type DynamoDBBatchWriteItem func(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
type DynamoDBBatchWriteItemWithContext func(context aws.Context, input *dynamodb.BatchWriteItemInput, options ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
type DynamoDBBatchWriteItemRequest func(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput)
type DynamoDBCreateBackup func(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error)
type DynamoDBCreateBackupWithContext func(context aws.Context, input *dynamodb.CreateBackupInput, options ...request.Option) (*dynamodb.CreateBackupOutput, error)
type DynamoDBCreateBackupRequest func(input *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput)
type DynamoDBCreateGlobalTable func(input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error)
type DynamoDBCreateGlobalTableWithContext func(context aws.Context, input *dynamodb.CreateGlobalTableInput, options ...request.Option) (*dynamodb.CreateGlobalTableOutput, error)
type DynamoDBCreateGlobalTableRequest func(input *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput)
type DynamoDBCreateTable func(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error)
type DynamoDBCreateTableWithContext func(context aws.Context, input *dynamodb.CreateTableInput, options ...request.Option) (*dynamodb.CreateTableOutput, error)
type DynamoDBCreateTableRequest func(input *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput)
type DynamoDBDeleteBackup func(input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error)
type DynamoDBDeleteBackupWithContext func(context aws.Context, input *dynamodb.DeleteBackupInput, options ...request.Option) (*dynamodb.DeleteBackupOutput, error)
type DynamoDBDeleteBackupRequest func(input *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput)
type DynamoDBDeleteItem func(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
type DynamoDBDeleteItemWithContext func(context aws.Context, input *dynamodb.DeleteItemInput, options ...request.Option) (*dynamodb.DeleteItemOutput, error)
type DynamoDBDeleteItemRequest func(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput)
type DynamoDBDeleteTable func(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error)
type DynamoDBDeleteTableWithContext func(context aws.Context, input *dynamodb.DeleteTableInput, options ...request.Option) (*dynamodb.DeleteTableOutput, error)
type DynamoDBDeleteTableRequest func(input *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput)
type DynamoDBDescribeBackup func(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error)
type DynamoDBDescribeBackupWithContext func(context aws.Context, input *dynamodb.DescribeBackupInput, options ...request.Option) (*dynamodb.DescribeBackupOutput, error)
type DynamoDBDescribeBackupRequest func(input *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput)
type DynamoDBDescribeContinuousBackups func(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error)
type DynamoDBDescribeContinuousBackupsWithContext func(context aws.Context, input *dynamodb.DescribeContinuousBackupsInput, options ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error)
type DynamoDBDescribeContinuousBackupsRequest func(input *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput)
type DynamoDBDescribeGlobalTable func(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error)
type DynamoDBDescribeGlobalTableWithContext func(context aws.Context, input *dynamodb.DescribeGlobalTableInput, options ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error)
type DynamoDBDescribeGlobalTableRequest func(input *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput)
type DynamoDBDescribeGlobalTableSettings func(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
type DynamoDBDescribeGlobalTableSettingsWithContext func(context aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, options ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
type DynamoDBDescribeGlobalTableSettingsRequest func(input *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput)
type DynamoDBDescribeLimits func(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error)
type DynamoDBDescribeLimitsWithContext func(context aws.Context, input *dynamodb.DescribeLimitsInput, options ...request.Option) (*dynamodb.DescribeLimitsOutput, error)
type DynamoDBDescribeLimitsRequest func(input *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput)
type DynamoDBDescribeTable func(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
type DynamoDBDescribeTableWithContext func(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.Option) (*dynamodb.DescribeTableOutput, error)
type DynamoDBDescribeTableRequest func(input *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput)
type DynamoDBDescribeTimeToLive func(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error)
type DynamoDBDescribeTimeToLiveWithContext func(context aws.Context, input *dynamodb.DescribeTimeToLiveInput, options ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error)
type DynamoDBDescribeTimeToLiveRequest func(input *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput)
type DynamoDBGetItem func(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
type DynamoDBGetItemWithContext func(context aws.Context, input *dynamodb.GetItemInput, options ...request.Option) (*dynamodb.GetItemOutput, error)
type DynamoDBGetItemRequest func(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput)
type DynamoDBListBackups func(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error)
type DynamoDBListBackupsWithContext func(context aws.Context, input *dynamodb.ListBackupsInput, options ...request.Option) (*dynamodb.ListBackupsOutput, error)
type DynamoDBListBackupsRequest func(input *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput)
type DynamoDBListGlobalTables func(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error)
type DynamoDBListGlobalTablesWithContext func(context aws.Context, input *dynamodb.ListGlobalTablesInput, options ...request.Option) (*dynamodb.ListGlobalTablesOutput, error)
type DynamoDBListGlobalTablesRequest func(input *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput)
type DynamoDBListTables func(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error)
type DynamoDBListTablesWithContext func(context aws.Context, input *dynamodb.ListTablesInput, options ...request.Option) (*dynamodb.ListTablesOutput, error)
type DynamoDBListTablesRequest func(input *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput)
type DynamoDBListTablesPages func(input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool) error
type DynamoDBListTablesPagesWithContext func(context aws.Context, input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool, options ...request.Option) error
type DynamoDBListTagsOfResource func(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error)
type DynamoDBListTagsOfResourceWithContext func(context aws.Context, input *dynamodb.ListTagsOfResourceInput, options ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error)
type DynamoDBListTagsOfResourceRequest func(input *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput)
type DynamoDBPutItem func(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
type DynamoDBPutItemWithContext func(context aws.Context, input *dynamodb.PutItemInput, options ...request.Option) (*dynamodb.PutItemOutput, error)
type DynamoDBPutItemRequest func(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput)
type DynamoDBQuery func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
type DynamoDBQueryWithContext func(context aws.Context, input *dynamodb.QueryInput, options ...request.Option) (*dynamodb.QueryOutput, error)
type DynamoDBQueryRequest func(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput)
type DynamoDBQueryPages func(input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool) error
type DynamoDBQueryPagesWithContext func(context aws.Context, input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool, options ...request.Option) error
type DynamoDBRestoreTableFromBackup func(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error)
type DynamoDBRestoreTableFromBackupWithContext func(context aws.Context, input *dynamodb.RestoreTableFromBackupInput, options ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error)
type DynamoDBRestoreTableFromBackupRequest func(input *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput)
type DynamoDBRestoreTableToPointInTime func(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error)
type DynamoDBRestoreTableToPointInTimeWithContext func(context aws.Context, input *dynamodb.RestoreTableToPointInTimeInput, options ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error)
type DynamoDBRestoreTableToPointInTimeRequest func(input *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput)
type DynamoDBScan func(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
type DynamoDBScanWithContext func(context aws.Context, input *dynamodb.ScanInput, options ...request.Option) (*dynamodb.ScanOutput, error)
type DynamoDBScanRequest func(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput)
type DynamoDBScanPages func(input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool) error
type DynamoDBScanPagesWithContext func(context aws.Context, input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool, options ...request.Option) error
type DynamoDBTagResource func(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error)
type DynamoDBTagResourceWithContext func(context aws.Context, input *dynamodb.TagResourceInput, options ...request.Option) (*dynamodb.TagResourceOutput, error)
type DynamoDBTagResourceRequest func(input *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput)
type DynamoDBUntagResource func(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error)
type DynamoDBUntagResourceWithContext func(context aws.Context, input *dynamodb.UntagResourceInput, options ...request.Option) (*dynamodb.UntagResourceOutput, error)
type DynamoDBUntagResourceRequest func(input *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput)
type DynamoDBUpdateContinuousBackups func(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error)
type DynamoDBUpdateContinuousBackupsWithContext func(context aws.Context, input *dynamodb.UpdateContinuousBackupsInput, options ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error)
type DynamoDBUpdateContinuousBackupsRequest func(input *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput)
type DynamoDBUpdateGlobalTable func(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error)
type DynamoDBUpdateGlobalTableWithContext func(context aws.Context, input *dynamodb.UpdateGlobalTableInput, options ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error)
type DynamoDBUpdateGlobalTableRequest func(input *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput)
type DynamoDBUpdateGlobalTableSettings func(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
type DynamoDBUpdateGlobalTableSettingsWithContext func(context aws.Context, input *dynamodb.UpdateGlobalTableSettingsInput, options ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
type DynamoDBUpdateGlobalTableSettingsRequest func(input *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput)
type DynamoDBUpdateItem func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error)
type DynamoDBUpdateItemWithContext func(context aws.Context, input *dynamodb.UpdateItemInput, options ...request.Option) (*dynamodb.UpdateItemOutput, error)
type DynamoDBUpdateItemRequest func(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput)
type DynamoDBUpdateTable func(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error)
type DynamoDBUpdateTableWithContext func(context aws.Context, input *dynamodb.UpdateTableInput, options ...request.Option) (*dynamodb.UpdateTableOutput, error)
type DynamoDBUpdateTableRequest func(input *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput)
type DynamoDBUpdateTimeToLive func(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error)
type DynamoDBUpdateTimeToLiveWithContext func(context aws.Context, input *dynamodb.UpdateTimeToLiveInput, options ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error)
type DynamoDBUpdateTimeToLiveRequest func(input *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput)
type DynamoDBWaitUntilTableExists func(input *dynamodb.DescribeTableInput) error
type DynamoDBWaitUntilTableExistsWithContext func(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error
type DynamoDBWaitUntilTableNotExists func(input *dynamodb.DescribeTableInput) error
type DynamoDBWaitUntilTableNotExistsWithContext func(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error

func (mdb *MockDynamoDBClient) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	return mdb.functions["BatchGetItem"].(DynamoDBBatchGetItem)(input)
}

func (mdb *MockDynamoDBClient) BatchGetItemWithContext(context aws.Context, input *dynamodb.BatchGetItemInput, options ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	return mdb.functions["BatchGetItemWithContext"].(DynamoDBBatchGetItemWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (*request.Request, *dynamodb.BatchGetItemOutput) {
	return mdb.functions["BatchGetItemRequest"].(DynamoDBBatchGetItemRequest)(input)
}

func (mdb *MockDynamoDBClient) BatchGetItemPages(input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	return mdb.functions["BatchGetItemPages"].(DynamoDBBatchGetItemPages)(input, callback)
}

func (mdb *MockDynamoDBClient) BatchGetItemPagesWithContext(context aws.Context, input *dynamodb.BatchGetItemInput, callback func(*dynamodb.BatchGetItemOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["BatchGetItemPagesWithContext"].(DynamoDBBatchGetItemPagesWithContext)(context, input, callback)
}

func (mdb *MockDynamoDBClient) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return mdb.functions["BatchWriteItem"].(DynamoDBBatchWriteItem)(input)
}

func (mdb *MockDynamoDBClient) BatchWriteItemWithContext(context aws.Context, input *dynamodb.BatchWriteItemInput, options ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	return mdb.functions["BatchWriteItemWithContext"].(DynamoDBBatchWriteItemWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (*request.Request, *dynamodb.BatchWriteItemOutput) {
	return mdb.functions["BatchWriteItemRequest"].(DynamoDBBatchWriteItemRequest)(input)
}

func (mdb *MockDynamoDBClient) CreateBackup(input *dynamodb.CreateBackupInput) (*dynamodb.CreateBackupOutput, error) {
	return mdb.functions["CreateBackup"].(DynamoDBCreateBackup)(input)
}

func (mdb *MockDynamoDBClient) CreateBackupWithContext(context aws.Context, input *dynamodb.CreateBackupInput, options ...request.Option) (*dynamodb.CreateBackupOutput, error) {
	return mdb.functions["CreateBackupWithContext"].(DynamoDBCreateBackupWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) CreateBackupRequest(input *dynamodb.CreateBackupInput) (*request.Request, *dynamodb.CreateBackupOutput) {
	return mdb.functions["CreateBackupRequest"].(DynamoDBCreateBackupRequest)(input)
}

func (mdb *MockDynamoDBClient) CreateGlobalTable(input *dynamodb.CreateGlobalTableInput) (*dynamodb.CreateGlobalTableOutput, error) {
	return mdb.functions["CreateGlobalTable"].(DynamoDBCreateGlobalTable)(input)
}

func (mdb *MockDynamoDBClient) CreateGlobalTableWithContext(context aws.Context, input *dynamodb.CreateGlobalTableInput, options ...request.Option) (*dynamodb.CreateGlobalTableOutput, error) {
	return mdb.functions["CreateGlobalTableWithContext"].(DynamoDBCreateGlobalTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) CreateGlobalTableRequest(input *dynamodb.CreateGlobalTableInput) (*request.Request, *dynamodb.CreateGlobalTableOutput) {
	return mdb.functions["CreateGlobalTableRequest"].(DynamoDBCreateGlobalTableRequest)(input)
}

func (mdb *MockDynamoDBClient) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	return mdb.functions["CreateTable"].(DynamoDBCreateTable)(input)
}

func (mdb *MockDynamoDBClient) CreateTableWithContext(context aws.Context, input *dynamodb.CreateTableInput, options ...request.Option) (*dynamodb.CreateTableOutput, error) {
	return mdb.functions["CreateTableWithContext"].(DynamoDBCreateTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) CreateTableRequest(input *dynamodb.CreateTableInput) (*request.Request, *dynamodb.CreateTableOutput) {
	return mdb.functions["CreateTableRequest"].(DynamoDBCreateTableRequest)(input)
}

func (mdb *MockDynamoDBClient) DeleteBackup(input *dynamodb.DeleteBackupInput) (*dynamodb.DeleteBackupOutput, error) {
	return mdb.functions["DeleteBackup"].(DynamoDBDeleteBackup)(input)
}

func (mdb *MockDynamoDBClient) DeleteBackupWithContext(context aws.Context, input *dynamodb.DeleteBackupInput, options ...request.Option) (*dynamodb.DeleteBackupOutput, error) {
	return mdb.functions["DeleteBackupWithContext"].(DynamoDBDeleteBackupWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DeleteBackupRequest(input *dynamodb.DeleteBackupInput) (*request.Request, *dynamodb.DeleteBackupOutput) {
	return mdb.functions["DeleteBackupRequest"].(DynamoDBDeleteBackupRequest)(input)
}

func (mdb *MockDynamoDBClient) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return mdb.functions["DeleteItem"].(DynamoDBDeleteItem)(input)
}

func (mdb *MockDynamoDBClient) DeleteItemWithContext(context aws.Context, input *dynamodb.DeleteItemInput, options ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return mdb.functions["DeleteItemWithContext"].(DynamoDBDeleteItemWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DeleteItemRequest(input *dynamodb.DeleteItemInput) (*request.Request, *dynamodb.DeleteItemOutput) {
	return mdb.functions["DeleteItemRequest"].(DynamoDBDeleteItemRequest)(input)
}

func (mdb *MockDynamoDBClient) DeleteTable(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	return mdb.functions["DeleteTable"].(DynamoDBDeleteTable)(input)
}

func (mdb *MockDynamoDBClient) DeleteTableWithContext(context aws.Context, input *dynamodb.DeleteTableInput, options ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	return mdb.functions["DeleteTableWithContext"].(DynamoDBDeleteTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DeleteTableRequest(input *dynamodb.DeleteTableInput) (*request.Request, *dynamodb.DeleteTableOutput) {
	return mdb.functions["DeleteTableRequest"].(DynamoDBDeleteTableRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeBackup(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	return mdb.functions["DescribeBackup"].(DynamoDBDescribeBackup)(input)
}

func (mdb *MockDynamoDBClient) DescribeBackupWithContext(context aws.Context, input *dynamodb.DescribeBackupInput, options ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	return mdb.functions["DescribeBackupWithContext"].(DynamoDBDescribeBackupWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeBackupRequest(input *dynamodb.DescribeBackupInput) (*request.Request, *dynamodb.DescribeBackupOutput) {
	return mdb.functions["DescribeBackupRequest"].(DynamoDBDescribeBackupRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeContinuousBackups(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return mdb.functions["DescribeContinuousBackups"].(DynamoDBDescribeContinuousBackups)(input)
}

func (mdb *MockDynamoDBClient) DescribeContinuousBackupsWithContext(context aws.Context, input *dynamodb.DescribeContinuousBackupsInput, options ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return mdb.functions["DescribeContinuousBackupsWithContext"].(DynamoDBDescribeContinuousBackupsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeContinuousBackupsRequest(input *dynamodb.DescribeContinuousBackupsInput) (*request.Request, *dynamodb.DescribeContinuousBackupsOutput) {
	return mdb.functions["DescribeContinuousBackupsRequest"].(DynamoDBDescribeContinuousBackupsRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeGlobalTable(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	return mdb.functions["DescribeGlobalTable"].(DynamoDBDescribeGlobalTable)(input)
}

func (mdb *MockDynamoDBClient) DescribeGlobalTableWithContext(context aws.Context, input *dynamodb.DescribeGlobalTableInput, options ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	return mdb.functions["DescribeGlobalTableWithContext"].(DynamoDBDescribeGlobalTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeGlobalTableRequest(input *dynamodb.DescribeGlobalTableInput) (*request.Request, *dynamodb.DescribeGlobalTableOutput) {
	return mdb.functions["DescribeGlobalTableRequest"].(DynamoDBDescribeGlobalTableRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeGlobalTableSettings(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return mdb.functions["DescribeGlobalTableSettings"].(DynamoDBDescribeGlobalTableSettings)(input)
}

func (mdb *MockDynamoDBClient) DescribeGlobalTableSettingsWithContext(context aws.Context, input *dynamodb.DescribeGlobalTableSettingsInput, options ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return mdb.functions["DescribeGlobalTableSettingsWithContext"].(DynamoDBDescribeGlobalTableSettingsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeGlobalTableSettingsRequest(input *dynamodb.DescribeGlobalTableSettingsInput) (*request.Request, *dynamodb.DescribeGlobalTableSettingsOutput) {
	return mdb.functions["DescribeGlobalTableSettingsRequest"].(DynamoDBDescribeGlobalTableSettingsRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeLimits(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	return mdb.functions["DescribeLimits"].(DynamoDBDescribeLimits)(input)
}

func (mdb *MockDynamoDBClient) DescribeLimitsWithContext(context aws.Context, input *dynamodb.DescribeLimitsInput, options ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	return mdb.functions["DescribeLimitsWithContext"].(DynamoDBDescribeLimitsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeLimitsRequest(input *dynamodb.DescribeLimitsInput) (*request.Request, *dynamodb.DescribeLimitsOutput) {
	return mdb.functions["DescribeLimitsRequest"].(DynamoDBDescribeLimitsRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	return mdb.functions["DescribeTable"].(DynamoDBDescribeTable)(input)
}

func (mdb *MockDynamoDBClient) DescribeTableWithContext(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	return mdb.functions["DescribeTableWithContext"].(DynamoDBDescribeTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeTableRequest(input *dynamodb.DescribeTableInput) (*request.Request, *dynamodb.DescribeTableOutput) {
	return mdb.functions["DescribeTableRequest"].(DynamoDBDescribeTableRequest)(input)
}

func (mdb *MockDynamoDBClient) DescribeTimeToLive(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return mdb.functions["DescribeTimeToLive"].(DynamoDBDescribeTimeToLive)(input)
}

func (mdb *MockDynamoDBClient) DescribeTimeToLiveWithContext(context aws.Context, input *dynamodb.DescribeTimeToLiveInput, options ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return mdb.functions["DescribeTimeToLiveWithContext"].(DynamoDBDescribeTimeToLiveWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) DescribeTimeToLiveRequest(input *dynamodb.DescribeTimeToLiveInput) (*request.Request, *dynamodb.DescribeTimeToLiveOutput) {
	return mdb.functions["DescribeTimeToLiveRequest"].(DynamoDBDescribeTimeToLiveRequest)(input)
}

func (mdb *MockDynamoDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return mdb.functions["GetItem"].(DynamoDBGetItem)(input)
}

func (mdb *MockDynamoDBClient) GetItemWithContext(context aws.Context, input *dynamodb.GetItemInput, options ...request.Option) (*dynamodb.GetItemOutput, error) {
	return mdb.functions["GetItemWithContext"].(DynamoDBGetItemWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) GetItemRequest(input *dynamodb.GetItemInput) (*request.Request, *dynamodb.GetItemOutput) {
	return mdb.functions["GetItemRequest"].(DynamoDBGetItemRequest)(input)
}

func (mdb *MockDynamoDBClient) ListBackups(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	return mdb.functions["ListBackups"].(DynamoDBListBackups)(input)
}

func (mdb *MockDynamoDBClient) ListBackupsWithContext(context aws.Context, input *dynamodb.ListBackupsInput, options ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	return mdb.functions["ListBackupsWithContext"].(DynamoDBListBackupsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) ListBackupsRequest(input *dynamodb.ListBackupsInput) (*request.Request, *dynamodb.ListBackupsOutput) {
	return mdb.functions["ListBackupsRequest"].(DynamoDBListBackupsRequest)(input)
}

func (mdb *MockDynamoDBClient) ListGlobalTables(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	return mdb.functions["ListGlobalTables"].(DynamoDBListGlobalTables)(input)
}

func (mdb *MockDynamoDBClient) ListGlobalTablesWithContext(context aws.Context, input *dynamodb.ListGlobalTablesInput, options ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	return mdb.functions["ListGlobalTablesWithContext"].(DynamoDBListGlobalTablesWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) ListGlobalTablesRequest(input *dynamodb.ListGlobalTablesInput) (*request.Request, *dynamodb.ListGlobalTablesOutput) {
	return mdb.functions["ListGlobalTablesRequest"].(DynamoDBListGlobalTablesRequest)(input)
}

func (mdb *MockDynamoDBClient) ListTables(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	return mdb.functions["ListTables"].(DynamoDBListTables)(input)
}

func (mdb *MockDynamoDBClient) ListTablesWithContext(context aws.Context, input *dynamodb.ListTablesInput, options ...request.Option) (*dynamodb.ListTablesOutput, error) {
	return mdb.functions["ListTablesWithContext"].(DynamoDBListTablesWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) ListTablesRequest(input *dynamodb.ListTablesInput) (*request.Request, *dynamodb.ListTablesOutput) {
	return mdb.functions["ListTablesRequest"].(DynamoDBListTablesRequest)(input)
}

func (mdb *MockDynamoDBClient) ListTablesPages(input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool) error {
	return mdb.functions["ListTablesPages"].(DynamoDBListTablesPages)(input, callback)
}

func (mdb *MockDynamoDBClient) ListTablesPagesWithContext(context aws.Context, input *dynamodb.ListTablesInput, callback func(*dynamodb.ListTablesOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["ListTablesPagesWithContext"].(DynamoDBListTablesPagesWithContext)(context, input, callback, options...)
}

func (mdb *MockDynamoDBClient) ListTagsOfResource(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	return mdb.functions["ListTagsOfResource"].(DynamoDBListTagsOfResource)(input)
}

func (mdb *MockDynamoDBClient) ListTagsOfResourceWithContext(context aws.Context, input *dynamodb.ListTagsOfResourceInput, options ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	return mdb.functions["ListTagsOfResourceWithContext"].(DynamoDBListTagsOfResourceWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) ListTagsOfResourceRequest(input *dynamodb.ListTagsOfResourceInput) (*request.Request, *dynamodb.ListTagsOfResourceOutput) {
	return mdb.functions["ListTagsOfResourceRequest"].(DynamoDBListTagsOfResourceRequest)(input)
}

func (mdb *MockDynamoDBClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return mdb.functions["PutItem"].(DynamoDBPutItem)(input)
}

func (mdb *MockDynamoDBClient) PutItemWithContext(context aws.Context, input *dynamodb.PutItemInput, options ...request.Option) (*dynamodb.PutItemOutput, error) {
	return mdb.functions["PutItemWithContext"].(DynamoDBPutItemWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) PutItemRequest(input *dynamodb.PutItemInput) (*request.Request, *dynamodb.PutItemOutput) {
	return mdb.functions["PutItemRequest"].(DynamoDBPutItemRequest)(input)
}

func (mdb *MockDynamoDBClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return mdb.functions["Query"].(DynamoDBQuery)(input)
}

func (mdb *MockDynamoDBClient) QueryWithContext(context aws.Context, input *dynamodb.QueryInput, options ...request.Option) (*dynamodb.QueryOutput, error) {
	return mdb.functions["QueryWithContext"].(DynamoDBQueryWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) QueryRequest(input *dynamodb.QueryInput) (*request.Request, *dynamodb.QueryOutput) {
	return mdb.functions["QueryRequest"].(DynamoDBQueryRequest)(input)
}

func (mdb *MockDynamoDBClient) QueryPages(input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool) error {
	return mdb.functions["QueryPages"].(DynamoDBQueryPages)(input, callback)
}

func (mdb *MockDynamoDBClient) QueryPagesWithContext(context aws.Context, input *dynamodb.QueryInput, callback func(*dynamodb.QueryOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["QueryPagesWithContext"].(DynamoDBQueryPagesWithContext)(context, input, callback)
}

func (mdb *MockDynamoDBClient) RestoreTableFromBackup(input *dynamodb.RestoreTableFromBackupInput) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return mdb.functions["RestoreTableFromBackup"].(DynamoDBRestoreTableFromBackup)(input)
}

func (mdb *MockDynamoDBClient) RestoreTableFromBackupWithContext(context aws.Context, input *dynamodb.RestoreTableFromBackupInput, options ...request.Option) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return mdb.functions["RestoreTableFromBackupWithContext"].(DynamoDBRestoreTableFromBackupWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) RestoreTableFromBackupRequest(input *dynamodb.RestoreTableFromBackupInput) (*request.Request, *dynamodb.RestoreTableFromBackupOutput) {
	return mdb.functions["RestoreTableFromBackupRequest"].(DynamoDBRestoreTableFromBackupRequest)(input)
}

func (mdb *MockDynamoDBClient) RestoreTableToPointInTime(input *dynamodb.RestoreTableToPointInTimeInput) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return mdb.functions["RestoreTableToPointInTime"].(DynamoDBRestoreTableToPointInTime)(input)
}

func (mdb *MockDynamoDBClient) RestoreTableToPointInTimeWithContext(context aws.Context, input *dynamodb.RestoreTableToPointInTimeInput, options ...request.Option) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return mdb.functions["RestoreTableToPointInTimeWithContext"].(DynamoDBRestoreTableToPointInTimeWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) RestoreTableToPointInTimeRequest(input *dynamodb.RestoreTableToPointInTimeInput) (*request.Request, *dynamodb.RestoreTableToPointInTimeOutput) {
	return mdb.functions["RestoreTableToPointInTimeRequest"].(DynamoDBRestoreTableToPointInTimeRequest)(input)
}

func (mdb *MockDynamoDBClient) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return mdb.functions["Scan"].(DynamoDBScan)(input)
}

func (mdb *MockDynamoDBClient) ScanWithContext(context aws.Context, input *dynamodb.ScanInput, options ...request.Option) (*dynamodb.ScanOutput, error) {
	return mdb.functions["ScanWithContext"].(DynamoDBScanWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) ScanRequest(input *dynamodb.ScanInput) (*request.Request, *dynamodb.ScanOutput) {
	return mdb.functions["ScanRequest"].(DynamoDBScanRequest)(input)
}

func (mdb *MockDynamoDBClient) ScanPages(input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool) error {
	return mdb.functions["ScanPages"].(DynamoDBScanPages)(input, callback)
}

func (mdb *MockDynamoDBClient) ScanPagesWithContext(context aws.Context, input *dynamodb.ScanInput, callback func(*dynamodb.ScanOutput, bool) bool, options ...request.Option) error {
	return mdb.functions["ScanPagesWithContext"].(DynamoDBScanPagesWithContext)(context, input, callback)
}

func (mdb *MockDynamoDBClient) TagResource(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	return mdb.functions["TagResource"].(DynamoDBTagResource)(input)
}

func (mdb *MockDynamoDBClient) TagResourceWithContext(context aws.Context, input *dynamodb.TagResourceInput, options ...request.Option) (*dynamodb.TagResourceOutput, error) {
	return mdb.functions["TagResourceWithContext"].(DynamoDBTagResourceWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) TagResourceRequest(input *dynamodb.TagResourceInput) (*request.Request, *dynamodb.TagResourceOutput) {
	return mdb.functions["TagResourceRequest"].(DynamoDBTagResourceRequest)(input)
}

func (mdb *MockDynamoDBClient) UntagResource(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	return mdb.functions["UntagResource"].(DynamoDBUntagResource)(input)
}

func (mdb *MockDynamoDBClient) UntagResourceWithContext(context aws.Context, input *dynamodb.UntagResourceInput, options ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	return mdb.functions["UntagResourceWithContext"].(DynamoDBUntagResourceWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UntagResourceRequest(input *dynamodb.UntagResourceInput) (*request.Request, *dynamodb.UntagResourceOutput) {
	return mdb.functions["UntagResourceRequest"].(DynamoDBUntagResourceRequest)(input)
}

func (mdb *MockDynamoDBClient) UpdateContinuousBackups(input *dynamodb.UpdateContinuousBackupsInput) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return mdb.functions["UpdateContinuousBackups"].(DynamoDBUpdateContinuousBackups)(input)
}

func (mdb *MockDynamoDBClient) UpdateContinuousBackupsWithContext(context aws.Context, input *dynamodb.UpdateContinuousBackupsInput, options ...request.Option) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return mdb.functions["UpdateContinuousBackupsWithContext"].(DynamoDBUpdateContinuousBackupsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateContinuousBackupsRequest(input *dynamodb.UpdateContinuousBackupsInput) (*request.Request, *dynamodb.UpdateContinuousBackupsOutput) {
	return mdb.functions["UpdateContinuousBackupsRequest"].(DynamoDBUpdateContinuousBackupsRequest)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTable(input *dynamodb.UpdateGlobalTableInput) (*dynamodb.UpdateGlobalTableOutput, error) {
	return mdb.functions["UpdateGlobalTable"].(DynamoDBUpdateGlobalTable)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableWithContext(context aws.Context, input *dynamodb.UpdateGlobalTableInput, options ...request.Option) (*dynamodb.UpdateGlobalTableOutput, error) {
	return mdb.functions["UpdateGlobalTableWithContext"].(DynamoDBUpdateGlobalTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableRequest(input *dynamodb.UpdateGlobalTableInput) (*request.Request, *dynamodb.UpdateGlobalTableOutput) {
	return mdb.functions["UpdateGlobalTableRequest"].(DynamoDBUpdateGlobalTableRequest)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableSettings(input *dynamodb.UpdateGlobalTableSettingsInput) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return mdb.functions["UpdateGlobalTableSettings"].(DynamoDBUpdateGlobalTableSettings)(input)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableSettingsWithContext(context aws.Context, input *dynamodb.UpdateGlobalTableSettingsInput, options ...request.Option) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return mdb.functions["UpdateGlobalTableSettingsWithContext"].(DynamoDBUpdateGlobalTableSettingsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateGlobalTableSettingsRequest(input *dynamodb.UpdateGlobalTableSettingsInput) (*request.Request, *dynamodb.UpdateGlobalTableSettingsOutput) {
	return mdb.functions["UpdateGlobalTableSettingsRequest"].(DynamoDBUpdateGlobalTableSettingsRequest)(input)
}

func (mdb *MockDynamoDBClient) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	return mdb.functions["UpdateItem"].(DynamoDBUpdateItem)(input)
}

func (mdb *MockDynamoDBClient) UpdateItemWithContext(context aws.Context, input *dynamodb.UpdateItemInput, options ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return mdb.functions["UpdateItemWithContext"].(DynamoDBUpdateItemWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateItemRequest(input *dynamodb.UpdateItemInput) (*request.Request, *dynamodb.UpdateItemOutput) {
	return mdb.functions["UpdateItemRequest"].(DynamoDBUpdateItemRequest)(input)
}

func (mdb *MockDynamoDBClient) UpdateTable(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	return mdb.functions["UpdateTable"].(DynamoDBUpdateTable)(input)
}

func (mdb *MockDynamoDBClient) UpdateTableWithContext(context aws.Context, input *dynamodb.UpdateTableInput, options ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	return mdb.functions["UpdateTableWithContext"].(DynamoDBUpdateTableWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateTableRequest(input *dynamodb.UpdateTableInput) (*request.Request, *dynamodb.UpdateTableOutput) {
	return mdb.functions["UpdateTableRequest"].(DynamoDBUpdateTableRequest)(input)
}

func (mdb *MockDynamoDBClient) UpdateTimeToLive(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return mdb.functions["UpdateTimeToLive"].(DynamoDBUpdateTimeToLive)(input)
}

func (mdb *MockDynamoDBClient) UpdateTimeToLiveWithContext(context aws.Context, input *dynamodb.UpdateTimeToLiveInput, options ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return mdb.functions["UpdateTimeToLiveWithContext"].(DynamoDBUpdateTimeToLiveWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) UpdateTimeToLiveRequest(input *dynamodb.UpdateTimeToLiveInput) (*request.Request, *dynamodb.UpdateTimeToLiveOutput) {
	return mdb.functions["UpdateTimeToLiveRequest"].(DynamoDBUpdateTimeToLiveRequest)(input)
}

func (mdb *MockDynamoDBClient) WaitUntilTableExists(input *dynamodb.DescribeTableInput) error {
	return mdb.functions["WaitUntilTableExists"].(DynamoDBWaitUntilTableExists)(input)
}

func (mdb *MockDynamoDBClient) WaitUntilTableExistsWithContext(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error {
	return mdb.functions["WaitUntilTableExistsWithContext"].(DynamoDBWaitUntilTableExistsWithContext)(context, input, options...)
}

func (mdb *MockDynamoDBClient) WaitUntilTableNotExists(input *dynamodb.DescribeTableInput) error {
	return mdb.functions["WaitUntilTableNotExists"].(DynamoDBWaitUntilTableNotExists)(input)
}

func (mdb *MockDynamoDBClient) WaitUntilTableNotExistsWithContext(context aws.Context, input *dynamodb.DescribeTableInput, options ...request.WaiterOption) error {
	return mdb.functions["WaitUntilTableNotExistsWithContext"].(DynamoDBWaitUntilTableNotExistsWithContext)(context, input, options...)
}

//NewMockDynamoDBClient creates a new mock dynamo db client
//example
//mockedDbFunctions := make(map[string]interface{})
//mockedDbFunctions["UpdateItem"] = testhelpers.DynamoDBUpdateItem(func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
//	return &dynamodb.UpdateItemOutput{}, nil
//})mockedDbFunctions["Query"] = testhelpers.DynamoDBQuery(func(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
//	return &dynamodb.QueryOutput{}, nil
//})
func NewMockDynamoDBClient(functions map[string]interface{}) dynamodbiface.DynamoDBAPI {
	return &MockDynamoDBClient{
		functions: functions,
	}
}
