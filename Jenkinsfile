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
                    env.IS_TAG = env.ref?.startsWith('refs/tags/') ? 'true' : 'false'
                    env.TAG_NAME = env.IS_TAG == 'true' ? env.ref.replace('refs/tags/', '') : ''
                    env.IS_RELEASE = env.IS_TAG

                    if (env.IS_TAG != 'true' && env.ref != 'refs/heads/main') {
                        echo "Skipping build for: ${env.ref}"
                        currentBuild.result = 'NOT_BUILT'
                        currentBuild.description = 'Skipped - not main or tag'
                        env.SKIP_BUILD = 'true'
                    }

                    if (env.IS_TAG == 'true') {
                        echo "Release build for tag: ${env.TAG_NAME}"
                        currentBuild.description = "Release ${env.TAG_NAME}"
                    }
                }
            }
        }

        stage('Checkout') {
            when { not { environment name: 'SKIP_BUILD', value: 'true' } }
            steps {
                git branch: 'main', url: 'https://github.com/evelinix/anagata.git', credentialsId: 'github-pat'
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

        stage('Release') {
            when {
                allOf {
                    environment name: 'IS_RELEASE', value: 'true'
                    not { environment name: 'SKIP_BUILD', value: 'true' }
                }
            }
            steps {
                withCredentials([usernamePassword(credentialsId: 'github-pat', usernameVariable: 'GH_USER', passwordVariable: 'GH_TOKEN')]) {
                    powershell """
                        \$tagName = '${env.TAG_NAME}'
                        \$releaseName = "AnagataSentinel \$tagName"
                        \$exePath = 'build\\bin\\AnagataSentinel.exe'
                        \$headers = @{
                            Authorization = "token \$env:GH_TOKEN"
                            Accept = 'application/vnd.github+json'
                        }

                        Write-Host "Creating release: \$releaseName"

                        \$body = @{
                            tag_name = \$tagName
                            name = \$releaseName
                            body = "## AnagataSentinel " + \$tagName + "`n`n### Changes`n- See commit history for details`n`n### Download`n- Download AnagataSentinel.exe below"
                            draft = \$false
                            prerelease = \$false
                        } | ConvertTo-Json -Depth 3

                        \$release = Invoke-RestMethod -Uri 'https://api.github.com/repos/evelinix/anagata/releases' -Method Post -Headers \$headers -Body \$body -ContentType 'application/json'
                        Write-Host "Release created: \$release.html_url"

                        Write-Host "Uploading AnagataSentinel.exe..."
                        \$baseUrl = \$release.upload_url.Split('{')[0]
                        \$uploadUrl = \$baseUrl + '?name=AnagataSentinel.exe'
                        \$fileBytes = [System.IO.File]::ReadAllBytes(\$exePath)
                        \$uploadHeaders = @{
                            Authorization = "token \$env:GH_TOKEN"
                            Content-Type = 'application/octet-stream'
                        }
                        Invoke-RestMethod -Uri \$uploadUrl -Method Post -Headers \$uploadHeaders -Body \$fileBytes

                        Write-Host "Upload complete!"
                        Write-Host "URL: \$release.html_url"
                    """
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
