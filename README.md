# 💻 devcode-llm-instructions

Centralized Repository for Coding LLM Instructions Leveraging `direnv` for Change Detection

This project is built on top of `direnv`, a shell extension that allows you to load and unload environment variables based on the current directory. The aim of this project is to centralize coding instructions for new projects. To advance or iterate on project specific instructions, you can simply disable the `direnv` hook, make your changes, and then re-enable it.

For more information on `direnv`, visit [direnv.net](https://direnv.net).

🧪 This project is experimental 🚧

## How to Use

1. Download/Install `direnv`:
   - For macOS: `brew install direnv`
   - For Linux: Follow instructions on [direnv's website](https://direnv.net/docs/installation.html)
   - For Windows: Use WSL or follow [direnv's Windows guide](https://direnv.net/docs/installation.html#windows)
2. Create an empty `.envrc` file in your project directory:
   ```bash
   touch .envrc
   ```
3. If you're using copilot, you can copy the contents of the `example-direnv.sh` file from this repository into your `.envrc` file:
   ```bash
   cp example-direnv.sh .envrc
   ```
4. Allow `direnv` to load the `.envrc` file:
   ```bash
   direnv allow .
   ```
5. Now, whenever you enter the directory, `direnv` will automatically load the environment variables defined in `.envrc` and check for new instructions in this repository.
6. If you want to make changes to the instructions, you can disable `direnv` from syncing these instructions by removing the `.envrc` file or commenting out the relevant lines in it. After making your changes, you can re-enable `direnv` by running:
   ```bash
   direnv allow .
   ```
