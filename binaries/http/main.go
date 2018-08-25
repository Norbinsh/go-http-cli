package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/visola/go-http-cli/cli"
	"github.com/visola/go-http-cli/daemon"
	"github.com/visola/go-http-cli/output"
	"github.com/visola/go-http-cli/request"
)

func main() {
	if daemonErr := daemon.EnsureDaemon(); daemonErr != nil {
		panic(daemonErr)
	}

	options, err := cli.ParseCommandLineOptions(os.Args[1:])

	if err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}

	executionOptions := request.ExecutionOptions{
		FileToUpload:    options.FileToUpload,
		FollowLocation:  options.FollowLocation,
		MaxRedirect:     options.MaxRedirect,
		PostProcessFile: options.PostProcessFile,
		ProfileNames:    options.Profiles,
		RequestName:     options.RequestName,
		Request: request.Request{
			Body:    options.Body,
			Headers: options.Headers,
			Method:  options.Method,
			URL:     options.URL,
			Values:  options.Values,
		},
		Variables: options.Variables,
	}

	requestExecution, requestError := daemon.ExecuteRequest(executionOptions)

	if requestError != nil {
		color.Red("Error while executing request: %s", requestError)
		os.Exit(10)
	}

	for _, requestResponse := range requestExecution.RequestResponses {
		output.PrintRequest(requestResponse.Request)
		fmt.Println("")
		output.PrintResponse(requestResponse.Response)
		if options.OutputFile != "" && requestResponse.Response.Body != "" {
			outWriteErr := ioutil.WriteFile(options.OutputFile, []byte(requestResponse.Response.Body), 0644)
			if outWriteErr != nil {
				color.Red("Error while writing to output file: %s", outWriteErr)
			}
		}
	}

	if requestExecution.ErrorMessage != "" {
		color.Red("Error while executing request: %s", requestExecution.ErrorMessage)
		os.Exit(20)
	}

	if requestExecution.PostProcessError != "" {
		color.Red("Error post processing request: %s", requestExecution.PostProcessError)
		os.Exit(30)
	}
}
