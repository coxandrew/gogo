#!/bin/bash

DG_DIR="${HOME}/.go"
PROJECTS_FILE="${DG_DIR}/projects"

function go_usage() {
  echo "
Usage

  go <project>                # go to project
  go [-l|--list]              # list all projects
  go [-a|--add] <project>     # add a new project at the current directory
  go [-d|--delete] <project>  # delete a project from the list
  go [-h|--help]              # show this help menu"
}

function go_add() {
  if [[ ! -d "${DG_DIR}" ]]; then
    mkdir -p "${DG_DIR}"
  fi
  local dir="$(pwd -L)"
  echo "${PROJECT}"="${dir}" >> "${DG_DIR}/projects"
  echo "Added ${PROJECT} for ${dir}"
}

function go_delete() {
  local pattern="^${PROJECT}="

  if [[ `cat ${PROJECTS_FILE} | grep "${pattern}"` ]]; then
    sed -i ".bak" "/${pattern}/d" "${PROJECTS_FILE}"
    echo "Deleted project: ${PROJECT}"
  else
    echo "No project named: ${PROJECT}"
    go_list
  fi
}

function go_list() {
  echo "
Available projects:
"

  if [[ -r "${PROJECTS_FILE}" ]]; then
    local IFS=$'\n'
    for line in $(cat "${DG_DIR}/projects"); do
      IFS=$'='
      local -a project=(${line})
      echo `printf '%-20s' "${project[0]}"` "- ${project[1]}"
    done
  else
    echo "You haven't added anything yet."
    return
  fi
}

function go_start() {
  local dir=''
  local project=''
  if [[ ! -r "${PROJECTS_FILE}" ]]; then
    echo "You haven't added anything yet."
    return
  fi
  local IFS=$'\n'
  for line in $(cat "${DG_DIR}/projects"); do
    IFS=$'='
    local -a project=(${line})

    if [[ "${project[0]}" = "${PROJECT}" ]]; then
      dir=${project[1]}
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

function argument_option() {
  local option=${1}

  if [ "${1:1:0}" != "-" ]; then
    if [ ${2} ]; then
      PROJECT=$2
      go_add
    else
      go_usage
    fi
  fi
}

function standalone_option() {
  local option=${1}

  if [ "${1:1:0}" != "-" ]; then
    if [ ! ${2} ]; then
      case $option in
        --list | -l )
          go_list
          ;;
        --help | -h )
          go_usage
          ;;
      esac
    else
      echo "Unexpected argument: ${2}"
      go_usage
    fi
  fi
}

function go() {
  PROJECT=""
  if [ ! ${1} ]; then
    go_usage
  fi
  until [ -z "$1" ]; do
    case $1 in
      --list | -l | --help | -h )
        standalone_option $*
        shift;;
      --add | -a )
        argument_option $*
        break;;
      --delete | -d )
        shift
        if [ "${1:1:0}" != "-" ]; then
          if [ ${1} ]; then
            PROJECT=$1
            go_delete $PROJECT
            shift
          else
            go_usage
            break
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
          go_usage
          break
        fi
    esac
  done
}
