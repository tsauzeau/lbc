### Build the project

```bash
go mod vendor
make
```

The binary will be in the ./bin/OS_ARCH/ folder

### Start the dev environment

```bash
docker-compose up -d
go mod vendor
make
./bin/OS_ARCH/lbc
```

Go on http://127.0.0.1:4242/v1/fizzbuzz?string1=toto&string2=tataa&limit=100&int1=5&int2=10
or http://127.0.0.1:4242/v1/stat

### Build the container for release

```bash
go mod vendor
make all-container
```

The container tag will be output

### Configure the application

A config file is stored in the ./config dir, you can directly edit it.
For release purpose with docker, you can use as-well env varibles (it will override the config file): 'LBC_APIPORT=":4343" LBC_REDISHOST="redis:6379" bin/darwin_amd64/lbc'

### Adding New Libraries/Dependencies

```bash
go mod vendor
```

### Using GitHub Registry

Create and Push:

```bash
docker login docker.pkg.github.com -u <USERNAME> -p <GITHUB_TOKEN>
docker build -t  docker.pkg.github.com/tsauzeau/lbc/lbc:latest .
# make container
docker push docker.pkg.github.com/tsauzeau/lbc/lbc:latest
# make push
```

Pull and Run:

```bash
docker pull docker.pkg.github.com/tsauzeau/lbc/lbc:latest
docker run docker.pkg.github.com/tsauzeau/lbc/lbc:latest
```

### Ananlyze the project with SonarQube

- On _SonarQube_:
  - Go on http://127.0.0.1:9000 (user: admin / password: admin)
  - Create a new project
  - In the sonar-project.properties change the projectKey, projectName and login keys by the keys output by sonar.
  - Run 'make ci'
  - Use the sonar-scanner 'sonar-scanner' command directly in the project root directory (brew install sonar-scanner on mac).
  - Check the result on the sonarQube interface (http://127.0.0.1:9000)
