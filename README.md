# ecpockman
Store realtime data of exchange centers


# Deployment

### Download [release](https://github.com/SpeculationFund/ecpockman/releases)

### Start mongodb
```
docker run --name dbtest -p 27017:27017 -d mongo
```

### Unzip, Config, Install
```
unzip ecpockman_amd64-YYYYMMDD.zip
cd .Build 
./ecpockman
```


### Validation Steps

* Execute `mongo` in CLI to start mongodb client
* Execute `show databases` in mongo client CLI to validate `SpeculationFund` was created 
* Execute `use SpeculationFund` to adopt `SpeculationFund` db
* Execute `db.Ticker.find()` to display the created data


# Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

* [Install Docker](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-16-04)

* Install golang  `sudo apt install golang-go`


### Installing

```
docker run --name dbtest -p 27017:27017 -d mongo
git clone https://github.com/SpeculationFund/ecpockman.git
cd ecpockman
go build
./ecpockman
```

### Running the tests

```
go test
```

### Documentation
* [GoDoc of ecpockman](https://godoc.org/github.com/SpeculationFund/ecpockman)
* [Development docs of ecpockman](https://github.com/SpeculationFund/ecpockman/wiki)


### Build

```
cd github.com/SpeculationFund/ecpockman
./build_release.sh
```
The output is `ecpockman-yyyymmdd.zip`, details are shown as follow,

```
Archive:  ecpockman-20171228.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  2017-12-28 17:39   .Build/
  9605928  2017-12-28 17:39   .Build/ecpockman_amd64-20171228
  7729176  2017-12-28 17:39   .Build/ecpockman_386-20171228
  7701280  2017-12-28 17:39   .Build/ecpockman_arm-20171228
---------                     -------
 25036384                     4 files
``` 

# Logistics

### Contributing

Please read [CONTRIBUTING.md](https://github.com/SpeculationFund/ecpockman/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository

### Authors

* **PhoenixAndMachine** - *Initial work* - [PhoenixAndMachine](https://github.com/PhoenixAndMachine)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

### Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc


### License

This project is licensed under the MIT License - see the [LICENSE.md](https://gist.github.com/Brownyuan/0b754b6009b7a4257bde9d1a23586678) file for details


