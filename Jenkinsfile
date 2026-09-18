pipeline {
    agent {
        label 'go'
    }

    triggers {
        pollSCM('* * * * *')
    }

    stages {
        stage('Prepare tools') {
            steps {
                sh '''
            mkdir -p "$WORKSPACE/.tools"
            GOBIN="$WORKSPACE/.tools" go install gotest.tools/gotestsum@v1.13.0
        '''
            }
        }

        stage('Dependencies') {
            steps {
                sh '''
  go mod download
  go mod verify
  '''}}
        stage('Test') {
            steps {
                sh '''
            mkdir -p reports
            rm -f reports/tests.xml

           "$WORKSPACE/.tools/gotestsum" \
                --junitfile reports/tests.xml \
                -- -count=1 ./...
        '''
            }

            post {
                always {
                    junit 'reports/tests.xml'
                }
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
                    Fingerprint:true
                )
            }
        }
    }

        post {
        always {
            echo "Build #${env.BUILD_NUMBER} has finished executing."
        }

        success {
            echo 'SUCCESS: All stages completed successfully.'
        }

        failure {
            echo 'FAILURE: Check Console Output for the first error.'
        }
        }
}
