# get current commit hash for tag
commithash=$(git rev-parse HEAD)

# define semantic version
semver=$(/bin/bash `dirname "$0"`/semver.sh)

# get repo name from git
remote=$(git config --get remote.origin.url)
repo=$(basename $remote .git)

# POST a new ref to repo via Github API
curl -s -X POST https://api.github.com/repos/${GITHUB_REPOSITORY,,}/git/refs \
	-H "Authorization: token ${GH_TOKEN}" \
	-d @- << EOF
{
  "ref": "refs/tags/$semver",
  "sha": "$commithash"
}
EOF
