rm -r bin
mkdir bin 
go build -o bin/KeePassSync  src/* 
chmod +x bin/KeePassSync