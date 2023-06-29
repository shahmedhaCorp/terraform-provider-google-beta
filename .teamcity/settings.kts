import GoogleBeta
import ClientConfiguration
import jetbrains.buildServer.configs.kotlin.*

version = "2023.05"

var custId = DslContext.getParameter("custId", "")
var org = DslContext.getParameter("org", "")
var org2 = DslContext.getParameter("org2", "")
var billingAccount = DslContext.getParameter("billingAccount", "")
var billingAccount2 = DslContext.getParameter("billingAccount2", "")
var masterBillingAccount = DslContext.getParameter("masterBillingAccount", "")
var credentials = DslContext.getParameter("credentials", "")
var environment = DslContext.getParameter("environment", "public")

var clientConfig = ClientConfiguration(custId, org, org2, billingAccount, billingAccount2, masterBillingAccount, credentials)

project(GoogleBeta(environment, clientConfig))