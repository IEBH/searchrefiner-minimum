# searchrefiner

_Systematic Review Query Visualisation and Understanding Interface_

searchrefiner is an interactive interface for visualising and understanding queries used to retrieve medical literature for
[systematic reviews](https://en.wikipedia.org/wiki/Systematic_review).

It is currently in development, however please find a demo link [on the project home page](https://searchrefinery.sr-accelerator.com/).

## Building

searchrefiner is built as a Go application.

 1. First, clone this repository.
 2. Configure the application. The application can then be configured via a `config.json` (a [sample](sample.minimal.config.json) is provided). In this minimal file, everything up to and including `Entrez` needs to be configured.
 3. Run `make run`. This will download all of the necessary dependencies and run the application.

At the moment, you still need to make an account to use searchrefiner, even locally. The account that you make is a local account and is not the same as the one you might create on another instance of searchrefiner.

## Docker build
searchrefiner can also be run from a preprepared [Dockerfile](./Dockerfile):
1. Create a `config.json` file (see above)
2. Setup the docker image with `make docker-build` or `make docker-build-force` to force a rebuild
3. Run the server with `make docker-run`
4. Or, to deploy on SRA in the `/sites/ecosystem` folder run `pm2 restart ecosystem.config.js`
5. Query site at [http://localhost:4853](http://localhost:4853)


## Documentation

API Endpoint TODO

## Citing

Please cite any references to the searchrefiner project as:

```
@inproceedings{scells2018searchrefiner,
    Author = {Scells, Harrisen and Zuccon, Guido},
    Booktitle = {Proceedings of the 27th ACM International Conference on Information and Knowledge Management},
    Organization = {ACM},
    Title = {searchrefiner: A Query Visualisation and Understanding Tool for Systematic Reviews},
    Year = {2018}
}
```

Please cite any references to any of the automation tools embedded in searchrefiner as:

```
@inproceedings{li2020systematic,
	Author = {Li, Hang and Scells, Harrisen and Zuccon, Guido},
	Booktitle = {Proceedings of the 43rd Internationa SIGIR Conference on Research and Development in Information Retrieval},
	Date-Added = {2020-06-09 13:11:19 +1000},
	Date-Modified = {2020-07-03 15:45:14 +1000},
	Month = {July},
	Pages = {25--30},
	Title = {Systematic Review Automation Tools for End-to-End Query Formulation},
	Year = {2020}
}
```
