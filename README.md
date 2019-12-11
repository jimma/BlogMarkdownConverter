# Overview
B2M is small utility tool to convert jboss developer blog like https://developer.jboss.org/en/resteasy/blog/2019/10/08/looking-to-the-future
to markdown format. It accepts the blog index url and convert all the blog entries in this index page to markdown file.
```
b2m run --help
Run converter to convert blog entry to Markdown file
Usage:
  b2m run <blog url> [flags]
Flags:
  -h, --help            help for run
  -o, --output string   Output directory
```

## Install 
b2m tool can be installed with GNU ***make install****
or execute this command  
```
go install github.com/jimma/blogmarkdownconverter/cmd/b2m
```
## Example
To convet blogs, simply gives ***b2m*** blog index url and output directory
```
b2m run https://developer.jboss.org/en/resteasy/blog -o ./output
b2m run https://developer.jboss.org/en/resteasy/blog?start=15 -o ./output
``` 

