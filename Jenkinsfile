pipeline {
    agent {
        label 'go'
    }

    triggers {
        pollSCM('* * * * *')
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
        } // Run stage ends here

        stage('Package') {
            steps {
                sh '''
                    mkdir -p dist
                    tar -czf "dist/go-ci-demo-${BUILD_NUMBER}.tar.gz" demo
                '''
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts(
                    artifacts: "dist/go-ci-demo-${env.BUILD_NUMBER}.tar.gz",
                    allowEmptyArchive: false
                )
            }
        }
    }
}
