# gha-docs

gha-docs is a command-line tool that generates a README.md file for your GitHub Actions by reading the action.yml file. This tool makes it easy to create a well-documented GitHub Action by automatically populating the README.md file with information such as the action name, description, inputs, and outputs.

To use gha-docs, simply run the following command in your terminal:

```text
gha-docs is a cli tool to generate documentation for github actions

Usage:
  gha-docs [flags]

Flags:
  -f, --action-yaml string   action file name (default "action.yml")
  -d, --debug                enable debug mode
  -h, --help                 help for gha-docs
```

Once you run the command, gha-docs will generate a README.md file in the same directory as your action.yml file. The README.md file will be populated with the following information:

The action name and description
The inputs and outputs of the action, including their descriptions and default values (if any)
This makes it easy for users of your GitHub Action to understand what your action does and how to use it.

Overall, gha-docs is a simple and effective tool for creating well-documented GitHub Actions.
