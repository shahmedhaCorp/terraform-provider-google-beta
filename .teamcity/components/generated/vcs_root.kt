// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot

// providerRepository is the config for a VCS root that's reused when defining the project and
// build configurations inside the project
// For GitVcsRoot docs, see https://teamcity.jetbrains.com/app/dsl-documentation/vcs/git-vcs-root/index.html
object providerRepository : GitVcsRoot({
    name = "terraform-provider-google-beta"
    url = "https://github.com/hashicorp/terraform-provider-google-beta.git"
    agentCleanPolicy = AgentCleanPolicy.ON_BRANCH_CHANGE
    agentCleanFilesPolicy = AgentCleanFilesPolicy.ALL_UNTRACKED
    branchSpec = "+:*"
    branch = "refs/heads/main"
    authMethod = anonymous()
})
