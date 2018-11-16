# go-testhelpers
Series of functions for testing in go

## AWS Helpers
These helpers are for mocking services provided by the AWS go SDK. While the SDK does describe how to mock the services
it's very inconvenient to do so consistently

### Dynamodb Helpers ###
```
mockedDbFunctions := make(map[string]interface{})
mockedDbFunctions["UpdateItem"] = testhelpers.UpdateItemFunc(func(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {

    key1 := input.Key["key1"]
    if *key1.S != "value1" {
        t.Errorf("Expected token to be 'value1', got '%s'", *key1.S)
    }
    return &dynamodb.UpdateItemOutput{}, nil
})
mockClient := testhelpers.NewMockDynamoDBClient(mockedDbFunctions)
```