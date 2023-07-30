# Mercury

[![Build](../../actions/workflows/build.yaml/badge.svg)](../../actions/workflows/build.yaml)

Mercury is a command-line tool designed to streamline your Git workflow. It uses generative AI to automate the process of creating Pull Request (PR) and commit messages. Additionally, it offers the unique ability to reword empty commits with meaningful descriptions, enhancing the quality of your commit history.

## Usage

```
Usage: mercury command [OPTIONS]

Commands:
  pr
  commit
  reword

Options:
  -u --set-upstream
  -v --verbose
  --repo
  --staged
  --create
```

## Commands

- `mercury pr`: Generate PR messages for your commits.
- `mercury commit`: Automatically create meaningful commit messages.
- `mercury reword`: Turn empty commits into informative ones.

For detailed options and usage, please refer to the command descriptions in the Usage section above.

## References

- [go-git](https://github.com/go-git/go-git)
- [go-openai](https://github.com/sashabaranov/go-openai)
- [godotenv](https://github.com/joho/godotenv)
