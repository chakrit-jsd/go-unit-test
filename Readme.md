#### Unit Test VS Code Configuration
```json
"go.coverOnSave": true,
"go.coverOnSingleTest": true,
"go.coverageDecorator": {
    "type": "gutter",
    "coveredHighlightColor": "rgba(64,128,128,0.5)",
    "uncoveredHighlightColor": "rgba(128,64,64,0.25)",        
    "coveredGutterStyle": "blockgreen",
    "uncoveredGutterStyle": "blockred"
}
```

#### command

* run test
```
go test module-name/package-name
go test module-name/package-name -v
```
* run subtest
```
go test module-name/package-name -run="func-name/sub-name"
```
* check cover
```
go test module-name/package-name -cover
```
