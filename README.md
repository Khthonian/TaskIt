# TaskIt
TaskIt is a command-line task management tool that allows users to create, complete, delete, and list tasks. Each task is given a name by the user, along with a generated ID number, unique hash value, and completion status. The user can interact with this program through intuitive commands, to enable simple but efficient task management.

## Features
- Creates Tasks
- Complete Tasks
- Delete Tasks 
- Delete All Tasks
- List Tasks
- Automated saving and loading of tasks
- Unique validation of tasks via hashing

## Dependencies
### Go Packages
- `crypto/sha256`
- `encoding/json`
- `flag`
- `fmt`
- `os`

## Usage
### Create Task
```bash
taskit -p create -t {name}
```

### Complete Task
```bash
taskit -p complete -t {name}
taskit -p complete -i {ID}
taskit -p complete -H {hash-value}
```

### Delete Task
```bash
taskit -p delete -t {name}
taskit -p delete -i {ID}
taskit -p delete -H {hash-value}
```

### Delete All Tasks
```bash
taskit -D
```

### List Tasks
```bash
taskit -p list
taskit -p list -s # List tasks with hash value
```

## Installation
1. Clone this repository or download the source code.
2. Run the `installTaskIt.sh` script with `sudo` permissions:
```bash
sudo bash installTaskIt.sh
```
3. Follow the on-screen prompts to complete installation.

## Uninstallation
Once `installTaskIt.sh` has been run, `uninstallTaskIt.sh` will automatically be generated. This script uses hash validation to authenticate that the script does not unintentionally remove directory content that happens to also share the name `taskit`. This script can be generated into the original download directory and can be run using:
```bash
sudo bash uninstallTaskIt.sh
```
## License
This project is licensed under the [MIT License](LICENSE.md).
