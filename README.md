# my-cli
exploring golang with cobra cli

# setup
After clone the repo, build the cli and to execute the cli as per the below instructions:

### to build the cli
```
$ go build -o my-cli main.go
```

## to execute the cli on terminal
```
$ ./my-cli --help // to get help
$ ./my-cli -m "kishor" // insert username for greeting
$ ./my-cli -m "kishor" -i // show username with a valid username through -i flag
$ ./my-cli -v  // check the version
$ ./my-cli echo working on echo // printing echo
$ ./my-cli uppercase kishor // make the letters into uppercase

$ ./mycli notepad -n "text is here" -d // create a default file
$ ./mycli notepad -n "text is here" -f "myfile.txt" // dynamic file 
```

