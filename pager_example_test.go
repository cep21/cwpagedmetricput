package cwpagedmetricput_test

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/cep21/cwpagedmetricput"
)

func ExamplePager_PutMetricData() {
	a := cwpagedmetricput.Pager{
		Client: cloudwatch.New(session.Must(session.NewSession())),
	}
	_, err := a.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("custom"),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("custom metric"),
				Value: aws.Float64(1.0),
			},
		},
	})
	if err != nil {
		// You'll need valid AWS credentials
		fmt.Println("error result")
	} else {
		fmt.Println("result")
	}
	// Output: error result
}
