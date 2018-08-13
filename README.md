# getweather
Small golang app which takes weather for certain city from Openweathermap.

## Prerequisites
### Install
* [Vagrant](https://www.vagrantup.com/)
* [VirtualBox](https://www.virtualbox.org/)

## Usage
### Run Vagrant
You can pass your to vagrant your Openweather api key anf name of city through environment variables
```
OPENWEATHER_API_KEY=XXXXXXXXXXXXXXXXXXXXX CITY_NAME=Tivat vagrant up 
```
Result will be in output of your console. It shows string from syslog. Docker container with app writes its output to syslog.

### Run Docker
App has own dockerfile, before running container you need to build an docker image. 
````
docker build -t [name]:[tag] .
````

Container must be executed with parameters.
```
docker run --rm -e OPENWEATHER_API_KEY="XXXXXXXXXXXX" -e CITY_NAME="XXXXXX" [name]:[tag]
```
### Run Go App without Docker and Vagrant
The application is written without using third-party libraries. There are no special requirements to run app, except environment variables. The application uses environment variables as parameters
````
cd [path to app folder]
OPENWEATHER_API_KEY=XXXXXXXXXXXXXXXXXXXXX CITY_NAME=XXXXXX go run main.go
````


