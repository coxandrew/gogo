#!/bin/bash

DG_DIR="${HOME}/.do"
PROJECTS_FILE="${DG_DIR}/projects"

function do_usage() {
  echo "
Usage

  do <project>                # go to project directory and execute the .dorc file
  do [-l|--list]              # list all projects
  do [-a|--add] <project>     # add a new project at the current directory
  do [-d|--delete] <project>  # delete a project from the list
  do [-h|--help]              # show this help menu"
}

function do_add() {
  if [[ ! -d "${DG_DIR}" ]]; then
    mkdir -p "${DG_DIR}"
  fi
  local dir="$(pwd -L)"
  echo "${PROJECT}"="${dir}" >> "${DG_DIR}/projects"
  echo "Added ${PROJECT} for ${dir}"
}

function do_delete() {
  local pattern="^${PROJECT}="

  if [[ `cat ${PROJECTS_FILE} | grep "${pattern}"` ]]; then
    sed -i ".bak" "/${pattern}/d" "${PROJECTS_FILE}"
    echo "Deleted project: ${PROJECT}"
  else
    echo "No project named: ${PROJECT}"
    do_list
  fi
}

function do_list() {
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

function do_start() {
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
  unset IFS

  if [[ ! -n "${dir}" ]]; then
    echo "Couldn't find '${PROJECT}'"
    return
  fi
  if [[ ! -d "${dir}" ]]; then
    echo "No such directory: ${dir}"
  fi

  echo "Launching ... ${PROJECT}!!"
  cd "${dir}"

  if [[ -r .dorc ]]; then
    source .dorc
  fi
}

function argument_option() {
  local option=${1}

  if [ "${1:1:0}" != "-" ]; then
    if [ ${2} ]; then
      PROJECT=$2
      do_add
    else
      do_usage
    fi
  fi
}

function standalone_option() {
  local option=${1}

  if [ "${1:1:0}" != "-" ]; then
    if [ ! ${2} ]; then
      case $option in
        --list | -l )
          do_list
          ;;
        --help | -h )
          do_usage
          ;;
      esac
    else
      echo "Unexpected argument: ${2}"
      do_usage
    fi
  fi
}

function do() {
  PROJECT=""
  if [ ! ${1} ]; then
    do_usage
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
            do_delete $PROJECT
            shift
          else
            do_usage
            break
          fi
        fi;;
      *)
        if [[ ${1} && ! ${2} ]]; then
          PROJECT=$1
          do_start $PROJECT
          shift
        else
          if [ ${2} ]; then
            echo "Unexpected argument: ${2}"
          fi
          do_usage
          break
        fi
    esac
  done
}
