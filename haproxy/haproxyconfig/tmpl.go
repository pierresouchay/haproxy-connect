package haproxyconfig

const tmplStr = `
global
    stats socket {{.SocketPath}} mode 600 level admin expose-fd listeners
    stats timeout 2m

{{range $fe := .Frontends}}
frontend {{$fe.Name}}
    mode tcp
    bind {{$fe.BindAddr}}:{{$fe.BindPort}}{{if $fe.TLS}} ssl crt {{$fe.ServerCRTPath}} ca-file {{$fe.ClientCAPath}} verify required{{end}}
    option tcplog
    timeout client 1m
	default_backend {{$fe.DefaultBackend}}
{{end}}

{{range $be := .Backends}}
backend {{$be.Name}}
    mode tcp
    option redispatch
    balance roundrobin
    timeout connect 10s
	timeout server 1m
	{{range $s := $be.Servers}}
	server {{$s.Name}} {{$s.Host}}:{{$s.Port}}{{if $s.TLS}} ssl crt {{$s.ClientCRTPath}} ca-file {{$s.ServerCAPath}} verify required{{end}}
	{{end}}
{{end}}
`
