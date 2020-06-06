# Generator
A mircoservice to generate stuff

The project boilerplate has been created using [go-wagen](https://github.com/groovili/go-wagen) web application generator.

#### Required Tools

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Build Tools
    * on [Mac OS X](https://osxdaily.com/2014/02/12/install-command-line-tools-mac-os-x/)  and compatible systems
    * on [Ubuntu](https://askubuntu.com/a/272020/900576) and compatible systems
    * on [Centos](https://unix.stackexchange.com/a/32439/91242)  and compatible systems
    * on [Windows](https://stackoverflow.com/a/32127632/6670698)  and compatible systems
    * or get some help from [`>_ command-not-found.com`](https://command-not-found.com/)

_Windows users will either need to have [WSL install](https://docs.microsoft.com/en-us/windows/wsl/install-win10)ed or install GNU `make` using [other ways](https://stackoverflow.com/a/32127632/6670698) _

#### Building

- Clone the repository
    ```
    git clone https://github.com/scalerstech/generator.git
    ```
- Change into the generator directory
    ```
    cd generator
    ```

- Run the application
    ```
    make run
    ```

## APIs

### GET `/`

Sample `Hello World` response as per boilerplate. See [go-wagen](https://github.com/groovili/go-wagen) for details.

### GET `/ping`

Sample `ping` response as per boilerplate. See [go-wagen](https://github.com/groovili/go-wagen) for details.

### GET `/urandom`

HTTP service to emulate *nix device `/dev/urandom`

#### Supported Query Parameters
    
*   `length` - Number of bytes of data to be returned. Default value: `1024`.

### GET `/fbsdrandom`

HTTP service to emulate the FreeBSD device `/dev/urandom`

#### Supported Query Parameters

* `length` - Number of bytes of data to be returned. Default value: `1024`.

  _Please Note: This API sends data as `application/octet-stream`._

### GET `/password`

HTTP service to generate a strong password

#### Supported Query Parameters

* `format` - Data output format (`json`, `raw`). Default: `json`

* `length` - Length of the generated password

* `digits` - Minimum number of digits to have in the generated password

* `symbol` - Minimum number of symbols to have in the generated password

* `lowercase` - Boolean value to ensure if the generated password needs to have lowercase alphabets only

* `repeated` - Boolean value to ensure if the generated password string will have repeated characters or not
