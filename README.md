# ssht

ssht is terminal based UI to manage SSH connections. Purpose of this app is to make it easier and faster to connect to SSH server.


[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)


## Usage/Examples
Running app without argument will run app in TUI mode (if connections exists)

### Add new connection
```bash
ssht add root@192.168.1.12 --port=22 --password=secret --connection-type=password
```

### List connections (no gui)
```bash
ssht list
```
Result:
```
+--------------+-------+--------------+------+----------+-----+----------+
| # Uuid       | ALIAS | HOST         | PORT | USERNAME | KEY | TYPE     |
+--------------+-------+--------------+------+----------+-----+----------+
| 916b9ff13aaf |       | 192.168.1.12 |   22 | root     |     | password |
+--------------+-------+--------------+------+----------+-----+----------+
```

### Set alias for connection
```bash
ssht 916b9ff13aaf someuniquename
```

### Connect
using uuid:
```bash
ssht connect 916b9ff13aaf
```
using alias:
```bash
ssht connect someuniquename
```

### Remove connection
```bash
ssht remove 916b9ff13aaf
```


## Running Tests

To run tests, run the following command

```bash
  make test
```


## Roadmap
- Add secret storage inegration (native to Linux/Windows/MacOS).

