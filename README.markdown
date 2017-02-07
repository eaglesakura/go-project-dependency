# prjdep

[![CircleCI](https://circleci.com/gh/eaglesakura/prjdep/tree/master.svg?style=svg&circle-token=a407077b50da0ed24be75694854e6ca3a4f747ee)](https://circleci.com/gh/eaglesakura/prjdep/tree/master)

`prjdep` is golang dependencies management tool.


## install

```
go get github.com/eaglesakura/prjdep
```

## Save golang dependencies

```
cd /path/to/project/dir
prjdep init

# generate `dependencies.json`
```

Generated `dependencies.json` file

```
{
  "Repositories": [
    {
      "ImportPath": "github.com/stretchr/testify",
      "Rev": "4d4bfba8f1d1027c4fdbe371823030df51419987",
      "Lang": "golang"
    },
    {
      "ImportPath": "github.com/urfave/cli",
      "Rev": "347a9884a87374d000eec7e6445a34487c1f4a2b",
      "Lang": "golang"
    }
  ]
}
```

## Restore golang dependencies

```
cd /path/to/project/dir
prjdep restore

# go get -> git checkout -> go install
```

## for Circle.yml

[sample circle.yml](./circle.yml)

```
dependencies:
  override:
    - go get github.com/eaglesakura/prjdep
    - prjdep restore
```

## LICENSE

see [LICENSE.txt](./LICENSE.txt)
