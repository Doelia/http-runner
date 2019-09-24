# cleanup
echo "Cleanup old files..."
rm -rf ~/.http-runner

# install
echo "Create .http-runner home folder..."
mkdir ~/.http-runner
curl https://raw.githubusercontent.com/Doelia/http-runner/master/installation/.http-runner/config.yaml > .http-runner/config.yaml
curl https://raw.githubusercontent.com/Doelia/http-runner/master/installation/http-runner > /usr/local/bin/http-runner
chmod o+x /usr/local/bin/http-runner
