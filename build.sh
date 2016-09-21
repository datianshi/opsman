if [ -z ${VERSION} ]; then
  echo "USAGE: VERSION NUMBER need to be provided"
  exit 1 
fi
GOOS=linux GOARCH=amd64 go build -o "out/linux/opsman-cli"
GOOS=darwin GOARCH=amd64 go build -o "out/darwin/opsman-cli"
GOOS=windows GOARCH=amd64 go build -o "out/windows/opsman-cli.exe"
tar -czvf opsman-${VERSION}.tgz out 
git tag ${VERSION}
git push origin ${VERSION}
