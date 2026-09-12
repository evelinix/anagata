pipeline {
    agent {
        label 'windows'
    }

    options {
        timeout(time: 30, unit: 'MINUTES')
        disableConcurrentBuilds()
        buildDiscarder(logRotator(numToKeepStr: '10'))
    }

    triggers {
        GenericTrigger(
            genericVariables: [
                [key: 'ref', value: '$.ref']
            ],
            token: 'sentinel-token',
            causeString: 'Triggered by GitHub push to $ref',
            printContributedVariables: true,
            printPostContent: true
        )
    }

    stages {
        stage('Filter') {
            steps {
                script {
                    if (env.ref != 'refs/heads/main') {
                        echo "Skipping build for branch: ${env.ref}"
                        currentBuild.result = 'NOT_BUILT'
                        currentBuild.description = 'Skipped - not main branch'
                        env.SKIP_BUILD = 'true'
                    }
                }
            }
        }

        stage('Checkout') {
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
            steps {
                git branch: 'main', url: 'https://github.com/evelinix/anagata.git'
            }
        }

        stage('Install') {
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
            steps {
                bat 'cd frontend && pnpm install --frozen-lockfile'
            }
        }

        stage('Build Frontend') {
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
            steps {
                bat 'cd frontend && pnpm build'
            }
        }

        stage('Lint') {
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
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
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
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
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
            steps {
                bat 'wails build'
            }
        }

        stage('Package') {
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
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
