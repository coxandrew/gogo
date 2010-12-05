# TODO
# * Add Bash TextMate bundle
# * Add to Github
# * Add a Readme.markdown
# * Add a --help section
# * Use flags instead of separate functions like `go --add recycling` and `go --list`

DG_DIR="${HOME}/.go"

# Shows the list of saved projects.
function show {
  if [[ -r "${DG_DIR}/projects" ]]; then
    cat "${DG_DIR}/projects"
  else
    echo "You haven't added anything yet."
    return
  fi
}

# Goes to the directory and performs any commands.
function go {
  local name="${1}"
  local dir=''
  if [[ -z "${name}" ]]; then
    echo "You need to specify a name."
    return
  fi
  if [[ ! -r "${DG_DIR}/projects" ]]; then
    echo "You haven't added anything yet."
    return
  fi
  local IFS=$'\n'
  for line in $(cat "${DG_DIR}/projects"); do
    IFS=$'\t'
    local -a line=(${line})
    if [[ "${line[0]}" = "${name}" ]]; then
      dir="${line[1]}"
    fi
  done

  if [[ ! -n "${dir}" ]]; then
    echo "Couldn't find '${name}'"
    return
  fi
  if [[ ! -d "${dir}" ]]; then
    echo "No such directory: ${dir}"
  fi

  cd "${dir}"

  if [[ -r .gorc ]]; then
    source .gorc
  fi
}

# Adds the current directory to your list.
function add {
  local name="${1:-}"
  if [[ -z "${name}" ]]; then
    echo "You need to specify a name."
    return
  fi
  if [[ ! -d "${DG_DIR}" ]]; then
    mkdir -p "${DG_DIR}"
  fi
  local dir="$(pwd -L)"
  echo "${name}"$'\t'"${dir}" >> "${DG_DIR}/projects"
  echo "Added ${name} for ${dir}"
}