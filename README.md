# yaml2json

A small tool for converting YAML files into JSON.

## Installation 

```
go get -u github.com/wakeful/yaml2json
```

## Usage

stdin pipe:
```
cat file.yml | yaml2json
```

or specify a file:
```
yaml2json path/file.yml
```

