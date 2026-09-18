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
        sh '''
            mkdir -p reports
            rm -f reports/tests.xml

            "$HOME/.local/bin/gotestsum" \
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
