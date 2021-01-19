# Its looks like semver, but its not, because we cannot release patch after major/minor release.
# There is no ability to release patch as backport
# magic number 4b825dc642cb6eb9a060e54bf8d69288fbee4904 is git commit hash of null (before init commit)
MAJOR_VERSION=0
MAJOR_LAST_COMMIT_HASH="4b825dc642cb6eb9a060e54bf8d69288fbee4904"
MINOR_LAST_COMMIT_HASH=$(git rev-list --invert-grep -i --grep="fix" ${MAJOR_LAST_COMMIT_HASH}..HEAD -n 1)
MINOR_VERSION=$(git rev-list --invert-grep -i --grep="fix" ${MAJOR_LAST_COMMIT_HASH}..HEAD --count)
PATCH_VERSION=$(git rev-list ${MINOR_LAST_COMMIT_HASH}..HEAD --count)
SEMVER=${MAJOR_VERSION}.${MINOR_VERSION}.${PATCH_VERSION}
HELMVER=$(grep -oP "^version: \K[^ ]*$" helm/Chart.yaml)
APPVER=$(grep -oP "^appVersion: \"\K[0-9\.]*" helm/Chart.yaml)
echo ${SEMVER}
[ ${SEMVER} == ${HELMVER} ] || exit 1
#(>&2 echo "package version \"${HELMVER}\" in Chart.yaml does not match semantic version"; exit 1)
[ ${SEMVER} == ${APPVER} ] || exit 1
#(>&2 echo "app version \"${APPVER}\" in Chart.yaml does not match semantic version in repository"; exit 1)
