// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.Project
import jetbrains.buildServer.configs.kotlin.AbsoluteId

const val providerName = "google-beta"

// GoogleBeta returns an instance of Project,
// which has multiple build configurations defined within it.
// See https://teamcity.jetbrains.com/app/dsl-documentation/root/project/index.html
fun GoogleBeta(environment: String, branchRef: String, manualVcsRoot: AbsoluteId, configuration: ClientConfiguration) : Project {

    return Project{

        var buildConfigs = buildConfigurationsForPackages(packages, providerName, "google-beta", environment, branchRef, manualVcsRoot, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
    }
}

fun buildConfigurationsForPackages(packages: Map<String, String>, providerName : String, path : String, environment: String, branchRef: String, manualVcsRoot: AbsoluteId, config: ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()

    packages.forEach { (packageName, displayName) ->
        if (packageName == "services") {
            // `services` is a folder containing packages, not a package itself; call buildConfigurationsForPackages to iterate through directories found within `services`
            var serviceList = buildConfigurationsForPackages(services, providerName, path+"/"+packageName, environment, branchRef, manualVcsRoot, config)
            list.addAll(serviceList)
        } else {
            // other folders assumed to be packages
            var testConfig = testConfiguration(environment)

            var pkg = packageDetails(packageName, displayName, environment, branchRef)
            var buildConfig = pkg.buildConfiguration(providerName, path, manualVcsRoot, true, testConfig.startHour, testConfig.parallelism, testConfig.daysOfWeek, testConfig.daysOfMonth)

            buildConfig.params.ConfigureGoogleSpecificTestParameters(config)

            list.add(buildConfig)
        }
    }

    return list
}

class testConfiguration(environment: String, parallelism: Int = defaultParallelism, startHour: Int = defaultStartHour, daysOfWeek: String = defaultDaysOfWeek, daysOfMonth: String = defaultDaysOfMonth) {

    // Default values are present if init doesn't change them
    var parallelism = parallelism
    var startHour = startHour
    var daysOfWeek = daysOfWeek
    var daysOfMonth = daysOfMonth

    init {
        // If the environment parameter is set to the value of MAJOR_RELEASE_TESTING, 
        // change the days of week to the day for v5.0.0 feature branch testing
        if (environment == MAJOR_RELEASE_TESTING) {
            this.parallelism = parallelism
            this.startHour = startHour
            this.daysOfWeek = "5" // Friday for Beta
            this.daysOfMonth = daysOfMonth
        }
    }

}