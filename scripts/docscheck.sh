#!/usr/bin/env bash

docs=$(ls website/docs/**/*.markdown)
error=false

for doc in $docs; do
  dirname=$(dirname "$doc")
  category=$(basename "$dirname")


  case "$category" in
    "guides")
      # Guides require a page_title
      if ! grep "^page_title: " "$doc" > /dev/null; then
        echo "Guide is missing a page_title: $doc"
        error=true
      fi
      ;;

    "d" | "r")
      # Resources and data sources require valid YAML front matter. The Registry
      # uses the front matter to place a page under its service subcategory.
      IFS= read -r first_line < "$doc"
      if [[ $first_line != "---" ]]; then
        echo "Doc must begin with YAML front matter (---): $doc"
        error=true
      fi

      # Resources and data sources require a subcategory.
      if ! grep "^subcategory: " "$doc" > /dev/null; then
        echo "Doc is missing a subcategory: $doc"
        error=true
      fi
      ;;

    *)
      error=true
      echo "Unknown category \"$category\". " \
        "Docs can only exist in r/, d/, or guides/ folders."
      ;;
  esac
done

# Terraform Registry only reads a page's subcategory when its complete YAML
# front matter is valid. Checking for individual lines is not sufficient.
if ! ruby -ryaml -e '
  failed = false
  ARGV.each do |path|
    content = File.binread(path)
    front_matter = content.match(/\A---\r?\n(.*?)\r?\n---(?:\r?\n|\z)/m)
    unless front_matter
      warn "Doc has invalid YAML front matter: #{path}"
      failed = true
      next
    end

    begin
      data = YAML.load(front_matter[1])
      unless data.is_a?(Hash) && data["subcategory"]
        warn "Doc is missing a front-matter subcategory: #{path}"
        failed = true
      end
    rescue Psych::Exception => error
      warn "Doc has invalid YAML front matter: #{path}: #{error.message.lines.first.strip}"
      failed = true
    end
  end
  exit 1 if failed
' website/docs/d/*.markdown website/docs/r/*.markdown; then
  error=true
fi

if $error; then
  exit 1
fi

exit 0