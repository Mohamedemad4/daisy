# daisy: `&&` chain commands across terminal sessions
but it also might extend to become a tool to make "dirty" N-terminal session sysadmin/hacking work cleaner and more
reproducible.

# this is DDD (documentation driven development lol; done features will be marked as such)



# Usage (normal command)
### Terminal 1
- `dchn cmd1 apt update` **OR** `apt update| dchn cmd1`

### Terminal 2
- `dchn cmd1 apt upgrade` **OR** `dchn -after="apt update" apt upgrade`

this will wait for `cmd1` to finish then execute the second command.

# Usage (ZSH plugin)
### Terminal 1
- `cmd1: apt update` **OR** `apt update @@cmd1`

### Terminal 2
- `cmd1:apt upgrade` **OR** `apt update @@cmd1` (since cmd1 already exists it will wait to run afterwards)

you can also pass arguments in plugin mode by adding them after **_** 
so for example

- `cmd1:apt upgrade` becomes `cmd1_xor:apt upgrade`



- just typing `dchn` will also print out all the crashed commands terminals (if they didn't get to terminate cleanly)


## Motivation
maybe this is a really specific use case
but I usually find myself using `sleep` to wait for a command to execute in another terminal session.



## Arguments
- `-m` Mode can be `and`,`or`,`xor`,`not` 
    - for example `and` mode and will only run the second command if the first command exited successfully
    - mode `not` will only run the 2nd command if the first command fails
    - mode `or` doesn't care which is the **default**

# Installation
// Todo some `wget https://install.bash | bash` script


## Manual compile and install 
- `go get github.com/fatih/color` 
- `go build .`
- `cp dchn /usr/bin`

## How it works
- it stores the cmd identifier in a file (how can we make this multi user friendly? or at least root user friendly)
- and just waits for it to be removed from the registry?

it looks like 
~/.dchn/\<cmdID\>.json
- if no cmdID the cmdID is just the md5 hash of the cmd


the cmdID.json contains:
- the cmd
- the cmd id 
- timestamp of start of the cmd
- termination state (if so who cleans up the fucking file? the `dchn` right after? yes...)
- state : EXECUTING,WAITING,DONE

and whenever dchn runs it just checks all the files for orphaned files older than 24 hours it alerts the user and deletes them


modes work by checking the termination flag from the cmdID.json file
and after they are done they delete the cmdID.json of their parent

### The zsh plugin
- it's just zsh script that reads the command and if it fits the regex parses the args and feeds them here

# todo
- manpage also


### color scheme 
blue -> info
red -> critical
purple -> help
