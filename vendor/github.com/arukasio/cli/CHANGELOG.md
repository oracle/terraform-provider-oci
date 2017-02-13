# CHANGELOG

### v0.2.0 (Feburary 7, 2017)

- [#48](https://github.com/arukasio/cli/pull/48) [DEPRECATION] `Name` field of `Container` struct.
  - Use `Container.ArukasDomain` instead.
  - `Container.Name` will be removed soon.
  - Deprecated `aruaks run --name`. Use `--arukas-domain`
- [#44](https://github.com/arukasio/cli/pull/44)[REFACTOR]
  Restructure files and directories to exclude unnecessary dependencies for clients.
