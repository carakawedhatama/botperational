# botperational

`botperational` is a service to update pps data

## Project status

It is an ongoing project, with the current maintainer is dev@runsystem.id

## Prerequisites
Before using this project, make sure you have the following applications or libraries installed:

- Go 1.21 or higher. You could download it [here](https://go.dev/dl/).
- [cosmtrek/air](https://github.com/cosmtrek/air) (to run watch mode)
- MariaDB 10.3 or higher. You could download it [here](https://mariadb.org/download/).

## Usage
To use `botperational`, follow these steps: 
1. Clone this project's `development` branch
```bash
git clone -b development [git_url]
```

2. Copy the config file, and adjust as it fits your needs
```bash
cp config.sample.yml config.yml
cp .env.sample .env
```
You could ask the maintainer about the sample config files.

3. Download the dependency library
```bash
go mod tidy
```

## License
This project is under [MIT License](http://botperational/-/blob/development/LICENSE).