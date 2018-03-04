# ghost-coder-iterm

[Ryan Michela](https://github.com/rmichela)'s Ghost Coder implementation in Golang.

## What is Ghost Coder?

Have you ever been in a situation where you have to demo something on the terminal in front of an audience and you start fat-fingering all of a sudden?

With ghost coder, you can save all those terminal commands in a text file and replay them on your terminal with a shortcut!

## Install

You can download the binary directly from the [releases page](https://github.com/sandeepraju/ghost-coder-iterm/releases).

Or you can install it from source using the command below:

```
go get github.com/sandeepraju/ghost-coder-iterm
```

## Usage

- Open iTerm2
- Press ⌥⌘R (or, Session -> Run Coprocess...)
- Enter `ghost-coder-iterm /path/to/file/containing/commands.txt` and click OK

## Credits

Thanks to [Ryan Michela](https://github.com/rmichela) for the idea. His original implementation in Java can be found [here](https://github.com/rmichela/GhostCoder-iTerm).

## License

This project is released under the [MIT License](./LICENSE).
