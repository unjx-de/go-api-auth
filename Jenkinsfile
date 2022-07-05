pipeline {
    agent any
    tools {
        go 'go-1.18'
    }
    environment {
        GOPATH = "${WORKSPACE}/go"
    }
    stages {
        stage('Test') {
            steps {
                sh 'go test -cover'
                sh 'go test -v 2>&1 ./... | /home/flohoss/go-junit-report -set-exit-code > report.xml'
            }
        }
        stage('Integration') {
            steps {
                junit 'report.xml'
            }
        }
    }
}