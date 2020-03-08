# goCfgParse
This is a small program made to simplify configuration file loading for golang based projects.

## Usage
All tunables are stored in nested structs (underneat the below parent) based on each toml field.
```
type tomlConfig struct {
}
```
### Adding additional sections
In order to add an additional section, you would create a struct
```
type api struct {
  Hostname string
  Address string
  Port string
}
```
and then add your struct to tomlConfig
```
type tomlConfig struct {
  Api api
}
```
To access "Hostname" in your program, simply reference `config.Api.Hostname`

### Loading into memory
In order to load the configration object into your program, simply execute the following (in main)
```
loadConfig(&config)
```

### Selecting your configuration file
By default, the configuration file path is stored in an environment variable called `GO_TOML_CONFIG` (see example below)

### Example
```
export GO_TOML_CONFIG='/home/skippy/repos/goCfgParse/config.toml'; go run .
2020/03/08 22:54:07 [Config] Successfully loaded /home/skippy/repos/goCfgParse/config.toml
2020/03/08 22:54:07 Hostname: testhost.example.domain
2020/03/08 22:54:07 Address: 127.0.0.1
2020/03/08 22:54:07 Port: 10000
```
