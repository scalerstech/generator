# generator
A mircoservice to generate stuff

The project boilerplate has been created using [go-wagen](https://github.com/groovili/go-wagen) web application generator.

#### Required tools
- Docker
- Docker Compose
- Make tools

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
---

#### GET `/`
Sample `Hello World` response as per boilerplate. See [go-wagen](https://github.com/groovili/go-wagen) for details.

#### GET `/ping`
Sample `ping` response as per boilerplate. See [go-wagen](https://github.com/groovili/go-wagen) for details.

#### GET `/urandom`
HTTP service to emulate *nix device `/dev/urandom`

Supported Query Parameters

`length` - Number of bytes of data to be returned. Default value: `1024`.

#### GET `/fbsdrandom`
HTTP service to emulate the FreeBSD device `/dev/urandom`

Supported Query Parameters

`length` - Number of bytes of data to be returned. Default value: `1024`.

_Please Note: This API sends data as `application/octet-stream`._

#### GET `/password`
HTTP service to generate a strong password

Supported Query Parameters

`format` - Data output format (`json`, `raw`). Default: `json`

`length` - Length of the generated password

`digits` - Minimum number of digits to have in the generated password

`symbol` - Minimum number of symbols to have in the generated password

`lowercase` - Boolean value to ensure if the generated password needs to have lowercase alphabets only

`repeated` - Boolean value to ensure if the generated password string will have repeated characters or not

