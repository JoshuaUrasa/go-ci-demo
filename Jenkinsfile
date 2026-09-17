pipeline {
    agent {
        label 'go'
    }

    stages {
        stage('Test') {
            steps {
                sh 'go test -v -count=1 ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o demo .'
            }
        }

        stage('Run') {
            steps {
                sh './demo'
            }
        }
    }
}
