Hacking on Knitter
==========================

## Local development

Knitter comes with a `Makefile` which defines following targets:

* `build` - is the default target responsible for building S2I binary, under the covers
it calls `hack/build-go.sh`. The resulting binary will be placed in `_output/local/go/bin/`.
* `all` - is synonym for `build`.
* `test` - is responsible for testing Knitter, under the covers it calls `hack/test-go.sh`.
Additionally you can pass `WHAT` or `TEST` variable specifying directory names to test,
eg. `make test WHAT=pkg/build`
* `check` - is synonym for `test`.
* `clean` - cleans environment by removing `_output` and `Godeps/_workspace/pkg` directories.

## Test Suites

Knitter uses two levels of testing - unit tests and integration tests, both of them are run on each
pull request, so make sure to run those before submitting one.


### Unit tests

Unit tests follow standard Go conventions and are intended to test the behavior and output of a
single package in isolation. All code is expected to be easily testable with mock interfaces and
stubs, and when they are not it usually means that there's a missing interface or abstraction in the
code. A unit test should focus on testing that branches and error conditions are properly returned
and that the interface and code flows work as described. Unit tests can depend on other packages but
should not depend on other components.

The unit tests for an entire package should not take more than 0.5s to run, and if they do, are
probably not really unit tests or need to be rewritten to avoid sleeps or pauses. Coverage on a unit
test should be above 70% unless the units are a special case.

Run the unit tests with:

    $ hack/test-go.sh

or an individual package unit test with:

    $ hack/test-go.sh pkg/adapter

To run only a certain regex of tests in a package, use:

    $ hack/test-go.sh pkg/SpecificPackage -test.run=TestSpecificCase

To get verbose output add `-v` to the end:

    $ hack/test-go.sh pkg/SpecificPackage -test.run=TestSpecificCase -v

To run all tests with verbose output:

    $ hack/test-go.sh "" -v

To run tests without the Go race detector, which is on by default, use:

    $ KNITTER_RACE="" hack/test-go.sh

A line coverage report is printed by default. To turn it off, use:

    $ KNITTER_COVER="" hack/test-go.sh

To create an HTML coverage report for all packages:

    $ OUTPUT_COVERAGE=/tmp/knitter-cover hack/test-go.sh


### Integration tests

The second category are integration tests which verify the whole Knitter flow. The integration tests
require a couple of images for testing, these can be built with `hack/build-test-images.sh`, if
integration tests don't find them it'll print appropriate information regarding running this command.

Run the integration tests with:

    $ hack/test-integration.sh


## Installing Godep

Knitter uses [Godep](https://github.com/tools/godep) for dependency management.
Godep allows versions of dependent packages to be locked at a specific commit by *vendoring* them
(checking a copy of them into `Godeps/_workspace/`).  This means that everything you need for
Knitter is checked into this repository.  To install `godep` locally run:

    $ go get github.com/tools/godep

If you are not updating packages you should not need godep installed.


## Building a Release

To build a Knitter release you run `make release` on a system with Docker,
which will create a build environment image and then execute a cross platform Go build within it. The build
output will be copied to `_output/releases` as a set of tars containing each version.

1. Create a new git tag `git tag vX.X.X -a -m "vX.X.X" HEAD`
2. Push the tag to GitHub `git push origin --tags` where `origin` is `github.com/HyperNetworks/knitter.git`
4. Run `make release`
5. Upload the binary artifacts generated by that build to GitHub release page.
6. Send an email to the dev list, including the important changes prior to the release.