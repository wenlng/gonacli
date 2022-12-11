
Step 1: Generate c/c++ code of nodejs addon bridging golang
> ./gonacli-darwin generate

Step 2: Generate golang library
> ./gonacli-darwin build

Step 3: Compile c/c++ napi to nodejs addon
Tip: Ensure that nodejs, npm and node-gyp are installed on the OS
> ./gonacli-darwin make