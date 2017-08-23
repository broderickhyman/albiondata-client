#!/usr/bin/env bash

set -e
echo $CIRCLE_TAG

IFS=$'\n'
badFiles=($(goimports -l $(go list -f {{.Dir}} ./... | grep -v /vendor/)))
unset IFS

if [ ${#badFiles[@]} -eq 0 ]; then
  echo "Congratulations! All Go source files are formatted correctly! :D"
else
  {
    echo "The following files are not formatted properly:"

    for f in "${badFiles[@]}"; do
      echo " - $f"
      echo "ERROR is: "
      goimports -d "${badFiles[@]}"
      echo
    done

    echo
    echo "Please reformat the above files. Use \"make fmt\" and commit the results."
    echo
  }

  false
fi
