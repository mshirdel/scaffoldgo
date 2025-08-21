package config

var _builtinConfig = `
server:
  address: "0.0.0.0:{{.Port}}"
  idle-timeout: 2m
  read-timeout: 1s
  write-timeout: 3s
logging:
  level: debug
`
