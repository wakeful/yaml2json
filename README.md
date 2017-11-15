# yaml2json

A small tool for converting YAML files into JSON.

## Installation 

macOS
```
$ brew tap wakeful/selection
$ brew install yaml2json
```

Linux
```
curl -Lo yaml2json https://github.com/wakeful/yaml2json/releases/download/0.1.0/yaml2json-linux-amd64 && chmod +x yaml2json && sudo mv yaml2json /usr/local/bin/
```

src
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

