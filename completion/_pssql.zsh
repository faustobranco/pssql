#compdef pssql

_pssql() {
    local -a servers
    # O uso de (@f) garante que cada linha do output do Go seja um item da lista
    servers=("${(@f)$((PSSQL_COMPLETION_MODE=true pssql 2>/dev/null))}")

    _arguments \
        '--connect[Server name]:server:compadd -a servers' \
        '--config[Config path]:file:_files -g "*.json"' \
        '--list[List all servers]' \
        '--help[Show help]' \
        '--version[Show version]'
}

_pssql "$@"