# Komiter

## What does it do?
it generates git commit message by reading git diff using AI, now its only support gemini by default

Komiter is Committer in indonesian



## How to use 

### Linux
1. set the env

```sh
API_URL=YOUR_LLM_API_URL
API_TOKEN=YOUR_LLM_TOKEN
```

2. clone the repo
```sh
git clone https://github.com/hmdnu/komiter.git
```
3. build 
```go
go build -o ./build/komiter 
```
4. make alias, it depends on what shell are you using
```sh
alias komiter="PATH_TO_BIN"
```
make sure your works is already staged

5. generate staged commit message
```sh
komiter
```
result example
```
feat(cli): Initialize AI-powered commit message generation tool
```

### Flags

1. copy to clipboard
```sh
komiter -c 
```

2. disply help
```sh
komiter -h
```

for windows and mac try to figure it out yourself
