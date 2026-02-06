_pssql_completions() {
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    
    opts="--connect --config --list --help --version"

    case "${prev}" in
        --connect)
            local IFS=$'\n'
            local servers=$(PSSQL_COMPLETION_MODE=true pssql 2>/dev/null)
            COMPREPLY=( $(compgen -W "${servers}" -- "${cur}") )
            return 0
            ;;
        --config)
            COMPREPLY=( $(compgen -f -X '!*.json' -- "${cur}") )
            return 0
            ;;
    esac

    if [[ ${cur} == -* ]] ; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    fi
}

complete -F _pssql_completions pssql