package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    //"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "fmt"
    "os"
    "io/ioutil"
    "log"
    "net/http"
	"encoding/json"
)
// snippet-end:[dynamodb.go.create_item.struct]


type Rates struct { 
  
    // defining struct variables 
	Cad float64 `json:"CAD"`
	Hkd float64 `json:"HKD"`
	Inr float64 `json:"INR"`
	Btc float64 `json:"BTC"`
	Usd float64 `json:"USD"`
	Jpy float64 `json:"JPY"`
	Kwd float64 `json:"KWD"`
	Byn float64 `json:"BYN"`
	Pkr float64 `json:"PKR"`
	Irr float64 `json:"IRR"`
}

type Crate struct { 
  
    // defining struct variables 
    Key Rates `json:"rates"`
}
// snippet-end:[dynamodb.go.create_item.struct]

func main() {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create DynamoDB client
    svc := dynamodb.New(sess)
	
	//code from previous file



	response, err := http.Get("http://data.fixer.io/api/latest?access_key=0122524645f7f4a2293b9bc3fc7c44d8")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
	var crate Crate 
	json.Unmarshal([]byte(responseData), &crate)

    tableName := "Currency"


    //HKD
    s := fmt.Sprintf("%f", crate.Key.Hkd)
    currencyname:= "HKD"
    currvalue:= s

    input := &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)
    //CAD
    s = fmt.Sprintf("%f", crate.Key.Cad)
    currencyname= "CAD"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)
    //INR
    s = fmt.Sprintf("%f", crate.Key.Inr)
    currencyname= "INR"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)
    //BTC
    s = fmt.Sprintf("%f", crate.Key.Btc)
    currencyname= "BTC"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)
    //USD
    s = fmt.Sprintf("%f", crate.Key.Usd)
    currencyname= "USD"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)

    //JPY
    s = fmt.Sprintf("%f", crate.Key.Jpy)
    currencyname= "JPY"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)
    //KWD
    s = fmt.Sprintf("%f", crate.Key.Kwd)
    currencyname= "KWD"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)

    //BYN
    s = fmt.Sprintf("%f", crate.Key.Byn)
    currencyname= "BYN"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)

    //PKR
    s = fmt.Sprintf("%f", crate.Key.Pkr)
    currencyname= "PKR"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)

    //IRR
    s = fmt.Sprintf("%f", crate.Key.Irr)
    currencyname= "IRR"
    currvalue= s

    input = &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(currvalue),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Currency": {
                S: aws.String(currencyname),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    svc.UpdateItem(input)

    fmt.Println("Successfully updated " + currvalue)
}