# Jenkins_Demo

This is a simple Jenkins Demo.

Use the following command in Jenkins build shell command to create and send log file with the date and time when the test result is unstable.

----------------------------------------------------------
Free Style:

1. go build
2. filename=log::"`date +"%Y-%m-%d-%I-%M-%S-%p"`".txt
3. go test >> $filename
4. errorMessage=$(go test)
4. ./Jenkins_Demo -filename $filename -error $errorMessage

----------------------------------------------------------

Pipeline:

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
                script {
                    FILENAME = new Date()
                    FILENAME = FILENAME.format("yyyyMMddHHmm")
                }
                echo "$FILENAME"
            }
        }
        stage('Test') {
            steps {
                sh 'go test > log.txt'
            }
        }
        stage('Deploy') {
            steps {
                sh 'go build'
                sh './Jenkins_Pipeline -filename log.txt -error FAIL'
            }
        }
        stage('Close') {
            steps {
                sh 'rm log.txt'
            }
        }
    }
}