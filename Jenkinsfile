pipeline {
    agent {
        label 'go'
    }

    triggers {
        pollSCM('* * * * *')
    }

    environment {
        APP_NAME = 'go-ci-demo'
        LAB_MODE = 'normal'
    }

    parameters {
        string(
            name: 'BUILD_MESSAGE',
            defaultValue: 'Learning Jenkins',
            description: 'A message to display for this build'
        )

        booleanParam(
            name: 'RUN_APP',
            defaultValue: true,
            description: 'Run the application after building'
        )

        choice(
            name: 'BUILD_MODE',
            choices: ['normal', 'diagnostic'],
            description: 'Choose whether to show agent diagnostics'
        )
    }

    stages {

    stage('Show parameters') {
        steps {
            echo "Build message: ${params.BUILD_MESSAGE}"
        }
    }

    stage('Stage environment') {
        environment {
            LAB_MODE = 'testing'
        }

        steps {
            echo "Application: ${env.APP_NAME}"
            sh 'echo "Inside first stage: $LAB_MODE"'
        }
    }

    stage('Pipeline environment') {
        steps {
            sh 'echo "Inside second stage: $LAB_MODE"'
        }
    }

    stage('Diagnostics') {
        when {
            expression { return params.BUILD_MODE == 'diagnostic' }
        }
        steps {
            sh '''
                whoami
                pwd
                go version
                git --version
            '''
        }
    }

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
            when {
                expression { return params.RUN_APP }
            }
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
                    allowEmptyArchive: false,
                    fingerprint:true
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

        aborted {
            echo 'ABORTED: Build was interrupted or timed out.'
        }
        }
}
