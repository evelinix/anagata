pipeline {
    agent {
        label 'windows'
    }

    options {
        timeout(time: 30, unit: 'MINUTES')
        disableConcurrentBuilds()
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Install') {
            steps {
                bat 'cd frontend && pnpm install --frozen-lockfile'
            }
        }

        stage('Build Frontend') {
            steps {
                bat 'cd frontend && pnpm build'
            }
        }

        stage('Generate Bindings') {
            steps {
                bat 'wails generate module'
            }
        }

        stage('Lint') {
            parallel {
                stage('Go Lint') {
                    steps {
                        bat 'go vet ./...'
                    }
                }
                stage('Frontend Lint') {
                    steps {
                        bat 'cd frontend && pnpm lint'
                    }
                }
            }
        }

        stage('Test') {
            parallel {
                stage('Go Test') {
                    steps {
                        bat 'go test ./internal/... -v'
                    }
                }
                stage('Frontend Test') {
                    steps {
                        bat 'cd frontend && pnpm test -- --run'
                    }
                }
            }
        }

        stage('Build') {
            steps {
                bat 'wails build'
            }
        }

        stage('Package') {
            steps {
                archiveArtifacts artifacts: 'build/bin/*.exe', fingerprint: true
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
