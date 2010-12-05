## Overview

Go is a simple bash script that allows you to quickly launch and initialize a project so you can start working immediately.

With Go, do away with the hassle of remembering where your project lives, firing up your text editor, opening a browser, etc.

## Installation

Add the following to your `.bashrc` or `.bash_profile`:

    source ~/.go/go        # adds to the 'go' script to your shell
    source ~/.go/projects  # adds aliases for each of your projects
    shopt -s cdable_vars   # set so that no '$' is required when cd'ing

## Usage

### Add a project

    $ cd ~/Projects/isitrecyclingweek
    $ go --add recycling

### Start working on a project

    $ go recycling

### Show existing projects

    $ go --list

## Custom project initialization steps

To customize what initialization steps your project performs, just create a `.gorc` file in your project's directory and this will be sourced after Go takes you to your project.