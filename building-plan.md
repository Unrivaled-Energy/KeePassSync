#Steps to build

## Setting up Minio

configure minio to create events for s3:ObjectCreated:Put

## Linux development

watch changes on local db
[ ] if local file changes, stop watching ,reupload to minio, resume watching
[ ] if remote file changes, stop watching,download db, resume watching

## Android Development

info - http://stackoverflow.com/questions/48034748/ddg#48238804

same logic as above

potential additions
[ ] encrypt db with gpg key before uploading
[ ] decrypt once downloaded
