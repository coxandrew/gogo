## Overview

Go is a simple bash script that allows you to quickly launch and initialize a project 
so you can start working immediately.

With GoGo, do away with the hassle of remembering where your project lives, firing up 
your text editor, opening a browser, etc.

## Installation

Add the following to your `.bashrc` or `.bash_profile`:

	. ~/.gogo/gogo
	. ~/.gogo/projects
	if [ -f ~/.gogo/projects ]; then
	    . ~/.gogo/projects
	fi
	
You may optionally want to alias 'go' for the 'gogo' script if you aren't already using the 'go' scripting language:

	alias go='gogo'

## Usage

### Add a project

    $ cd ~/Projects/skillbonsai
    $ gogo --add skillbonsai

### Start working on a project

    $ gogo skillbonsai

### Show existing projects

    $ gogo --list

### Go to a project directory

    $ cd skillbonsai    # once it has been added to your list, you'll have an alias to it

## Custom project initialization steps

To customize what initialization steps your project performs, just create a `.gorc` file in your project's directory and this will be sourced after Go takes you to your project. For example, if you just wanted to open up Textmate in your project directory, you could simply have the line:

    mate .
    
Then when you launch go, you'll navigate to the directory and open your text editor.

## Coming soon

* Paging
* Sorting by alpha or last used

