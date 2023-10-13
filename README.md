# Colosseum Project - Arena

Microservice application part of the Colosseum Project and responsible for engaging gladiators to battle.

- [Colosseum Project - Arena](#colosseum-project---arena)
  - [License](#license)
  - [Contributing](#contributing)
  - [Build OCI image](#build-oci-image)

---

## License

This application is released under the terms of the MIT license.
See [LICENSE](LICENSE) for more information or see <https://opensource.org/licenses/MIT>.

## Contributing

If you want to contribute to the project, please read the [CONTRIBUTING file](CONTRIBUTING.md).

## Build OCI image

Images can be built using Docker and the provided Dockerfile.
To build an image, run:

```sh
docker build -t arena .

# or build minimal x86-64 container image
docker build -f Dockerfile.amd64-tiny -t arena:amd64-tiny .
```
