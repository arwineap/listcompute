package main

import (
    "fmt"
    "os"
    "flag"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ec2"
)


func main() {

    flag.Usage = func() {
        fmt.Println("\nUsage:")
        fmt.Println(os.Args[0] + " [-i|--internalip] [-e|--externalip] server name")
        fmt.Println("  -n -- Show only names")
        fmt.Println("  -e -- Show only external ips")
        fmt.Println("  -i -- Show only internal ips")
        fmt.Println("\n")
    }

    flag_name := flag.Bool("n", false, "only show names")
    flag_externalip := flag.Bool("e", false, "only show external ips")
    flag_internalip := flag.Bool("i", false, "only show internal ips")

    flag.Parse()

    // Note that you can also configure your region globally by
    // exporting the AWS_REGION environment variable
    svc := ec2.New(&aws.Config{Region: aws.String("us-west-1")})

    // Call the DescribeInstances Operation
    resp, err := svc.DescribeInstances(nil)
    if err != nil {
        panic(err)
    }


    var matched_hosts []*ec2.Instance

    for idx, _ := range resp.Reservations {
        for _, inst := range resp.Reservations[idx].Instances {
            inst_name := ""
            for _, tag := range inst.Tags {
                if *tag.Key == "Name" {
                    inst_name = *tag.Value
                }
            }

            grep_score := 0
            for _, grep_var := range flag.Args() {
                if strings.Contains(inst_name, grep_var) {
                    grep_score += 1
                }
            }

            if grep_score == len(flag.Args()) {
                matched_hosts = append(matched_hosts, inst)
            }
            
        }
    }


    flag_filter_count := 0
    if *flag_name {
        flag_filter_count += 1
    }
    if *flag_externalip {
        flag_filter_count += 1
    }
    if *flag_internalip {
        flag_filter_count += 1
    }
    for _, inst := range matched_hosts {
        inst_name := ""
        for _, tag := range inst.Tags {
            if *tag.Key == "Name" {
                inst_name = *tag.Value
            }
        }
        var output_line []string
        if *flag_name {
            output_line = append(output_line, inst_name)
        }
        if *flag_externalip {
            output_line = append(output_line, *inst.PublicIpAddress)
        }
        if *flag_internalip {
            output_line = append(output_line, *inst.PrivateIpAddress)
        }
        if flag_filter_count == 0 {
            fmt.Printf("%s %s %s\n", inst_name, *inst.PublicIpAddress, *inst.PrivateIpAddress)
        } else {
            fmt.Println(strings.Join(output_line, " "))
        }
    }

}
