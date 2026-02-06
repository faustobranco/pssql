#compdef pssql

_pssql() {
    local -a servers
    servers=(${(f)"$(PSSQL_COMPLETION_MODE=true pssql 2>/dev/null)"})

    _arguments \
        '--connect[Server name]:server:compadd -a servers' \
        '--config[Config path]:file:_files -g "*.json"' \
        '--list[List all servers]' \
        '--help[Show help]' \
        '--version[Show version]'
}

_pssql "$@"