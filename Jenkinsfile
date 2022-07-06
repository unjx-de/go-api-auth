pipeline {
    agent any
    tools {
        go 'go-1.18'
    }
    stages {
        stage('Test') {
            steps {
                sh 'go mod download'
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