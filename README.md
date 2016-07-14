# openclan-server
Server implementation for OpenClan

## Building OpenClan Server
### Requirements
- Ubuntu 16.04 or later
- snapcraft
- golang

### Creating a snap
- Run `snapcraft stage` within the directory containing the source
- If successful, run `snapcraft snap` to create a new snap
- There will now by a snap with a name similar to `openclan-server_0_amd64.snap`

### Installing the snap locally
- Run `sudo snap install <filename of snap here>` or `sudo snap install *.snap`
- The resulting snap should be running on your system and the server will be accessible via `http://localhost:8080/v1/`
