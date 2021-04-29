# SP-Structured Processor
In some bash scripts, structured data must be created or manipulated. For this there great tools like [jq](https://github.com/stedolan/jq) for JSON and [yq](https://github.com/mikefarah/yq) for Yaml.

In some cases, I was not completely satisfied with the handling. In addition, in some cases I had to install both tools in the pipeline, because I did not have access to the image. Therefore, I created my own tool, which supports multiple formats at the same time and whose handling is more to my taste.

Maybe someone else can also do something with it.

## Table of content
* [Write queries](#write-queries)
* [Commands](#commands)
  * [Concatenate](#concatenate)
  * [Create](#create)
  * [Delete](#delete)
  * [Difference](#difference)
  * [Equal](#equal)
  * [Get](#get)
  * [Merge](#merge)
  * [Put](#put)
  * [Validate](#validate)

## Documentation
This is a MVD 😅.

### Write queries
Writing queries is pretty simple.
Let's take the following example:
```yaml
name: steven
nested: 
  a: A
  b: B
  c: C
  arr: 
    - year: 2019
    - year: 2020
    - year: 2021
arr:
  - 1
  - 2
  - 3
```

To get a value we can use a key path separated by `.` (dots).
E.g. `"name"` will result in `steven` and `"nested.c"` in `C`.

Array indices are defined by `[]` (brackets).
E.g. `"arr.[1]"` will result in `2` and `"nested.arr.[2]"` in `2021`.
It is also possible to get the whole array by using the array key `"arr"` or using a `*` (star) e.g. `"arr.[*]"`.

### Commands

#### Concatenate
This command concatenate structure file's (equal to `cat xyz`).
```bash
sp concatenate --help
```

##### Example
```bash
## Input file1.yaml
# count:
# - 1
# - 2
# - 3
# year: 2021
## Input file2.yaml
# count:
# - 3
# - 2
# - 1
# year: 2020

sp concatenate --file-input "test1.yaml" --file-input "test2.yaml" --output-format "j"

## Output
# {"count":[1,2,3],"year":2021}
# {"count":[3,2,1],"year":2020}
```

#### Create
This command create empty structure file's.
```bash
sp create --help
```

##### Example
```bash
sp create --file-output "z.json" --output-format "j"
sp create --file-output "z.toml" --output-format "t"
sp create --file-output "z.yaml" --output-format "y"
```

#### Delete
This command create empty structure file's.
```bash
sp delete --help
```

#### Example delete field
```bash
## Input
# count:
# - 1
# - 2
# - 3
# year: 2021

sp delete --file-input "file.yaml" --query "year" --file-output "file.yaml" --output-format "y"

## Output
# count:
# - 1
# - 2
# - 3
```
#### Example delete array item
```bash
## Input
# count:
# - 1
# - 2
# - 3
# year: 2021

sp delete --file-input "file.yaml" --query "count.[0]" --file-output "file.yaml" --output-format "y"

## Output
# count:
# - 2
# - 3
# year: 2021
```
#### Example clear array
```bash
## Input
# count:
# - 1
# - 2
# - 3
# year: 2021

sp delete --file-input "file.yaml"  --query "count.[*]" --file-output "file.yaml" --output-format "y"

## Output
# count: [] 
# year: 2021
```

#### Difference
This command compare multiple inputs.
```bash
sp difference --help
```

##### Example
```bash
## Input file1.yaml
# count:
# - 3
# - 2
# - 1
# year: 2020

## Input file2.yaml
# count:
# - 1
# - 2
# - 3
# year: 2021

sp difference --file-input "file1.yaml" --file-input "file2.yaml" --output-format "y"

## Output
#          count:
# ---        - 3
# +++        - 1
#            - 2
# ---        - 1
# +++        - 3
# ---      year: 2020
# +++      year: 2021
```

#### Equal
This command compare multiple inputs.
```bash
sp equal --help
```

##### Example
```bash
sp equal --standard-streams-input '{"year": 2021}' --standard-streams-input '{"year": 2021}'

## Output
# true


sp equal --standard-streams-input '{"year": 2020}' --standard-streams-input '{"year": 2021}'

## Output
# false
```

#### Get
This command return a value/values based on given query
```bash
sp get --help
```

##### Example
```bash
## Input
# override_me: aaa
# count:
# - 1
# - 2
# - 3
# year: 2021

sp get --file-input "file.yaml" --query "year"

## Output
# 2021


sp get --file-input "file.yaml" --query "count"

## Output
# [1 2 3]


sp get --file-input "file.yaml" --query "count.[1]"

## Output
# 2
```

#### Merge
This command merge multiple inputs to single output (previous value is prioritized).
Do not use std with file input, this will result in as mess.
```bash
sp merge --help
```

##### Merge without array append
```bash
## Input file1.yaml
# override_me: aaa
# count:
# - 1
# - 2
# - 3
# year: 2021

## Input file2.yaml
# override_me: bbb
# count:
# - 4
# - 5
# - 6
# year: 2021

sp merge --file-input "file1.yaml" --file-input "file2.yaml" --output-format "y"

## Output
# count:
# - 4
# - 5
# - 6
# override_me: aaa
# year: 2021
```

##### Merge with array append
```bash
## Input file1.yaml
# override_me: aaa
# count:
# - 1
# - 2
# - 3
# year: 2021

## Input file2.yaml
# override_me: bbb
# count:
# - 4
# - 5
# - 6
# year: 2021

sp merge --file-input "file1.yaml" --file-input "file2.yaml" --append-array --output-format "y"

## Output
# count:
# - 1
# - 2
# - 3
# - 4
# - 5
# - 6
# override_me: aaa
# year: 2021
```

#### Put
This command can be used to insert new or override existing fields.
Get argument information by running:
```bash
sp put --help
```

##### Example add key-value
```bash
## Input file.yaml
# year: 2021

sp put --standard-streams-input "{}" --query "year" --item "2021" --file-output "file.yaml" --output-format "y"

## Output file.yaml
# year: 2021
```
```bash
## Input file.yaml
# year: 2021

sp put --file-input "file.yaml" --query "name" --item "steven" --file-output "file.yaml" --output-format "y"

## Output file.yaml
# name: steven
# year: 2021
```

##### Example append item on array
```bash
## Input file.yaml
# year: 2021

sp put --file-input "file.yaml" --query "count.[+]" --item "1" --file-output "file.yaml" --output-format "y"
sp put --file-input "file.yaml" --query "count.[+]" --item "3" --file-output "file.yaml" --output-format "y"
sp put --file-input "file.yaml" --query "count.[+]" --item "5" --file-output "file.yaml" --output-format "y"

## Output file.yaml
# count:
# - 1
# - 3
# - 5
# year: 2021
```

##### Example insert item before after index on array
```bash
## Input file.yaml
# count:
# - 1
# - 3
# - 5
# year: 2021

sp put --file-input "file.yaml" --query "count.[1:]" --item "2" --file-output "file.yaml" --output-format "y"

## Output file.yaml
# count:
# - 1
# - 2
# - 3
# - 5
# year: 2021


sp put --file-input "file.yaml" --query "count.[:2]" --item "4" --file-output "file.yaml" --output-format "y"

## Output file.yaml
# count:
# - 1
# - 2
# - 3
# - 4
# - 5
# year: 2021
```

#### Validate
This command validate if the input has a valid format, specified by argument.
Get argument information by running:
```bash
sp validate --help
```

##### Example JSON stdin
```bash
sp validate --standard-streams-input "{}" --input-format "j"

## Output:
# true
```
##### Example Yaml input file
```bash
## Input file.yaml
# year: 2021

sp validate --file-input "file.yaml" --input-format "y"

## Output:
# true
```
