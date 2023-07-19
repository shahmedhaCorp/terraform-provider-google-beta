// this file is auto-generated with mmv1, any changes made here will be overwritten

package tests

import GoogleBeta
import org.junit.Assert.assertTrue
import org.junit.Test

class VcsTests {
    @Test
    fun buildsHaveCleanCheckOut() {
        val project = GoogleBeta("default", "refs/heads/main", TestVcsRootId(), TestConfiguration())
        project.buildTypes.forEach { bt ->
            assertTrue("Build '${bt.id}' doesn't use clean checkout", bt.vcs.cleanCheckout)
        }
    }
}