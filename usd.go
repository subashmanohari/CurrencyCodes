package main

// snippet-start:[dynamodb.go.read_item.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-lambda-go/lambda"
	"fmt"
	//"context"
)
type Item struct {
    Currency   string
    Rating string
}

func Handle(item Item)(string,error) {
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create DynamoDB client
    svc := dynamodb.New(sess)
    tableName := "Currency"
    curr := "USD"

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(curr),
            },
        },
    })
    if err != nil {
        return fmt.Sprintf(err.Error()),nil
    }
    if result.Item == nil {
        msg := "Could not find "
        return fmt.Sprintf(msg),nil
    }
        
    item = Item{}

    err = dynamodbattribute.UnmarshalMap(result.Item, &item)
    if err != nil {
        panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
    }
    return fmt.Sprintf(item.Rating),nil
    // snippet-end:[dynamodb.go.read_item.unmarshall]
}
func main(){
	lambda.Start(Handle)
}
