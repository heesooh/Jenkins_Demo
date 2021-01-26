# Jenkins_Demo

This is a simple Jenkins Demo.

Use the following command in Jenkins build shell command to create and send log file with the date and time when the test result is unstable.

1. go build
2. filename=log::"`date +"%Y-%m-%d-%I-%M-%S-%p"`".txt
3. go test >> $filename
4. ./Jenkins_Demo -filenme $filename