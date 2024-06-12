#!/bin/bash

if [ -f ~/.git-completion.bash ]; then
  . ~/.git-completion.bash
fi

PS1="\u:\w\$ \[\]"
