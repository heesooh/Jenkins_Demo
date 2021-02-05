# Jenkins_Demo

This is a simple Jenkins Demo.

Use the following command in Jenkins build shell command to create and send log file when the test result is unstable.
Both Freestyle and Pipeline versions are supported. 

----------------------------------------------------------
Free Style:

1. go build
2. filename=log::"`date +"%Y-%m-%d-%I-%M-%S-%p"`".txt
3. set +e
4. go test >> $filename
5. ./Jenkins_Demo -filename $filename

----------------------------------------------------------

Pipeline: 

```
def FILENAME
def ERROR

pipeline {
    agent any
    tools {
        go 'go-1.15.7'
    }
    environment {
        GO1157MODULE = 'on'
    }
    stages {
        stage('SCM Checkout') {
            steps {
                git 'https://github.com/heesooh/Jenkins_Demo'
            }
        }
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh 'go test > log.txt'
                }
            }
        }
        stage('Deploy') {
            steps {
                sh './Jenkins_Pipeline -filename log.txt'
            }
        }
        stage('Close') {
            steps {
                sh 'rm log.txt'
            }
        }
    }
}

```