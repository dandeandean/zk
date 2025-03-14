_zk()
{
    local cur prev simple_opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    simple_opts="edit graph index init list lsp new tag"

    local notes_list=$(ls $ZK_NOTEBOOK_DIR)

    if [[ ${prev} == "edit" ]] ; then
        COMPREPLY=( $(compgen -W "${notes_list}" ${cur}) )
        return 0
    fi

    # if cur is "", then compgen w/ simple
    if [[ ${cur} == "" ]] ; then
        COMPREPLY=( $(compgen -W "${simple_opts}" ${cur}) )
        return 0
    fi


}

complete -F _zk zk
