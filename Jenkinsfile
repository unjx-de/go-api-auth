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
                sh 'go mod download'
                sh 'go test -cover'
            }
        }
    }
}