<div align="center">
  <h1>Convert csv and json</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/convert-csv-json"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/convert-csv-json"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/convert-csv-json"/>

<a href="#project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run-project">Run Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#technologies">Technologies</a>

</div>

## Project

Convert files csv in json and files json in csv

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)

## Run Project

Run this commands in your terminal:

```sh
## run this command to install dependencies:
❯ go mod tidy

## run this command to convert files csv to json
❯ go run main.go convert-to-json -p="../test.csv" -s=";" -d="../tmp/test.json"

## run this command to convert files json to csv
❯ go run main.go convert-to-csv -p="../test.json" -d="../test.csv"

```

## Technologies

- golang
- cobra-cli

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
