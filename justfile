default: deps build

version := `
export TAG=$(git describe --tags)
if [[ -z "$TAG" ]]; then
    echo "dev+$(git rev-parse --short HEAD)"
else
    echo $TAG
fi
`
versionFlag := "-X 'main.Version=" + version + "'"

gitCommit := `git rev-parse --short HEAD`
commitFlag := "-X 'main.Commit=" + gitCommit + "'"

buildDate := `TZ=UTC LC_TIME=en_US.UTF-8 date +"%a %b %d %Y, %T %Z"`
dateFlag := "-X 'main.Compiled=" + buildDate + "'"

ldFlags := "-s -w" + " " + commitFlag + " " + dateFlag + " " + versionFlag

@build:
  if ! test -d ./build; then \
    mkdir build; \
  fi
  echo "Building..."
  go build -ldflags "{{ldFlags}}" -o ./build/frg .

@deps:
  go mod download

@lint:
  golangci-lint run --fix

@run ARGS='--help':
  ./build/frg {{ARGS}}
