pipeline {
    agent none
    stages {
        stage('Build') {
            when {
                beforeAgent true
                branch 'master'
            }
            agent any
            steps {
                sh """
                    docker build \
                        -f ./scripts/build/Dockerfile \
                        -t duck-api:latest \
                        -t duck-api:dev-${env.GIT_COMMIT[0..7]} \
                        .
                """
            }
        }
        stage('Deploy') {
            parallel {
                stage('Development') {
                    agent any
                    when {
                        branch 'master'
                    }
                    steps {
                        withCredentials(bindings: [sshUserPrivateKey(credentialsId: 'DEV_SERVER', keyFileVariable: 'SSH_KEY_FOR_DEV')]) {
                            sh """
                                    ssh -i $SSH_KEY_FOR_DEV -t -oStrictHostKeyChecking=no nc-user@nc-machine \"
                                    sed -i 's/API_TAG=.*/API_TAG=dev-${env.GIT_COMMIT[0..7]}/g' /home/nc-user/app/dev/.env
                                    cd /home/nc-user/app
                                    cd dev && ./start-app.sh
                                \"
                            """
                        }
                    }
                }
            }
        }
    }
}
