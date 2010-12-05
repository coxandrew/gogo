#!/bin/bash

# TODO
# * Add a --help section

DG_DIR="${HOME}/.go"
PROJECTS_FILE="${DG_DIR}/projects"

function usage() {
  echo "
Usage

  go <project>             # go to project
  go [-l|--list]           # list all projects
  go [-a|--add] <project>  # add a new project at the current directory"
}

function go_add() {
  if [[ ! -d "${DG_DIR}" ]]; then
    mkdir -p "${DG_DIR}"
  fi
  local dir="$(pwd -L)"
  echo "${PROJECT}"="${dir}" >> "${DG_DIR}/projects"
  echo "Added ${PROJECT} for ${dir}"
}

function go_list() {
  echo "Listing projects ...
"

  if [[ -r "${PROJECTS_FILE}" ]]; then
    cat ${PROJECTS_FILE}
  else
    echo "You haven't added anything yet."
    return
  fi
}

function go_start() {
  local dir=''
  local project=''
  if [[ ! -r "${DG_DIR}/projects" ]]; then
    echo "You haven't added anything yet."
    return
  fi
  local IFS=$'\n'
  for line in $(cat "${DG_DIR}/projects"); do
    IFS=$'\t'
    local -a line=(${line})

    project=`echo ${line} | cut -d "=" -f 1`
    if [[ "${project}" = "${PROJECT}" ]]; then
      dir=`echo ${line} | cut -d "=" -f 2`
    fi
  done

  if [[ ! -n "${dir}" ]]; then
    echo "Couldn't find '${PROJECT}'"
    return
  fi
  if [[ ! -d "${dir}" ]]; then
    echo "No such directory: ${dir}"
  fi

  echo "Go go gadget ${PROJECT}!!"
  cd "${dir}"

  if [[ -r .gorc ]]; then
    source .gorc
  fi
}

function go() {
  PROJECT=""
  if [ ! ${1} ]; then
    usage
  fi
  until [ -z "$1" ]; do
    case $1 in
      --list | -l )
        shift
        if [ "${1:1:0}" != "-" ]; then
          if [ ! ${1} ]; then
            go_list
          else
            echo "Unexpected argument: ${1}"
            usage
            exit
          fi
        fi;;
      --add | -a )
        shift
        if [ "${1:1:0}" != "-" ]; then
          if [ ${1} ]; then
            PROJECT=$1
            go_add $PROJECT
            shift
          else
            usage
            exit
          fi
        fi;;
      *)
        if [[ ${1} && ! ${2} ]]; then
          PROJECT=$1
          go_start $PROJECT
          shift
        else
          if [ ${2} ]; then
            echo "Unexpected argument: ${2}"
          fi
          usage
          exit 1
        fi
    esac
  done
}